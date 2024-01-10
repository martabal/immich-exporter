# immich-exporter

[![Publish Release](https://github.com/martabal/immich-exporter/actions/workflows/docker.yml/badge.svg)](https://github.com/martabal/immich-exporter/actions/workflows/docker.yml)
[![Build](https://github.com/martabal/immich-exporter/actions/workflows/build.yml/badge.svg)](https://github.com/martabal/immich-exporter/actions/workflows/build.yml)
[![Test](https://github.com/martabal/immich-exporter/actions/workflows/test.yml/badge.svg)](https://github.com/martabal/immich-exporter/actions/workflows/test.yml)

<p align="center">
<img src="img/immich.png" width=100> &nbsp; <img src="img/prometheus.png" width=100><img src="img/golang.png" width=100>
</p>

This app is a Prometheus exporter for immich.  
This app is made to be integrated with the [immich-grafana-dashboard](https://github.com/martabal/immich-exporter/blob/main/grafana/dashboard.json)  

## Run it

Create an API key in your Immich settings and set `IMMICH_API_KEY` to is value.

### Docker-cli ([click here for more info](https://docs.docker.com/engine/reference/commandline/cli/))

```sh
docker run --name=immich-exporter \
    -e IMMICH_BASE_URL=http://192.168.1.10:8080 \
    -e IMMICH_API_KEY=<your_api_key> \
    -p 8090:8090 \
    ghcr.io/martabal/immich-exporter
```

### Docker-compose

```yaml
version: "2.1"
services:
  immich:
    image: ghcr.io/martabal/immich-exporter:latest
    container_name: immich-exporter
    environment:
      - IMMICH_BASE_URL=http://192.168.1.10:8080
      - IMMICH_API_KEY=<your_api_key>
    ports:
      - 8090:8090
    restart: unless-stopped
```

### Without docker

```sh
git clone https://github.com/martabal/immich-exporter.git
cd immich-exporter/
go get -d -v
cd src
go build -o ./immich-exporter
./immich-exporter
```

If you want to use an .env file, edit `.env.example` to match your setup, rename it `.env` then run it with :

```sh
./immich-exporter -e
```

## Parameters

### Environment variables

| Parameters | Function |
| :-----: | ----- |
| `-p 8090` | Webservice port |
| `-e IMMICH_BASE_URL` | Immich base URL |
| `-e IMMICH_API_KEY` | Immich API key  |
| `-e EXPORTER_PORT` | qbittorrent export port (optional) | `8090` |
| `-e LOG_LEVEL` | App log level (`TRACE`, `DEBUG`, `INFO`, `WARN` and `ERROR`) | `INFO` |

### Arguments

| Arguments | Function |
| :-----: | ----- |
| -e | If qbittorrent-exporter detects a .env file in the same directory, the values in the .env will be used, `-e` forces the usage of environment variables |

### Setup

Add the target to your `scrape_configs` in your `prometheus.yml` file of your Prometheus instance.

```yaml
scrape_configs:
  - job_name: 'immich'
    static_configs:
      - targets: [ '<your_ip_address>:8090' ]
```
