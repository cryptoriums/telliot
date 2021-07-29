// Copyright (c) The Tellor Authors.
// Licensed under the MIT License.

package profit

import (
	"context"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/bluele/gcache"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/event"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/tellor-io/telliot/pkg/contracts"
	"github.com/tellor-io/telliot/pkg/contracts/tellor"
	"github.com/tellor-io/telliot/pkg/ethereum"
	"github.com/tellor-io/telliot/pkg/gas_estimator"
	"github.com/tellor-io/telliot/pkg/logging"
)

const ComponentName = "profitTracker"

const DefaultRetry = 30 * time.Second

type Config struct {
	LogLevel string
}

type ProfitTracker struct {
	netID            *big.Int
	client           *ethclient.Client
	logger           log.Logger
	contractInstance *contracts.ITellor
	abi              abi.ABI
	ctx              context.Context
	stop             context.CancelFunc

	// All registered addresses to track in different formats.
	// They hold the same addresses, just in a different format
	// to make it easyer to pass around.
	accounts []*ethereum.Account
	addrs    []common.Address
	addrsMap map[common.Address]struct{} // The same as above but used for quick matching.

	gasEstimator    gas_estimator.GasEstimator
	lastGasEstimate uint64

	lastSubmit time.Time

	cacheTXsProfit     gcache.Cache
	cacheTXsCost       gcache.Cache
	cacheTXsCostFailed gcache.Cache
	lastFailedBlock    int64

	submitProfit *prometheus.GaugeVec
	submitCost   *prometheus.GaugeVec
	balances     *prometheus.GaugeVec

	submitCount     *prometheus.CounterVec
	submitFailCount *prometheus.CounterVec
	submitValue     *prometheus.GaugeVec

	submitGasUsageEstimated *prometheus.GaugeVec
	submitGasUsageActual    *prometheus.GaugeVec
}

func NewProfitTracker(
	logger log.Logger,
	ctx context.Context,
	cfg Config,
	client *ethclient.Client,
	gasEstimator gas_estimator.GasEstimator,
	contractInstance *contracts.ITellor,
	accounts []*ethereum.Account,
) (*ProfitTracker, error) {
	logger, err := logging.ApplyFilter(cfg.LogLevel, logger)
	if err != nil {
		return nil, errors.Wrap(err, "apply filter logger")
	}
	logger = log.With(logger, "component", ComponentName)

	if gasEstimator == nil {
		gasEstimator = gas_estimator.NewDefault(contractInstance, client)
	}

	abi, err := abi.JSON(strings.NewReader(tellor.TellorABI))
	if err != nil {
		return nil, errors.Wrap(err, "abi read")
	}

	addrsMap := make(map[common.Address]struct{})
	var addrs []common.Address
	for _, acc := range accounts {
		addrsMap[acc.Address] = struct{}{}
		addrs = append(addrs, acc.Address)
	}

	netID, err := client.NetworkID(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "get nerwork ID")
	}

	ctx, cncl := context.WithCancel(ctx)

	return &ProfitTracker{
		netID:            netID,
		client:           client,
		logger:           logger,
		contractInstance: contractInstance,
		abi:              abi,
		accounts:         accounts,
		addrs:            addrs,
		addrsMap:         addrsMap,
		ctx:              ctx,
		stop:             cncl,
		gasEstimator:     gasEstimator,

		cacheTXsProfit:     gcache.New(50).LRU().Build(),
		cacheTXsCost:       gcache.New(50).LRU().Build(),
		cacheTXsCostFailed: gcache.New(20).LRU().Build(),

		submitProfit: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "submit_profit",
			Help:      "Accumulated TRB amount from rewards for all registered addresses",
		},
			[]string{"addr"},
		),
		submitCost: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "submit_cost",
			Help:      "Accumulated ETH cost from the submits for all registered addresses",
		},
			[]string{"addr"},
		),
		balances: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "balances",
			Help:      "Current token balances for all registered addresses",
		},
			[]string{"addr", "token"},
		),
		submitCount: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "submit_total",
			Help:      "The total number of submitted solutions",
		},
			[]string{"addr"},
		),
		submitFailCount: promauto.NewCounterVec(prometheus.CounterOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "submit_fails_total",
			Help:      "The total number of failed submits",
		},
			[]string{"addr"},
		),
		submitValue: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "submit_value",
			Help:      "The submitted value",
		},
			[]string{"id", "addr"},
		),
		submitGasUsageEstimated: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "submit_gas_usage_estimated",
			Help:      "The submitted estimation for the gas usage",
		},
			[]string{"slot"},
		),
		submitGasUsageActual: promauto.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "telliot",
			Subsystem: ComponentName,
			Name:      "submit_gas_usage_actual",
			Help:      "The submitted actual result for the gas usage",
		},
			[]string{"slot"},
		),
	}, nil
}

