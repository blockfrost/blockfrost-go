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

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return metrics, err
	}
	req.WithContext(ctx)
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

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return metricsEndpoints, err
	}

	req.Header.Add("project_id", c.projectId)
	req.WithContext(ctx)

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

func handleAPIErrorResponse(res *http.Response) error {
	var err error
	switch res.StatusCode {
	case 400:
		br := BadRequest{}
		if err = json.NewDecoder(res.Body).Decode(&br); err != nil {
			return err
		}
		return &APIError{
			Response: br,
		}
	case 403:
		ua := UnauthorizedError{}
		if err = json.NewDecoder(res.Body).Decode(&ua); err != nil {
			return err
		}
		return &APIError{
			Response: ua,
		}
	case 404:
		nf := NotFound{}
		if err = json.NewDecoder(res.Body).Decode(&nf); err != nil {
			return err
		}
		return &APIError{
			Response: nf,
		}
	case 429:
		ol := OverusageLimit{}
		if err = json.NewDecoder(res.Body).Decode(&ol); err != nil {
			return err
		}
		return &APIError{
			Response: ol,
		}
	case 418:
		ab := AutoBanned{}
		if err = json.NewDecoder(res.Body).Decode(&ab); err != nil {
			return err
		}
		return &APIError{
			Response: ab,
		}
	case 500:
		ise := InternalServerError{}
		if err = json.NewDecoder(res.Body).Decode(&ise); err != nil {
			return err
		}
		return &APIError{
			Response: ise,
		}
	default:
		return &APIError{}
	}
}
