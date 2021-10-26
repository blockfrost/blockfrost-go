package blockfrost_test

import (
	"context"
	"path/filepath"
	"strings"
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

func TestTransactionIntegration(t *testing.T) {
	hash := "6e5f825c82c1c6d6b77f2a14092f3b78c8f1b66db6f4cf8caec1555b6f967b3b"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.Transaction(context.TODO(), hash)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	want := blockfrost.TransactionContent{}
	testIntUtil(t, fp, &got, &want)
}

func TestTransactionUTXOs(t *testing.T) {
	hash := "6d619f41ba2e11b78c0d5647fb71ee5df05622fda1748a9124446226e54e6b50"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.TransactionUTXOs(context.TODO(), hash)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	want := blockfrost.TransactionUTXOs{}
	testIntUtil(t, fp, &got, &want)
}

func TestTransactionStakeAddressCertsIntegration(t *testing.T) {
	hash := "6e5f825c82c1c6d6b77f2a14092f3b78c8f1b66db6f4cf8caec1555b6f967b3b"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})
	got, err := api.TransactionStakeAddressCerts(context.TODO(), hash)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	want := []blockfrost.TransactionStakeAddressCert{}
	testIntUtil(t, fp, &got, &want)
}

func TestTransactionWithdrawlsIntegration(t *testing.T) {
	hash := "6d619f41ba2e11b78c0d5647fb71ee5df05622fda1748a9124446226e54e6b50"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.TransactionWithdrawals(context.TODO(), hash)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	want := []blockfrost.TransactionWidthrawal{}
	testIntUtil(t, fp, &got, &want)
}

func TestTransactionMIRsIntegration(t *testing.T) {
	hash := "6d619f41ba2e11b78c0d5647fb71ee5df05622fda1748a9124446226e54e6b50"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.TransactionMIRs(context.TODO(), hash)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	want := []blockfrost.TransactionMIR{}
	testIntUtil(t, fp, &got, &want)
}

func TestTransactionMetadata(t *testing.T) {
	hash := "6d619f41ba2e11b78c0d5647fb71ee5df05622fda1748a9124446226e54e6b50"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.TransactionMetadata(context.TODO(), hash)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	want := []blockfrost.TransactionMetadata{}
	testIntUtil(t, fp, &got, &want)
}

func TestTransactionMetadataInCBORsIntegration(t *testing.T) {
	hash := "6d619f41ba2e11b78c0d5647fb71ee5df05622fda1748a9124446226e54e6b50"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.TransactionMetadataInCBORs(context.TODO(), hash)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	want := []blockfrost.TransactionMetadataCbor{}
	testIntUtil(t, fp, &got, &want)
}

func TestTransactionRedeemersIntegration(t *testing.T) {
	hash := "6d619f41ba2e11b78c0d5647fb71ee5df05622fda1748a9124446226e54e6b50"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})
	got, err := api.TransactionRedeemers(context.TODO(), hash)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	want := []blockfrost.TransactionRedeemer{}
	testIntUtil(t, fp, &got, &want)
}

func TestTransactionDelegationCertsIntegration(t *testing.T) {
	hash := "6d619f41ba2e11b78c0d5647fb71ee5df05622fda1748a9124446226e54e6b50"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})
	got, err := api.TransactionDelegationCerts(context.TODO(), hash)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	want := []blockfrost.TransactionDelegation{}
	testIntUtil(t, fp, &got, &want)
}

func TestTransactionPoolUpdatesIntegration(t *testing.T) {
	hash := "6d619f41ba2e11b78c0d5647fb71ee5df05622fda1748a9124446226e54e6b50"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.TransactionPoolUpdates(context.TODO(), hash)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	want := []blockfrost.TransactionPoolCert{}
	testIntUtil(t, fp, &got, &want)
}

func TestTransactionPoolUpdateCertsIntegration(t *testing.T) {
	hash := "6d619f41ba2e11b78c0d5647fb71ee5df05622fda1748a9124446226e54e6b50"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.TransactionPoolUpdateCerts(context.TODO(), hash)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	want := []blockfrost.TransactionPoolCert{}
	testIntUtil(t, fp, &got, &want)
}

func TestTransactionPoolRetirementsIntegration(t *testing.T) {
	hash := "6d619f41ba2e11b78c0d5647fb71ee5df05622fda1748a9124446226e54e6b50"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})
	got, err := api.TransactionPoolRetirementCerts(context.TODO(), hash)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	want := []blockfrost.TransactionPoolCert{}
	testIntUtil(t, fp, &got, &want)
}
