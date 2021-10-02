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
	resourceAccount                            = "accounts"
	resourceAccountHistory                     = "history"
	resourceAccountRewardsHistory              = "rewards"
	resourceAccountDelegationHistory           = "delegations"
	resourceAccountRegistrationHistory         = "registrations"
	resourceAccountWithdrawalHistory           = "withdrawals"
	resourceAccountMIRHistory                  = "mirs"
	resourceAccountAssociatedAddress           = "addresses"
	resourceAccountAddressWithAssetsAssociated = "addresses/assets"
)

// Account return Specific account address
// Obtain information about a specific stake account.
type Account struct {
	StakeAddress       string `json:"stake_address,omitempty"`
	Active             bool   `json:"active,omitempty"`
	ActiveEpoch        int64  `json:"active_epoch,omitempty"`
	ControlledAmount   string `json:"controlled_amount,omitempty"`
	RewardsSum         string `json:"rewards_sum,omitempty"`
	WithdrawalsSum     string `json:"withdrawals_sum,omitempty"`
	ReservesSum        string `json:"reserves_sum,omitempty"`
	TreasurySum        string `json:"treasury_sum,omitempty"`
	WithdrawableAmount string `json:"withdrawable_amount,omitempty"`
	PoolID             string `json:"pool_id,omitempty"`
}

// AccountRewardsHist return Account reward history
// Obtain information about the reward history of a specific account.
type AccountRewardsHistory struct {
	Epoch  int32  `json:"epoch,omitempty"`
	Amount string `json:"amount,omitempty"`
	PoolID string `json:"pool_id,omitempty"`
}

// AccountHistory return Account history
// Obtain information about the history of a specific account.
type AccountHistory struct {
	ActiveEpoch int32  `json:"active_epoch,omitempty"`
	Amount      string `json:"amount,omitempty"`
	PoolID      string `json:"pool_id,omitempty"`
}

// AccountDelegationHistory return Account delegation history
// Obtain information about the delegation of a specific account.
type AccountDelegationHistory struct {
	ActiveEpoch int32  `json:"active_epoch,omitempty"`
	TXHash      string `json:"tx_hash,omitempty"`
	Amount      string `json:"amount,omitempty"`
	PoolID      string `json:"pool_id,omitempty"`
}

// AccountRegistrationHistory return Account registration history
// Obtain information about the registrations and deregistrations of a specific account.
type AccountRegistrationHistory struct {
	TXHash string `json:"tx_hash,omitempty"`
	Action string `json:"action,omitempty"`
}

// AccountWithdrawalHistory return Account withdrawal history
// Obtain information about the withdrawals of a specific account.
type AccountWithdrawalHistory struct {
	TXHash string `json:"tx_hash,omitempty"`
	Amount string `json:"amount,omitempty"`
}

// AccountMIRHistory return Account MIR history
// Obtain information about the MIRs of a specific account.
type AccountMIRHistory struct {
	TXHash string `json:"tx_hash,omitempty"`
	Amount string `json:"amount,omitempty"`
}

// AccountAssociatedAddress return Account associated addresses
// Obtain information about the addresses of a specific account.
type AccountAssociatedAddress struct {
	Address string `json:"address,omitempty"`
}

// AccountAssociatedAsset return Assets associated with the account addresses
// Obtain information about assets associated with addresses of a specific account.
// Be careful, as an account could be part of a mangled address and does not necessarily mean the addresses are owned by user as the account.
type AccountAssociatedAsset struct {
	Unit     string `json:"unit,omitempty"`
	Quantity string `json:"quantity,omitempty"`
}

type AccountHistoryResult struct {
	Res []AccountHistory
	Err error
}

type AccountRewardHisResult struct {
	Res []AccountRewardsHistory
	Err error
}

type AccDelegationHistoryResult struct {
	Res []AccountDelegationHistory
	Err error
}

type AccountRegistrationHistoryResult struct {
	Res []AccountRegistrationHistory
	Err error
}

type AccountMIRHistoryResult struct {
	Res []AccountMIRHistory
	Err error
}

