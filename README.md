# Client Server API - Go Expert
[![Go](https://img.shields.io/badge/go-1.21.0-informational?logo=go)](https://go.dev)

This project implements the first challenge - Client Server API - for the Postgraduate in Go Expert.

# Index
- [Client Server API - Go Expert](#client-server-api---go-expert)
- [Index](#index)
- [Stack](#stack)
  - [Start project](#start-project)
    - [Start the server](#start-the-server)
      - [Makefile](#makefile)
      - [Go command](#go-command)
    - [Start client](#start-client)
      - [Makefile](#makefile-1)
      - [Go command](#go-command-1)
  - [Clean project](#clean-project)
    - [Makefile](#makefile-2)
    - [Command](#command)
- [Endpoint](#endpoint)
  - [Getting cotacao](#getting-cotacao)
    - [Request example](#request-example)
    - [Body example](#body-example)
    - [Response example](#response-example)

# Stack
- [Golang](https://go.dev/)

## Start project

### Start the server
Start the Server with the following command:

#### Makefile
```sh
make run-server
```

#### Go command
```sh
go run ./server/server.go
```

### Start client
Once our server is up, start the Client with the following command:

#### Makefile
```sh
make run-client
```

#### Go command
```sh
go run ./client/client.go
```

## Clean project
### Makefile
```sh
make clean-project
```

### Command
```sh
rm quote.db cotacao.txt
```

# Endpoint
| Method | Resource         |
|:------:|:-----------------|
| GET    | /cotacao         |

## Getting cotacao

 - GET - /cotacao

### Request example

```sh
curl --location --request GET 'http://localhost:8080/cotacao'
```

### Body example
```sh
No body
```

### Response example
```json
{
  "code": "USD",
  "codein": "BRL",
  "name": "Dólar Americano/Real Brasileiro",
  "high": "4.9552",
  "low": "4.9008",
  "varBid": "-0.0069",
  "pctChange": "-0.14",
  "bid": "4.946",
  "ask": "4.949",
  "timestamp": "1693601998",
  "create_date": "2023-09-01 17:59:58"
}
```