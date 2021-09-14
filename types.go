package blockfrostgo

import (
	"context"
	"io"
	"net/http"
)

const (
	CardanoMainnet = "https://cardano-mainnet.blockfrost.io/api/v0"
	CardanoTestnet = "https://cardano-testnet.blockfrost.io/api/v0"
	IpfsBlockfrost = "https://ipfs.blockfrost.io/api/v0"
)

var (
	healthStatusPath     = "health"
	healthClockPath      = "health/clock"
	metricsPath          = "metrics/"
	metricsEndpointsPath = "metrics/endpoints"
)

// Client is an interface that implements https://blockfrost.io
type Client interface {
	// Health endpoints

	Info(ctx context.Context) (AppInfo, error)
	Health(ctx context.Context) (HealthStatus, error)
	HealthClock(ctx context.Context) (HealthClock, error)

	// Metrics
	Metrics(ctx context.Context) ([]Metric, error)
	MetricsEndpoints(ctx context.Context) ([]MetricEndpoint, error)
}

// Client
// client represents a blockfrost client. If the Debug field is set to an io.Writer
// (for example os.Stdout), then the client will dump API requests and responses
// to it.  To use a non-default HTTP client (for example, for testing, or to set
// a timeout), assign to the HTTPClient field. To set a non-default URL (for
// example, for testing), assign to the URL field.
type client struct {
	apiKey     string
	url        string
	httpClient *http.Client
	debug      io.Writer
}

// Health endpoints

// AppInfo return the Root endpoint
// has no other function than to point end users to documentation.
type AppInfo struct {
	URL     string `json:"url,omitempty"`
	Version string `json:"version,omitempty"`
}

// HealthStatus return the Backend health status
// Return backend status as a boolean. Your application
// should handle situations when backend for the given chain is unavailable.
type HealthStatus struct {
	IsHealthy bool `json:"is_healthy,omitempty"`
}

// HealthClock return the Current backend time
// This endpoint provides the current UNIX time.
// Your application might use this to verify
// if the client clock is not out of sync.
type HealthClock struct {
	ServerTime int64 `json:"server_time,omitempty"`
}

// Metrics return the Blockfrost usage metrics
// History of your Blockfrost usage metrics in the past 30 days.
type Metric struct {
	Time  int64 `json:"time,omitempty"`
	Calls int32 `json:"calls,omitempty"`
}
type MetricEndpoint struct {
	Time     int64  `json:"time,omitempty"`
	Calls    int32  `json:"calls,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
}
