package blockfrostgo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

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

func (c *client) Info(ctx context.Context) (AppInfo, error) {
	requestUrl, err := url.Parse(c.url)
	if err != nil {
		return AppInfo{}, err
	}

	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		requestUrl.String(),
		nil,
	)
	if err != nil {
		return AppInfo{}, err
	}
	if status != http.StatusOK {
		return AppInfo{}, fmt.Errorf("unexpected response status %d: %q", status, res)
	}
	result := AppInfo{}
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return AppInfo{}, fmt.Errorf("decoding error for data %s: %v", res, err)
	}
	return result, nil
}

func (c *client) Health(ctx context.Context) (HealthStatus, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.url, healthPath))
	if err != nil {
		return HealthStatus{}, err
	}

	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		requestUrl.String(),
		nil,
	)
	if err != nil {
		return HealthStatus{}, err
	}
	if status != http.StatusOK {
		return HealthStatus{}, fmt.Errorf("unexpected response status %d: %q", status, res)
	}
	result := HealthStatus{}
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return HealthStatus{}, fmt.Errorf("decoding error for data %s: %v", res, err)
	}
	return result, nil
}

func (c *client) HealthClock(ctx context.Context) (HealthClock, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.url, healthPath, "clock"))
	if err != nil {
		return HealthClock{}, err
	}

	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		requestUrl.String(),
		nil,
	)
	if err != nil {
		return HealthClock{}, err
	}
	if status != http.StatusOK {
		return HealthClock{}, fmt.Errorf("unexpected response status %d: %q", status, res)
	}
	result := HealthClock{}
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return HealthClock{}, fmt.Errorf("decoding error for data %s: %v", res, err)
	}
	return result, nil
}
