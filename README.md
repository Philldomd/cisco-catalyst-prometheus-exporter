# cisco-catalyst-prometheus-exporter
This repository is a direct reimplementation of [Hawar Koyi](https://hawar.no/2021/05/cisco-dna-center-with-grafana-dashboard/) DNA dashboard but in Golang. This does not require any database, instead it will call the healt and metric ports upon scraping and deliver the information in prometheus format.
## Contribution
---
## Setup
---
## Configuration
| Key | Value | Description |
| --- | --- | --- |
| server.name | 127.0.0.1 | Hostname of the server |
| server.port | 9000 | Port serving metrics |
| certificate.crt |  \<path>/\<pub> | Path to public server key |
| certificate.key | \<path>/\<key> | Path to private server key |
| cisco.dna_url | /\<dna_url> | Url to cisco service |
| cisco.auth_url | /dna/system/api/v1/auth/token | Url to Authentication service |
| cisco.token | \<auth_token> | API authentication token |
| cisco.site_health | /dna/intent/api/v1/site-health | Url for site health |
| cisco.network_health | /dna/intent/api/v1/network-health | Url for network heatlth |
| cisco.client_health | /dna/intent/api/v1/client-health | Url for client health |
| cisco.devices_list | /dna/intent/api/v1/network-device | Url for network devices |
---