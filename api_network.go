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

// NetworkSupply contains information on network supply
type NetworkSupply struct {
	Max         string `json:"max,"`
	Total       string `json:"total,"`
	Circulating string `json:"circulating,"`
	Locked      string `json:"locked,"`
}

// NetworkStake contains information on the cardano network stake
type NetworkStake struct {
	Live   string `json:"live,"`
	Active string `json:"active,"`
}

// NetworkInfo contains network stake and supply information on the network
type NetworkInfo struct {
	Supply NetworkSupply `json:"supply"`
	Stake  NetworkStake  `json:"stake"`
}

// Network returns detailed network information.
func (c *apiClient) Network(ctx context.Context) (ni NetworkInfo, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceNetwork))
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
	if err = json.NewDecoder(res.Body).Decode(&ni); err != nil {
		return
	}

	return ni, nil
}
