# Database as a Service

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/database-as-a-service)](https://pkg.go.dev/github.com/go-zoox/database-as-a-service)
[![Build Status](https://github.com/go-zoox/database-as-a-service/actions/workflows/release.yml/badge.svg?branch=master)](https://github.com/go-zoox/database-as-a-service/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/database-as-a-service)](https://goreportcard.com/report/github.com/go-zoox/database-as-a-service)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/database-as-a-service/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/database-as-a-service?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/database-as-a-service.svg)](https://github.com/go-zoox/database-as-a-service/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/database-as-a-service.svg?label=Release)](https://github.com/go-zoox/database-as-a-service/tags)


## Installation
To install the package, run:

```bash
# with go
go install github.com/go-zoox/database-as-a-service/cmd/daas@latest
```

if you dont have go installed, you can use the install script (zmicro package manager):

```bash
curl -o- https://raw.githubusercontent.com/zcorky/zmicro/master/install | bash

zmicro package install daas
```

## Features
* Engine
  * [x] MySQL
  * [x] PostgreSQL
  * [x] SQLite3
  * [ ] SQL Server
  * [ ] Oracle
  * [ ] MongoDB
  * [ ] Redis

## Quick Start

### Start DaaS Server

```bash
database-as-a-service server
```

### Connect DaaS with Client

```bash
database-as-a-service client --server http://127.0.0.1:8838
```

### Connect DaaS with API

```bash
# MySQL / PostgreSQL:
curl --location 'http://127.0.0.1:9998' \
--header 'Content-Type: application/json' \
--data-raw '{
    "engine": "postgres",
    "dsn": "postgres://user:pass@10.0.0.1:5432/daas?sslmode=disable",
    "statement": "select name as label, id as value from my_table limit 1000"
}'
```

```bash
# Sqlite3:
curl --location 'http://127.0.0.1:9998' \
--header 'Content-Type: application/json' \
--data '{
    "engine": "sqlite3",
    "dsn": "https://sqliteviewer.app/Chinook_Sqlite.sqlite",
    "statement": "SELECT * FROM Album"
}'
```

## Usage

### Server

```bash
database-as-a-service server --help

NAME:
   daas server - database as a service server

USAGE:
   daas server [command options] [arguments...]

OPTIONS:
   --port value, -p value  server port (default: 8080) [$PORT]
   --path value            api path
   --username value        Username for Basic Auth [$USERNAME]
   --password value        Password for Basic Auth [$PASSWORD]
   --help, -h              show help
```

### Client

```bash
database-as-a-service client --help

NAME:
   daas client - database as a service client

USAGE:
   daas client [command options] [arguments...]

OPTIONS:
   --server value, -s value  server url [$SERVER]
   --engine value            database engine, e.g. mysql, postgres, sqlite3 [$ENGINE]
   --dsn value               database dsn [$DSN]
   --statement value         database statement [$STATEMENT]
   --username value          Username for Basic Auth [$USERNAME]
   --password value          Password for Basic Auth [$PASSWORD]
   --help, -h                show help
```


## License
GoZoox is released under the [MIT License](./LICENSE).