type AccountWithdrawalHistoryResult struct {
	Res []AccountWithdrawalHistory
	Err error
}

type AccountAssociatedAddressesAll struct {
	Res []AccountAssociatedAddress
	Err error
}

type AccountAssociatedAssetsAll struct {
	Res []AccountAssociatedAsset
	Err error
}

// Account returns the content of a requested Account by the specific stake account.
func (c *apiClient) Account(ctx context.Context, stakeAddress string) (acc Account, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourceAccount, stakeAddress))
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

	if err = json.NewDecoder(res.Body).Decode(&acc); err != nil {
		return
	}
	return acc, nil
}

// AccountRewardsHistory returns the content of a requested Account by the specific stake account.
// Obtain information about the reward history.
func (c *apiClient) AccountRewardsHistory(ctx context.Context, stakeAddress string, query APIPagingParams) (ah []AccountRewardsHistory, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountRewardsHistory))
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

	err = json.NewDecoder(res.Body).Decode(&ah)
	if err != nil {
		return ah, err
	}
	return ah, nil
}

func (c *apiClient) AccountRewardsHistoryAll(ctx context.Context, stakeAddress string) <-chan AccountRewardHisResult {
	ch := make(chan AccountRewardHisResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan AccountRewardHisResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				his, err := c.AccountRewardsHistory(j.ctx, stakeAddress, j.query)
				if len(his) != j.query.Count || err != nil {
					quit <- true
				}
				res := AccountRewardHisResult{Res: his, Err: err}
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

// AccountHistory returns the content of a requested Account by the specific stake account.
// Obtain information about the history.
func (c *apiClient) AccountHistory(ctx context.Context, stakeAddress string, query APIPagingParams) (ah []AccountHistory, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountHistory))
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

	if err = json.NewDecoder(res.Body).Decode(&ah); err != nil {
		return []AccountHistory{}, err
	}
	return ah, nil
}

func (c *apiClient) AccountHistoryAll(ctx context.Context, address string) <-chan AccountHistoryResult {
	ch := make(chan AccountHistoryResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan AccountHistoryResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				his, err := c.AccountHistory(j.ctx, address, j.query)
				if len(his) != j.query.Count || err != nil {
					quit <- true
				}
				res := AccountHistoryResult{Res: his, Err: err}
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

// AccountDelegationHistory returns the content of a requested Account by the specific stake account.
// Obtain information about the delegations.
func (c *apiClient) AccountDelegationHistory(ctx context.Context, stakeAddress string, query APIPagingParams) (adh []AccountDelegationHistory, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountDelegationHistory))
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

	if err = json.NewDecoder(res.Body).Decode(&adh); err != nil {
		return []AccountDelegationHistory{}, err
	}
	return adh, nil
}

func (c *apiClient) AccountDelegationHistoryAll(ctx context.Context, stakeAddress string) <-chan AccDelegationHistoryResult {
	ch := make(chan AccDelegationHistoryResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan AccDelegationHistoryResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				his, err := c.AccountDelegationHistory(j.ctx, stakeAddress, j.query)
				if len(his) != j.query.Count || err != nil {
					quit <- true
				}
				res := AccDelegationHistoryResult{Res: his, Err: err}
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

// AccountRegistrationHistory returns the content of a requested Account by the specific stake account.
// Obtain information about the Registrations.
func (c *apiClient) AccountRegistrationHistory(ctx context.Context, stakeAddress string, query APIPagingParams) (arh []AccountRegistrationHistory, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountRegistrationHistory))
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

	if err = json.NewDecoder(res.Body).Decode(&arh); err != nil {
		return
	}
	return arh, nil
}

