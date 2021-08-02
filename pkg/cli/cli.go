// Copyright (c) The Cryptorium Authors.
// Licensed under the MIT License.

package cli

import (
	"context"
	"io/ioutil"
	"strings"
	"time"

	"github.com/cryptoriums/telliot/pkg/config"
	"github.com/cryptoriums/telliot/pkg/contracts"
	"github.com/cryptoriums/telliot/pkg/ethereum"
	"github.com/cryptoriums/telliot/pkg/logging"
	"github.com/cryptoriums/telliot/pkg/math"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"
	"github.com/go-kit/kit/log/level"
	"github.com/google/go-github/v35/github"
	"github.com/pkg/errors"
)

const VersionMessage = `
    The unofficial Tellor cli tool %s (%s)
    -----------------------------------------
	Website: https://tellor.io
	Github:  https://github.com/cryptoriums/telliot
`

var CLIDefault = CLI{

	Dataserver: DataserverCmd{},
	Report:     ReportCmd{},

	Transfer: TransferCmd{},
	Approve:  ApproveCmd{},
	Accounts: AccountsCmd{},
	Data:     DataCmd{},
	Balance:  BalanceCmd{},
	Stake: stakeCmd{
		Deposit:  DepositCmd{},
		Request:  RequestCmd{},
		Withdraw: WithdrawCmd{},
		Status:   StatusCmd{},
	},
	Dispute: disputeCmd{
		New:       NewDisputeCmd{},
		Vote:      VoteCmd{},
		List:      ListCmd{},
		Tally:     TallyCmd{},
		UnlockFee: UnlockFeeCmd{},
	},
	Encrypt: EncryptCmd{},
	Decrypt: DecryptCmd{},
	Version: VersionCmd{},
}

type CLI struct {
	Cfg

	Dataserver DataserverCmd `cmd:"" help:"launch only a dataserver instance"`
	Report     ReportCmd     `cmd:"" help:"Submit data to the oracle contracts"`
	Submit     SubmitCmd     `cmd:"" help:"Make a single manual submit to the oracle contracts"`

	Transfer TransferCmd `cmd:"" help:"Transfer tokens"`
	Approve  ApproveCmd  `cmd:"" help:"Approve tokens"`
	Accounts AccountsCmd `cmd:"" help:"Show accounts"`
	Data     DataCmd     `cmd:"" help:"Retrieve data from the contract"`
	Balance  BalanceCmd  `cmd:"" help:"Check the balance of an address"`
	Stake    stakeCmd    `cmd:"" help:"Perform one of the stake operations"`
	Dispute  disputeCmd  `cmd:"" help:"Perform commands related to disputes"`
	Encrypt  EncryptCmd  `cmd:"" help:"Encrypts a file to be securely stored on disk and later used only with a password prompt"`
	Decrypt  DecryptCmd  `cmd:"" help:"Decrypts an ecrypted file and write the decrytped version to disk"`
	Events   eventsCmd   `cmd:"" help:"Subscribe to watch logs from the network."`
	Version  VersionCmd  `cmd:"" help:"Show the CLI version information"`
}

func (self *CLI) AfterApply() error {
	logger := logging.NewLogger()

	cfg, err := config.Parse(logger, string(self.Config))
	if err != nil {
		return errors.Wrap(err, "creating config")
	}

	err = config.LoadEnvFile(logger, cfg.EnvFile)
	if err != nil {
		return errors.Wrapf(err, "loading the enf file:%v", cfg.EnvFile)
	}

	self.Dataserver.SetConfig(cfg)
	self.Report.SetConfig(cfg)
	self.Submit.SetConfig(cfg)
	self.Dispute.List.SetConfig(cfg)

	return nil
}

type disputeCmd struct {
	New       NewDisputeCmd `cmd:"" help:"start a new dispute"`
	Vote      VoteCmd       `cmd:"" help:"vote on a open dispute"`
	List      ListCmd       `cmd:"" help:"list open disputes"`
	Tally     TallyCmd      `cmd:"" help:"tally votes for a dispute ID, need to run unlock fee after that"`
	TallyList TallyListCmd  `cmd:"" help:"list tally for disputes"`
	UnlockFee UnlockFeeCmd  `cmd:"" help:"after tallying the votes this command transfers the fee to the reporter or reported based on the voting"`
}

type stakeCmd struct {
	Deposit  DepositCmd  `cmd:"" help:"deposit a stake"`
	Request  RequestCmd  `cmd:"" help:"request to withdraw stake"`
	Withdraw WithdrawCmd `cmd:"" help:"withdraw stake"`
	Status   StatusCmd   `cmd:"" help:"show stake status"`
}

type VersionCmd struct {
}

func (cmd VersionCmd) Run() error {
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
	Contract string `optional:"" help:"provide valid hex address"`
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
	GasBaseFee float64 `optional:"" help:"gas base fee to use when running the command"`
	GasTip     float64 `optional:"" help:"gas tip fee to use when running the command"`
}

func (self *Gas) Validate() error {
	if self.GasBaseFee > 300 {
		return errors.Errorf("gase base fee too high:%v", self.GasBaseFee)
	}
	if self.GasTip > 300 {
		return errors.Errorf("gase tip too high:%v", self.GasTip)
	}
	return nil
}

type GasAccount struct {
	Gas
	AccountArg
}

type AccountsCmd struct {
}

func (self *AccountsCmd) Run() error {
	logger := logging.NewLogger()
	ctx := context.Background()

	accounts, err := ethereum.GetAccounts()
	if err != nil {
		return errors.Wrap(err, "getting accounts")
	}

	client, netID, err := ethereum.NewClient(logger, ctx)
	if err != nil {
		return errors.Wrap(err, "creating ethereum client")
	}

	contract, err := contracts.NewITellor(logger, ctx, client, netID, contracts.DefaultParams)
	if err != nil {
		return errors.Wrap(err, "create tellor contract instance")
	}

	for i, account := range accounts {
		ethBalance, err := client.BalanceAt(ctx, account.Address, nil)
		if err != nil {
			return errors.Wrap(err, "get eth balance")
		}
		trbBalance, err := contract.BalanceOf(&bind.CallOpts{Context: ctx}, account.Address)
		if err != nil {
			return errors.Wrapf(err, "getting trb balance")
		}
		level.Info(logger).Log("msg", "account",
			"index", i,
			"address", account.Address.Hex(),
			"trb", math.BigIntToFloatDiv(trbBalance, params.Ether),
			"eth", math.BigIntToFloatDiv(ethBalance, params.Ether),
		)
	}

	return nil
}

func CheckNewVersion(repo string, current string) (string, error) {
	if current == "" { // Can be empty when using `go run`.
		return "", nil
	}
	ctx, cncl := context.WithTimeout(context.Background(), 3*time.Second)
	defer cncl()
	client := github.NewClient(nil)
	release, resp, err := client.Repositories.GetLatestRelease(ctx, repo, "telliot")
	if err != nil {
		return "", errors.Wrap(err, "checking for a new release")
	}
	defer resp.Body.Close()

	if resp.StatusCode/100 != 2 {
		rbody, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", errors.Errorf("bad response status %v", resp.Status)
		}
		return "", errors.Errorf("bad response status %v from %q", resp.Status, string(rbody))
	}

	parts := strings.Split(current, "-")

	if len(parts) == 0 {
		return "", errors.New("failed to process current tag name")
	}

	if parts[0] != release.GetTagName() {
		return release.GetHTMLURL(), nil
	}

	return "", nil
}
