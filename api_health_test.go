package blockfrost_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestResourceInfo(t *testing.T) {
	t.Parallel()
	sUrl := "https://blockfrost.io/"
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(
			[]byte(
				`{
					"url": "https://blockfrost.io/",
					"version": "0.1.0"
				  }`,
			))
	}))

	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)

	info, err := api.Info(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	if info.Url != sUrl {
		t.Fatalf("Expected info.Url to be %s got %s", sUrl, info.Url)
	}
}

func TestResourceHealth(t *testing.T) {
	t.Parallel()
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(
			[]byte(
				`{
					"is_healthy": true
				  }`,
			))
	}))

	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)

	health, err := api.Health(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	if health.IsHealthy != true {
		t.Fatalf("Expected health.IsHealthy to be %t, got %t", true, health.IsHealthy)
	}
}

func TestResourceHealthClock(t *testing.T) {
	t.Parallel()
	serverTime := 1603400958947
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(
			[]byte(
				`{
					"server_time": 1603400958947
				  }`,
			))
	}))

	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)

	healthClock, err := api.HealthClock(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	if healthClock.ServerTime != int64(serverTime) {
		t.Fatalf("Expected healthClock.ServerTime to be %v, got %v", serverTime, healthClock.ServerTime)
	}
}
