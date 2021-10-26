package blockfrost_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestResourceMetadataTxLabels(t *testing.T) {
	t.Parallel()

	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{
		Count: 1,
	}
	got, err := api.MetadataTxLabels(context.TODO(), q)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if reflect.DeepEqual(got, []blockfrost.MetadataTxLabel{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourceMetadataTxContentInJSON(t *testing.T) {
	t.Parallel()

	inputLabel := "1990"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{
		Count: 1,
	}
	got, err := api.MetadataTxContentInJSON(context.TODO(), inputLabel, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if reflect.DeepEqual(got, []blockfrost.MetadataTxContentInJSON{}) {
		t.Fatalf("got null %+v", got)
	}

}

func TestResourceMetadataTxContentInCBOR(t *testing.T) {
	t.Parallel()

	inputLabel := "1990"

	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{
		Count: 1,
	}
	got, err := api.MetadataTxContentInCBOR(context.TODO(), inputLabel, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if reflect.DeepEqual(got, []blockfrost.MetadataTxContentInCBOR{}) {
		t.Fatalf("got null %+v", got)
	}

}
