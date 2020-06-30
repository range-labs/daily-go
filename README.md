# Go Daily.co

[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/range-labs/daily-go)
[![License](https://img.shields.io/github/license/range-labs/daily-go.svg)](https://github.com/range-labs/daily-go/blob/master/LICENSE)
[![Twitter](https://img.shields.io/twitter/follow/rangelabs.svg?style=social)](https://twitter.com/rangelabs)

Unofficial [Daily.co](https://daily.co) Go client library.

## Documentation

Refer to [Daily.co's API reference](https://docs.daily.co/reference) for details
about the underlying REST API, requests, and responses.

## Usage

```go
client := daily.New(daily.WithAuth(API_KEY))
cfg, err := client.GetDomainConfig(context.Background())
```
