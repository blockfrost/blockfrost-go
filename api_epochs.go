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
	resourceEpochs          = "epochs"
	resourceEpochsLatest    = "latest"
	resourceEpochsNext      = "next"
	resourceEpochsPrevious  = "previous"
	resourceEpochsStakes    = "stakes"
	resourceEpochsBlocks    = "blocks"
	resourceEpochParameters = "parameters"
)

type EpochStake struct {
	StakeAddress string `json:"stake_address,omitempty"`
	PoolID       string `json:"pool_id,omitempty"`
	Amount       string `json:"amount,omitempty"`
}

type Epoch struct {
	// Sum of all the active stakes within the epoch in Lovelaces
	ActiveStake string `json:"active_stake"`

	// Number of blocks within the epoch
	BlockCount int `json:"block_count"`

	// Unix time of the end of the epoch
	EndTime int `json:"end_time"`

	// Epoch number
	Epoch int `json:"epoch"`

	// Sum of all the fees within the epoch in Lovelaces
	Fees string `json:"fees"`

	// Unix time of the first block of the epoch
	FirstBlockTime int `json:"first_block_time"`

	// Unix time of the last block of the epoch
	LastBlockTime int `json:"last_block_time"`

	// Sum of all the transactions within the epoch in Lovelaces
	Output string `json:"output"`

	// Unix time of the start of the epoch
	StartTime int `json:"start_time"`

	// Number of transactions within the epoch
	TxCount int `json:"tx_count"`
}

type EpochParameters struct {
	// Pool pledge influence
	A0 float32 `json:"a0"`

	// Percentage of blocks produced by federated nodes
	DecentralisationParam float32 `json:"decentralisation_param"`

	// Epoch bound on pool retirement
	EMax int `json:"e_max"`

	// Epoch number
	Epoch int `json:"epoch"`

	// Seed for extra entropy
	ExtraEntropy *map[string]interface{} `json:"extra_entropy"`

	// The amount of a key registration deposit in Lovelaces
	KeyDeposit string `json:"key_deposit"`

	// Maximum block header size
	MaxBlockHeaderSize int `json:"max_block_header_size"`

	// Maximum block body size in Bytes
	MaxBlockSize int `json:"max_block_size"`

	// Maximum transaction size
	MaxTxSize int `json:"max_tx_size"`

	// The linear factor for the minimum fee calculation for given epoch
	MinFeeA int `json:"min_fee_a"`

	// The constant factor for the minimum fee calculation
	MinFeeB int `json:"min_fee_b"`

	// Minimum stake cost forced on the pool
	MinPoolCost string `json:"min_pool_cost"`

	// Minimum UTXO value
	MinUtxo string `json:"min_utxo"`

	// Desired number of pools
	NOpt int `json:"n_opt"`

	// Epoch number only used once
	Nonce string `json:"nonce"`

	// The amount of a pool registration deposit in Lovelaces
	PoolDeposit string `json:"pool_deposit"`

	// Accepted protocol major version
	ProtocolMajorVer int `json:"protocol_major_ver"`

	// Accepted protocol minor version
	ProtocolMinorVer int `json:"protocol_minor_ver"`

	// Monetary expansion
	Rho float32 `json:"rho"`

	// Treasury expansion
	Tau float32 `json:"tau"`
}

type EpochResult struct {
	Res []Epoch
	Err error
}

type EpochStakeResult struct {
	Res []EpochStake
	Err error
}

type BlockDistributionResult struct {
	Res []string
	Err error
}

func (c *apiClient) EpochLatest(ctx context.Context) (ep Epoch, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceEpochs, resourceEpochsLatest, resourceEpochParameters))
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

	if err = json.NewDecoder(res.Body).Decode(&ep); err != nil {
		return
	}
	return ep, nil
}

func (c *apiClient) LatestEpochParameters(ctx context.Context) (epr EpochParameters, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourceEpochs, resourceEpochsLatest))
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

	if err = json.NewDecoder(res.Body).Decode(&epr); err != nil {
		return
	}
	return epr, nil
}

func (c *apiClient) Epoch(ctx context.Context, epochNumber int) (ep Epoch, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%d", c.server, resourceEpochs, epochNumber))
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

	if err = json.NewDecoder(res.Body).Decode(&ep); err != nil {
		return
	}
	return ep, nil
}

func (c *apiClient) EpochsNext(ctx context.Context, epochNumber int, query APIPagingParams) (eps []Epoch, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%d/%s", c.server, resourceEpochs, epochNumber, resourceEpochsNext))
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

	if err = json.NewDecoder(res.Body).Decode(&eps); err != nil {
		return
	}
	return eps, nil
}

