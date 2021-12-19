// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"fmt"
	"math/big"
	"os"
	"sort"
	"text/tabwriter"
	"time"

	"github.com/cryptoriums/packages/ethereum"
	"github.com/cryptoriums/packages/math"
	"github.com/cryptoriums/telliot/pkg/config"
	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/pkg/errors"
	"github.com/willabides/kongplete"
)

const VersionMessage = `
    The unofficial Tellor cli tool %s (%s)
    -----------------------------------------
	Website: https://tellor.io
	Github:  https://github.com/cryptoriums/telliot
`

var CLIDefault = CLI{
	Dataserver: DataserverCmd{},

	Transfer: TransferCmd{},
	Approve:  ApproveCmd{},
	Accounts: AccountsCmd{},
	Data:     DataCmd{},
	Balance:  BalanceCmd{},
	Stake: StakeCmd{
		Deposit:  DepositCmd{},
		Request:  RequestCmd{},
		Withdraw: WithdrawCmd{},
		Status:   StatusCmd{},
	},
	Dispute: DisputeCmd{
		New:     NewDisputeCmd{},
		Vote:    VoteCmd{},
		List:    ListCmd{},
		Tally:   TallyCmd{},
		Execute: ExecuteCmd{},
	},
	Encrypt: EncryptCmd{},
	Decrypt: DecryptCmd{},
	Version: VersionCmd{},
}

type CLI struct {
	Cfg
	ConfigStrictParsing bool `help:"whether to return an error when the config contains unknown fields"`
	ContractFlag

	Dataserver DataserverCmd `cmd:"" help:"launch only a dataserver instance"`
	Submit     SubmitCmd     `cmd:"" help:"Make a single manual submit to the oracle contracts"`

	Transfer TransferCmd `cmd:"" help:"Transfer tokens"`
	Approve  ApproveCmd  `cmd:"" help:"Approve tokens"`
	Accounts AccountsCmd `cmd:"" help:"Show accounts"`
	Data     DataCmd     `cmd:"" help:"Retrieve data from the contract"`
	Balance  BalanceCmd  `cmd:"" help:"Check the balance of an address"`
	Stake    StakeCmd    `cmd:"" help:"Perform one of the stake operations"`
	Dispute  DisputeCmd  `cmd:"" help:"Perform commands related to disputes"`
	Encrypt  EncryptCmd  `cmd:"" help:"Encrypts a file to be securely stored on disk and later used only with a password prompt"`
	Decrypt  DecryptCmd  `cmd:"" help:"Decrypts an ecrypted file and write the decrytped version to disk"`
	Events   EventsCmd   `cmd:"" help:"Subscribe to watch logs from the network."`
	Version  VersionCmd  `cmd:"" help:"Show the CLI version information"`

	InstallCompletions kongplete.InstallCompletions `cmd:"" help:"install shell completions"`
}

type DisputeCmd struct {
	New       NewDisputeCmd `cmd:"" help:"start a new dispute"`
	Vote      VoteCmd       `cmd:"" help:"vote on a open dispute"`
	List      ListCmd       `cmd:"" help:"list open disputes"`
	Tally     TallyCmd      `cmd:"" help:"tally votes for a dispute ID, need to run unlock fee after that"`
	TallyList TallyListCmd  `cmd:"" help:"list tally for disputes"`
	Execute   ExecuteCmd    `cmd:"" help:"after tallying the votes this command transfers the fee to the reporter or reported based on the voting"`
}

type StakeCmd struct {
	Deposit  DepositCmd  `cmd:"" help:"deposit a stake"`
	Request  RequestCmd  `cmd:"" help:"request to withdraw stake"`
	Withdraw WithdrawCmd `cmd:"" help:"withdraw stake"`
	Status   StatusCmd   `cmd:"" help:"show stake status"`
}

type VersionCmd struct {
}

func (cmd VersionCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	// The main entry point prints the version message so
	// here just return nil and the message will be printed.
	return nil
}

type NoChks struct {
	NoChecks bool `help:"whether to run some additional checks like balance, did vote etc. Useful to disable when running against a test contract"`
}

type Cfg struct {
	Config string `type:"existingfile" help:"path to config file"`
}

type AccountArg struct {
	Account string `arg:"" required:""`
}

type AccountArgOptional struct {
	Account string `arg:"" optional:""`
}

func (self *AccountArgOptional) Validate() error {
	return ValidateHex(self.Account)
}

func (self *AccountArg) Validate() error {
	return ValidateHex(self.Account)
}

func ValidateHex(account string) error {
	if account == "" {
		return nil
	}
	if !common.IsHexAddress(account) {
		return errors.Errorf("the address is not a hex string:%v", account)
	}
	return nil
}

type ContractFlag struct {
	Contract string `optional:"" help:"hex address of the contract to interract with"`
}

func (self *ContractFlag) Validate() error {
	if self.Contract == "" {
		return nil
	}
	if !common.IsHexAddress(self.Contract) {
		return errors.Errorf("the address is not a hex string:%v", self.Contract)
	}
	return nil
}

