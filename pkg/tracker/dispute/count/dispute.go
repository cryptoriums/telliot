// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package count

import (
	"context"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "disputeCountTracker"

const reorgEventWait = 3 * time.Minute

type Config struct {
	LogLevel string
}

type Dispute struct {
	logger        log.Logger
	ctx           context.Context
	close         context.CancelFunc
	cfg           Config
	client        *ethclient.Client
	contract      *contracts.ITellor
	disputes      *prometheus.GaugeVec
	pendingAppend map[string]context.CancelFunc
	mtx           sync.Mutex
}

func New(
	logger log.Logger,
	ctx context.Context,
	cfg Config,
	client *ethclient.Client,
	contract *contracts.ITellor,
) (*Dispute, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)
	ctx, close := context.WithCancel(ctx)

	return &Dispute{
		client:        client,
		contract:      contract,
		cfg:           cfg,
		ctx:           ctx,
		close:         close,
		logger:        logger,
		pendingAppend: make(map[string]context.CancelFunc),
		disputes: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "disputes",
			Help:      "All open disputes indexed by the disputed addresses",
		},
			[]string{"addr"},
		),
	}, nil
}

func (self *Dispute) Start() error {
	queryDays := int64(10)
	logs, err := contracts.GetDisputeLogs(self.ctx, self.client, self.contract.Address, int(queryDays))
	if err != nil {
		return errors.Wrap(err, "getting latest disputes")
	}

	for _, log := range logs {
		self.disputes.With(prometheus.Labels{"addr": log.Miner.String()}).Inc()
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	logger := log.With(self.logger, "contract", "tellor")

	var sub event.Subscription
	events := make(chan *tellor.TellorNewDispute)

	for {
		select {
		case <-self.ctx.Done():
			return nil
		default:
		}
		sub, err = self.newSub(events)
		if err != nil {
			level.Error(logger).Log("msg", "initial subscribing to events failed")
			<-ticker.C
			continue
		}
		break
	}

	for {
		select {
		case <-self.ctx.Done():
			return nil
		case err := <-sub.Err():
			if err != nil {
				level.Error(logger).Log(
					"msg",
					"subscription error",
					"err", err)
			}

			// Trying to resubscribe until it succeeds.
			for {
				select {
				case <-self.ctx.Done():
					return nil
				default:
				}
				sub, err = self.newSub(events)
				if err != nil {
					level.Error(logger).Log("msg", "re-subscribing to events failed", "err", err)
					<-ticker.C
					continue
				}
				break
			}
			level.Info(logger).Log("msg", "re-subscribed to events")
		case event := <-events:
			level.Debug(self.logger).Log(
				"msg", "new event",
				"removed", event.Raw.Removed,
				"hash", event.Raw.TxHash.String()[:8],
				"miner", event.Miner.String()[:8],
			)
			if event.Raw.Removed {
				self.removePending(event)
			}

			ctx, cncl := context.WithCancel(self.ctx)
			self.mtx.Lock()
			self.pendingAppend[event.Raw.TxHash.String()] = cncl
			self.mtx.Unlock()

			go func(ctx context.Context, event *tellor.TellorNewDispute) {
				ticker := time.NewTicker(reorgEventWait) // Wait this long for any re-org events that can cancel this append.
				defer ticker.Stop()

				select {
				case <-ticker.C:
					self.disputes.With(prometheus.Labels{"addr": event.Miner.String()}).Inc()
					self.removePending(event)
				case <-ctx.Done():
					level.Debug(self.logger).Log("msg", "append canceled", "hash", event.Raw.TxHash.String()[:8])
					return
				}
			}(ctx, event)
		}
	}
}

// removePending is extracted in a separate function to use defer for unlocking the mutex and
// avoid forgetting to unlock it for early returns.
func (self *Dispute) removePending(event *tellor.TellorNewDispute) {
	self.mtx.Lock()
	defer self.mtx.Unlock()
	pendingCncl, ok := self.pendingAppend[event.Raw.TxHash.String()]
	if !ok {
		level.Error(self.logger).Log("msg", "missing pending TX for removed event")
		return
	}
	pendingCncl()
	delete(self.pendingAppend, event.Raw.TxHash.String())
}

func (self *Dispute) Stop() {
	self.close()
}

func (self *Dispute) newSub(output chan *tellor.TellorNewDispute) (event.Subscription, error) {
	tellorFilterer, err := tellor.NewTellorFilterer(self.contract.Address, self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting instance")
	}
	sub, err := tellorFilterer.WatchNewDispute(&bind.WatchOpts{Context: self.ctx}, output, nil, nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting channel")
	}
	return sub, nil
}
