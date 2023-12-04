package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Block defines content of a block
type Block struct {
	// Block creation time in UNIX time
	Time int `json:"time"`

	// Block number
	Height int `json:"height"`

	// Hash of the block
	Hash string `json:"hash"`

	// Slot number
	Slot int `json:"slot"`

	// Epoch number
	Epoch int `json:"epoch"`

	// Slot within the epoch
	EpochSlot int `json:"epoch_slot"`

	// Bech32 ID of the slot leader or specific block description in case there is no slot leader
	SlotLeader string `json:"slot_leader"`

	// Block size in Bytes
	Size int `json:"size"`

	// Number of transactions in the block
	TxCount int `json:"tx_count"`

	// Total output within the block in Lovelaces
	Output string `json:"output"`

	// Total fees within the block in Lovelaces
	Fees string `json:"fees"`

	// VRF key of the block
	BlockVRF string `json:"block_vrf"`

	// Hash of the previous block
	PreviousBlock string `json:"previous_block"`

	// Hash of the next block
	NextBlock *string `json:"next_block"`

	// Number of block confirmations
	Confirmations int `json:"confirmations"`
}

// BlocksLatest Return the latest block available to the backends, also known as the
// tip of the blockchain.
func (c *apiClient) BlockLatest(ctx context.Context) (b Block, err error) {
	requestUrl, err := url.Parse((fmt.Sprintf("%s/%s", c.server, resourceBlocksLatest)))
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

	err = json.NewDecoder(res.Body).Decode(&b)
	if err != nil {
		return
	}
	return b, nil
}

// Block returns the content of a requested block by the hash or block number
func (c *apiClient) Block(ctx context.Context, hashOrNumber string) (bl Block, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", c.server, resourceBlock, hashOrNumber))
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

	if err = json.NewDecoder(res.Body).Decode(&bl); err != nil {
		return
	}
	return bl, nil
}

// BlocksNext returns the list of blocks following a specific block
// denoted by the hash or block number
func (c *apiClient) BlocksNext(ctx context.Context, hashorNumber string) (bls []Block, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceBlock, hashorNumber, "next"))
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

	if err = json.NewDecoder(res.Body).Decode(&bls); err != nil {
		return
	}
	return bls, nil
}

// BlocksPrevious returns the list of blocks preceding a specific block
// specified by a hash or block number
func (c *apiClient) BlocksPrevious(ctx context.Context, hashorNumber string) (bls []Block, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceBlock, hashorNumber, "previous"))
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

	if err = json.NewDecoder(res.Body).Decode(&bls); err != nil {
		return
	}
	return bls, nil
}

// BlocksTransactions returns slice of Transaction within the block specified
// by a hash or block number
func (c *apiClient) BlockTransactions(ctx context.Context, hashOrNumber string) (txs []Transaction, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%s", c.server, resourceBlock, hashOrNumber, "txs"))
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

	err = json.NewDecoder(res.Body).Decode(&txs)
	if err != nil {
		return
	}
	return txs, nil
}

// BlockLatestTransactions returns the transactions within the latest block.
func (c *apiClient) BlockLatestTransactions(ctx context.Context) (txs []Transaction, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceBlocksLatestTransactions))
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

	if err = json.NewDecoder(res.Body).Decode(&txs); err != nil {
		return
	}
	return txs, nil
}

// BlocksBySlot returns the content of a requested block for a specific slot.
func (c *apiClient) BlockBySlot(ctx context.Context, slotNumber int) (bl Block, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%d", c.server, resourceBlocksSlot, slotNumber))
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

	err = json.NewDecoder(res.Body).Decode(&bl)
	if err != nil {
		return
	}
	return bl, nil

}

// BlocksBySlotAndEpoch returns a Block for a specific slot and epoch
func (c *apiClient) BlocksBySlotAndEpoch(ctx context.Context, slotNumber int, epochNumber int) (bl Block, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s/%d/%s/%d", c.server, "blocks", "epoch", epochNumber, "slot", slotNumber))
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

	err = json.NewDecoder(res.Body).Decode(&bl)
	if err != nil {
		return
	}
	return bl, nil
}
