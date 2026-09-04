package blockfrost_test

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestAssetTransactionsQueryParams(t *testing.T) {
	const asset = "asset"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != fmt.Sprintf("/assets/%s/transactions", asset) {
			t.Errorf("unexpected path: %s", r.URL.Path)
		}
		if got := r.URL.Query().Get("from"); got != "8929261" {
			t.Errorf("unexpected from query parameter: %s", got)
		}
		if got := r.URL.Query().Get("to"); got != "9999269:10" {
			t.Errorf("unexpected to query parameter: %s", got)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("[]"))
	}))
	defer server.Close()

	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{Server: server.URL})
	if _, err := api.AssetTransactions(context.Background(), asset, blockfrost.APIQueryParams{
		From: "8929261",
		To:   "9999269:10",
	}); err != nil {
		t.Fatal(err)
	}
}

func TestAssetUnmarshal(t *testing.T) {

	var onchainMetadata interface{} = map[string]interface{}{
		"image":       "ipfs://ipfs/QmfKyJ4tuvHowwKQCbCHj4L5T3fSj8cjs7Aau8V7BWv226",
		"name":        "My NFT token",
		"description": "optional",
	}

	fp := filepath.Join(testdata, "json", "assets", "asset.json")
	want := blockfrost.Asset{
		Asset:             "b0d07d45fe9514f80213f4020e5a61241458be626841cde717cb38a76e7574636f696e",
		AssetName:         "6e7574636f696e",
		PolicyId:          "b0d07d45fe9514f80213f4020e5a61241458be626841cde717cb38a7",
		Fingerprint:       "asset1pkpwyknlvul7az0xx8czhl60pyel45rpje4z8w",
		InitialMintTxHash: "6804edf9712d2b619edb6ac86861fe93a730693183a262b165fcc1ba1bc99cad",
		MintOrBurnCount:   1,
		Quantity:          "12000",
		OnchainMetadata:   &onchainMetadata,
		Metadata: &blockfrost.AssetMetadata{
			Name:        "nutcoin",
			Description: "The Nut Coin",
			Ticker:      "nutc",
			Decimals:    6,
			URL:         "https://www.stakenuts.com/",
			Logo:        "42",
		},
	}
	got := blockfrost.Asset{}
	testStructGotWant(t, fp, &got, &want)
}

func TestResourceAssetsIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.Assets(context.TODO(), blockfrost.APIQueryParams{})
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")
	want := []blockfrost.AssetByPolicy{}
	testIntUtil(t, fp, &got, &want)
}

func TestResourceAssetIntegration(t *testing.T) {
	asset := "3a9241cd79895e3a8d65261b40077d4437ce71e9d7c8c6c00e3f658e4669727374636f696e"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.Asset(context.TODO(), asset)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")
	want := blockfrost.Asset{}
	testIntUtil(t, fp, &got, &want)
}
func TestResourceAssetMetadataIntegration(t *testing.T) {
	asset := "14696a4676909f4e3cb1f2e60e2e08e5abed70caf5c02699be97113943554259"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.Asset(context.TODO(), asset)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")
	want := blockfrost.Asset{}
	testIntUtil(t, fp, &got, &want)
}

func TestResourceAssetHistoryIntegration(t *testing.T) {
	asset := "3a9241cd79895e3a8d65261b40077d4437ce71e9d7c8c6c00e3f658e4669727374636f696e"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.AssetHistory(context.TODO(), asset, blockfrost.APIQueryParams{})
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")
	want := []blockfrost.AssetHistory{}
	testIntUtil(t, fp, &got, &want)
}

func TestResourceAssetTransactionIntegration(t *testing.T) {
	asset := "3a9241cd79895e3a8d65261b40077d4437ce71e9d7c8c6c00e3f658e4669727374636f696e"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.AssetTransactions(context.TODO(), asset, blockfrost.APIQueryParams{})
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")
	want := []blockfrost.AssetTransaction{}
	testIntUtil(t, fp, &got, &want)
}

func TestResourceAssetddressesIntegration(t *testing.T) {
	asset := "3a9241cd79895e3a8d65261b40077d4437ce71e9d7c8c6c00e3f658e4669727374636f696e"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.AssetAddresses(context.TODO(), asset, blockfrost.APIQueryParams{})
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")
	want := []blockfrost.AssetAddress{}
	testIntUtil(t, fp, &got, &want)
}

func TestResourceAssetsByPolicyIntegration(t *testing.T) {
	policy := "3a9241cd79895e3a8d65261b40077d4437ce71e9d7c8c6c00e3f658e"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.AssetsByPolicy(context.TODO(), policy, blockfrost.APIQueryParams{})
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")

	want := []blockfrost.AssetByPolicy{}
	testIntUtil(t, fp, &got, &want)
}

type withCustomTest func() func(got interface{}, want interface{})

func testIntUtil(t *testing.T, fp string, got interface{}, want interface{}, opts ...withCustomTest) {
	if *update {
		data, err := json.Marshal(got)
		if err != nil {
			t.Fatal(err)
		}
		WriteGoldenFile(t, fp, data)
	}
	bytes := ReadOrGenerateGoldenFile(t, fp, got)
	if err := json.Unmarshal(bytes, want); err != nil {
		t.Fatal(err)
	}

	for _, opt := range opts {
		opt()(&got, &want)
	}

	if !reflect.DeepEqual(got, want) {
		gotJSON, err := json.Marshal(got)
		wantJSON, err := json.Marshal(want)
		if err != nil {
			t.Fatalf("\nexpected %v \ngot %v", want, got)
		}

		t.Fatalf("\nexpected %s \ngot %s", string(wantJSON), string(gotJSON))

	}
}
