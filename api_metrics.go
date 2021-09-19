package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Metrics returns the history of your Blockfrost usage metrics in the past 30 days.
func (c *apiClient) Metrics(ctx context.Context) ([]Metric, error) {
	var metrics []Metric
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceMetrics))
	if err != nil {
		return metrics, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return metrics, err
	}
	req.Header.Add("project_id", c.projectId)

	res, err := c.client.Do(req)
	if err != nil {
		return metrics, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return metrics, handleAPIErrorResponse(res)
	}

	err = json.NewDecoder(res.Body).Decode(&metrics)
	if err != nil {
		return []Metric{}, err
	}
	return metrics, nil
}

// MetricsEndpoints returns history of your blockfrost usage metrics
// History of your Blockfrost usage metrics per endpoint in the past 30 days.
func (c *apiClient) MetricsEndpoints(ctx context.Context) ([]MetricsEndpoint, error) {
	var metricsEndpoints []MetricsEndpoint
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceMetricsEndpoint))
	if err != nil {
		return metricsEndpoints, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return metricsEndpoints, err
	}

	req.Header.Add("project_id", c.projectId)

	res, err := c.client.Do(req)
	if err != nil {
		return metricsEndpoints, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return metricsEndpoints, handleAPIErrorResponse(res)
	}

	err = json.NewDecoder(res.Body).Decode(&metricsEndpoints)
	if err != nil {
		return []MetricsEndpoint{}, err
	}
	return metricsEndpoints, nil
}
