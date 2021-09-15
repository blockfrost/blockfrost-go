package blockfrostgo

import (
	"bytes"
	"context"
	"crypto/tls"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
)

const (
	CardanoMainnet = "https://cardano-mainnet.blockfrost.io/api/v0"
	CardanoTestnet = "https://cardano-testnet.blockfrost.io/api/v0"
	IpfsBlockfrost = "https://ipfs.blockfrost.io/api/v0"
)

var (
	healthStatusPath      = "health"
	healthClockPath       = "health/clock"
	metricsPath           = "metrics/"
	metricsEndpointsPath  = "metrics/endpoints"
	accountsByAddressPath = "accounts" // accounts/{stake_address} parameter
)

// Client is an interface that implements https://blockfrost.io
type BlockfrostAPI interface {
	// Health endpoints

	Info(ctx context.Context) (AppInfo, error)
	Health(ctx context.Context) (HealthStatus, error)
	HealthClock(ctx context.Context) (HealthClock, error)

	// Metrics
	Metrics(ctx context.Context) ([]Metric, error)
	MetricsEndpoints(ctx context.Context) ([]MetricEndpoint, error)

	// Accounts
	Account(ctx context.Context, stakeAddr string) (Account, error)
}

// Client represents a blockfrost client. If the Debug field is set to an io.Writer
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

func NewBlockfrostAPI(
	// API_KEY
	apiKey string,
	// mtss API's base url
	cardanoUrlNET string,
	// skipVerify
	skipVerify bool,
	//optional, defaults to http.DefaultClient
	httpClient *http.Client,
	debug io.Writer,
) BlockfrostAPI {

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
