---
description: Telliot tweaks and settings to keep your rig running smoothly.
---

# Configuration reference

## CLI reference

Telliot commands and config file options are as the following:

#### Required Flags <a id="docs-internal-guid-d1a57725-7fff-a753-9236-759dd3f42eed"></a>

* `--config` \(path to your config file.\)

#### Telliot Commands

* `accounts`

```
Usage: telliot accounts

Show accounts

Flags:
  -h, --help             Show context-sensitive help.
      --config=STRING    path to config file

```

* `approve`

```
Usage: telliot approve --from=STRING --to=STRING <amount>

Approve tokens

Arguments:
  <amount>

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file

      --gas-base-fee=FLOAT-64    gas base fee to use when running the command
      --gas-tip=FLOAT-64         gas tip fee to use when running the command
      --from=STRING
      --to=STRING

```

* `balance`

```
Usage: telliot balance <account>

Check the balance of an address

Arguments:
  <account>

Flags:
  -h, --help             Show context-sensitive help.
      --config=STRING    path to config file

```

* `data`

```
Usage: telliot data

Retrieve data from the contract

Flags:
  -h, --help             Show context-sensitive help.
      --config=STRING    path to config file

      --from=UINT-64     the unix timestamp to use as a starting point for the
                         data retrieval
      --look-back=2h     how far to lookback

```

* `dataserver`

```
Usage: telliot dataserver

launch only a dataserver instance

Flags:
  -h, --help             Show context-sensitive help.
      --config=STRING    path to config file

```

* `decrypt`

```
Usage: telliot decrypt <file>

Decrypts an ecrypted file and write the decrytped version to disk

Arguments:
  <file>    the file to encrypt

Flags:
  -h, --help             Show context-sensitive help.
      --config=STRING    path to config file

```

* `dispute`

```
Usage: telliot dispute <command>

Perform commands related to disputes

Flags:
  -h, --help             Show context-sensitive help.
      --config=STRING    path to config file

Commands:
  dispute new --request-id=INT-64 --timestamp=INT-64 --slot=INT-64 <account>
    start a new dispute

  dispute vote --dispute-id=INT-64 --support <account>
    vote on a open dispute

  dispute list
    list open disputes

  dispute tally --dispute-id=INT-64
    tally votes for a dispute ID, need to run unlock fee after that

  dispute tally-list
    list tally for disputes

  dispute unlock-fee --dispute-id=INT-64
    after tallying the votes this command transfers the fee to the reporter or
    reported based on the voting

```

* `dispute list`

```
Usage: telliot dispute list

list open disputes

Flags:
  -h, --help              Show context-sensitive help.
      --config=STRING     path to config file

      --show-closed       also show executed disputes
      --look-back=120h    how far to lookback, the default only few days since
                          disputes can be voted only for 2 days.

```

* `dispute new`

```
Usage: telliot dispute new --request-id=INT-64 --timestamp=INT-64 --slot=INT-64 <account>

start a new dispute

Arguments:
  <account>

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file

      --gas-base-fee=FLOAT-64    gas base fee to use when running the command
      --gas-tip=FLOAT-64         gas tip fee to use when running the command
      --no-checks                whether to run some additional checks like
                                 balance, did vote etc. Useful to disable when
                                 running against a test contract
      --request-id=INT-64        the request id to dispute
      --timestamp=INT-64         the submitted timestamp to dispute
      --slot=INT-64              the reporter index to dispute

```

* `dispute tally`

```
Usage: telliot dispute tally --dispute-id=INT-64

tally votes for a dispute ID, need to run unlock fee after that

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file

      --gas-base-fee=FLOAT-64    gas base fee to use when running the command
      --gas-tip=FLOAT-64         gas tip fee to use when running the command
      --dispute-id=INT-64        the dispute id

```

* `dispute tally-list`

```
Usage: telliot dispute tally-list

list tally for disputes

Flags:
  -h, --help              Show context-sensitive help.
      --config=STRING     path to config file

      --look-back=120h    how far to lookback, the default only few days since
                          disputes can be voted only for 2 days.

```

