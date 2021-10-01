package main

import (
	"context"
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/blockfrost/blockfrost-go"
)

var (
	fp = flag.String("file", "", "path to export csv")
)

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	api, err := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})
	handleError(err)

	file, err := os.Create(*fp)
	handleError(err)

	wr := csv.NewWriter(file)
	defer wr.Flush()

	err = wr.Write([]string{"batch_id", "script_hash"})
	handleError(err)

	for r := range api.ScriptsAll(context.TODO()) {
		for k, v := range r.Scripts {
			arr := []string{fmt.Sprintf("%d", k), v.ScriptHash}
			err := wr.Write(arr)
			handleError(err)
		}
	}
}
