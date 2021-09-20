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

	scripts, err := api.Scripts(context.TODO(), blockfrost.APIPagingParams{})
	if err != nil {
		log.Fatal(err)
	}
	if len(scripts) == 0 {
		return
	}

	addr := scripts[0].ScriptHash

	script, err := api.Script(context.TODO(), addr)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Script: %+v", script)

}