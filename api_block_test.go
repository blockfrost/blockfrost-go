package blockfrost_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestResourceBlockLatest(t *testing.T) {
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)
	got, err := api.BlockLatest(context.TODO())
	testErrorHelper(t, err)
	nullGot := blockfrost.Block{}
	if reflect.DeepEqual(got, nullGot) {
		t.Fatal("got null struct")
	}
}

func TestResourceBlock(t *testing.T) {
	hash := "b2466aac25c5d620fd99ae6e0faa72bec3acb9aeb29961363f464dab51b5ac05"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)
	got, err := api.Block(context.TODO(), hash)
	testErrorHelper(t, err)
	nullGot := blockfrost.Block{}
	if reflect.DeepEqual(got, nullGot) {
		t.Fatal("got null struct")
	}
}

func TestResourceBlockBySlot(t *testing.T) {
	slot := 30895909
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)
	got, err := api.BlockBySlot(context.TODO(), slot)
	testErrorHelper(t, err)
	nullGot := blockfrost.Block{}
	if reflect.DeepEqual(got, nullGot) {
		t.Fatal("got null struct")
	}
}

func TestResourceBlockBySlotAndEpoch(t *testing.T) {
	slot := 106397
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)
	got, err := api.BlocksBySlotAndEpoch(context.TODO(), slot, 297)
	testErrorHelper(t, err)
	nullGot := blockfrost.Block{}
	if reflect.DeepEqual(got, nullGot) {
		t.Fatal("got null struct")
	}
}

func TestBlockLatestTransactionsIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)
	got, err := api.BlockLatestTransactions(context.TODO())
	testErrorHelper(t, err)
	nullGot := []blockfrost.Transaction{}
	if reflect.DeepEqual(got, nullGot) {
		t.Fatal("got null struct")
	}
}

func TestBlockTransactions(t *testing.T) {
	block := "4873401"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)
	got, err := api.BlockTransactions(context.TODO(), block)
	testErrorHelper(t, err)
	nullGot := []blockfrost.Transaction{}
	if reflect.DeepEqual(got, nullGot) {
		t.Fatal("got null struct")
	}
}

func TestBlockNextIntegration(t *testing.T) {
	hash := "5ea1ba291e8eef538635a53e59fddba7810d1679631cc3aed7c8e6c4091a516a"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)
	got, err := api.BlocksNext(context.TODO(), hash)
	testErrorHelper(t, err)
	nullGot := []blockfrost.Block{}
	if reflect.DeepEqual(got, nullGot) {
		t.Fatal("got null struct")
	}
}

func TestBlockPreviousIntegration(t *testing.T) {
	hash := "5ea1ba291e8eef538635a53e59fddba7810d1679631cc3aed7c8e6c4091a516a"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)
	got, err := api.BlocksPrevious(context.TODO(), hash)
	testErrorHelper(t, err)
	nullGot := []blockfrost.Block{}
	if reflect.DeepEqual(got, nullGot) {
		t.Fatal("got null struct")
	}
}
