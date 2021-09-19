package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blockfrost/blockfrost-go"
)

func main() {
	api, err := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})
	if err != nil {
		log.Fatal(err)
	}

	gen, err := api.Genesis(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Genesis Block: \n \t%+v", gen)
}
