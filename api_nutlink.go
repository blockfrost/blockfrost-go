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

type TickerRecord struct {
	TxHash      string `json:"tx_hash,omitempty"`
	BlockHeight int    `json:"block_height,omitempty"`
	TxIndex     int    `json:"tx_index,omitempty"`
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

	res, err := c.handleRequest(req)
	if err != nil {
		return nu, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&nu); err != nil {
		return nu, err
	}
	return nu, nil
}

func (c *apiClient) Tickers(ctx context.Context, address string, query APIQueryParams) (ti []Ticker, err error) {
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

	res, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&ti); err != nil {
		return ti, err
	}
	return ti, nil
}

func (c *apiClient) TickerRecords(ctx context.Context, ticker string, query APIQueryParams) (trs []TickerRecord, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceNutLink, resourceTickers, ticker))
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

	res, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&trs); err != nil {
		return
	}
	return trs, nil
}

func (c *apiClient) AddressTickerRecords(ctx context.Context, address string, ticker string, query APIQueryParams) (trs []TickerRecord, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s/%s", c.server, resourceNutLink, address, resourceTickers, ticker))
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
	res, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&trs); err != nil {
		return
	}
	return trs, nil
}
