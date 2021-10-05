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
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

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
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

      --gas-price=FLOAT-64       gas max fee to use when running the command
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
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

```

* `data`

```
Usage: telliot data

Retrieve data from the contract

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

      --from=UINT-64             the unix timestamp to use as a starting point
                                 for the data retrieval
      --look-back=2h             how far to lookback

```

* `dataserver`

```
Usage: telliot dataserver

launch only a dataserver instance

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

```

* `decrypt`

```
Usage: telliot decrypt <file>

Decrypts an ecrypted file and write the decrytped version to disk

Arguments:
  <file>    the file to encrypt

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

```

* `dispute`

```
Usage: telliot dispute <command>

Perform commands related to disputes

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

Commands:
  dispute new --data-id=INT-64 --timestamp=INT-64 --slot=INT-64 <account>
    start a new dispute

  dispute vote --dispute-id=INT-64 --support <account>
    vote on a open dispute

  dispute list
    list open disputes

  dispute tally
    tally votes for a dispute ID, need to run unlock fee after that

  dispute tally-list
    list tally for disputes

  dispute unlock-fee
    after tallying the votes this command transfers the fee to the reporter or
    reported based on the voting

```

* `dispute list`

```
Usage: telliot dispute list

list open disputes

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

      --show-closed              also show executed disputes
      --look-back=120h           how far to lookback, the default only few days
                                 since disputes can be voted only for 2 days.

```

* `dispute new`

```
Usage: telliot dispute new --data-id=INT-64 --timestamp=INT-64 --slot=INT-64 <account>

start a new dispute

Arguments:
  <account>

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

      --gas-price=FLOAT-64       gas max fee to use when running the command
      --no-checks                whether to run some additional checks like
                                 balance, did vote etc. Useful to disable when
                                 running against a test contract
      --data-id=INT-64           the request id to dispute
      --timestamp=INT-64         the submitted timestamp to dispute
      --slot=INT-64              the reporter index to dispute

```

* `dispute tally`

```
Usage: telliot dispute tally

tally votes for a dispute ID, need to run unlock fee after that

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

      --gas-price=FLOAT-64       gas max fee to use when running the command
      --dispute-id=INT-64        the dispute id
      --all

```

* `dispute tally-list`

```
Usage: telliot dispute tally-list

list tally for disputes

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

      --look-back=120h           how far to lookback, the default only few days
                                 since disputes can be voted only for 2 days.

```

* `dispute unlock-fee`

```
Usage: telliot dispute unlock-fee

after tallying the votes this command transfers the fee to the reporter or
reported based on the voting

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

      --no-checks                whether to run some additional checks like
                                 balance, did vote etc. Useful to disable when
                                 running against a test contract
      --gas-price=FLOAT-64       gas max fee to use when running the command
      --dispute-id=INT-64        the dispute id
      --all

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
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

      --gas-price=FLOAT-64       gas max fee to use when running the command
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
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

```

* `events`

```
Usage: telliot events --event-name=STRING

Subscribe to watch logs from the network.

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

      --look-back=DURATION       how far to look for the initiali qyery
      --event-name=STRING        the name of the log to watch
      --reorg-wait=3s            how long to wait for removed logs from reorg
                                 events

```

* `report`

```
Usage: telliot report

Submit data to the oracle contracts

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

```

* `stake`

```
Usage: telliot stake <command>

Perform one of the stake operations

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

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
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

      --gas-price=FLOAT-64       gas max fee to use when running the command
      --no-checks                whether to run some additional checks like
                                 balance, did vote etc. Useful to disable when
                                 running against a test contract

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
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

      --gas-price=FLOAT-64       gas max fee to use when running the command

```

* `stake status`

```
Usage: telliot stake status <account>

show stake status

Arguments:
  <account>

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

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
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

      --gas-price=FLOAT-64       gas max fee to use when running the command

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
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

      --gas-price=FLOAT-64       gas max fee to use when running the command
      --custom-submit
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
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

      --gas-price=FLOAT-64       gas max fee to use when running the command
      --from=STRING
      --to=STRING

```

* `version`

```
Usage: telliot version

Show the CLI version information

Flags:
  -h, --help                     Show context-sensitive help.
      --config=STRING            path to config file
      --config-strict-parsing    whether to return an error when the config
                                 contains unknown fields
      --contract=STRING          hex address of the contract to interract with

```

#### .env file options:


* `ETH_PRIVATE_KEYS` \(required\) - list of private keys separated by `,`

* `NODE_URL` \(required\) - websocket node URL \(e.g [wss://mainnet.infura.io/bbbb](wss://mainnet.infura.io/bbbb) or [wss://localhost:8546](ws://localhost:8546) if own node\)


#### Config file options:
```json
{
	"Aggregator": {
		"LogLevel": "Required:false, Default:"
	},
	"Db": {
		"LogLevel": "Required:false, Default:",
		"Path": "Required:false, Default:db",
		"RemoteHost": "Required:false, Default:",
		"RemotePort": "Required:false, Default:9090",
		"RemoteTimeout": {
			"Duration": "Required:false, Default:5s"
		}
	},
	"Mining": {
		"Heartbeat": "Required:false, Default:1m0s",
		"LogLevel": "Required:false, Default:"
	},
	"PsrTellor": {
		"MinConfidenceDefault": "Required:false, Default:80",
		"MinConfidencePerSymbol": "Required:false, Default:map[41:100]"
	},
	"SubmitterTellor": {
		"LogLevel": "Required:false, Default:",
		"MinSubmitPeriod": {
			"Duration": "Required:false, Default:15m30s"
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
		"Transact": "Required:false, Default:true"
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
		"LogLevel": ""
	},
	"Db": {
		"LogLevel": "",
		"Path": "db",
		"RemoteHost": "",
		"RemotePort": 9090,
		"RemoteTimeout": "5s"
	},
	"Mining": {
		"Heartbeat": 60000000000,
		"LogLevel": ""
	},
	"PsrTellor": {
		"MinConfidenceDefault": 80,
		"MinConfidencePerSymbol": {
			"41": 100
		}
	},
	"SubmitterTellor": {
		"LogLevel": "",
		"MinSubmitPeriod": "15m30s"
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
		"Transact": true
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