func (self *ProfitTracker) Start() error {
	level.Info(self.logger).Log("msg", "starting")

	for _, addr := range self.addrs {
		balance, err := self.getTRBBalance(addr)
		if err != nil {
			level.Error(self.logger).Log("msg", "getting initial TRB balance", "addr", addr.String(), "err", err)
		}
		level.Info(self.logger).Log("msg", "initial TRB balance", "addr", addr.String(), "balance", balance)
		self.balances.With(prometheus.Labels{"addr": addr.String(), "token": "TRB"}).Set(balance)
	}

	for _, addr := range self.addrs {
		balance, err := self.getETHBalance(addr)
		if err != nil {
			level.Error(self.logger).Log("msg", "getting initial ETH balance", "addr", addr.String(), "err", err)
		}
		level.Info(self.logger).Log("msg", "initial ETH balance", "addr", addr.String(), "balance", balance)
		self.balances.With(prometheus.Labels{"addr": addr.String(), "token": "ETH"}).Set(balance)
	}

	go self.monitorCost()
	go self.monitorReward()
	go self.monitorCostFailed()

	<-self.ctx.Done()
	return nil
}

func (self *ProfitTracker) Stop() {
	self.stop()
}

func (self *ProfitTracker) LastSubmit() time.Time {
	return self.lastSubmit
}

func (self *ProfitTracker) monitorReward() {
	var err error
	ticker := time.NewTicker(DefaultRetry)
	defer ticker.Stop()
	var sub event.Subscription
	events := make(chan *tellor.TellorTransferred)

	logger := log.With(self.logger, "event", "Transfer")

	for {
		select {
		case <-self.ctx.Done():
			return
		default:
		}
		sub, err = self.transferSub(events)
		if err != nil {
			level.Error(logger).Log("msg", "initial subscribing to events", "err", err)
			<-ticker.C
			continue
		}
		break
	}

	for {
		select {
		case <-self.ctx.Done():
			return
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
					return
				default:
				}
				sub, err = self.transferSub(events)
				if err != nil {
					level.Error(logger).Log("msg", "re-subscribing to events failed")
					<-ticker.C
					continue
				}
				break
			}
			level.Info(logger).Log("msg", "re-subscribed to events")
		case event := <-events:
			logger := log.With(logger, "addr", event.To.String()[:6], "tx", event.Raw.TxHash)

			if event.Raw.Removed {
				val, err := self.cacheTXsProfit.Get(txIDTransfer(event))
				if err != nil {
					level.Error(logger).Log("msg", "getting cache amount for removed event", "err", err)
					continue
				}
				level.Debug(logger).Log("msg", "removing cost from dropped event", "amount", val.(float64))
				self.submitProfit.With(prometheus.Labels{"addr": event.To.String()}).Sub(val.(float64))
				continue
			}

			self.setProfitWhenConfirmed(logger, event)
		}
	}
}

