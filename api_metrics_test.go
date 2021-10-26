package blockfrost_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestResourceMetricsIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	got, err := api.Metrics(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(got, []blockfrost.Metric{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourceMetricsEndpointsIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	got, err := api.MetricsEndpoints(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(got, []blockfrost.MetricsEndpoint{}) {
		t.Fatalf("got null %+v", got)
	}
}
