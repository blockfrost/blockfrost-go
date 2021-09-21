package blockfrost_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/blockfrost/blockfrost-go"
	"github.com/stretchr/testify/assert"
)

func TestResourceMetadataTxLabels(t *testing.T) {
	t.Parallel()

	expectedElement := blockfrost.MetadataTxLabel{
		Label: "1967",
		Cip10: "nut.link metadata oracles registry",
		Count: "3",
	}
	expectedElementNull := blockfrost.MetadataTxLabel{
		Label: "1990",
		Count: "1",
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			fp := filepath.Join(testdata, "json", "metadata", "metadata.json")
			file, err := ioutil.ReadFile(fp)
			if err != nil {
				t.Fatalf("an error ocurred while trying to read json test file %s", fp)
			}
			w.Write(file)
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{
		Count: 1,
	}
	accountArray, err := api.MetadataTxLabels(context.TODO(), q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElementNull, accountArray[0])
	assert.Equal(t, expectedElement, accountArray[1])

}

func TestResourceMetadataTxContentInJSON(t *testing.T) {
	t.Parallel()

	inputLabel := "1990"
	expectedElement := blockfrost.MetadataTxContentInJSON{
		TxHash: "257d75c8ddb0434e9b63e29ebb6241add2b835a307aa33aedba2effe09ed4ec8",
		JSONMetadata: map[string]interface{}{
			"ADAUSD": []interface{}{
				map[string]interface{}{
					"value":  "0.10409800535729975",
					"source": "ergoOracles",
				},
			},
		},
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			fp := filepath.Join(testdata, "json", "metadata", "metadata_in_json.json")
			file, err := ioutil.ReadFile(fp)
			if err != nil {
				t.Fatalf("an error ocurred while trying to read json test file %s", fp)
			}
			w.Write(file)

		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{
		Count: 1,
	}
	accountArray, err := api.MetadataTxContentInJSON(context.TODO(), inputLabel, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])

}

func TestResourceMetadataTxContentInCBOR(t *testing.T) {
	t.Parallel()

	inputLabel := "1990"
	expectedElement := blockfrost.MetadataTxContentInCBOR{
		TxHash:       "4237501da3cfdd53ade91e8911e764bd0699d88fd43b12f44a1f459b89bc91be",
		CborMetadata: "\\xa100a16b436f6d62696e6174696f6e8601010101010c",
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			fp := filepath.Join(testdata, "json", "metadata", "metadata_in_cbor.json")
			file, err := ioutil.ReadFile(fp)
			if err != nil {
				t.Fatalf("an error ocurred while trying to read json test file %s", fp)
			}
			w.Write(file)

		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{
		Count: 1,
	}
	accountArray, err := api.MetadataTxContentInCBOR(context.TODO(), inputLabel, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[2])

}
