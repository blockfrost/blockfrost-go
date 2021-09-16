package blockfrostgo

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

// Cardano Endpoints

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
type AccountRewardsHist struct {
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

// AccountDelegation return Account delegation history
// Obtain information about the delegation of a specific account.
type AccountDelegation struct {
	ActiveEpoch int32  `json:"active_epoch,omitempty"`
	TXHash      string `json:"tx_hash,omitempty"`
	Amount      string `json:"amount,omitempty"`
	PoolID      string `json:"pool_id,omitempty"`
}

// AccountRegistration return Account registration history
// Obtain information about the registrations and deregistrations of a specific account.
type AccountRegistration struct {
	TXHash string `json:"tx_hash,omitempty"`
	Action string `json:"action,omitempty"`
}

// AccountWithdrawal return Account withdrawal history
// Obtain information about the withdrawals of a specific account.
type AccountWithdrawal struct {
	TXHash string `json:"tx_hash,omitempty"`
	Amount string `json:"amount,omitempty"`
}

// AccountMIR return Account MIR history
// Obtain information about the MIRs of a specific account.
type AccountMIR struct {
	TXHash string `json:"tx_hash,omitempty"`
	Amount string `json:"amount,omitempty"`
}

// AccountAssociated return Account associated addresses
// Obtain information about the addresses of a specific account.
type AccountAssociated struct {
	Address string `json:"address,omitempty"`
}

// AccountAssets return Assets associated with the account addresses
// Obtain information about assets associated with addresses of a specific account.
// Be careful, as an account could be part of a mangled address and does not necessarily mean the addresses are owned by user as the account.
type AccountAssets struct {
	Unit     string `json:"unit,omitempty"`
	Quantity string `json:"quantity,omitempty"`
}

func (c *client) Account(ctx context.Context, stakeAddr string) (Account, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s", c.url, accountsPath, stakeAddr),
	)
	if err != nil {
		return Account{}, err
	}

	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		requestUrl.String(),
		nil,
	)
	if err != nil {
		return Account{}, err
	}
	if status != http.StatusOK {
		return Account{}, fmt.Errorf("unexpected response status %d: %q", status, res)
	}
	result := Account{}
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return Account{}, fmt.Errorf("decoding error for data %s: %v", res, err)
	}
	return result, nil
}

func (c *client) AccountRewards(
	ctx context.Context,
	stakeAddr string,
	query QueryParmasAPI,
) ([]AccountRewardsHist, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.url, accountsPath, stakeAddr, "rewards"),
	)
	if err != nil {
		return []AccountRewardsHist{}, err
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

	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		requestUrl.String(),
		nil,
	)
	if err != nil {
		return []AccountRewardsHist{}, err
	}
	if status != http.StatusOK {
		return []AccountRewardsHist{}, fmt.Errorf("unexpected response status %d: %q", status, res)
	}
	result := []AccountRewardsHist{}
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return []AccountRewardsHist{}, fmt.Errorf("decoding error for data %s: %v", res, err)
	}
	return result, nil
}

func (c *client) AccountHistory(
	ctx context.Context,
	stakeAddr string,
	query QueryParmasAPI,
) ([]AccountHistory, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.url, accountsPath, stakeAddr, "history"),
	)
	if err != nil {
		return []AccountHistory{}, err
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

	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		requestUrl.String(),
		nil,
	)
	if err != nil {
		return []AccountHistory{}, err
	}
	if status != http.StatusOK {
		return []AccountHistory{}, fmt.Errorf("unexpected response status %d: %q", status, res)
	}
	result := []AccountHistory{}
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return []AccountHistory{}, fmt.Errorf("decoding error for data %s: %v", res, err)
	}
	return result, nil
}

func (c *client) AccountDelegations(
	ctx context.Context,
	stakeAddr string,
	query QueryParmasAPI,
) ([]AccountDelegation, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.url, accountsPath, stakeAddr, "delegations"),
	)
	if err != nil {
		return []AccountDelegation{}, err
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

	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		requestUrl.String(),
		nil,
	)
	if err != nil {
		return []AccountDelegation{}, err
	}
	if status != http.StatusOK {
		return []AccountDelegation{}, fmt.Errorf("unexpected response status %d: %q", status, res)
	}
	result := []AccountDelegation{}
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return []AccountDelegation{}, fmt.Errorf("decoding error for data %s: %v", res, err)
	}
	return result, nil
}

