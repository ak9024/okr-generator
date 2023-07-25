<div align="center">
    <img src="./assets/sq-backend.png" />
    <h2>Backend (okr-generator)</h2>
</div>

### Prerequisite

- Go

### Getting Started

> How to run for development mode

```bash
# please add config, and suitable with your preference
cp .config.toml.example .config.toml

# run the server
go run . server
```

> How to run with Docker

```bash
# Please create .config.toml.prod
cp .config.toml.example .config.toml.prod

# Please edit `.config.toml.prod` suitable with your configuration

# build docker image
docker build -t backend:latest .
# run the server
docker run -d -p <port_external>:<port_internal> --name <container_name> backend:latest
# check with
docker ps
```

> How to run with Docker Compose

```bash
cd ../
docker-compose up -d
```

> The configuration

```toml
[app]
port =
host = ""
version = ""

[chatgpt]
token = ""
```
