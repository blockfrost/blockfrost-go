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
func (c *apiClient) BlockLatest(ctx context.Context) (Block, error) {
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

	if res.StatusCode != http.StatusOK {
		return Block{}, handleAPIErrorResponse(res)
	}
	block := Block{}
	err = json.NewDecoder(res.Body).Decode(&block)
	if err != nil {
		return Block{}, err
	}
	return block, nil
}

// Block returns the content of a requested block by the hash or block number
func (c *apiClient) Block(ctx context.Context, hashOrNumber string) (Block, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourceBlock, hashOrNumber))
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

	if res.StatusCode != http.StatusOK {
		return Block{}, handleAPIErrorResponse(res)
	}
	block := Block{}
	err = json.NewDecoder(res.Body).Decode(&block)
	if err != nil {
		return Block{}, err
	}
	return block, nil
}

// BlocksNext returns the list of blocks following a specific block
// denoted by the hash or block number
func (c *apiClient) BlocksNext(ctx context.Context, hashorNumber string) ([]Block, error) {
	var blocks []Block
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceBlock, hashorNumber, "next"))
	if err != nil {
		return blocks, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	req.Header.Add("project_id", c.projectId)
	req.WithContext(ctx)
	if err != nil {
		return blocks, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return blocks, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return blocks, handleAPIErrorResponse(res)
	}
	err = json.NewDecoder(res.Body).Decode(&blocks)
	if err != nil {
		return []Block{}, err
	}
	return blocks, nil
}

// BlocksPrevious returns the list of blocks preceding a specific block
// specified by a hash or block number
func (c *apiClient) BlocksPrevious(ctx context.Context, hashorNumber string) ([]Block, error) {
	var blocks []Block
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceBlock, hashorNumber, "previous"))
	if err != nil {
		return blocks, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	req.Header.Add("project_id", c.projectId)
	req.WithContext(ctx)
	if err != nil {
		return blocks, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return blocks, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return blocks, handleAPIErrorResponse(res)
	}
	err = json.NewDecoder(res.Body).Decode(&blocks)
	if err != nil {
		return []Block{}, err
	}
	return blocks, nil
}

func (c *apiClient) BlockTransactions(ctx context.Context, hashOrNumber string) ([]Transaction, error) {
	var txs []Transaction
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceBlock, hashOrNumber, "txs"))
	if err != nil {
		return txs, err
	}
	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	req.Header.Add("project_id", c.projectId)
	req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return txs, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return txs, handleAPIErrorResponse(res)
	}

	err = json.NewDecoder(res.Body).Decode(&txs)
	if err != nil {
		return []Transaction{}, err
	}
	return txs, nil
}

func (c *apiClient) BlockLatestTransactions(ctx context.Context) ([]Transaction, error) {
	var txs []Transaction
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceBlocksLatestTransactions))
	if err != nil {
		return txs, err
	}
	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	req.Header.Add("project_id", c.projectId)
	req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return txs, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return txs, handleAPIErrorResponse(res)
	}

	err = json.NewDecoder(res.Body).Decode(&txs)
	if err != nil {
		return []Transaction{}, err
	}
	return txs, nil
}

func (c *apiClient) BlockBySlot(ctx context.Context, slotNumber int) (Block, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%d", c.server, resourceBlocksSlot, slotNumber))
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
	if res.StatusCode != http.StatusOK {
		return Block{}, handleAPIErrorResponse(res)
	}

	block := Block{}
	err = json.NewDecoder(res.Body).Decode(&block)
	if err != nil {
		return Block{}, err
	}
	return block, nil

}

func (c *apiClient) BlocksBySlotAndEpoch(ctx context.Context, slotNumber int, epochNumber int) (Block, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%d/%s/%d", c.server, epochNumber, "slot", slotNumber))
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

	if res.StatusCode != http.StatusOK {
		return Block{}, handleAPIErrorResponse(res)
	}

	block := Block{}
	err = json.NewDecoder(res.Body).Decode(&block)
	if err != nil {
		return Block{}, err
	}
	return block, nil
}
