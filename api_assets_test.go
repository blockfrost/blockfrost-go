package blockfrost_test

import (
	"path/filepath"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestAssetUnmarshal(t *testing.T) {
	fp := filepath.Join(testdata, "json", "assets", "asset.json")
	want := blockfrost.Asset{
		Asset:             "b0d07d45fe9514f80213f4020e5a61241458be626841cde717cb38a76e7574636f696e",
		AssetName:         "6e7574636f696e",
		PolicyId:          "b0d07d45fe9514f80213f4020e5a61241458be626841cde717cb38a7",
		Fingerprint:       "asset1pkpwyknlvul7az0xx8czhl60pyel45rpje4z8w",
		InitialMintTxHash: "6804edf9712d2b619edb6ac86861fe93a730693183a262b165fcc1ba1bc99cad",
		MintOrBurnCount:   1,
		Quantity:          "12000",
		OnchainMetadata: blockfrost.AssetOnchainMetadata{
			Image: "ipfs://ipfs/QmfKyJ4tuvHowwKQCbCHj4L5T3fSj8cjs7Aau8V7BWv226",
			Name:  "My NFT token",
		},
		Metadata: blockfrost.AssetMetadata{
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
