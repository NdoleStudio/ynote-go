# ynote-go

[![Build](https://github.com/NdoleStudio/ynote-go/actions/workflows/main.yml/badge.svg)](https://github.com/NdoleStudio/ynote-go/actions/workflows/main.yml)
[![codecov](https://codecov.io/gh/NdoleStudio/ynote-go/branch/main/graph/badge.svg)](https://codecov.io/gh/NdoleStudio/ynote-go)
[![Scrutinizer Code Quality](https://scrutinizer-ci.com/g/NdoleStudio/ynote-go/badges/quality-score.png?b=main)](https://scrutinizer-ci.com/g/NdoleStudio/ynote-go/?branch=main)
[![Go Report Card](https://goreportcard.com/badge/github.com/NdoleStudio/ynote-go)](https://goreportcard.com/report/github.com/NdoleStudio/ynote-go)
[![GitHub contributors](https://img.shields.io/github/contributors/NdoleStudio/ynote-go)](https://github.com/NdoleStudio/ynote-go/graphs/contributors)
[![GitHub license](https://img.shields.io/github/license/NdoleStudio/ynote-go?color=brightgreen)](https://github.com/NdoleStudio/ynote-go/blob/master/LICENSE)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/NdoleStudio/ynote-go)](https://pkg.go.dev/github.com/NdoleStudio/ynote-go)


This package provides a generic `go` client for the Y-Note API

## Installation

`ynote-go` is compatible with modern Go releases in module mode, with Go installed:

```bash
go get github.com/NdoleStudio/ynote-go
```

Alternatively the same can be achieved if you use `import` in a package:

```go
import "github.com/NdoleStudio/ynote-go"
```


## Implemented

- **Token**
  - `POST /oauth2/token`: Get Access Token
- **Refund**
  - `POST /prod/refund`: Refund a transaction
  - `GET /prod/refund/status/{transactionID}`: Get the status of a refund transaction

## Usage

### Initializing the Client

An instance of the client can be created using `New()`.

```go
package main

import (
    "github.com/NdoleStudio/ynote-go"
)

func main() {
    client := ynote.New(
        ynote.WithUsername(""),
        ynote.WithPassword(""),
    )
}
```

### Error handling

All API calls return an `error` as the last return object. All successful calls will return a `nil` error.

```go
transaction, response, err := client.Refund.Status(context.Background(), "")
if err != nil {
    //handle error
}
```

## Testing

You can run the unit tests for this client from the root directory using the command below:

```bash
go test -v
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details
