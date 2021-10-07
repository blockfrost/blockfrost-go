package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	resourceNetwork = "network"
)

type NetworkSupply struct {
	Max         string `json:"max,omitempty"`
	Total       string `json:"total,omitempty"`
	Circulating string `json:"circulating,omitempty"`
	Locked      string `json:"locked,omitempty"`
}

type NetworkStake struct {
	Live   string `json:"live,omitempty"`
	Active string `json:"active,omitempty"`
}
type NetworkInfo struct {
	Supply NetworkSupply `json:"supply,omitempty"`
	Stake  NetworkStake  `json:"stake,omitempty"`
}

func (c *apiClient) Network(ctx context.Context) (ni NetworkInfo, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceNetwork))
	if err != nil {
		return ni, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return ni, err
	}

	res, err := c.handleRequest(req)
	if err != nil {
		return ni, err
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&ni); err != nil {
		return ni, err
	}

	return
}
