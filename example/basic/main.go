package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blockfrost/blockfrost-go"
)

func main() {
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{}, // Add ProjectID or exclude to load from env:BLOCKFROST_PROJECT_ID
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
