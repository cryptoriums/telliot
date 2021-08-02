# telliot

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.0.2](https://img.shields.io/badge/AppVersion-0.0.2-informational?style=flat-square)

A Helm chart for Kubernetes

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| command | list | `["dataserver"]` | Array of commands to spawn separate instances of telliot instances with |
| container.image | string | `"cryptoriums/telliot:master"` | Docker image for telliot |
| container.port | int | `9090` |  |
| namespace | string | `"tellor"` |  |
| service.port | int | `9090` |  |
| storage | int | `1` | telliot persistent storage size in GB |

