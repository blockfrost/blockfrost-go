package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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
	req.Header.Add("project_id", c.projectId)

	res, err := c.client.Do(req)
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

func (c *apiClient) Script(ctx context.Context, address string) (script Script, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourceScripts, address))
	if err != nil {
		return
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}

	req.Header.Add("project_id", c.projectId)

	res, err := c.client.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return script, handleAPIErrorResponse(res)
	}

	if err := json.NewDecoder(res.Body).Decode(&script); err != nil {
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

	req.Header.Add("project_id", c.projectId)
	res, err := c.client.Do(req)
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
