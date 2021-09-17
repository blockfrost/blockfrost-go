package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

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
		return []AccountRewardsHistory{}, err
	}

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
		return []AccountHistory{}, err
	}

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
		return []AccountDelegationHistory{}, err
	}

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
		return []AccountRegistrationHistory{}, err
	}

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
		return []AccountWithdrawalHistory{}, err
	}

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
		return []AccountMIRHistory{}, err
	}

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

// AccountAssociatedAddress returns the content of a requested Account by the specific stake account.
// Obtain information about the addresses of a specific account.
func (c *apiClient) AccountAssociatedAddress(
	ctx context.Context,
	stakeAddress string,
	query APIPagingParams,
) ([]AccountAssociatedAddress, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountAssociatedAddress))
	if err != nil {
		return []AccountAssociatedAddress{}, err
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
		return []AccountAssociatedAddress{}, err
	}

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

// AccountAssetsWithAddress returns the content of a requested Account by the specific stake account.
// Obtain information about the addresses of a specific account.
func (c *apiClient) AccountAssetsWithAddress(
	ctx context.Context,
	stakeAddress string,
	query APIPagingParams,
) ([]AccountAssetsWithAddress, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAccount, stakeAddress, resourceAccountAddressWithAssetsAssociated))
	if err != nil {
		return []AccountAssetsWithAddress{}, err
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
		return []AccountAssetsWithAddress{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []AccountAssetsWithAddress{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []AccountAssetsWithAddress{}, handleAPIErrorResponse(res)
	}
	accounts := []AccountAssetsWithAddress{}
	err = json.NewDecoder(res.Body).Decode(&accounts)
	if err != nil {
		return []AccountAssetsWithAddress{}, err
	}
	return accounts, nil
}
