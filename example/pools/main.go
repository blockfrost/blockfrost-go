package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blockfrost/blockfrost-go"
)

func main() {
	poolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"

	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{}, // Add ProjectID or exclude to load from env:BLOCKFROST_PROJECT_ID
	)
	if err != nil {
		log.Fatal(err)
	}
	q := blockfrost.APIQueryParams{}

	pool, err := api.PoolHistory(context.TODO(), poolID, q)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v", pool)
}
