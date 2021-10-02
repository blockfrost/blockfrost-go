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
	q := blockfrost.APIPagingParams{}
	m, err := api.MetadataTxLabels(context.TODO(), q)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", m)
}