func (c *client) AccountRegistrations(
	ctx context.Context,
	stakeAddr string,
	query QueryParmasAPI,
) ([]AccountRegistration, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.url, accountsPath, stakeAddr, "registrations"),
	)
	if err != nil {
		return []AccountRegistration{}, err
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

	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		requestUrl.String(),
		nil,
	)
	if err != nil {
		return []AccountRegistration{}, err
	}
	if status != http.StatusOK {
		return []AccountRegistration{}, fmt.Errorf("unexpected response status %d: %q", status, res)
	}
	result := []AccountRegistration{}
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return []AccountRegistration{}, fmt.Errorf("decoding error for data %s: %v", res, err)
	}
	return result, nil
}

func (c *client) AccountWithdrawals(
	ctx context.Context,
	stakeAddr string,
	query QueryParmasAPI,
) ([]AccountWithdrawal, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.url, accountsPath, stakeAddr, "withdrawals"),
	)
	if err != nil {
		return []AccountWithdrawal{}, err
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

	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		requestUrl.String(),
		nil,
	)
	if err != nil {
		return []AccountWithdrawal{}, err
	}
	if status != http.StatusOK {
		return []AccountWithdrawal{}, fmt.Errorf("unexpected response status %d: %q", status, res)
	}
	result := []AccountWithdrawal{}
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return []AccountWithdrawal{}, fmt.Errorf("decoding error for data %s: %v", res, err)
	}
	return result, nil
}

func (c *client) AccountMIRHistory(
	ctx context.Context,
	stakeAddr string,
	query QueryParmasAPI,
) ([]AccountMIR, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.url, accountsPath, stakeAddr, "mirs"),
	)
	if err != nil {
		return []AccountMIR{}, err
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

	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		requestUrl.String(),
		nil,
	)
	if err != nil {
		return []AccountMIR{}, err
	}
	if status != http.StatusOK {
		return []AccountMIR{}, fmt.Errorf("unexpected response status %d: %q", status, res)
	}
	result := []AccountMIR{}
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return []AccountMIR{}, fmt.Errorf("decoding error for data %s: %v", res, err)
	}
	return result, nil
}
func (c *client) AccountAssociatedAddress(
	ctx context.Context,
	stakeAddr string,
	query QueryParmasAPI,
) ([]AccountAssociated, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.url, accountsPath, stakeAddr, "addresses"),
	)
	if err != nil {
		return []AccountAssociated{}, err
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

	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		requestUrl.String(),
		nil,
	)
	if err != nil {
		return []AccountAssociated{}, err
	}
	if status != http.StatusOK {
		return []AccountAssociated{}, fmt.Errorf("unexpected response status %d: %q", status, res)
	}
	result := []AccountAssociated{}
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return []AccountAssociated{}, fmt.Errorf("decoding error for data %s: %v", res, err)
	}
	return result, nil
}

func (c *client) AccountAssetsWithAddress(
	ctx context.Context,
	stakeAddr string,
	query QueryParmasAPI,
) ([]AccountAssets, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.url, accountsPath, stakeAddr, "addresses/assets"),
	)
	if err != nil {
		return []AccountAssets{}, err
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

	status, res, err := c.apiCall(
		ctx,
		http.MethodGet,
		requestUrl.String(),
		nil,
	)
	if err != nil {
		return []AccountAssets{}, err
	}
	if status != http.StatusOK {
		return []AccountAssets{}, fmt.Errorf("unexpected response status %d: %q", status, res)
	}
	result := []AccountAssets{}
	err = json.NewDecoder(strings.NewReader(res)).Decode(&result)
	if err != nil {
		return []AccountAssets{}, fmt.Errorf("decoding error for data %s: %v", res, err)
	}
	return result, nil
}
