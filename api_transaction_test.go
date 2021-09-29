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

func TestTransactionStakeAddressCertUnmarshall(t *testing.T) {
	fp := filepath.Join(testdata, "json", "transactions", "tx_stakeaddr_cert.json")
	want := []blockfrost.TransactionStakeAddressCert{
		{
			Address:      "stake1u9t3a0tcwune5xrnfjg4q7cpvjlgx9lcv0cuqf5mhfjwrvcwrulda",
			CertIndex:    0,
			Registration: true,
		},
	}
	got := []blockfrost.TransactionStakeAddressCert{}
	testStructGotWant(t, fp, &got, &want)
}

func TestTransactionDelegationUnmarshal(t *testing.T) {
	fp := filepath.Join(testdata, "json", "transactions", "tx_delegations.json")
	want := []blockfrost.TransactionDelegation{
		{
			Index:       0,
			CertIndex:   0,
			Address:     "stake1u9r76ypf5fskppa0cmttas05cgcswrttn6jrq4yd7jpdnvc7gt0yc",
			PoolId:      "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy",
			ActiveEpoch: 210,
		},
	}
	got := []blockfrost.TransactionDelegation{}
	testStructGotWant(t, fp, &got, &want)
}

func TestTransactionWithdrawalsUnmarshal(t *testing.T) {
	fp := filepath.Join(testdata, "json", "transactions", "tx_withdrawals.json")
	want := []blockfrost.TransactionWidthrawal{
		{
			Address: "stake1u9r76ypf5fskppa0cmttas05cgcswrttn6jrq4yd7jpdnvc7gt0yc",
			Amount:  "431833601",
		},
	}
	got := []blockfrost.TransactionWidthrawal{}
	testStructGotWant(t, fp, &got, &want)
}

func TestTransactionMIRsUnmarshal(t *testing.T) {
	fp := filepath.Join(testdata, "json", "transactions", "tx_mirs.json")
	want := []blockfrost.TransactionMIR{
		{
			Pot:       "reserve",
			CertIndex: 0,
			Address:   "stake1u9r76ypf5fskppa0cmttas05cgcswrttn6jrq4yd7jpdnvc7gt0yc",
			Amount:    "431833601",
		},
	}
	got := []blockfrost.TransactionMIR{}
	testStructGotWant(t, fp, &got, &want)
}

func TestTransactionMetadataCborUnmarshal(t *testing.T) {
	fp := filepath.Join(testdata, "json", "transactions", "tx_cbor.json")
	want := []blockfrost.TransactionMetadataCbor{
		{
			Label:        "1968",
			CborMetadata: "\\xa100a16b436f6d62696e6174696f6e8601010101010c",
		},
	}
	got := []blockfrost.TransactionMetadataCbor{}
	testStructGotWant(t, fp, &got, &want)
}
