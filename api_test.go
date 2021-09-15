package blockfrost_test

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

var api blockfrost.APIClient

func init() {
	c, err := blockfrost.NewAPIClient()
	if err != nil {
		log.Fatal("Failed to create client", err)
	}
	api = c
}

func TestResourceInfo(t *testing.T) {
	t.Parallel()
	info, err := api.Info(context.TODO())

	if err != nil {
		t.Fatal("Failed to fetch `/` with error,", err)
	}

	if info == (blockfrost.Info{}) {
		t.Fatalf("got nil info %+v", info)
	}
}

func TestResourceHealth(t *testing.T) {
	t.Parallel()

	health, err := api.Health(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("health: %v\n", health)
}