func (self *ProfitTracker) monitorCost() {
	var err error
	ticker := time.NewTicker(DefaultRetry)
	defer ticker.Stop()

	logger := log.With(self.logger, "event", "NonceSubmitted")

	var sub event.Subscription
	events := make(chan *tellor.TellorNonceSubmitted)

	for {
		select {
		case <-self.ctx.Done():
			return
		default:
		}
		sub, err = self.nonceSubmittedSub(events)
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
			return
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
					return
				default:
				}
				sub, err = self.nonceSubmittedSub(events)
				if err != nil {
					level.Error(logger).Log("msg", "re-subscribing to events failed", "err", err)
					<-ticker.C
					continue
				}
				break
			}
			level.Info(logger).Log("msg", "re-subscribed to events")
		case event := <-events:
			logger := log.With(logger, "addr", event.Miner.String()[:6], "tx", event.Raw.TxHash)

			if event.Raw.Removed {
				val, err := self.cacheTXsCost.Get(txIDNonceSubmit(event))
				if err != nil {
					level.Error(logger).Log("msg", "getting cache amount for removed event", "err", err)
					continue
				}
				level.Debug(logger).Log("msg", "removed event", "amount", val.(float64))
				self.submitCost.With(prometheus.Labels{"addr": event.Miner.String()}).Sub(val.(float64))
				continue
			}

			// When the next slot is 4 also record the gas estimation.
			if event.Slot.Int64() == 3 {
				var gasEstimate uint64
				var ok bool
				for _, acc := range self.accounts {
					gasEstimate, err = self.gasEstimator.EstimateGas(self.ctx, acc, event.Slot.Uint64())
					if err == nil {
						ok = true
						break
					}
				}
				if !ok {
					level.Error(self.logger).Log("msg", "gas estimation didn't work with any of the registered addresses", "lastErr", err)

				} else {
					self.lastGasEstimate = gasEstimate
					self.submitGasUsageEstimated.With(
						prometheus.Labels{
							"slot": strconv.Itoa(int(event.Slot.Uint64())),
						},
					).Set(float64(gasEstimate))
				}
			}

			self.setCostWhenConfirmed(logger, event)
		}
	}
}

func (self *ProfitTracker) monitorCostFailed() {
	var err error
	ticker := time.NewTicker(DefaultRetry)
	defer ticker.Stop()

	logger := log.With(self.logger, "event", "NewHead")

	var sub event.Subscription
	events := make(chan *types.Header)

	for {
		select {
		case <-self.ctx.Done():
			return
		default:
		}

		sub, err = self.headSub(events)
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
			return
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
					return
				default:
				}
				sub, err = self.headSub(events)
				if err != nil {
					level.Error(logger).Log("msg", "re-subscribing to events failed", "err", err)
					<-ticker.C
					continue
				}
				break
			}
			level.Info(logger).Log("msg", "re-subscribed to events")
		case event := <-events:
			if event.Bloom.Test(self.abi.Events["NonceSubmitted"].ID.Bytes()) {
				logger := log.With(logger, "block", event.Number)

				block, err := self.client.BlockByNumber(self.ctx, event.Number)
				if err != nil {
					level.Error(logger).Log("msg", "get block by hash", "err", err)
					continue
				}

				level.Debug(logger).Log("msg", "new block")

				for _, tx := range block.Transactions() {
					logger := log.With(logger, "tx", tx.Hash())
					level.Debug(logger).Log("msg", "processing TX")

					addr, err := types.Sender(types.LatestSignerForChainID(self.netID), tx)
					if err != nil {
						level.Error(logger).Log("msg", "get tx sender", "err", err)
						continue
					}
					if _, ok := self.addrsMap[addr]; !ok {
						level.Debug(logger).Log("msg", "skipping TX for unregistered address", "addr", addr)
						continue
					}
					receipt, err := self.client.TransactionReceipt(self.ctx, tx.Hash())
					if err != nil {
						level.Error(logger).Log("msg", "receipt retrieval", "err", err)
						continue
					} else if receipt != nil && receipt.Status != types.ReceiptStatusSuccessful { // Track only the failed TXs. All other TXs are tracked from the emitted logs.
						// When it is a reorg event, remove the cost for the dropped blocks.
						if self.lastFailedBlock >= event.Number.Int64() {
							level.Debug(logger).Log("msg", "reorg head")
							for cachedBlock, _cost := range self.cacheTXsCostFailed.GetALL(false) {
								if cachedBlock.(int64) >= event.Number.Int64() {
									cost := _cost.(float64)
									level.Debug(logger).Log("msg", "removing cost from dropped block", "amount", cost)
									self.submitCost.With(prometheus.Labels{"addr": addr.String()}).Sub(cost)
								}
							}
						}
						self.lastFailedBlock = event.Number.Int64()
						cost, _ := big.NewFloat(0).Mul(big.NewFloat(float64(tx.GasPrice().Int64())), big.NewFloat(float64(receipt.GasUsed))).Float64()
						cost = cost / 1e18
						level.Debug(logger).Log("msg", "adding cost", "amount", cost)
						self.submitCost.With(prometheus.Labels{"addr": addr.String()}).Add(cost)

						if err := self.cacheTXsCostFailed.Set(event.Number.Int64(), cost); err != nil {
							level.Error(logger).Log("msg", "adding cost to the cache", "err", err)
						}

						self.submitFailCount.With(
							prometheus.Labels{
								"addr": addr.String(),
							},
						).Inc()

						balance, err := self.getETHBalance(addr)
						if err != nil {
							level.Error(logger).Log("msg", "getting ETH balance", "err", err)
							continue
						}
						level.Debug(logger).Log("msg", "new ETH balance", "balance", balance)
						self.balances.With(prometheus.Labels{"addr": addr.String(), "token": "ETH"}).Set(balance)
					}
				}
			}
		}
	}
}

