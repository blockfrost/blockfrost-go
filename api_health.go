package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Info returns the root endpoint `/`. Root endpoint has
// no other function than to point end users to documentation.
func (c *apiClient) Info(ctx context.Context) (Info, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, c.server, nil)
	if err != nil {
		return Info{}, err
	}
	req.Header.Add("project_id", c.projectId)

	res, err := c.client.Do(req)

	if err != nil {
		return Info{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return Info{}, handleAPIErrorResponse(res)
	}

	info := Info{}
	err = json.NewDecoder(res.Body).Decode(&info)
	if err != nil {
		return Info{}, err
	}
	return info, nil

}

// Health returns the backend health status
// Return backend status as a boolean. Your application
// should handle situations when backend for the given chain is unavailable.
func (c *apiClient) Health(ctx context.Context) (Health, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceHealth))
	if err != nil {
		return Health{}, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return Health{}, err
	}
	req.Header.Add("project_id", c.projectId)

	res, err := c.client.Do(req)
	if err != nil {
		return Health{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return Health{}, handleAPIErrorResponse(res)
	}

	health := Health{}
	err = json.NewDecoder(res.Body).Decode(&health)
	if err != nil {
		return Health{}, err
	}
	return health, nil
}

// HealthClock returns the current UNIX time. Your application might
// use this to verify if the client clock is not out of sync.
func (c *apiClient) HealthClock(ctx context.Context) (HealthClock, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceHealthClock))
	if err != nil {
		return HealthClock{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return HealthClock{}, err
	}
	req.Header.Add("project_id", c.projectId)

	res, err := c.client.Do(req)

	if err != nil {
		return HealthClock{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return HealthClock{}, handleAPIErrorResponse(res)
	}

	healthClock := HealthClock{}
	err = json.NewDecoder(res.Body).Decode(&healthClock)
	if err != nil {
		return HealthClock{}, err
	}
	return healthClock, nil
}
