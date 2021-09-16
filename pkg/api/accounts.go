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
		fmt.Sprintf("%s/%s/%s/%s", c.url, accountsPath, stakeAddr, accountsRewardsHistPath),
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