func (self *ProfitTracker) setCostWhenConfirmed(logger log.Logger, event *tellor.TellorNonceSubmitted) {
	ticker := time.NewTicker(DefaultRetry)
	defer ticker.Stop()

	// Give up checking after this time.
	ctx, cncl := context.WithTimeout(self.ctx, time.Hour)
	defer cncl()

	for {
		select {
		case <-ctx.Done():
			level.Debug(logger).Log("msg", "transaction confirmation check canceled")
			return
		case <-ticker.C:
		}
		receipt, err := self.client.TransactionReceipt(ctx, event.Raw.TxHash)
		if err != nil {
			level.Error(logger).Log("msg", "receipt retrieval", "err", err)
		} else if receipt != nil && receipt.Status == types.ReceiptStatusSuccessful { // Failed transactions cost is monitored in a different process.
			tx, _, err := self.client.TransactionByHash(self.ctx, event.Raw.TxHash)
			if err != nil {
				level.Error(logger).Log("msg", "get transaction by hash", "err", err)
				return
			}
			cost, _ := big.NewFloat(0).Mul(big.NewFloat(float64(tx.GasPrice().Int64())), big.NewFloat(float64(receipt.GasUsed))).Float64()
			cost = cost / 1e18
			level.Debug(logger).Log("msg", "adding cost", "amount", cost)
			self.submitCost.With(prometheus.Labels{"addr": event.Miner.String()}).Add(cost)

			if err := self.cacheTXsCost.Set(txIDNonceSubmit(event), cost); err != nil {
				level.Error(logger).Log("msg", "adding cost to the cache", "err", err)
			}

			balance, err := self.getETHBalance(event.Miner)
			if err != nil {
				level.Error(logger).Log("msg", "getting ETH balance", "err", err)
				return
			}
			level.Debug(logger).Log("msg", "new ETH balance", "balance", balance)
			self.balances.With(prometheus.Labels{"addr": event.Miner.String(), "token": "ETH"}).Set(balance)

			self.submitCount.With(
				prometheus.Labels{
					"addr": event.Miner.String(),
				},
			).Inc()

			self.lastSubmit = time.Now()

			self.submitGasUsageActual.With(
				prometheus.Labels{
					"slot": strconv.Itoa(int(event.Slot.Uint64())),
				},
			).Set(float64(receipt.GasUsed))

			for i, id := range event.RequestId {
				self.submitValue.With(
					prometheus.Labels{
						"addr": event.Miner.String(),
						"id":   id.String(),
					},
				).Set(float64(event.Value[i].Int64()))
			}

			return
		}

		level.Debug(logger).Log("msg", "transaction not yet mined")
	}
}