* `dispute unlock-fee`

```
Usage: telliot dispute unlock-fee --dispute-id=INT-64

after tallying the votes this command transfers the fee to the reporter or
reported based on the voting

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file

      --no-checks                whether to run some additional checks like
                                 balance, did vote etc. Useful to disable when
                                 running against a test contract
      --gas-base-fee=FLOAT-64    gas base fee to use when running the command
      --gas-tip=FLOAT-64         gas tip fee to use when running the command
      --dispute-id=INT-64        the dispute id

```

* `dispute vote`

```
Usage: telliot dispute vote --dispute-id=INT-64 --support <account>

vote on a open dispute

Arguments:
  <account>

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file

      --gas-base-fee=FLOAT-64    gas base fee to use when running the command
      --gas-tip=FLOAT-64         gas tip fee to use when running the command
      --dispute-id=INT-64        the dispute id
      --no-checks                whether to run some additional checks like
                                 balance, did vote etc. Useful to disable when
                                 running against a test contract
      --support                  true or false

```

* `encrypt`

```
Usage: telliot encrypt <file>

Encrypts a file to be securely stored on disk and later used only with a
password prompt

Arguments:
  <file>    the file to encrypt

Flags:
  -h, --help             Show context-sensitive help.
      --config=STRING    path to config file

```

* `events`

```
Usage: telliot events --event-name=STRING

Subscribe to watch logs from the network.

Flags:
  -h, --help                  Show context-sensitive help.
      --config=STRING         path to config file

      --look-back=DURATION    how far to look for the initiali qyery
      --contract=STRING       provide valid hex address
      --event-name=STRING     the name of the log to watch
      --reorg-wait=3s         how long to wait for removed logs from reorg
                              events

```

* `report`

```
Usage: telliot report

Submit data to the oracle contracts

Flags:
  -h, --help             Show context-sensitive help.
      --config=STRING    path to config file

```

* `stake`

```
Usage: telliot stake <command>

Perform one of the stake operations

Flags:
  -h, --help             Show context-sensitive help.
      --config=STRING    path to config file

Commands:
  stake deposit <account>
    deposit a stake

  stake request <account>
    request to withdraw stake

  stake withdraw <account>
    withdraw stake

  stake status <account>
    show stake status

```

* `stake deposit`

```
Usage: telliot stake deposit <account>

deposit a stake

Arguments:
  <account>

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file

      --gas-base-fee=FLOAT-64    gas base fee to use when running the command
      --gas-tip=FLOAT-64         gas tip fee to use when running the command
      --contract=STRING          provide valid hex address

```

* `stake request`

```
Usage: telliot stake request <account>

request to withdraw stake

Arguments:
  <account>

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file

      --gas-base-fee=FLOAT-64    gas base fee to use when running the command
      --gas-tip=FLOAT-64         gas tip fee to use when running the command

```

* `stake status`

```
Usage: telliot stake status <account>

show stake status

Arguments:
  <account>

Flags:
  -h, --help               Show context-sensitive help.
      --config=STRING      path to config file

      --contract=STRING    provide valid hex address

```

* `stake withdraw`

```
Usage: telliot stake withdraw <account>

withdraw stake

Arguments:
  <account>

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file

      --gas-base-fee=FLOAT-64    gas base fee to use when running the command
      --gas-tip=FLOAT-64         gas tip fee to use when running the command

```

* `submit`

```
Usage: telliot submit [<account>]

Make a single manual submit to the oracle contracts

Arguments:
  [<account>]

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file

      --gas-base-fee=FLOAT-64    gas base fee to use when running the command
      --gas-tip=FLOAT-64         gas tip fee to use when running the command
      --contract=STRING          provide valid hex address
      --skip-confirm             submit without confirming, useful for testing

```

* `transfer`

```
Usage: telliot transfer --from=STRING --to=STRING <amount>

Transfer tokens

Arguments:
  <amount>

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file

      --gas-base-fee=FLOAT-64    gas base fee to use when running the command
      --gas-tip=FLOAT-64         gas tip fee to use when running the command
      --from=STRING
      --to=STRING

```

