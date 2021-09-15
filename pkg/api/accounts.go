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

func (c *client) Account(ctx context.Context, stakeAddr string) (Account, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s", c.url, accountsByAddressPath, stakeAddr),
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
