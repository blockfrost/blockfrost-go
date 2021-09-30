package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

// Account returns the content of a requested Account by the specific stake account.
func (c *apiClient) Account(ctx context.Context, stakeAddress string) (Account, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourceAccount, stakeAddress))
	if err != nil {
		return Account{}, err
	}
	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return Account{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return Account{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Account{}, handleAPIErrorResponse(res)
	}
	account := Account{}
	err = json.NewDecoder(res.Body).Decode(&account)
	if err != nil {
		return Account{}, err
	}
	return account, nil
}

// AccountRewardsHistory returns the content of a requested Account by the specific stake account.
// Obtain information about the reward history.
func (c *apiClient) AccountRewardsHistory(
	ctx context.Context,
	stakeAddress string,
	query APIPagingParams,
) ([]AccountRewardsHistory, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountRewardsHistory))
	if err != nil {
		return []AccountRewardsHistory{}, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []AccountRewardsHistory{}, err
	}

	v := req.URL.Query()
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()
	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []AccountRewardsHistory{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []AccountRewardsHistory{}, handleAPIErrorResponse(res)
	}
	accounts := []AccountRewardsHistory{}
	err = json.NewDecoder(res.Body).Decode(&accounts)
	if err != nil {
		return []AccountRewardsHistory{}, err
	}
	return accounts, nil
}

// AccountHistory returns the content of a requested Account by the specific stake account.
// Obtain information about the history.
func (c *apiClient) AccountHistory(
	ctx context.Context,
	stakeAddress string,
	query APIPagingParams,
) ([]AccountHistory, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountHistory))
	if err != nil {
		return []AccountHistory{}, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []AccountHistory{}, err
	}

	v := req.URL.Query()
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()
	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []AccountHistory{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []AccountHistory{}, handleAPIErrorResponse(res)
	}
	accounts := []AccountHistory{}
	err = json.NewDecoder(res.Body).Decode(&accounts)
	if err != nil {
		return []AccountHistory{}, err
	}
	return accounts, nil
}

// AccountDelegationHistory returns the content of a requested Account by the specific stake account.
// Obtain information about the delegations.
func (c *apiClient) AccountDelegationHistory(
	ctx context.Context,
	stakeAddress string,
	query APIPagingParams,
) ([]AccountDelegationHistory, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountDelegationHistory))
	if err != nil {
		return []AccountDelegationHistory{}, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []AccountDelegationHistory{}, err
	}

	v := req.URL.Query()
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()
	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []AccountDelegationHistory{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []AccountDelegationHistory{}, handleAPIErrorResponse(res)
	}
	accounts := []AccountDelegationHistory{}
	err = json.NewDecoder(res.Body).Decode(&accounts)
	if err != nil {
		return []AccountDelegationHistory{}, err
	}
	return accounts, nil
}

// AccountRegistrationHistory returns the content of a requested Account by the specific stake account.
// Obtain information about the Registrations.
func (c *apiClient) AccountRegistrationHistory(
	ctx context.Context,
	stakeAddress string,
	query APIPagingParams,
) ([]AccountRegistrationHistory, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountRegistrationHistory))
	if err != nil {
		return []AccountRegistrationHistory{}, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []AccountRegistrationHistory{}, err
	}

	v := req.URL.Query()
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()
	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []AccountRegistrationHistory{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []AccountRegistrationHistory{}, handleAPIErrorResponse(res)
	}
	accounts := []AccountRegistrationHistory{}
	err = json.NewDecoder(res.Body).Decode(&accounts)
	if err != nil {
		return []AccountRegistrationHistory{}, err
	}
	return accounts, nil
}

// AccountWithdrawalHistory returns the content of a requested Account by the specific stake account.
// Obtain information about the Withdrawals.
func (c *apiClient) AccountWithdrawalHistory(
	ctx context.Context,
	stakeAddress string,
	query APIPagingParams,
) ([]AccountWithdrawalHistory, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountWithdrawalHistory))
	if err != nil {
		return []AccountWithdrawalHistory{}, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []AccountWithdrawalHistory{}, err
	}

	v := req.URL.Query()
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()
	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []AccountWithdrawalHistory{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []AccountWithdrawalHistory{}, handleAPIErrorResponse(res)
	}
	accounts := []AccountWithdrawalHistory{}
	err = json.NewDecoder(res.Body).Decode(&accounts)
	if err != nil {
		return []AccountWithdrawalHistory{}, err
	}
	return accounts, nil
}

// AccountMIRHistory returns the content of a requested Account by the specific stake account.
// Obtain information about the MIRs.
func (c *apiClient) AccountMIRHistory(
	ctx context.Context,
	stakeAddress string,
	query APIPagingParams,
) ([]AccountMIRHistory, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountMIRHistory))
	if err != nil {
		return []AccountMIRHistory{}, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []AccountMIRHistory{}, err
	}

	v := req.URL.Query()
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()
	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []AccountMIRHistory{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []AccountMIRHistory{}, handleAPIErrorResponse(res)
	}
	accounts := []AccountMIRHistory{}
	err = json.NewDecoder(res.Body).Decode(&accounts)
	if err != nil {
		return []AccountMIRHistory{}, err
	}
	return accounts, nil
}

// AccountAssociatedAddresses returns the content of a requested Account by the specific stake account.
// Obtain information about the addresses of a specific account.
func (c *apiClient) AccountAssociatedAddresses(
	ctx context.Context,
	stakeAddress string,
	query APIPagingParams,
) ([]AccountAssociatedAddress, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountAssociatedAddress))
	if err != nil {
		return []AccountAssociatedAddress{}, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []AccountAssociatedAddress{}, err
	}
	v := req.URL.Query()
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()
	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []AccountAssociatedAddress{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []AccountAssociatedAddress{}, handleAPIErrorResponse(res)
	}
	accounts := []AccountAssociatedAddress{}
	err = json.NewDecoder(res.Body).Decode(&accounts)
	if err != nil {
		return []AccountAssociatedAddress{}, err
	}
	return accounts, nil
}

// AccountAssociatedAssets returns the content of a requested Account by the specific stake account.
// Obtain information about the addresses of a specific account.
func (c *apiClient) AccountAssociatedAssets(
	ctx context.Context,
	stakeAddress string,
	query APIPagingParams,
) ([]AccountAssociatedAsset, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountAddressWithAssetsAssociated))
	if err != nil {
		return []AccountAssociatedAsset{}, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []AccountAssociatedAsset{}, err
	}
	v := req.URL.Query()
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()
	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []AccountAssociatedAsset{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []AccountAssociatedAsset{}, handleAPIErrorResponse(res)
	}
	accounts := []AccountAssociatedAsset{}
	err = json.NewDecoder(res.Body).Decode(&accounts)
	if err != nil {
		return []AccountAssociatedAsset{}, err
	}
	return accounts, nil
}
