package blockfrost_test

import (
	"context"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestMempoolUnmarshal(t *testing.T) {
	want := []blockfrost.Mempool{
		{TxHash: "abc"},
		{TxHash: "def"},
	}
	fp := filepath.Join(testdata, "json", "mempool", "mempool.json")
	got := []blockfrost.Mempool{}
	testStructGotWant(t, fp, &got, &want)
}

func TestMempoolTransactionContentUnmarshal(t *testing.T) {
	want := blockfrost.MempoolTransactionContent{
		Tx: blockfrost.MempoolTransaction{
			Hash: "1f96d3824eb2aeeb0b09b99748bb70ac681e0cae6e37e01c43958b79ca69c986",
			OutputAmount: []struct {
				Quantity string `json:"quantity"`

				// The unit of the value
				Unit string `json:"unit"`
			}{
				{
					Unit:     "lovelace",
					Quantity: "2837715",
				},
			},
			Fees:             "369133",
			Deposit:          "0",
			Size:             4683,
			InvalidBefore:    "",
			InvalidHereafter: "109798439",
			UtxoCount:        2,
			ValidContract:    true,
		},
		Inputs: []blockfrost.MempoolTransactionInput{{
			Address:     "addr1vx9wkegx062xmmdzfd69jz6dt48p5mse4v35mml5h6ceznq8ap8fz",
			TxHash:      "7aa461f4f924586864c74141d457e70cfb26b2b5b9cfea4c5d5580f037ef41da",
			OutputIndex: 0,
			Collateral:  false,
			Reference:   false,
		}},
		Outputs: []blockfrost.MempoolTransactionOutput{
			{
				Address: "addr1vx9wkegx062xmmdzfd69jz6dt48p5mse4v35mml5h6ceznq8ap8fz",
				Amount: []blockfrost.TxAmount{{
					Unit:     "lovelace",
					Quantity: "2837715",
				}, {
					Unit:     "6787a47e9f73efe4002d763337140da27afa8eb9a39413d2c39d4286524144546f6b656e73",
					Quantity: "15000",
				}},
				OutputIndex: 0,
				// DataHash:            nil,
				// InlineDatum:         nil,
				// ReferenceScriptHash: nil,
				DataHash:            "",
				InlineDatum:         "",
				ReferenceScriptHash: "",
				Collateral:          false,
			}},
		Redeemers: nil,
	}
	fp := filepath.Join(testdata, "json", "mempool", "transaction.json")
	got := blockfrost.MempoolTransactionContent{}
	testStructGotWant(t, fp, &got, &want)
}

func TestMempoolIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.Mempool(context.TODO(), blockfrost.APIQueryParams{})
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Got response: %+v", got)

	var want []blockfrost.Mempool
	if reflect.TypeOf(got) != reflect.TypeOf(want) {
		t.Fatalf("Expected type []blockfrost.Mempool, got type %T", got)
	}
}

// MempoolTx and MempoolByAddress need mocked server to provide static mempool transaction
