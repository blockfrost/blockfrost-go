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
	resourceNutLink = "nutlink"
	resourceTickers = "tickers"
)

type NutlinkAddressMeta struct {
	Ticker      string `json:"ticker,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	HomePage    string `json:"homepage,omitempty"`
	Address     string `json:"address,omitempty"`
}

type NutlinkAddress struct {
	Address      string             `json:"address,omitempty"`
	MetadataUrl  string             `json:"metadata_url,omitempty"`
	MetadataHash string             `json:"metadata_hash,omitempty"`
	Metadata     NutlinkAddressMeta `json:"metadata,omitempty"`
}

type Ticker struct {
	Name        string `json:"name,omitempty"`
	Count       int    `json:"count,omitempty"`
	LatestBlock int    `json:"latest_block,omitempty"`
}

type TickerRecord struct {
	TxHash      string `json:"tx_hash,omitempty"`
	BlockHeight int    `json:"block_height,omitempty"`
	TxIndex     int    `json:"tx_index,omitempty"`
}

type TickerResult struct {
	Res []Ticker
	Err error
}

type TickerRecordResult struct {
	Res []TickerRecord
	Err error
}

// Nutlink returns list metadata about specific address.
func (c *apiClient) Nutlink(ctx context.Context, address string) (nu NutlinkAddress, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourceNutLink, address))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}
	res, err := c.handleRequest(req)
	if err != nil {
		return nu, err
	}
	defer res.Body.Close()

	if err := json.NewDecoder(res.Body).Decode(&nu); err != nil {
		return nu, err
	}
	return nu, nil
}

// Tickers returns paginated list tickers for a specific metadata oracle.
func (c *apiClient) Tickers(ctx context.Context, address string, query APIQueryParams) (ti []Ticker, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceNutLink, address, resourceTickers))
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

	if err := json.NewDecoder(res.Body).Decode(&ti); err != nil {
		return ti, err
	}
	return ti, nil
}

// TickersAll returns all tickers for a specific metadata oracle.
func (c *apiClient) TickersAll(ctx context.Context, address string) <-chan TickerResult {
	ch := make(chan TickerResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan TickerResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				as, err := c.Tickers(j.ctx, address, j.query)
				if len(as) != j.query.Count || err != nil {
					quit <- true
				}
				res := TickerResult{Res: as, Err: err}
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

// TickerRecords returns list of records of a specific ticker.
func (c *apiClient) TickerRecords(ctx context.Context, ticker string, query APIQueryParams) (trs []TickerRecord, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceNutLink, resourceTickers, ticker))
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

	if err = json.NewDecoder(res.Body).Decode(&trs); err != nil {
		return
	}
	return trs, nil
}

// TickerRecordsAll returns list of all records of a specific ticker.
func (c *apiClient) TickerRecordsAll(ctx context.Context, ticker string) <-chan TickerRecordResult {
	ch := make(chan TickerRecordResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan TickerRecordResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				as, err := c.TickerRecords(j.ctx, ticker, j.query)
				if len(as) != j.query.Count || err != nil {
					quit <- true
				}
				res := TickerRecordResult{Res: as, Err: err}
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

// AddressTickeRecords returns list of records of a specific ticker by address.
func (c *apiClient) AddressTickerRecords(ctx context.Context, address string, ticker string, query APIQueryParams) (trs []TickerRecord, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s/%s", c.server, resourceNutLink, address, resourceTickers, ticker))
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

	if err = json.NewDecoder(res.Body).Decode(&trs); err != nil {
		return
	}
	return trs, nil
}

// AddressTickerRecordsAll returns list of all records of a specific ticker by address.
func (c *apiClient) AddressTickerRecordsAll(ctx context.Context, address string, ticker string) <-chan TickerRecordResult {
	ch := make(chan TickerRecordResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan TickerRecordResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				as, err := c.AddressTickerRecords(j.ctx, address, ticker, j.query)
				if len(as) != j.query.Count || err != nil {
					quit <- true
				}
				res := TickerRecordResult{Res: as, Err: err}
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
