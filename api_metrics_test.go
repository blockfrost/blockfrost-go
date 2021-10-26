package blockfrost_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestResourceMetrics(t *testing.T) {
	t.Parallel()
	metric := blockfrost.Metric{Time: 1612543884, Calls: 42}
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(
			[]byte(
				`[
					{
					"time": 1612543884,
					"calls": 42
					},
					{
					"time": 1614523884,
					"calls": 6942
					}
				]`,
			))
	}))

	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)

	metrics, err := api.Metrics(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	if mlen := len(metrics); mlen != 2 {
		t.Fatalf("Expected metrics to be of len %v got %v", 2, mlen)
	}

	if metrics[0] != metric {
		t.Fatalf("Expected metrics[0] to be %v got %v", metric, metrics)
	}
}

func TestResourceMetricsEndpoints(t *testing.T) {
	t.Parallel()
	metricEndpoint := blockfrost.MetricsEndpoint{Time: 1612543814, Calls: 182, Endpoint: "block"}
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(
			[]byte(
				`[
					{
					  "time": 1612543814,
					  "calls": 182,
					  "endpoint": "block"
					},
					{
					  "time": 1612553884,
					  "calls": 89794,
					  "endpoint": "block"
					}
				  ]`,
			))
	}))
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)

	metricsEndpoints, err := api.MetricsEndpoints(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	if mlen := len(metricsEndpoints); mlen != 2 {
		t.Fatalf("Expected metricsEnpoints to be of len %v got %v", 2, mlen)
	}

	if metricsEndpoints[0] != metricEndpoint {
		t.Fatalf("Expected metricsEndpoints[0] to be %v, got %v", metricEndpoint, metricsEndpoints[0])
	}
}
