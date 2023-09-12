# go-api

[![Go Reference](https://pkg.go.dev/badge/github.com/khulnasoft-lab/go-api.svg)](https://pkg.go.dev/github.com/khulnasoft-lab/go-api)
![Test](https://github.com/khulnasoft-lab/go-api/workflows/Test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/khulnasoft-lab/go-api?style=flat-square)](https://goreportcard.com/report/github.com/khulnasoft-lab/go-api)

> **Note**: This library is under active development as we expand it to cover
> our (expanding!) API. Consider the public API of this package a little
> unstable as we work towards a v1.0.

A Go library for interacting with
[Khulnasoft's API v4](https://api.khulnasoft.com/). This library allows you to:

- Manage and automate changes to your DNS records within Khulnasoft
- Manage and automate changes to your zones (domains) on Khulnasoft, including
  adding new zones to your account
- List and modify the status of WAF (Web Application Firewall) rules for your
  zones
- Fetch Khulnasoft's IP ranges for automating your firewall whitelisting

A command-line client, [flarectl](cmd/flarectl), is also available as part of
this project.

## Installation

You need a working Go environment. We officially support only currently supported Go versions according to [Go project's release policy](https://go.dev/doc/devel/release#policy).

```
go get github.com/khulnasoft-lab/go-api
```

## Getting Started

```go
package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/khulnasoft-lab/go-api"
)

func main() {
	// Construct a new API object using a global API key
	api, err := khulnasoft.New(os.Getenv("KHULNASOFT_API_KEY"), os.Getenv("KHULNASOFT_API_EMAIL"))
	// alternatively, you can use a scoped API token
	// api, err := khulnasoft.NewWithAPIToken(os.Getenv("KHULNASOFT_API_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Most API calls require a Context
	ctx := context.Background()

	// Fetch user details on the account
	u, err := api.UserDetails(ctx)
	if err != nil {
		log.Fatal(err)
	}
	// Print user details
	fmt.Println(u)
}
```

Also refer to the
[API documentation](https://pkg.go.dev/github.com/khulnasoft-lab/go-api) for
how to use this package in-depth.

## Experimental improvements

This library is starting to ship with experimental improvements that are not yet
ready for production but will be introduced before the next major version. See
[experimental README](/docs/experimental.md) for full details.

## Contributing

Pull Requests are welcome, but please open an issue (or comment in an existing
issue) to discuss any non-trivial changes before submitting code.

## License

BSD licensed. See the [LICENSE](LICENSE) file for details.
