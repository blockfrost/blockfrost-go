package blockfrost

import (
	"context"
	"net/http"
	"os"
	"time"
)

type apiClient struct {
	server    string
	projectId string
	client    HttpRequestDoer
}

type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type ClientOption func(*apiClient) error

func WithProjectID(projectId string) ClientOption {
	return func(c *apiClient) error {
		c.projectId = projectId
		return nil
	}
}

func WithTestNet() ClientOption {
	return func(c *apiClient) error {
		c.server = CardanoTestNet
		return nil
	}
}

func WithHTTPServer(server string) ClientOption {
	return func(c *apiClient) error {
		c.server = server
		return nil
	}
}

func WithHTTPClient(doer HttpRequestDoer) ClientOption {
	return func(c *apiClient) error {
		c.client = doer
		return nil
	}
}
func NewAPIClient(opts ...ClientOption) (APIClient, error) {
	client := &apiClient{
		server: CardanoMainNet,
	}

	for _, opt := range opts {
		if err := opt(client); err != nil {
			return nil, err
		}
	}

	// TODO: Make configurable. Timeout, retries ...
	client.client = &http.Client{Timeout: time.Second * 10}

	if client.projectId == "" {
		client.projectId = os.Getenv("BLOCKFROST_PROJECT_ID")
	}
	return client, nil
}

type APIClient interface {
	Info(ctx context.Context) (Info, error)
	Health(ctx context.Context) (Health, error)
	HealthClock(ctx context.Context) (HealthClock, error)
	Metrics(ctx context.Context) ([]Metric, error)
	MetricsEndpoints(ctx context.Context) ([]MetricsEndpoint, error)
	Block(ctx context.Context, hashOrNumber string) (Block, error)
	BlockLatest(ctx context.Context) (Block, error)
	BlockLatestTransactions(ctx context.Context) ([]Transaction, error)
	BlockTransactions(ctx context.Context, hashOrNumer string) ([]Transaction, error)
	BlocksNext(ctx context.Context, hashOrNumber string) ([]Block, error)
	BlocksPrevious(ctx context.Context, hashOrNumber string) ([]Block, error)
}
