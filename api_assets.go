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

// Contains metadata information about an asset.
type AssetMetadata struct {
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Ticker      string `json:"ticker,omitempty"`
	URL         string `json:"url,omitempty"`
	Logo        string `json:"logo,omitempty"`
	Decimals    int    `json:"decimals,omitempty"`
}

// Assets contains information on an asset.
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

// AssetHistory contains history of an asset.
type AssetHistory struct {
	TxHash string `json:"tx_hash,omitempty"`
	Action string `json:"action,omitempty"`
	Amount string `json:"amount,omitempty"`
}

// AssetTransaction contains information on transactions belonging
// to an asset.
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

// Assets returns a paginated list of assets.
func (c *apiClient) Assets(ctx context.Context, query APIQueryParams) (a []Asset, err error) {
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

	res, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&a); err != nil {
		return
	}
	return a, nil
}

// AssetsAll returns all assets.
func (c *apiClient) AssetsAll(ctx context.Context) <-chan AssetResult {
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
				jobs <- methodOptions{ctx: ctx, query: APIQueryParams{Count: 100, Page: i}}
			}
		}

		wg.Wait()
	}()
	return ch
}

// Asset returns information about a specific asset.
func (c *apiClient) Asset(ctx context.Context, asset string) (a Asset, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourceAssets, asset))
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
	if err = json.NewDecoder(res.Body).Decode(&a); err != nil {
		return
	}
	return a, nil
}

// AssetHistory returns history of a specific asset.
func (c *apiClient) AssetHistory(ctx context.Context, asset string) (hist []AssetHistory, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAssets, asset, resourceAssetHistory))
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

	if err = json.NewDecoder(res.Body).Decode(&hist); err != nil {
		return
	}
	return hist, nil
}

// AssetTransactions returns list of a specific asset transactions.
func (c *apiClient) AssetTransactions(ctx context.Context, asset string) (trs []AssetTransaction, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAssets, asset, resourceAssetTransactions))
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

	if err = json.NewDecoder(res.Body).Decode(&trs); err != nil {
		return
	}
	return trs, nil
}

// AssetAddresses returns list of a addresses containing a specific asset.
func (c *apiClient) AssetAddresses(ctx context.Context, asset string) (addrs []AssetAddress, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAssets, asset, resourceAssetHistory))
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

	if err = json.NewDecoder(res.Body).Decode(&addrs); err != nil {
		return
	}
	return addrs, nil
}

// AssetsByPolicy returns list of assets minted under a specific policy.
func (c *apiClient) AssetsByPolicy(ctx context.Context, policyId string) (a []Asset, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourcePolicyAssets, policyId))
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

	if err = json.NewDecoder(res.Body).Decode(&a); err != nil {
		return
	}
	return a, nil
}