* `version`

```
Usage: telliot version

Show the CLI version information

Flags:
  -h, --help             Show context-sensitive help.
      --config=STRING    path to config file

```

#### .env file options:


* `ETH_PRIVATE_KEYS` \(required\) - list of private keys separated by `,`

* `NODE_URL` \(required\) - websocket node URL \(e.g [wss://mainnet.infura.io/bbbb](wss://mainnet.infura.io/bbbb) or [wss://localhost:8546](ws://localhost:8546) if own node\)


#### Config file options:
```json
{
	"Aggregator": {
		"LogLevel": "Required:false, Default:",
		"ManualDataFile": "Required:false, Default:configs/manualData.json"
	},
	"Db": {
		"LogLevel": "Required:false, Default:",
		"Path": "Required:false, Default:db",
		"RemoteHost": "Required:false, Default:",
		"RemotePort": "Required:false, Default:0",
		"RemoteTimeout": {
			"Duration": "Required:false, Default:5s"
		}
	},
	"Mining": {
		"Heartbeat": "Required:false, Default:1m0s",
		"LogLevel": "Required:false, Default:"
	},
	"PsrTellor": {
		"MinConfidence": "Required:false, Default:80"
	},
	"SubmitterTellor": {
		"Enabled": "Required:false, Default:true",
		"LogLevel": "Required:false, Default:",
		"MinSubmitPeriod": {
			"Duration": "Required:false, Default:15m1s"
		}
	},
	"Tasker": {
		"LogLevel": "Required:false, Default:"
	},
	"TrackerIndex": {
		"IndexFile": "Required:false, Default:configs/index.json",
		"Interval": {
			"Duration": "Required:false, Default:1m0s"
		},
		"LogLevel": "Required:false, Default:"
	},
	"TransactorTellor": {
		"GasMaxTipGwei": "Required:false, Default:10, Description:Hard limit of the gas tip in Gwei.",
		"LogLevel": "Required:false, Default:",
		"ProfitThreshold": "Required:false, Default:0, Description:Minimum percent of profit when submitting a solution. For example if the tx cost is 0.01 ETH and current reward is 0.02 ETH a ProfitThreshold of 200% or more will wait until the reward is increased or the gas cost is lowered a ProfitThreshold of 199% or less will submit."
	},
	"Web": {
		"ListenHost": "Required:false, Default:",
		"ListenPort": "Required:false, Default:9090",
		"LogLevel": "Required:false, Default:",
		"ReadTimeout": {
			"Duration": "Required:false, Default:0s"
		}
	},
	"envFile": "Required:false, Default:configs/.env"
}
```
Here are the config defaults in json format:
```json
{
	"Aggregator": {
		"LogLevel": "",
		"ManualDataFile": "configs/manualData.json"
	},
	"Db": {
		"LogLevel": "",
		"Path": "db",
		"RemoteHost": "",
		"RemotePort": 0,
		"RemoteTimeout": "5s"
	},
	"Mining": {
		"Heartbeat": 60000000000,
		"LogLevel": ""
	},
	"PsrTellor": {
		"MinConfidence": 80
	},
	"SubmitterTellor": {
		"Enabled": true,
		"LogLevel": "",
		"MinSubmitPeriod": "15m1s"
	},
	"Tasker": {
		"LogLevel": ""
	},
	"TrackerIndex": {
		"IndexFile": "configs/index.json",
		"Interval": "1m0s",
		"LogLevel": ""
	},
	"TransactorTellor": {
		"GasMaxTipGwei": 10,
		"LogLevel": "",
		"ProfitThreshold": 0
	},
	"Web": {
		"ListenHost": "",
		"ListenPort": 9090,
		"LogLevel": "",
		"ReadTimeout": "0s"
	},
	"envFile": "configs/.env"
}
```
### Log levels
Note the default level is "INFO", so to turn down the number of logs, enter "WARN" or "ERROR".

DEBUG - logs everything in INFO and additional developer logs

INFO - logs most information about the reporting operation

WARN - logs all warnings and errors

ERROR - logs only serious errors
