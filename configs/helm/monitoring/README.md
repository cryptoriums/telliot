# monitoring

![Version: 1.0.0](https://img.shields.io/badge/Version-1.0.0-informational?style=flat-square) ![Type: application](https://img.shields.io/badge/Type-application-informational?style=flat-square) ![AppVersion: 0.0.2](https://img.shields.io/badge/AppVersion-0.0.2-informational?style=flat-square)

A Helm chart for Kubernetes

## Values

| Key | Type | Default | Description |
|-----|------|---------|-------------|
| grafana.container.image | string | `"grafana/grafana:8.1.0"` | Docker image for grafana |
| grafana.container.port | int | `3000` |  |
| grafana.ingress.class | string | `"nginx"` | Ingress class to use for grafana |
| grafana.ingress.hostname | string | `"monitor"` | Hostname to use for accessing grafana |
| grafana.ingress.path | string | `"/"` | Subpath to access grafana |
| grafana.ingress.tls.enabled | bool | `false` | Enable/Disable TLS for grafana |
| grafana.ingress.tls.secret | string | `"grafana-tls-secret"` | Name of TLS secret to use for grafana |
| grafana.loadBalancing | bool | `false` |  |
| grafana.serveFromSubPath | bool | `false` |  |
| grafana.service.port | int | `80` |  |
| grafana.storage | int | `1` | Grafana storage size in GB |
| namespace | string | `"tellor"` |  |
| prometheus.container.image | string | `"prom/prometheus:v2.28.1"` | Docker image for prometheus |
| prometheus.container.port | int | `9090` |  |
| prometheus.enabled | bool | `false` |  |
| prometheus.service.port | int | `9090` |  |
| prometheus.storage | int | `10` | Prometheus storage in GB |

