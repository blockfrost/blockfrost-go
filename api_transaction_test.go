package blockfrost_test

import (
	"path/filepath"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestTransactionContentUnmarshal(t *testing.T) {
	want := blockfrost.TransactionContent{
		Hash:        "1e043f100dce12d107f679685acd2fc0610e10f72a92d412794c9773d11d8477",
		Block:       "356b7d7dbb696ccd12775c016941057a9dc70898d87a63fc752271bb46856940",
		BlockHeight: 123456,
		Slot:        42000000,
		Index:       1,
		OutputAmount: []struct {
			Quantity string `json:"quantity"`

			// The unit of the value
			Unit string `json:"unit"`
		}{},
		Fees:             "182485",
		Deposit:          "0",
		Size:             433,
		InvalidBefore:    "",
		InvalidHereafter: "13885913",
		UtxoCount:        4,
	}
	fp := filepath.Join(testdata, "json", "transactions", "transaction.json")
	got := blockfrost.TransactionContent{}
	testStructGotWant(t, fp, &got, &want)
}
