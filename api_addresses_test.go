package blockfrost_test

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

var (
	update   = flag.Bool("update", false, "update .golden files")
	generate = flag.Bool("gen", false, "generate .golden files")
)

const testdata = "testdata"

func TestAddressUnMarshall(t *testing.T) {
	stakeAddress := "stake1ux3u6x5cs388djqz6awnyuvez2f6n8jzjhqq59s4yxhm8jskeh0t9"

	want := blockfrost.Address{
		Address:      "addr1qxqs59lphg8g6qndelq8xwqn60ag3aeyfcp33c2kdp46a09re5df3pzwwmyq946axfcejy5n4x0y99wqpgtp2gd0k09qsgy6pz",
		Amount:       []blockfrost.AddressAmount{{Unit: "lovelace", Quantity: "0"}},
		StakeAddress: &stakeAddress,
		Type:         "shelley",
		Script:       false,
	}
	fp := filepath.Join(testdata, "json", "address", "address.json")
	got := blockfrost.Address{}
	testStructGotWant(t, fp, &got, &want)
}

func TestAddressDetailsUnMarshall(t *testing.T) {
	want := blockfrost.AddressDetails{
		Address: "addr1qxqs59lphg8g6qndelq8xwqn60ag3aeyfcp33c2kdp46a09re5df3pzwwmyq946axfcejy5n4x0y99wqpgtp2gd0k09qsgy6pz",
		ReceivedSum: []blockfrost.AddressAmount{
			{Unit: "lovelace", Quantity: "3324671000"},
		},
		SentSum: []blockfrost.AddressAmount{
			{Unit: "lovelace", Quantity: "3324671000"},
		},
		TxCount: 2,
	}

	fp := filepath.Join(testdata, "json", "address", "address_details.json")
	got := blockfrost.AddressDetails{}
	testStructGotWant(t, fp, &got, &want)
}

func WriteGoldenFile(t *testing.T, path string, bytes []byte) {
	t.Helper()
	err := os.MkdirAll(filepath.Dir(path), 0777)
	if err != nil {
		t.Fatal(err)
	}
	err = ioutil.WriteFile(path, bytes, 0666)
	if err != nil {
		t.Fatal(err)
	}
}

func ReadOrGenerateGoldenFile(t *testing.T, path string, v interface{}) []byte {
	t.Helper()
	b, err := os.ReadFile(path)
	switch {
	case errors.Is(err, os.ErrNotExist):
		if *generate {
			var buf bytes.Buffer
			enc := json.NewEncoder(&buf)
			enc.SetIndent("", "  ")
			err := enc.Encode(v)
			if err != nil {
				t.Fatal("Failed to marshal to json")
			}
			WriteGoldenFile(t, path, buf.Bytes())
			return buf.Bytes()
		}
		t.Fatalf("Missing golden file. Run `go test -args -gen` to generate it.")
	case err != nil:
		t.Fatal(err)
	}
	return b
}

func TestResourceAddress(t *testing.T) {
	addr := "addr1qxqs59lphg8g6qndelq8xwqn60ag3aeyfcp33c2kdp46a09re5df3pzwwmyq946axfcejy5n4x0y99wqpgtp2gd0k09qsgy6pz"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.Address(context.TODO(), addr)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")

	want := blockfrost.Address{}
	testIntUtil(t, fp, &got, &want)
}

func TestResourceAddressDetails(t *testing.T) {
	addr := "addr1qxqs59lphg8g6qndelq8xwqn60ag3aeyfcp33c2kdp46a09re5df3pzwwmyq946axfcejy5n4x0y99wqpgtp2gd0k09qsgy6pz"
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.AddressDetails(context.TODO(), addr)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")
	want := blockfrost.AddressDetails{}

	testIntUtil(t, fp, &got, &want)
}

func TestResourceAddressTransactions(t *testing.T) {
	addr := "addr1qxqs59lphg8g6qndelq8xwqn60ag3aeyfcp33c2kdp46a09re5df3pzwwmyq946axfcejy5n4x0y99wqpgtp2gd0k09qsgy6pz"

	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.AddressTransactions(
		context.TODO(),
		addr,
		blockfrost.APIQueryParams{},
	)
	if err != nil {
		t.Fatal(err)
	}

	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")
	var want []blockfrost.AddressTransactions
	testIntUtil(t, fp, &got, &want)
}

func TestAddressUTXOs(t *testing.T) {
	addr := "addr1qxqs59lphg8g6qndelq8xwqn60ag3aeyfcp33c2kdp46a09re5df3pzwwmyq946axfcejy5n4x0y99wqpgtp2gd0k09qsgy6pz"

	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.AddressUTXOs(
		context.TODO(),
		addr,
		blockfrost.APIQueryParams{},
	)
	if err != nil {
		t.Fatal(err)
	}

	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")
	var want []blockfrost.AddressUTXO
	testIntUtil(t, fp, &got, &want)
}

func TestAddressUTXOsAsset(t *testing.T) {
	addr := "addr1q8zsjx7vxkl4esfejafhxthyew8c54c9ch95gkv3nz37sxrc9ty742qncmffaesxqarvqjmxmy36d9aht2duhmhvekgq3jd3w2"
	asset := "d436d9f6b754582f798fe33f4bed12133d47493f78b944b9cc55fd1853756d6d69744c6f64676534393539"

	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.AddressUTXOsAsset(
		context.TODO(),
		addr,
		asset,
		blockfrost.APIQueryParams{},
	)
	if err != nil {
		t.Fatal(err)
	}

	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")
	var want []blockfrost.AddressUTXO
	testIntUtil(t, fp, &got, &want)
}
