package blockfrostgo

import (
	"bytes"
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

func NewClient(
	// API_KEY
	apiKey string,
	// mtss API's base url
	cardanoUrlNET string,
	// skipVerify
	skipVerify bool,
	//optional, defaults to http.DefaultClient
	httpClient *http.Client,
	debug io.Writer,
) Client {

	c := &client{
		apiKey:     apiKey,
		url:        cardanoUrlNET,
		httpClient: httpClient,
		debug:      debug,
	}
	if httpClient != nil {
		c.httpClient = httpClient
	} else {
		c.httpClient = http.DefaultClient
	}
	if skipVerify {
		// #nosec
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		c.httpClient.Transport = tr
	}
	return c
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
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.url, healthStatusPath))
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
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.url, healthClockPath))
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

func (c *client) Metrics(ctx context.Context) ([]Metric, error) {
	var metrics []Metric
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.url, metricsPath))
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
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.url, metricsEndpointsPath))
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

// Helpers functions

// dumpResponse writes the raw response data to the debug output, if set, or
// standard error otherwise.
func (c *client) dumpResponse(resp *http.Response) {
	// ignore errors dumping response - no recovery from this
	responseDump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		log.Fatalf("dumpResponse: " + err.Error())
	}
	fmt.Fprintln(c.debug, string(responseDump))
	fmt.Fprintln(c.debug)
}

// apiCall define how you can make a call to Mtss API
func (c *client) apiCall(
	ctx context.Context,
	method string,
	URL string,
	data []byte,
) (statusCode int, response string, err error) {

	req, err := http.NewRequest(method, URL, bytes.NewBuffer(data))
	if err != nil {
		return 0, "", fmt.Errorf("failed to create HTTP request: %v", err)
	}
	req.Header.Add("project_id", c.apiKey)
	req.Header.Add("content-type", "application/json")
	req.Header.Set("User-Agent", "blokfrost-client/0.0")
	if c.debug != nil {
		requestDump, err := httputil.DumpRequestOut(req, true)
		if err != nil {
			return 0, "", fmt.Errorf("error dumping HTTP request: %v", err)
		}
		fmt.Fprintln(c.debug, string(requestDump))
		fmt.Fprintln(c.debug)
	}
	req = req.WithContext(ctx)
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return 0, "", fmt.Errorf("HTTP request failed with: %v", err)
	}
	defer resp.Body.Close()
	if c.debug != nil {
		c.dumpResponse(resp)
	}
	res, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, "", fmt.Errorf("HTTP request failed: %v", err)
	}
	return resp.StatusCode, string(res), nil
}
