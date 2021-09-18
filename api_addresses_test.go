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
	"reflect"
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
	want := blockfrost.Address{
		Address:      "addr1qxqs59lphg8g6qndelq8xwqn60ag3aeyfcp33c2kdp46a09re5df3pzwwmyq946axfcejy5n4x0y99wqpgtp2gd0k09qsgy6pz",
		Amount:       []blockfrost.AddressAmount{{Unit: "lovelace", Quantity: "0"}},
		StakeAddress: "stake1ux3u6x5cs388djqz6awnyuvez2f6n8jzjhqq59s4yxhm8jskeh0t9",
		Type:         "shelley",
		Script:       false,
	}
	fp := filepath.Join(testdata, "json", "address", "address.json")
	file, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatalf("an error ocurred while trying to read json test file %s", fp)
	}
	got := blockfrost.Address{}
	if err = json.Unmarshal([]byte(file), &got); err != nil {
		t.Fatalf("failed to unmarshal %s with err %v", fp, err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v got %v", want, got)
	}

}

func TestAddressDetailsMarshall(t *testing.T) {
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
	file, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatalf("an error ocurred while trying to read json test file %s", fp)
	}
	got := blockfrost.AddressDetails{}
	if err = json.Unmarshal([]byte(file), &got); err != nil {
		t.Fatalf("failed to unmarshal %s with err %v", fp, err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected  %v got %v", want, got)
	}
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
	b, err := ioutil.ReadFile(path)
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
	api, err := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})
	if err != nil {
		t.Fatal(err)
	}
	got, err := api.Address(context.TODO(), addr)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")
	if *update {
		data, err := json.Marshal(got)
		if err != nil {
			t.Fatal(err)
		}
		WriteGoldenFile(t, fp, data)
	}
	bytes := ReadOrGenerateGoldenFile(t, fp, got)
	want := blockfrost.Address{}
	if err = json.Unmarshal(bytes, &want); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v got %v", want, got)
	}
}

func TestResourceAddressDetails(t *testing.T) {
	addr := "addr1qxqs59lphg8g6qndelq8xwqn60ag3aeyfcp33c2kdp46a09re5df3pzwwmyq946axfcejy5n4x0y99wqpgtp2gd0k09qsgy6pz"
	api, err := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})
	if err != nil {
		t.Fatal(err)
	}
	got, err := api.AddressDetails(context.TODO(), addr)
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")
	if *update {
		data, err := json.Marshal(got)
		if err != nil {
			t.Fatal(err)
		}
		WriteGoldenFile(t, fp, data)
	}
	bytes := ReadOrGenerateGoldenFile(t, fp, got)
	want := blockfrost.AddressDetails{}
	if err = json.Unmarshal(bytes, &want); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v got %v", want, got)
	}

	t.Logf("Got: %v", got)
	t.Logf("Want: %v", want)
}
