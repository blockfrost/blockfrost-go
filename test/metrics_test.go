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

func Test_Metrics(t *testing.T) {
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.Write([]byte(
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
	metrics, err := client.Metrics(context.TODO())
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, 2, len(metrics))
}

func Test_Metrics_Endpoints(t *testing.T) {
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.Write([]byte(
				`[
					{
					"time": 1612543814,
					"calls": 182,
					"endpoint": "block"
					},
					{
					"time": 1612543814,
					"calls": 42,
					"endpoint": "epoch"
					}
					]`,
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
	metrics, err := client.MetricsEndpoints(context.TODO())
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, 2, len(metrics))
	assert.Equal(t, "block", metrics[0].Endpoint)
	assert.Equal(t, "epoch", metrics[1].Endpoint)
}
