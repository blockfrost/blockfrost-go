package main

import (
	"context"
	"fmt"
	"log"
	"os"

	blockfrostgo "github.com/blockfrost/blockfrost-go/pkg/api"
)

func main() {

	client := blockfrostgo.NewBlockfrostAPI(
		os.Getenv("API_KEY"),
		blockfrostgo.CardanoMainnet,
		false,
		nil,
		os.Stdout,
	)

	appinfo, err := client.Info(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(appinfo)

}