type Gas struct {
	GasPrice float64 `optional:"" help:"gas max fee to use when running the command"`
}

func (self *Gas) Validate() error {
	if self.GasPrice > 300 {
		return errors.Errorf("gase base fee too high:%v", self.GasPrice)
	}
	return nil
}

type GasAccount struct {
	Gas
	AccountArg
}

type AccountsCmd struct {
}

func (self *AccountsCmd) Run(cli *CLI, ctx context.Context, logger log.Logger) error {
	cfg, client, master, oracle, _, err := ConfigClientContract(ctx, logger, cli.Config, cli.ConfigStrictParsing, cli.Contract, contracts.DefaultParams)
	if err != nil {
		return err
	}
	defer client.Close()

	accounts, err := ethereum.GetAccounts(logger, cfg.EnvVars)
	if err != nil {
		return errors.Wrap(err, "getting accounts")
	}

	PrintAccounts(ctx, logger, accounts, client, master, oracle)

	return nil
}

type accountWithDetails struct {
	common.Address
	StakedStatus       string
	BalanceTRB         float64
	BalanceETH         float64
	TimeTillNextSubmit time.Duration
}

func PrintAccounts(
	ctx context.Context,
	logger log.Logger,
	accounts []*ethereum.Account,
	client ethereum.EthClient,
	master contracts.TellorMasterCaller,
	oracle contracts.TellorOracleCaller,
) {
	level.Info(logger).Log("msg", "getting all account details so take a chill pill and wait....")

	_minSubmitPeriod, err := oracle.ReportingLock(&bind.CallOpts{Context: ctx})
	if err != nil {
		level.Error(logger).Log("msg", "couln't get reporting lock so will use a default", "err", err)
		_minSubmitPeriod = big.NewInt(int64((12 * time.Hour).Seconds()))
	}
	minSubmitPeriod := time.Duration(_minSubmitPeriod.Int64()) * time.Second

	var accountsAll []accountWithDetails

	for _, account := range accounts {
		logger := log.With(logger, "acc", account.Address.Hex())

		ethBalance, err := client.BalanceAt(ctx, account.Address, nil)
		if err != nil {
			level.Error(logger).Log("msg", "get eth balance")
		}
		trbBalance, err := master.BalanceOf(&bind.CallOpts{Context: ctx}, account.Address)
		if err != nil {
			level.Error(logger).Log("msg", "get TRB balance")
		}

		timeSincelastSubmit, _, err := contracts.LastSubmit(ctx, oracle, account.Address)
		if err != nil {
			level.Error(logger).Log("msg", "checking last submit time", "err", err)
		}

		status, _, err := master.GetStakerInfo(&bind.CallOpts{Context: ctx}, account.Address)
		if err != nil {
			level.Error(logger).Log("msg", "getting stake status", "err", err)
		}

		accountsAll = append(accountsAll, accountWithDetails{
			Address:            account.Address,
			BalanceTRB:         math.BigIntToFloatDiv(trbBalance, params.Ether),
			BalanceETH:         math.BigIntToFloatDiv(ethBalance, params.Ether),
			TimeTillNextSubmit: -time.Since(time.Now().Add(minSubmitPeriod - timeSincelastSubmit)),
			StakedStatus:       contracts.ReporterStatusName(status.Int64()),
		})
	}

	sort.Slice(accountsAll, func(i, j int) bool {
		return accountsAll[i].TimeTillNextSubmit < accountsAll[j].TimeTillNextSubmit
	})

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)

	for _, account := range accountsAll {
		//lint:ignore faillint looks cleaner with print instead of logs
		fmt.Fprintf(w, "%v, \ttrb:%v \teth:%v \tnextSubmit:%v \tstakeStatus:%v \n",
			account.Address.Hex()[:10],
			account.BalanceTRB,
			account.BalanceETH,
			account.TimeTillNextSubmit,
			account.StakedStatus,
		)
	}
	w.Flush()
}

func ConfigClientContract(
	ctx context.Context,
	logger log.Logger,
	configPath string,
	configStrictParsing bool,
	contractAddr string,
	params contracts.Params,
) (
	*config.Config,
	ethereum.EthClient,
	contracts.TellorMasterCaller,
	contracts.TellorOracleCaller,
	contracts.TellorGovernCaller,
	error,
) {
	cfg, err := config.LoadConfig(ctx, logger, configPath, configStrictParsing)
	if err != nil {
		return nil, nil, nil, nil, nil, err
	}

	client, err := ethereum.NewClient(ctx, logger, cfg.EnvVars)
	if err != nil {
		return nil, nil, nil, nil, nil, errors.Wrap(err, "creating ethereum client")
	}

	if params == (contracts.Params{}) {
		params = contracts.DefaultParams
	}

	master, oracle, govern, err := contracts.NewContracts(ctx, logger, client, common.HexToAddress(contractAddr), params)
	if err != nil {
		return nil, nil, nil, nil, nil, errors.Wrap(err, "create tellor contract instance")
	}

	return cfg, client, master, oracle, govern, nil

}
