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
	Unit     string `json:"unit,"`
	Quantity string `json:"quantity,"`
}

type Address struct {
	Address      string          `json:"address,"`
	Amount       []AddressAmount `json:"amount,"`
	StakeAddress string          `json:"stake_address,"`
	Type         string          `json:"type,"`
	Script       bool            `json:"script,"`
}

type AddressDetails struct {
	Address     string          `json:"address,"`
	ReceivedSum []AddressAmount `json:"received_sum,"`
	SentSum     []AddressAmount `json:"sent_sum,"`
	TxCount     int             `json:"tx_count"`
}

type AddressTransactions struct {
	TxHash      string `json:"tx_hash,"`
	TxIndex     int    `json:"tx_index,"`
	BlockHeight int    `json:"block_height,"`
}

type AddressUTXO struct {
	TxHash      string          `json:"tx_hash,"`
	OutputIndex int             `json:"output_index,"`
	Amount      []AddressAmount `json:"amount,"`
	Block       string          `json:"block,"`
	DataHash    string          `json:"data_hash,"`
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

	if err = json.NewDecoder(res.Body).Decode(&addr); err != nil {
		return
	}
	return addr, nil
}

func (c *apiClient) AddressTransactions(ctx context.Context, address string, query APIQueryParams) (txs []AddressTransactions, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAddresses, address, resourceTransactions))
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

	if err = json.NewDecoder(res.Body).Decode(&txs); err != nil {
		return
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
				jobs <- methodOptions{ctx: ctx, query: APIQueryParams{Count: 100, Page: i}}
			}
		}

		wg.Wait()
	}()
	return ch
}

func (c *apiClient) AddressDetails(ctx context.Context, address string) (ad AddressDetails, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAddresses, address, resourceTotal))
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

	if err = json.NewDecoder(res.Body).Decode(&ad); err != nil {
		return
	}
	return ad, nil
}

func (c *apiClient) AddressUTXOs(ctx context.Context, address string, query APIQueryParams) (utxos []AddressUTXO, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceAddresses, address, resourceUTXOs))
	if err != nil {
		return
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}
	v := req.URL.Query()
	query.From = ""
	query.To = ""
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()

	res, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&utxos); err != nil {
		return
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
				jobs <- methodOptions{ctx: ctx, query: APIQueryParams{Count: 100, Page: i}}
			}
		}

		wg.Wait()
	}()
	return ch
}
