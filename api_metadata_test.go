package blockfrost_test

import (
	"context"
	"net/http"
	"net/http/httptest"
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
			w.Write([]byte(
				`[
					{
					"label": "1990",
					"cip10": null,
					"count": "1"
					},
					{
					"label": "1967",
					"cip10": "nut.link metadata oracles registry",
					"count": "3"
					}
					]`,
			))
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
		JSONMetadata: blockfrost.JSONMetadata{
			Adausd: []blockfrost.Adausd{
				{
					Value:  "0.10409800535729975",
					Source: "ergoOracles",
				},
			},
		},
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.Write([]byte(
				`[
					{
						"tx_hash": "257d75c8ddb0434e9b63e29ebb6241add2b835a307aa33aedba2effe09ed4ec8",
						"json_metadata": {
							"ADAUSD": [
								{
									"value": "0.10409800535729975",
									"source": "ergoOracles"
								}
							]
						}
					},
					{
						"tx_hash": "e865f2cc01ca7381cf98dcdc4de07a5e8674b8ea16e6a18e3ed60c186fde2b9c",
						"json_metadata": {
							"ADAUSD": [
								{
									"value": "0.15409850555139935",
									"source": "ergoOracles"
								}
							]
						}
					}
				]`,
			))
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

func TestResourceMetadataTxContentInJSONRaw(t *testing.T) {
	t.Parallel()

	inputLabel := "1990"
	expectedElement := blockfrost.MetadataTxContentInJSONRaw{
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
			w.Write([]byte(
				`[
					{
						"tx_hash": "257d75c8ddb0434e9b63e29ebb6241add2b835a307aa33aedba2effe09ed4ec8",
						"json_metadata": {
							"ADAUSD": [
								{
									"value": "0.10409800535729975",
									"source": "ergoOracles"
								}
							]
						}
					},
					{
						"tx_hash": "e865f2cc01ca7381cf98dcdc4de07a5e8674b8ea16e6a18e3ed60c186fde2b9c",
						"json_metadata": {
							"ADAUSD": [
								{
									"value": "0.15409850555139935",
									"source": "ergoOracles"
								}
							]
						}
					}
				]`,
			))
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
	accountArray, err := api.MetadataTxContentInJSONRaw(context.TODO(), inputLabel, q)
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
			w.Write([]byte(
				`[
					{
						"tx_hash": "257d75c8ddb0434e9b63e29ebb6241add2b835a307aa33aedba2effe09ed4ec8",
						"cbor_metadata": null
					},
					{
						"tx_hash": "e865f2cc01ca7381cf98dcdc4de07a5e8674b8ea16e6a18e3ed60c186fde2b9c",
						"cbor_metadata": null
					},
					{
						"tx_hash": "4237501da3cfdd53ade91e8911e764bd0699d88fd43b12f44a1f459b89bc91be",
						"cbor_metadata": "\\xa100a16b436f6d62696e6174696f6e8601010101010c"
					}
				]`,
			))
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
