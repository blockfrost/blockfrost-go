package blockfrostgo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

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

func (c *client) Metrics(ctx context.Context) ([]Metric, error) {
	var metrics []Metric
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/", c.url, metricsPath))
	if err != nil {
		return metrics, err
	}

	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		requestUrl.String(),
		nil,
	)
	if err != nil {
		return metrics, err
	}
	if status != http.StatusOK {
		return metrics, fmt.Errorf("unexpected response status %d: %q", status, res)
	}
	result := metrics
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return metrics, fmt.Errorf("decoding error for data %s: %v", res, err)
	}
	return result, nil
}

func (c *client) MetricsEndpoints(ctx context.Context) ([]MetricEndpoint, error) {
	var metrics []MetricEndpoint
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.url, metricsPath, "endpoints"))
	if err != nil {
		return metrics, err
	}

	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		requestUrl.String(),
		nil,
	)
	if err != nil {
		return metrics, err
	}
	if status != http.StatusOK {
		return metrics, fmt.Errorf("unexpected response status %d: %q", status, res)
	}
	result := metrics
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return metrics, fmt.Errorf("decoding error for data %s: %v", res, err)
	}
	return result, nil
}
