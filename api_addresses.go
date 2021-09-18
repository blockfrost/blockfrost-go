package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	resourceAddresses = "addresses"
	resourceTotal     = "total"
)

type AddressAmount struct {
	Unit     string `json:"unit,omitempty"`
	Quantity string `json:"quantity,omitempty"`
}

type Address struct {
	Address      string          `json:"address,omitempty"`
	Amount       []AddressAmount `json:"amount,omitempty"`
	StakeAddress string          `json:"stake_address,omitempty"`
	Type         string          `json:"type,omitempty"`
	Script       bool            `json:"script,omitempty"`
}

type AddressDetails struct {
	Address     string          `json:"address,omitempty"`
	ReceivedSum []AddressAmount `json:"received_sum,omitempty"`
	SentSum     []AddressAmount `json:"sent_sum,omitempty"`
	TxCount     int             `json:"tx_count"`
}

type AddressTransactions struct {
	TxHash      string `json:"tx_hash,omitempty"`
	TxIndex     string `json:"tx_index,omitempty"`
	BlockHeight string `json:"block_height,omitempty"`
}

type AddressUTXO struct {
}

// Address ret
func (c *apiClient) Address(ctx context.Context, address string) (Address, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourceAddresses, address))
	if err != nil {
		return Address{}, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return Address{}, err
	}
	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return Address{}, nil
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Address{}, handleAPIErrorResponse(res)
	}

	txs := Address{}
	if err := json.NewDecoder(res.Body).Decode(&txs); err != nil {
		return Address{}, err
	}
	return txs, nil
}

func (c *apiClient) AddressTransactions(ctx context.Context, address string, query APIPagingParams) ([]AddressTransactions, error) {
	var txs []AddressTransactions
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, address, resourceAddresses))
	if err != nil {
		return txs, err
	}
	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return txs, err
	}
	v := req.URL.Query()
	v = formatParams(v, APIPagingParams{})
	req.URL.RawQuery = v.Encode()
	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return txs, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return txs, handleAPIErrorResponse(res)
	}

	if err := json.NewDecoder(res.Body).Decode(&txs); err != nil {
		return txs, err
	}
	return txs, nil
}

func (c *apiClient) AddressDetails(ctx context.Context, address string) (AddressDetails, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAddresses, address, resourceTotal))
	if err != nil {
		return AddressDetails{}, err
	}
	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return AddressDetails{}, err
	}
	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return AddressDetails{}, err
	}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return AddressDetails{}, err
	}

	det := AddressDetails{}
	if err = json.NewDecoder(res.Body).Decode(&det); err != nil {
		return det, err
	}
	return det, nil
}
