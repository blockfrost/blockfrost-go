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
	resourcePool          = "pools"
	resourcePoolRetired   = "pools/retired"
	resourcePoolRetiring  = "pools/retiring"
	resourcePoolHistory   = "history"
	resourcePoolMetadata  = "metadata"
	resourcePoolRelay     = "relays"
	resourcePoolDelegator = "delegators"
	resourcePoolBlocks    = "blocks"
	resourcePoolUpdate    = "updates"
)

// List of stake pools
// List of registered stake pools.
type Pools []string

// List of retired stake pools
// List of already retired pools.
type PoolRetired struct {
	PoolID string `json:"pool_id,omitempty"`
	Epoch  int    `json:"epoch,omitempty"`
}

type PoolRetiring struct {
	PoolID string `json:"pool_id,omitempty"`
	Epoch  int    `json:"epoch,omitempty"`
}

// Pool information.
type Pool struct {
	PoolID         string   `json:"pool_id,omitempty"`
	Hex            string   `json:"hex,omitempty"`
	VrfKey         string   `json:"vrf_key,omitempty"`
	BlocksMinted   int      `json:"blocks_minted,omitempty"`
	LiveStake      string   `json:"live_stake,omitempty"`
	LiveSize       float64  `json:"live_size,omitempty"`
	LiveSaturation float64  `json:"live_saturation,omitempty"`
	LiveDelegators int      `json:"live_delegators,omitempty"`
	ActiveStake    string   `json:"active_stake,omitempty"`
	ActiveSize     float64  `json:"active_size,omitempty"`
	DeclaredPledge string   `json:"declared_pledge,omitempty"`
	LivePledge     string   `json:"live_pledge,omitempty"`
	MarginCost     float64  `json:"margin_cost,omitempty"`
	FixedCost      string   `json:"fixed_cost,omitempty"`
	RewardAccount  string   `json:"reward_account,omitempty"`
	Owners         []string `json:"owners,omitempty"`
	Registration   []string `json:"registration,omitempty"`
	Retirement     []string `json:"retirement,omitempty"`
}

// PoolHistory return Stake pool history
// History of stake pool parameters over epochs.
type PoolHistory struct {
	Epoch           int     `json:"epoch,omitempty"`
	Blocks          int     `json:"blocks,omitempty"`
	ActiveStake     string  `json:"active_stake,omitempty"`
	ActiveSize      float64 `json:"active_size,omitempty"`
	DelegatorsCount int     `json:"delegators_count,omitempty"`
	Rewards         string  `json:"rewards,omitempty"`
	Fees            string  `json:"fees,omitempty"`
}

