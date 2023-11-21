# CoinGecko API

[![Go Version](https://img.shields.io/badge/go-v1.21.1-green.svg)](https://golang.org/dl/)
[![Code Lint](https://github.com/bufdata/coingecko-api/actions/workflows/code-lint.yml/badge.svg)](https://github.com/bufdata/coingecko-api/actions/workflows/code-lint.yml)
[![Unit Test](https://github.com/bufdata/coingecko-api/actions/workflows/unit-test.yml/badge.svg)](https://github.com/bufdata/coingecko-api/actions/workflows/unit-test.yml)
[![GoDoc](https://pkg.go.dev/github.com/bufdata/coingecko-api?status.svg)](https://pkg.go.dev/github.com/bufdata/coingecko-api)
[![goreports](https://goreportcard.com/badge/github.com/bufdata/coingecko-api)](https://goreportcard.com/report/github.com/bufdata/coingecko-api)
[![Codecov](https://codecov.io/gh/bufdata/coingecko-api/branch/master/graph/badge.svg)](https://codecov.io/gh/bufdata/coingecko-api)
[![license](https://img.shields.io/badge/license-MIT-blue)](https://github.com/bufdata/coingecko-api/blob/main/LICENSE)
[![Build Status](https://travis-ci.com/bufdata/coingecko-api.svg?branch=main)](https://travis-ci.com/bufdata/coingecko-api)

[![](https://static.coingecko.com/s/coingecko-logo-d13d6bcceddbb003f146b33c2f7e8193d72b93bb343d38e392897c3df3e78bdd.png)](https://coingecko.com)
[![](https://www.geckoterminal.com/_next/static/media/logo_dark.6b1547fe.png)](https://www.geckoterminal.com)

This library contains two apis:

* [CoinGecko API](https://apiguide.coingecko.com/getting-started/introduction)
* [GeckoTerminal API](https://apiguide.geckoterminal.com/)

## Installation

```shell
go get -u github.com/bufdata/coingecko-api
```

## Usage

### CoinGecko

For `free API`, you can provide an API Key or not:

```go
package main

import (
	"context"
	"log/slog"

	"github.com/bufdata/coingecko-api/coingecko"
)

func main() {
	api := coingecko.NewCoinGecko("your_api_key", false, nil)
	data, err := api.ListCoinsInfo(context.Background(), true)
	if err != nil {
		slog.Error("failed to call ListCoinsInfo", "error", err)
	}
	slog.Info("call ListSupportedCoinsInfo successfully", "response data", (*data)[0])
}
```

For users with `Pro API Key`:

```go
package main

import (
	"context"
	"log/slog"

	"github.com/bufdata/coingecko-api/coingecko"
)

func main() {
	api := coingecko.NewCoinGecko("your_api_key", true, nil)
	data, err := api.ListCoinsInfo(context.Background(), true)
	if err != nil {
		slog.Error("failed to call ListCoinsInfo", "error", err)
	}
	slog.Info("call ListSupportedCoinsInfo successfully", "response data", (*data)[0])
}
```

If you use `coingecko` library in production, you might need to set your own `http.Client` param.

This library has covered all APIs. For detailed APIs info, you can read [CoinGecko docs](https://www.coingecko.com/api/documentation).

**Note**

The methods in `coingecko/paid_methods.go` and `coingecko/enterprise_methods.go` are not tested, so if you find there 
are some bugs in them, you can raise pr or issue.

### GeckoTerminal

```go
package main

import (
	"context"
	"log/slog"

	"github.com/bufdata/coingecko-api/geckoterminal"
)

func main() {
	api := geckoterminal.NewGeckoTerminal(nil)
	data, err := api.GetNetworks(context.Background(), 0)
	if err != nil {
		slog.Error("failed to call GetNetworks", "error", err)
	}
	slog.Info("call GetNetworks successfully", "response data", *data)
}
```

If you use `geckoterminal` library in production, you might need to set your own `http.Client` param.

This library has covered all APIs. For detailed APIs info, you can read [GeckoTerminal API](https://apiguide.geckoterminal.com/).

## License

[MIT](https://choosealicense.com/licenses/mit/)
