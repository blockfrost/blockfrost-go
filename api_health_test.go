package blockfrost_test

import (
	"context"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestResourceInfoIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	got, err := api.Info(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	if got.Url != "https://blockfrost.io/" {
		t.Fatalf("unexpected url: %s", got.Url)
	}
	if got.Version == "" {
		t.Fatal("got empty version")
	}
}

func TestResourceHealthIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	got, err := api.Health(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	want := blockfrost.Health{}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")

	testIntUtil(t, fp, &got, &want)
}

func TestResourceHealthClockIntegration(t *testing.T) {
	t.Parallel()

	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	got, err := api.HealthClock(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	nullGot := blockfrost.HealthClock{}

	if reflect.DeepEqual(got, nullGot) {
		t.Fatalf("got null healthclock, %+v", got)
	}

}
