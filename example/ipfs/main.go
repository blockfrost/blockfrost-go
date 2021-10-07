package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"github.com/blockfrost/blockfrost-go"
)

var (
	fp = flag.String("file", "", "Path to file")
)

func main() {
	flag.Parse()
	ipfs, err := blockfrost.NewIPFSClient(blockfrost.IPFSClientOptions{})
	if err != nil {
		log.Fatal(err)
	}

	ipo, err := ipfs.Add(context.TODO(), *fp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("IPFS Object: %+v\n", ipo)

	pin, err := ipfs.Pin(context.TODO(), ipo.IPFSHash)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Pin: %+v", pin)
}
