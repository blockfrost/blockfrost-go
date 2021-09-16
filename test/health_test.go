package blockfrostgo_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	blockfrostgo "github.com/blockfrost/blockfrost-go/pkg/api"

	"github.com/stretchr/testify/assert"
)

const (
	ApiKey = "apiKey"
)

func Test_Invalid_URL(t *testing.T) {

	client := blockfrostgo.NewBlockfrostAPI(
		ApiKey,
		"ht&@-tp://:aa",
		false,
		nil,
		nil,
	)
	actual, err := client.Info(context.Background())
	assert.Error(t, err)
	assert.Empty(t, actual)

}

func Test_Health_Info(t *testing.T) {
	expectedVersion := "0.1.0"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.Write([]byte(
				`{
					"url": "https://blockfrost.io/",
					"version": "0.1.0"
					}`,
			))
		}),
	)
	defer s.Close()
	client := blockfrostgo.NewBlockfrostAPI(
		ApiKey,
		s.URL,
		false,
		nil,
		nil,
	)
	appinfo, err := client.Info(context.TODO())
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedVersion, appinfo.Version)
}

func Test_Health(t *testing.T) {
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.Write([]byte(
				`{
					"is_healthy": true
				}`,
			))
		}),
	)
	defer s.Close()
	client := blockfrostgo.NewBlockfrostAPI(
		ApiKey,
		s.URL,
		false,
		nil,
		nil,
	)
	health, err := client.Health(context.TODO())
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, true, health.IsHealthy)
}

func Test_Health_Clock(t *testing.T) {
	expectedtTime := int64(1603400958947)
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.Write([]byte(
				`{
					"server_time": 1603400958947
					}`,
			))
		}),
	)
	defer s.Close()
	client := blockfrostgo.NewBlockfrostAPI(
		ApiKey,
		s.URL,
		false,
		nil,
		nil,
	)
	healthClock, err := client.HealthClock(context.TODO())
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedtTime, healthClock.ServerTime)
}
