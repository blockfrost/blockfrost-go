package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Metrics returns the history of your Blockfrost usage metrics in the past 30 days.
func (c *apiClient) Metrics(ctx context.Context) (mes []Metric, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceMetrics))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}
	res, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&mes); err != nil {
		return
	}
	return mes, nil
}

// MetricsEndpoints returns history of your blockfrost usage metrics
// History of your Blockfrost usage metrics per endpoint in the past 30 days.
func (c *apiClient) MetricsEndpoints(ctx context.Context) (mes []MetricsEndpoint, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceMetricsEndpoint))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}

	res, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&mes)
	if err != nil {
		return
	}
	return mes, nil
}
