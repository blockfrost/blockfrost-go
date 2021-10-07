<p align="center">
  <a href="https://blockfrost.io" target="_blank" align="center">
    <img src="https://blockfrost.io/images/logo.svg" width="280">
  </a>
  <br />
</p>

# Official Blockfrost SDK Client

[![Build status](https://github.com/blockfrost/blockfrost-go/actions/workflows/test.yml/badge.svg?branch=staging)](https://github.com/blockfrost/blockfrost-go/actions/workflows/test.yml)
[![Go report card](https://goreportcard.com/badge/github.com/blockfrost/blockfrost-go)](https://goreportcard.com/report/github.com/blockfrost/blockfrost-go)


# Installation

`blockfrost-go` can be installed through go get

```console
$ go get https://github.com/blockfrost/blockfrost-go
```

## Usage

### Cardano API
```go
package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blockfrost/blockfrost-go"
)

func main() {
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{
            ProjectID: "YOUR_PROJECT_ID_HERE", // Exclude to load from env:BLOCKFROST_PROJECT_ID
        },
	)
	if err != nil {
		log.Fatal(err)
	}

	info, err := api.Info(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("API Info:\n\tUrl: %s\n\tVersion: %s", info.Url, info.Version)
}
```

## License

Licensed under the [Apache License 2.0](https://opensource.org/licenses/Apache-2.0), see [`LICENSE`](https://github.com/blockfrost/blockfrost-go/blob/master/LICENSE)
