package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	resourceNutLink = "nutlink"
	resourceTickers = "tickers"
)

type NutlinkAddressMeta struct {
	Ticker      string `json:"ticker,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	HomePage    string `json:"homepage,omitempty"`
	Address     string `json:"address,omitempty"`
}

type NutlinkAddress struct {
	Address      string             `json:"address,omitempty"`
	MetadataUrl  string             `json:"metadata_url,omitempty"`
	MetadataHash string             `json:"metadata_hash,omitempty"`
	Metadata     NutlinkAddressMeta `json:"metadata,omitempty"`
}

type Ticker struct {
	Name        string `json:"name,omitempty"`
	Count       int    `json:"count,omitempty"`
	LatestBlock int    `json:"latest_block,omitempty"`
}

func (c *apiClient) Nutlink(ctx context.Context, address string) (nu NutlinkAddress, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourceNutLink, address))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}
	req.Header.Add("project_id", c.projectId)

	res, err := c.client.Do(req)
	if err != nil {
		return nu, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return nu, handleAPIErrorResponse(res)
	}

	if err := json.NewDecoder(res.Body).Decode(&nu); err != nil {
		return nu, err
	}
	return nu, nil
}

func (c *apiClient) Tickers(ctx context.Context, address string, query APIPagingParams) (ti []Ticker, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceNutLink, address, resourceTickers))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}

	v := req.URL.Query()
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()
	req.Header.Add("project_id", c.projectId)

	res, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return ti, handleAPIErrorResponse(res)
	}

	if err := json.NewDecoder(res.Body).Decode(&ti); err != nil {
		return ti, err
	}
	return ti, nil
}