// PoolMetadata return Stake pool metadata
// Stake pool registration metadata.
type PoolMetadata struct {
	PoolID      string `json:"pool_id,omitempty"`
	Hex         string `json:"hex,omitempty"`
	URL         string `json:"url,omitempty"`
	Hash        string `json:"hash,omitempty"`
	Ticker      string `json:"ticker,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Homepage    string `json:"homepage,omitempty"`
}

// PoolRelay return Stake pool relays
// Relays of a stake pool.
type PoolRelay struct {
	Ipv4   string `json:"ipv4,omitempty"`
	Ipv6   string `json:"ipv6,omitempty"`
	DNS    string `json:"dns,omitempty"`
	DNSSrv string `json:"dns_srv,omitempty"`
	Port   int    `json:"port,omitempty"`
}

// PoolDelegator return Stake pool delegators
// List of current stake pools delegators.
type PoolDelegator struct {
	Address   string `json:"address"`
	LiveStake string `json:"live_stake"`
}

// PoolBlocks return Stake pool blocks
// List of stake pools blocks.
type PoolBlocks []string

// PoolUpdate return Stake pool updates
// List of certificate updates to the stake pool.
type PoolUpdate struct {
	TxHash    string `json:"tx_hash,omitempty"`
	CERTIndex int    `json:"cert_index,omitempty"`
	Action    string `json:"action,omitempty"`
}

type PoolsResult struct {
	Res Pools
	Err error
}

type PoolsRetiredResult struct {
	Res []PoolRetired
	Err error
}

type PoolsRetiringResult struct {
	Res []PoolRetiring
	Err error
}

type PoolHistoryResult struct {
	Res []PoolHistory
	Err error
}

type PoolDelegatorsResult struct {
	Res []PoolDelegator
	Err error
}

type PoolBlocksResult struct {
	Res PoolBlocks
	Err error
}

type PoolUpdateResult struct {
	Res []PoolUpdate
	Err error
}

// Pools returns the List of stake pools
// List of registered stake pools.
func (c *apiClient) Pools(ctx context.Context, query APIPagingParams) (Pools, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourcePool))
	if err != nil {
		return Pools{}, err
	}

	v := url.Values{}
	if query.Count > 0 {
		v.Set("count", fmt.Sprintf("%d", query.Count))
		requestUrl.RawQuery = v.Encode()
	}
	if query.Page > 0 {
		v.Set("page", fmt.Sprintf("%d", query.Page))
		requestUrl.RawQuery = v.Encode()
	}
	if query.Order != "" {
		v.Set("order", query.Order)
		requestUrl.RawQuery = v.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return Pools{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return Pools{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Pools{}, handleAPIErrorResponse(res)
	}
	pools := Pools{}
	err = json.NewDecoder(res.Body).Decode(&pools)
	if err != nil {
		return Pools{}, err
	}
	return pools, nil
}

func (c *apiClient) PoolsAll(ctx context.Context) <-chan PoolsResult {
	ch := make(chan PoolsResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan PoolsResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				pools, err := c.Pools(j.ctx, j.query)
				if len(pools) != j.query.Count || err != nil {
					quit <- true
				}
				res := PoolsResult{Res: pools, Err: err}
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

// PoolsRetired returns the List of retired stake pools
// List of already retired pools.
func (c *apiClient) PoolsRetired(ctx context.Context, query APIPagingParams) ([]PoolRetired, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourcePoolRetired))
	if err != nil {
		return []PoolRetired{}, err
	}

	v := url.Values{}
	if query.Count > 0 {
		v.Set("count", fmt.Sprintf("%d", query.Count))
		requestUrl.RawQuery = v.Encode()
	}
	if query.Page > 0 {
		v.Set("page", fmt.Sprintf("%d", query.Page))
		requestUrl.RawQuery = v.Encode()
	}
	if query.Order != "" {
		v.Set("order", query.Order)
		requestUrl.RawQuery = v.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []PoolRetired{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []PoolRetired{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []PoolRetired{}, handleAPIErrorResponse(res)
	}
	pools := []PoolRetired{}
	err = json.NewDecoder(res.Body).Decode(&pools)
	if err != nil {
		return []PoolRetired{}, err
	}
	return pools, nil
}

func (c *apiClient) PoolsRetiredAll(ctx context.Context) <-chan PoolsRetiredResult {
	ch := make(chan PoolsRetiredResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan PoolsRetiredResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				pools, err := c.PoolsRetired(j.ctx, j.query)
				if len(pools) != j.query.Count || err != nil {
					quit <- true
				}
				res := PoolsRetiredResult{Res: pools, Err: err}
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

// PoolsRetiring returns the List of retiring stake pools
// List of stake pools retiring in the upcoming epochs
func (c *apiClient) PoolsRetiring(ctx context.Context, query APIPagingParams) (pr []PoolRetiring, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourcePoolRetiring))
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

	if err = json.NewDecoder(res.Body).Decode(&pr); err != nil {
		return
	}

	return pr, nil
}

func (c *apiClient) PoolsRetiringAll(ctx context.Context) <-chan PoolsRetiringResult {
	ch := make(chan PoolsRetiringResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan PoolsRetiringResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				pools, err := c.PoolsRetiring(j.ctx, j.query)
				if len(pools) != j.query.Count || err != nil {
					quit <- true
				}
				res := PoolsRetiringResult{Res: pools, Err: err}
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

// Pool returns the Specific Stake Pool
func (c *apiClient) Pool(ctx context.Context, poolID string) (pool Pool, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourcePool, poolID))
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
	if err = json.NewDecoder(res.Body).Decode(&pool); err != nil {
		return
	}
	return pool, nil
}

// PoolHistory returns the Stake pool history
// History of stake pool parameters over epochs.
func (c *apiClient) PoolHistory(ctx context.Context, poolID string, query APIPagingParams) (ph []PoolHistory, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourcePool, poolID, resourcePoolHistory))
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

	if err := json.NewDecoder(res.Body).Decode(&ph); err != nil {
		return []PoolHistory{}, err
	}
	return ph, nil
}

func (c *apiClient) PoolHistoryAll(ctx context.Context, poolId string) <-chan PoolHistoryResult {
	ch := make(chan PoolHistoryResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan PoolHistoryResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				pools, err := c.PoolHistory(j.ctx, poolId, j.query)
				if len(pools) != j.query.Count || err != nil {
					quit <- true
				}
				res := PoolHistoryResult{Res: pools, Err: err}
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

// PoolMetadata returns the Stake pool metadata
// Stake pool registration metadata.
func (c *apiClient) PoolMetadata(ctx context.Context, poolID string) (pm PoolMetadata, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourcePool, poolID, resourcePoolMetadata))
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
	err = json.NewDecoder(res.Body).Decode(&pm)
	if err != nil {
		return
	}
	return pm, nil
}

// PoolRelay returns the Stake pool relays
// Relays of a stake pool.
func (c *apiClient) PoolRelays(ctx context.Context, poolID string) (prs []PoolRelay, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourcePool, poolID, resourcePoolRelay))
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

	if err = json.NewDecoder(res.Body).Decode(&prs); err != nil {
		return prs, err
	}
	return prs, nil
}

// PoolDelegator returns the Stake pool delegators
// List of current stake pools delegators.
func (c *apiClient) PoolDelegators(ctx context.Context, poolID string, query APIPagingParams) (pd []PoolDelegator, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourcePool, poolID, resourcePoolDelegator))
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

	if res.StatusCode != http.StatusOK {
		return pd, handleAPIErrorResponse(res)
	}
	if err = json.NewDecoder(res.Body).Decode(&pd); err != nil {
		return pd, err
	}
	return pd, nil
}

func (c *apiClient) PoolDelegatorsAll(ctx context.Context, poolId string) <-chan PoolDelegatorsResult {
	ch := make(chan PoolDelegatorsResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan PoolDelegatorsResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				pools, err := c.PoolDelegators(j.ctx, poolId, j.query)
				if len(pools) != j.query.Count || err != nil {
					quit <- true
				}
				res := PoolDelegatorsResult{Res: pools, Err: err}
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

// PoolBlocks returns the Stake pool blocks
// List of stake pools blocks.
func (c *apiClient) PoolBlocks(ctx context.Context, poolID string, query APIPagingParams) (pb PoolBlocks, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourcePool, poolID, resourcePoolBlocks))
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

	if res.StatusCode != http.StatusOK {
		return pb, handleAPIErrorResponse(res)
	}
	if err = json.NewDecoder(res.Body).Decode(&pb); err != nil {
		return pb, err
	}
	return pb, nil
}

func (c *apiClient) PoolBlocksAll(ctx context.Context, poolId string) <-chan PoolBlocksResult {
	ch := make(chan PoolBlocksResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan PoolBlocksResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				pools, err := c.PoolBlocks(j.ctx, poolId, j.query)
				if len(pools) != j.query.Count || err != nil {
					quit <- true
				}
				res := PoolBlocksResult{Res: pools, Err: err}
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

// PoolUpdate returns the Stake pool updates
// List of certificate updates to the stake pool.
func (c *apiClient) PoolUpdates(ctx context.Context, poolID string, query APIPagingParams) (pu []PoolUpdate, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourcePool, poolID, resourcePoolUpdate))
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

	if res.StatusCode != http.StatusOK {
		return pu, handleAPIErrorResponse(res)
	}
	if err = json.NewDecoder(res.Body).Decode(&pu); err != nil {
		return
	}
	return pu, nil
}

func (c *apiClient) PoolUpdateAll(ctx context.Context, poolId string) <-chan PoolUpdateResult {
	ch := make(chan PoolUpdateResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan PoolUpdateResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				pools, err := c.PoolUpdates(j.ctx, poolId, j.query)
				if len(pools) != j.query.Count || err != nil {
					quit <- true
				}
				res := PoolUpdateResult{Res: pools, Err: err}
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
