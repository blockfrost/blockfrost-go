package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// BlocksLatest returns the latest block available to the backends,
// also known as the tip of the blockchain.
func (c *apiClient) BlocksLatest(ctx context.Context) (Block, error) {
	requestUrl, err := url.Parse((fmt.Sprintf("%s/%s", c.server, resourceBlocksLatest)))
	if err != nil {
		return Block{}, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return Block{}, err
	}
	req.Header.Add("project_id", c.projectId)
	req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return Block{}, err
	}
	defer res.Body.Close()

	block := Block{}
	if res.StatusCode != http.StatusOK {
		return Block{}, handleAPIErrorResponse(res)
	}

	err = json.NewDecoder(res.Body).Decode(&block)
	if err != nil {
		return Block{}, err
	}
	return block, nil
}
