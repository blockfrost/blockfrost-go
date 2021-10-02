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
	resourceAddresses    = "addresses"
	resourceTotal        = "total"
	resourceTransactions = "transactions"
	resourceUTXOs        = "utxos"
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
	TxIndex     int    `json:"tx_index,omitempty"`
	BlockHeight int    `json:"block_height,omitempty"`
}

type AddressUTXO struct {
	TxHash      string          `json:"tx_hash,omitempty"`
	OutputIndex int             `json:"output_index,omitempty"`
	Amount      []AddressAmount `json:"amount,omitempty"`
	Block       string          `json:"block,omitempty"`
	DataHash    string          `json:"data_hash,omitempty"`
}

type AddressTxResult struct {
	Res []AddressTransactions
	Err error
}

type AddressUTXOResult struct {
	Res []AddressUTXO
	Err error
}

// Address ret
func (c *apiClient) Address(ctx context.Context, address string) (addr Address, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourceAddresses, address))
	if err != nil {
		return Address{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return Address{}, err
	}

	res, err := c.handleRequest(req)
	if err != nil {
		return Address{}, nil
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&addr); err != nil {
		return Address{}, err
	}
	return addr, nil
}

func (c *apiClient) AddressTransactions(ctx context.Context, address string, query APIPagingParams) ([]AddressTransactions, error) {
	var txs []AddressTransactions
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAddresses, address, resourceTransactions))
	if err != nil {
		return txs, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return txs, err
	}
	v := req.URL.Query()
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()
	res, err := c.handleRequest(req)
	if err != nil {
		return txs, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&txs); err != nil {
		return txs, err
	}
	return txs, nil
}

func (c *apiClient) AddressTransactionsAll(ctx context.Context, address string) <-chan AddressTxResult {
	ch := make(chan AddressTxResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan AddressTxResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				atx, err := c.AddressTransactions(j.ctx, address, j.query)
				if len(atx) != j.query.Count || err != nil {
					quit <- true
				}
				res := AddressTxResult{Res: atx, Err: err}
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

func (c *apiClient) AddressDetails(ctx context.Context, address string) (AddressDetails, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAddresses, address, resourceTotal))
	if err != nil {
		return AddressDetails{}, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return AddressDetails{}, err
	}
	res, err := c.handleRequest(req)
	if err != nil {
		return AddressDetails{}, err
	}
	defer res.Body.Close()

	det := AddressDetails{}
	if err = json.NewDecoder(res.Body).Decode(&det); err != nil {
		return det, err
	}
	return det, nil
}

func (c *apiClient) AddressUTXOs(ctx context.Context, address string, query APIPagingParams) ([]AddressUTXO, error) {
	var utxos []AddressUTXO
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAddresses, address, resourceUTXOs))
	if err != nil {
		return utxos, err
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return utxos, err
	}
	v := req.URL.Query()
	query.From = ""
	query.To = ""
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()

	res, err := c.handleRequest(req)
	if err != nil {
		return utxos, err
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&utxos); err != nil {
		return utxos, err
	}
	return utxos, nil
}

func (c *apiClient) AddressUTXOsAll(ctx context.Context, address string) <-chan AddressUTXOResult {
	ch := make(chan AddressUTXOResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan AddressUTXOResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				autxo, err := c.AddressUTXOs(j.ctx, address, j.query)
				if len(autxo) != j.query.Count || err != nil {
					quit <- true
				}
				res := AddressUTXOResult{Res: autxo, Err: err}
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
