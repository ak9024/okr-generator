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