func (self *ProfitTracker) setProfitWhenConfirmed(logger log.Logger, event *tellor.TellorTransferred) {
	ticker := time.NewTicker(DefaultRetry)
	defer ticker.Stop()
	for {
		select {
		case <-self.ctx.Done():
			level.Debug(logger).Log("msg", "transaction confirmation check canceled")
			return
		case <-ticker.C:
		}
		receipt, err := self.client.TransactionReceipt(self.ctx, event.Raw.TxHash)
		if err != nil {
			level.Error(logger).Log("msg", "receipt retrieval", "err", err)
			return
		} else if receipt != nil {
			if receipt.Status != types.ReceiptStatusSuccessful {
				level.Error(logger).Log("msg", "event status not success so no profit added", "status", receipt.Status)
				return
			}

			trb, _ := big.NewFloat(float64(event.Value.Int64())).Float64()
			trb = trb / 1e18
			level.Debug(logger).Log("msg", "adding profit", "amount", trb)
			self.submitProfit.With(prometheus.Labels{"addr": event.To.String()}).Add(trb)

			if err := self.cacheTXsProfit.Set(txIDTransfer(event), trb); err != nil {
				level.Error(logger).Log("msg", "adding amount to the cache", "err", err)
			}

			balance, err := self.getTRBBalance(event.To)
			if err != nil {
				level.Error(logger).Log("msg", "getting TRB balance", "err", err)
				return
			}
			level.Debug(logger).Log("msg", "new TRB balance", "balance", balance)
			self.balances.With(prometheus.Labels{"addr": event.To.String(), "token": "TRB"}).Set(balance)
			return
		}
		level.Debug(logger).Log("msg", "transaction not yet mined")
	}
}

func (self *ProfitTracker) nonceSubmittedSub(output chan *tellor.TellorNonceSubmitted) (event.Subscription, error) {
	tellorFilterer, err := tellor.NewTellorFilterer(self.contractInstance.Address, self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting instance")
	}

	sub, err := tellorFilterer.WatchNonceSubmitted(&bind.WatchOpts{Context: self.ctx}, output, self.addrs, nil)
	if err != nil {
		return nil, errors.Wrap(err, "getting channel")
	}
	return sub, nil
}

func (self *ProfitTracker) transferSub(output chan *tellor.TellorTransferred) (event.Subscription, error) {
	tellorFilterer, err := tellor.NewTellorFilterer(self.contractInstance.Address, self.client)
	if err != nil {
		return nil, errors.Wrap(err, "getting instance")
	}
	sub, err := tellorFilterer.WatchTransferred(
		&bind.WatchOpts{Context: self.ctx},
		output,
		[]common.Address{
			common.HexToAddress("0x0000000000000000000000000000000000000000"),
		},
		self.addrs,
	)
	if err != nil {
		return nil, errors.Wrap(err, "getting channel")
	}
	return sub, nil
}

func (self *ProfitTracker) headSub(output chan *types.Header) (event.Subscription, error) {
	sub, err := self.client.SubscribeNewHead(self.ctx, output)
	if err != nil {
		return nil, errors.Wrap(err, "getting channel")
	}
	return sub, nil
}

func (self *ProfitTracker) getTRBBalance(addr common.Address) (float64, error) {
	balance, err := self.contractInstance.BalanceOf(nil, addr)
	if err != nil {
		return 0, errors.Wrap(err, "retrieving trb balance")
	}
	_balanceH, _ := big.NewFloat(1).SetString(balance.String())
	decimals, _ := big.NewFloat(1).SetString("1000000000000000000")
	if decimals != nil {
		_balanceH = _balanceH.Quo(_balanceH, decimals)
	}
	balanceH, _ := _balanceH.Float64()
	return balanceH, nil
}

func (self *ProfitTracker) getETHBalance(addr common.Address) (float64, error) {
	balance, err := self.client.BalanceAt(self.ctx, addr, nil)
	if err != nil {
		return 0, errors.Wrap(err, "retrieving balance")
	}
	_balanceH, _ := big.NewFloat(1).SetString(balance.String())
	decimals, _ := big.NewFloat(1).SetString("1000000000000000000")
	if decimals != nil {
		_balanceH = _balanceH.Quo(_balanceH, decimals)
	}
	balanceH, _ := _balanceH.Float64()
	return balanceH, nil
}

func txIDTransfer(event *tellor.TellorTransferred) string {
	return event.Raw.TxHash.String() + event.To.String()
}

func txIDNonceSubmit(event *tellor.TellorNonceSubmitted) string {
	return event.Raw.TxHash.String() + event.Miner.String()
}
