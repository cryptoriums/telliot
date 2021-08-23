---
description: Here are the nuts and bolts for using the CLI
---

# Setup and usage

## Get the CLI

The CLI support only linux and is provided as a pre-built binary with every release and also as a docker image.

[Github releases](https://github.com/cryptoriums/telliot/releases)

[https://hub.docker.com/r/cryptoriums/telliot](https://hub.docker.com/r/cryptoriums/telliot)

## Config files.
 - `.env` - keeps private information(private keys, api keys etc.). Most commands require some secrets and these are kept in this file as a precaution against accidental exposure. For a working setup it is required to at least add one private key in your `"ETH_PRIVATE_KEYS"` environment variable. Multiple private keys are supported separated by `,`.
 - `index.json` - all api endpoint for data providers. The cli uses these provider endpoints to gather data which is then used to submit to the onchain oracle.
 - `manualdata.json` - for providing data manually.
 For testing purposes, or if you want to hardcode in a specific value, you can use the file to add manual data for a given requestID. Add the request ID, a given value \(with granularity\), and a date on which the manual data expires.
The following example shows request ID 4, inputting a value of 9000 with 6 digits granularity. Note the date is a unix timestamp.
```bash
"4":{
    "VALUE":9000.123456,
    "DATE":1596153600
}
```
 - `config.json` - optional config file to override any of the defaults. See the [configuration page](configuration.md) for full reference.


> by default the cli looks for these in the `./configs` folder relative to the cli folder.

### Here is a quick reference how to run the cli with the default configs.

```
mkdir ./configs
cd ./configs
wget https://raw.githubusercontent.com/cryptoriums/telliot/master/configs/index.json
wget https://raw.githubusercontent.com/cryptoriums/telliot/master/configs/manualData.json
wget https://raw.githubusercontent.com/cryptoriums/telliot/master/configs/.env.example
mv .env.example .env
cd ../
wget https://github.com/cryptoriums/telliot/releases/latest/download/telliot
chmod +x telliot
```

## Deposit or withdraw a stake

As of now, mining requires you to deposit 500 TRB to be allowed to submit values to the oracle and earn rewards. This is a security deposit. If you are a malicious actor \(aka submit a bad value\), the community can vote to slash your 500 tokens.
Your stake is locked for a minimum of 7 days after you run the command to request withdrawal.

Run the following command to deposit your stake:

```bash
./telliot stake deposit
```

To unstake your tokens, you need to request a withdraw:

```bash
./telliot stake request
```

One week after the request, the tokens are free to move at your discretion after running the command:

```bash
./telliot stake withdraw
```

## Start mining.
{% hint style="info" %}
The same instance can be used with multiple private keys in the `.env` file separated by a comma.
{% endhint %}

```bash
./telliot report
```


## DataServer - a shared data API feeds.

{% hint style="info" %}
Advanced usage! If you are setting up a Tellor reporter for the first time, it might be a good idea to skip this section and come back after you're up and running with one reporter. See the [configuration page](configuration.md) for the required configs.
{% endhint %}

Some oracle feeds require 24h avarages and for these enough historical data is needed. Running a dataserver is the solution to always have enough historical data to generate these averages.

The network topology of this setup looks like the diagram below.
One ore more reporters are connected to the same data server for fetching current or historical data to submit to the oracle.
The data server pulls data from the API providers, the 5 staked reporters pull data from the data server and submit on-chain to the Tellor Core smart contracts.

```bash
                            /(0xE037)\
                Reporter   | (0xcdd8) |
Tellor     <-> (multiple   | (0xb9dD) | <-> Data Server <-> Data APIs
(on chain)      keys)      | (0x2305) |
                            \(0x3233)/
```


## Run with Docker - [https://hub.docker.com/u/tellor](https://hub.docker.com/u/tellor)

```bash
cp configs/.env.example configs/.env # Edit the file after the copy.
docker run -v $(pwd)/configs:/configs cryptoriums/telliot:master report
```

### To build a custom docker image.

```bash
export REPO= # Your docker repository name.
docker build . -t $REPO/telliot:master
docker push $REPO/telliot:master

```

## Run on Kubernetes with Helm Chart

A Helm chart for installing telliot on Kubernetes

[A guide for installing helm can be found here](https://helm.sh/docs/intro/install/)

## telliot Configuration

Include telliot configuration files in the files directory of this chart.

These files should be:

- .env
- index.json

Run the following to use the default configuration files:

```bash
cp configs/.env.example configs/helm/telliot/files/.env # you will need to edit this file with your own secrets after copying
```

Optionally you can also include config.json if you would like to override any default config values.

If you would like to use separate .env or config.json files for the mining and dataserver instances, copy those files under their respective directory in config/helm/telliot/files.
For example to use a unique config.json and .env for mining and dataserver instances run:

```bash
# Mining instance
cp configs/config.json configs/helm/telliot/files/reporter/
cp configs/.env.example configs/helm/telliot/files/reporter/.env
# dataserver instance
cp configs/config.json configs/helm/telliot/files/dataserver/
cp configs/.env.example configs/helm/telliot/files/dataserver/.env
```

## Usage

After you have moved your configuration files to config/helm/telliot/files, you can install this chart using the following command:

```bash
export INSTANCE_NAME=lat
helm install $INSTANCE_NAME configs/helm/telliot/ --create-namespace
```

INSTANCE_NAME being a string you would use to denote this instance of telliot.

Keep in mind this command is using all default values.

## Values

The default helm values will install a mining instance of telliot.

To override these values during installation include `--set "key=value"` in the helm upgrade command.

Here is an example, to run a dataserver instance of telliot using some custom settings:

```bash
export INSTANCE_NAME=lat
helm install $INSTANCE_NAME configs/helm/telliot \
    --create-namespace \
    --set "storage=1Gi" \
    --set "command={dataserver}"
```

Or to run telliot to report values:

```bash
export INSTANCE_NAME=lat
helm install $INSTANCE_NAME configs/helm/telliot --create-namespace --set "command={report}"
```

A full list of values and their description can be found [here](https://github.com/cryptoriums/telliot/blob/master/configs/helm/telliot/README.md)

## Monitoring

Monitoring is recommended and can be installed using a separate helm chart.

To install monitoring with the default values run:

```bash
helm install mon configs/helm/monitoring --set="grafana.ingress.hostname=monitor-telliot.mywire.org"
```
A full list of values and their description can be found [here](https://github.com/cryptoriums/telliot/blob/master/configs/helm/monitoring/README.md)


## Upgrade

To update a telliot instance run

```bash
helm upgrade $INSTANCE_NAME configs/helm/telliot
```
and for the monitoring
```bash
helm upgrade mon configs/helm/monitoring
```

With `$key=$value` being the desired value changes you would like to make.

If you would like to make any config changes, move the updated configuration file to configs/helm/files/ and run

```bash
helm upgrade $INSTANCE_NAME configs/helm/
```

## Removal

Uninstalling an instance of telliot using helm is as simple as

```bash
helm uninstall $INSTANCE_NAME
```

Where instance name was the name for the release you specified during installation. You can find this value by running

```bash
helm list
```
