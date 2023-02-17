# immich-exporter

[![Publish Release](https://github.com/martabal/immich-exporter/actions/workflows/push_docker.yml/badge.svg)](https://github.com/martabal/immich-exporter/actions/workflows/push_docker.yml)

<p align="center">
<img src="img/immich.png" width=100> &nbsp; <img src="img/prometheus.png" width=100><img src="img/golang.png" width=100>
</p>

This app is a Prometheus exporter for immich.  
This app is made to be integrated with the [immich-grafana-dashboard](https://github.com/martabal/immich-exporter/blob/main/grafana/dashboard.json)  

## Run it

### Docker-cli ([click here for more info](https://docs.docker.com/engine/reference/commandline/cli/))

```sh
docker run --name=immich-exporter \
    -e IMMICH_URL=http://192.168.1.10:8080 \
    -e IMMICH_PASSWORD='<your_password>' \
    -e IMMICH_USERNAME=admin \
    -p 8090:8090 \
    martabal/immich-exporter
```

### Docker-compose

```yaml
version: "2.1"
services:
  immich:
    image: martabal/immich-exporter:latest
    container_name: immich-exporter
    environment:
      - IMMICH_URL=http://192.168.1.10:8080
      - IMMICH_PASSWORD='<your_password>'
      - IMMICH_USERNAME=admin
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
| `-e IMMICH_USERNAME` | Immich username |
| `-e IMMICH_PASSWORD` | Immich password |
| `-e IMMICH_BASE_URL` | Immich base URL |

### Arguments

| Arguments | Function |
| :-----: | ----- |
| -e | Use a .env file containing environment variables (.env file must be placed in the same directory) |
