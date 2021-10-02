package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"sync"
)

const (
	resourceAssets            = "assets"
	resourceAssetHistory      = "history"
	resourceAssetTransactions = "transactions"
	resourceAssetAddresses    = "addresses"
	resourcePolicyAssets      = "assets/policy"
)

type AssetOnchainMetadata struct {
	Name  string `json:"name,omitempty"`
	Image string `json:"image,omitempty"`
}

type AssetMetadata struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Ticker      string `json:"ticker,omitempty"`
	URL         string `json:"url,omitempty"`
	Logo        string `json:"logo,omitempty"`
	Decimals    int    `json:"decimals,omitempty"`
}

type Asset struct {
	Asset             string               `json:"asset,omitempty"`
	PolicyId          string               `json:"policy_id,omitempty"`
	AssetName         string               `json:"asset_name,omitempty"`
	Fingerprint       string               `json:"fingerprint,omitempty"`
	Quantity          string               `json:"quantity,omitempty"`
	InitialMintTxHash string               `json:"initial_mint_tx_hash,omitempty"`
	MintOrBurnCount   int                  `json:"mint_or_burn_count,omitempty"`
	OnchainMetadata   AssetOnchainMetadata `json:"onchain_metadata,omitempty"`
	Metadata          AssetMetadata        `json:"metadata,omitempty"`
}

type AssetHistory struct {
	TxHash string `json:"tx_hash,omitempty"`
	Action string `json:"action,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type AssetTransaction struct {
	TxHash      string `json:"tx_hash,omitempty"`
	TxIndex     int    `json:"tx_index,omitempty"`
	BlockHeight int    `json:"block_height,omitempty"`
}

type AssetAddress struct {
	Address  string `json:"address,omitempty"`
	Quantity string `json:"quantity,omitempty"`
}

type AssetResult struct {
	Res []Asset
	Err error
}

func (c *apiClient) Assets(ctx context.Context, query APIPagingParams) (a []Asset, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceAssets))
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

	res, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return a, handleAPIErrorResponse(res)
	}

	if err = json.NewDecoder(res.Body).Decode(&a); err != nil {
		return
	}
	return a, nil
}

func (c *apiClient) AssetsAll(ctx context.Context, poolId string) <-chan AssetResult {
	ch := make(chan AssetResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan AssetResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				assets, err := c.Assets(j.ctx, j.query)
				if len(assets) != j.query.Count || err != nil {
					quit <- true
				}
				res := AssetResult{Res: assets, Err: err}
				ch <- res
			}

		}(jobs, ch, &wg)
	}
	go func() {
		defer close(ch)
		fetchScripts := true
		for i := 1; fetchScripts; i++ {
			select {
			case <-quit:
				fetchScripts = false
				return
			default:
				jobs <- methodOptions{ctx: ctx, query: APIPagingParams{Count: 100, Page: i}}
			}
		}

		wg.Wait()
	}()
	return ch
}

func (c *apiClient) Asset(ctx context.Context, asset string) (a Asset, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourceAssets, asset))
	if err != nil {
		return
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}
	req.Header.Add("project_id", c.projectId)

	res, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return a, handleAPIErrorResponse(res)
	}

	if err = json.NewDecoder(res.Body).Decode(&a); err != nil {
		return
	}
	return a, nil
}

func (c *apiClient) AssetHistory(ctx context.Context, asset string) (hist []AssetHistory, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAssets, asset, resourceAssetHistory))
	if err != nil {
		return
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}

	req.Header.Add("project_id", c.projectId)

	res, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return hist, handleAPIErrorResponse(res)
	}

	if err = json.NewDecoder(res.Body).Decode(&hist); err != nil {
		return
	}
	return hist, nil
}

func (c *apiClient) AssetTransactions(ctx context.Context, asset string) (trs []AssetTransaction, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAssets, asset, resourceAssetTransactions))
	if err != nil {
		return
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}

	req.Header.Add("project_id", c.projectId)

	res, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return trs, handleAPIErrorResponse(res)
	}

	if err = json.NewDecoder(res.Body).Decode(&trs); err != nil {
		return
	}
	return trs, nil
}

func (c *apiClient) AssetAddresses(ctx context.Context, asset string) (addrs []AssetAddress, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAssets, asset, resourceAssetHistory))
	if err != nil {
		return
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}

	req.Header.Add("project_id", c.projectId)

	res, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return addrs, handleAPIErrorResponse(res)
	}

	if err = json.NewDecoder(res.Body).Decode(&addrs); err != nil {
		return
	}
	return addrs, nil
}

func (c *apiClient) AssetsByPolicy(ctx context.Context, policyId string) (a []Asset, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAssets, resourceAssets, policyId))
	if err != nil {
		return
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}

	req.Header.Add("project_id", c.projectId)

	res, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return a, handleAPIErrorResponse(res)
	}

	if err = json.NewDecoder(res.Body).Decode(&a); err != nil {
		return
	}
	return a, nil
}