func (c *apiClient) EpochNextAll(ctx context.Context, epochNumber int) <-chan EpochResult {
	ch := make(chan EpochResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan EpochResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				as, err := c.EpochsNext(j.ctx, epochNumber, j.query)
				if len(as) != j.query.Count || err != nil {
					quit <- true
				}
				res := EpochResult{Res: as, Err: err}
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

func (c *apiClient) EpochsPrevious(ctx context.Context, epochNumber int, query APIPagingParams) (eps []Epoch, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%d/%s", c.server, resourceEpochs, epochNumber, resourceEpochsPrevious))
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

	if err = json.NewDecoder(res.Body).Decode(&eps); err != nil {
		return
	}
	return eps, nil
}

func (c *apiClient) EpochPreviousAll(ctx context.Context, epochNumber int) <-chan EpochResult {
	ch := make(chan EpochResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan EpochResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				as, err := c.EpochsPrevious(j.ctx, epochNumber, j.query)
				if len(as) != j.query.Count || err != nil {
					quit <- true
				}
				res := EpochResult{Res: as, Err: err}
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

func (c *apiClient) EpochStakeDistribution(ctx context.Context, epochNumber int, query APIPagingParams) (eps []EpochStake, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%d/%s", c.server, resourceEpochs, epochNumber, resourceEpochsStakes))
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

	if err = json.NewDecoder(res.Body).Decode(&eps); err != nil {
		return
	}
	return eps, nil
}

func (c *apiClient) EpochStakeDistributionAll(ctx context.Context, epochNumber int) <-chan EpochStakeResult {
	ch := make(chan EpochStakeResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan EpochStakeResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				eps, err := c.EpochStakeDistribution(j.ctx, epochNumber, j.query)
				if len(eps) != j.query.Count || err != nil {
					quit <- true
				}
				res := EpochStakeResult{Res: eps, Err: err}
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

func (c *apiClient) EpochStakeDistributionByPool(ctx context.Context, epochNumber int, poolId string, query APIPagingParams) (eps []EpochStake, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%d/%s/%s", c.server, resourceEpochs, epochNumber, resourceEpochsStakes, poolId))
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

	if err = json.NewDecoder(res.Body).Decode(&eps); err != nil {
		return
	}
	return eps, nil
}

func (c *apiClient) EpochStakeDistributionByPoolAll(ctx context.Context, epochNumber int, poolId string) <-chan EpochStakeResult {
	ch := make(chan EpochStakeResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan EpochStakeResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				eps, err := c.EpochStakeDistributionByPool(j.ctx, epochNumber, poolId, j.query)
				if len(eps) != j.query.Count || err != nil {
					quit <- true
				}
				res := EpochStakeResult{Res: eps, Err: err}
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

func (c *apiClient) EpochBlockDistribution(ctx context.Context, epochNumber int, query APIPagingParams) (bd []string, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%d/%s", c.server, resourceEpochs, epochNumber, resourceEpochsStakes))
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

	if err = json.NewDecoder(res.Body).Decode(&bd); err != nil {
		return
	}
	return bd, nil
}

func (c *apiClient) EpochBlockDistributionAll(ctx context.Context, epochNumber int) <-chan BlockDistributionResult {
	ch := make(chan BlockDistributionResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan BlockDistributionResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				eps, err := c.EpochBlockDistribution(j.ctx, epochNumber, j.query)
				if len(eps) != j.query.Count || err != nil {
					quit <- true
				}
				res := BlockDistributionResult{Res: eps, Err: err}
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

func (c *apiClient) EpochBlockDistributionByPool(ctx context.Context, epochNumber int, poolId string, query APIPagingParams) (bd []string, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%d/%s", c.server, resourceEpochs, epochNumber, resourceEpochsStakes))
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

	if err = json.NewDecoder(res.Body).Decode(&bd); err != nil {
		return
	}
	return bd, nil
}

func (c *apiClient) EpochBlockDistributionByPoolAll(ctx context.Context, epochNumber int, poolId string) <-chan BlockDistributionResult {
	ch := make(chan BlockDistributionResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan BlockDistributionResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				eps, err := c.EpochBlockDistributionByPool(j.ctx, epochNumber, poolId, j.query)
				if len(eps) != j.query.Count || err != nil {
					quit <- true
				}
				res := BlockDistributionResult{Res: eps, Err: err}
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

func (c *apiClient) EpochParameters(ctx context.Context, epochNumber int) (eps EpochParameters, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%d/%s", c.server, resourceEpochs, epochNumber, resourceEpochsStakes))
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

	if err = json.NewDecoder(res.Body).Decode(&eps); err != nil {
		return
	}
	return eps, nil
}
