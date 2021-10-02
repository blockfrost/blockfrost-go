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
	resourceScripts   = "scripts"
	resourceRedeemers = "redeemers"
)

type Script struct {
	ScriptHash     string `json:"script_hash,omitempty"`
	Type           string `json:"type,omitempty"`
	SerialisedSize int    `json:"serialised_size,omitempty"`
}

type ScriptRedeemer struct {
	TxHash    string `json:"tx_hash,omitempty"`
	TxIndex   int    `json:"tx_index,omitempty"`
	Purpose   string `json:"purpose,omitempty"`
	UnitMem   string `json:"unit_mem,omitempty"`
	UnitSteps string `json:"unit_steps,omitempty"`
	Fee       string `json:"fee,omitempty"`
}

type methodOptions struct {
	ctx   context.Context
	query APIPagingParams
}

type ScriptAllResult struct {
	Scripts []Script
	Err     error
}

type ScriptRedeemerResult struct {
	ScriptRedeemers []ScriptRedeemer
	Err             error
}

func (c *apiClient) Scripts(ctx context.Context, query APIPagingParams) (scripts []Script, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceScripts))
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

	if res.StatusCode != http.StatusOK {
		return scripts, handleAPIErrorResponse(res)
	}

	if err := json.NewDecoder(res.Body).Decode(&scripts); err != nil {
		return scripts, err
	}
	return scripts, nil

}

func (c *apiClient) ScriptsAll(ctx context.Context) <-chan ScriptAllResult {
	ch := make(chan ScriptAllResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan ScriptAllResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				sc, err := c.Scripts(j.ctx, j.query)
				if len(sc) != j.query.Count || err != nil {
					quit <- true
				}
				res := ScriptAllResult{Scripts: sc, Err: err}
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

func (c *apiClient) Script(ctx context.Context, address string) (script Script, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourceScripts, address))
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
	if res.StatusCode != http.StatusOK {
		return script, handleAPIErrorResponse(res)
	}

	if err = json.NewDecoder(res.Body).Decode(&script); err != nil {
		return script, err
	}
	return script, nil
}

func (c *apiClient) ScriptRedeemers(ctx context.Context, address string, query APIPagingParams) (sr []ScriptRedeemer, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceScripts, address, resourceRedeemers))
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

	if res.StatusCode != http.StatusOK {
		return sr, handleAPIErrorResponse(res)
	}

	if err := json.NewDecoder(res.Body).Decode(&sr); err != nil {
		return sr, err
	}
	return sr, nil
}

func (c *apiClient) ScriptRedeemersAll(ctx context.Context, address string) <-chan ScriptRedeemerResult {
	ch := make(chan ScriptRedeemerResult, c.routines)
	jobs := make(chan methodOptions, c.routines)
	quit := make(chan bool, c.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < c.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan ScriptRedeemerResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				sr, err := c.ScriptRedeemers(j.ctx, address, j.query)
				if len(sr) != j.query.Count || err != nil {
					quit <- true
				}
				res := ScriptRedeemerResult{ScriptRedeemers: sr, Err: err}
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
