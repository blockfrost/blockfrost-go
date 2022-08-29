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
	Name  string `json:"name"`
	Image string `json:"image"`
}

// Contains metadata information about an asset.
type AssetMetadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Ticker      string `json:"ticker"`
	URL         string `json:"url"`
	Logo        string `json:"logo"`
	Decimals    int    `json:"decimals"`
}

// Assets contains information on an asset.
type Asset struct {
	// Hex-encoded asset full name
	Asset string `json:"asset"`

	// Policy ID of the asset
	PolicyId string `json:"policy_id"`

	// Hex-encoded asset name of the asset
	AssetName string `json:"asset_name"`

	// CIP14 based user-facing fingerprint
	Fingerprint string `json:"fingerprint"`

	// Current asset quantity
	Quantity string `json:"quantity"`

	// ID of the initial minting transaction
	InitialMintTxHash string `json:"initial_mint_tx_hash"`

	// Count of mint and burn transactions
	MintOrBurnCount int `json:"mint_or_burn_count"`

	// On-chain metadata stored in the minting transaction under label 721,
	// community discussion around the standard ongoing at https://github.com/cardano-foundation/CIPs/pull/85
	OnchainMetadata AssetOnchainMetadata `json:"onchain_metadata"`
	Metadata        AssetMetadata        `json:"metadata"`
}

// AssetHistory contains history of an asset.
type AssetHistory struct {
	// Hash of the transaction containing the asset actio
	TxHash string `json:"tx_hash"`

	// Action executed upon the asset policy.
	// Enum: "minted" "burned"
	Action string `json:"action"`

	// Asset amount of the specific action
	Amount string `json:"amount"`
}

// AssetTransaction contains information on transactions belonging
// to an asset.
type AssetTransaction struct {
	// Hash of the transaction
	TxHash string `json:"tx_hash"`

	// Transaction index within the block
	TxIndex int `json:"tx_index"`

	// Block height
	BlockHeight int `json:"block_height"`

	// Block time
	BlockTime int `json:"block_time"`
}

type AssetAddress struct {
	// Address containing the specific asset
	Address string `json:"address"`

	// Asset quantity on the specific address
	Quantity string `json:"quantity"`
}

type AssetResult struct {
	Res []Asset
	Err error
}

type AssetAddressesAll struct {
	Res []AssetAddress
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
	quit := make(chan bool, 1)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan AssetResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				assets, err := c.Assets(j.ctx, j.query)
				if len(assets) != j.query.Count || err != nil {
					select {
					case quit <- true:
					default:
					}
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
			default:
				jobs <- methodOptions{ctx: ctx, query: APIQueryParams{Count: 100, Page: i}}
			}
		}

		close(jobs)
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

func (c *apiClient) AssetAddresses(ctx context.Context, asset string, query APIQueryParams) (addrs []AssetAddress, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAssets, asset, resourceAddresses))
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

	if err = json.NewDecoder(res.Body).Decode(&addrs); err != nil {
		return
	}
	return addrs, nil
}

// AssetAddresses returns list of a addresses containing a specific asset.
func (c *apiClient) AssetAddressesAll(ctx context.Context, asset string) <-chan AssetAddressesAll {
	ch := make(chan AssetAddressesAll, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, 1)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan AssetAddressesAll, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				ad, err := c.AssetAddresses(j.ctx, asset, j.query)
				if len(ad) != j.query.Count || err != nil {
					select {
					case quit <- true:
					default:
					}
				}
				res := AssetAddressesAll{Res: ad, Err: err}
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
			default:
				jobs <- methodOptions{ctx: ctx, query: APIQueryParams{Count: 100, Page: i}}
			}
		}

		close(jobs)
		wg.Wait()
	}()
	return ch
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