func (c *apiClient) AccountRegistrationHistoryAll(ctx context.Context, stakeAddress string) <-chan AccountRegistrationHistoryResult {
	ch := make(chan AccountRegistrationHistoryResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan AccountRegistrationHistoryResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				his, err := c.AccountRegistrationHistory(j.ctx, stakeAddress, j.query)
				if len(his) != j.query.Count || err != nil {
					quit <- true
				}
				res := AccountRegistrationHistoryResult{Res: his, Err: err}
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

// AccountWithdrawalHistory returns the content of a requested Account by the specific stake account.
// Obtain information about the Withdrawals.
func (c *apiClient) AccountWithdrawalHistory(ctx context.Context, stakeAddress string, query APIPagingParams) (awh []AccountWithdrawalHistory, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountWithdrawalHistory))
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

	err = json.NewDecoder(res.Body).Decode(&awh)
	if err != nil {
		return awh, err
	}
	return awh, nil
}

func (c *apiClient) AccountWithdrawalHistoryAll(ctx context.Context, stakeAddress string) <-chan AccountWithdrawalHistoryResult {
	ch := make(chan AccountWithdrawalHistoryResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan AccountWithdrawalHistoryResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				his, err := c.AccountWithdrawalHistory(j.ctx, stakeAddress, j.query)
				if len(his) != j.query.Count || err != nil {
					quit <- true
				}
				res := AccountWithdrawalHistoryResult{Res: his, Err: err}
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

// AccountMIRHistory returns the content of a requested Account by the specific stake account.
// Obtain information about the MIRs.
func (c *apiClient) AccountMIRHistory(ctx context.Context, stakeAddress string, query APIPagingParams) (amh []AccountMIRHistory, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountMIRHistory))
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

	if err = json.NewDecoder(res.Body).Decode(&amh); err != nil {
		return amh, err
	}
	return amh, nil
}

func (c *apiClient) AccountMIRHistoryAll(ctx context.Context, stakeAddress string) <-chan AccountMIRHistoryResult {
	ch := make(chan AccountMIRHistoryResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan AccountMIRHistoryResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				his, err := c.AccountMIRHistory(j.ctx, stakeAddress, j.query)
				if len(his) != j.query.Count || err != nil {
					quit <- true
				}
				res := AccountMIRHistoryResult{Res: his, Err: err}
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

// AccountAssociatedAddresses returns the content of a requested Account by the specific stake account.
// Obtain information about the addresses of a specific account.
func (c *apiClient) AccountAssociatedAddresses(ctx context.Context, stakeAddress string, query APIPagingParams) (aas []AccountAssociatedAddress, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountAssociatedAddress))
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

	if err = json.NewDecoder(res.Body).Decode(&aas); err != nil {
		return []AccountAssociatedAddress{}, err
	}
	return aas, nil
}

func (c *apiClient) AccountAssociatedAddressesAll(ctx context.Context, stakeAddress string) <-chan AccountAssociatedAddressesAll {
	ch := make(chan AccountAssociatedAddressesAll, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan AccountAssociatedAddressesAll, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				addrs, err := c.AccountAssociatedAddresses(j.ctx, stakeAddress, j.query)
				if len(addrs) != j.query.Count || err != nil {
					quit <- true
				}
				res := AccountAssociatedAddressesAll{Res: addrs, Err: err}
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

// AccountAssociatedAssets returns the content of a requested Account by the specific stake account.
// Obtain information about the addresses of a specific account.
func (c *apiClient) AccountAssociatedAssets(ctx context.Context, stakeAddress string, query APIPagingParams) (aaa []AccountAssociatedAsset, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountAddressWithAssetsAssociated))
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

	err = json.NewDecoder(res.Body).Decode(&aaa)
	if err != nil {
		return aaa, err
	}
	return aaa, nil
}

func (c *apiClient) AccountAssociatedAssetsAll(ctx context.Context, stakeAddress string) <-chan AccountAssociatedAssetsAll {
	ch := make(chan AccountAssociatedAssetsAll, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan AccountAssociatedAssetsAll, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				as, err := c.AccountAssociatedAssets(j.ctx, stakeAddress, j.query)
				if len(as) != j.query.Count || err != nil {
					quit <- true
				}
				res := AccountAssociatedAssetsAll{Res: as, Err: err}
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
