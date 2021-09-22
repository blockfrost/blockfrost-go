package blockfrost

import "context"

type Asset struct {
	Asset             string      `json:"asset,omitempty"`
	PolicyId          string      `json:"policy_id,omitempty"`
	AssetName         string      `json:"asset_name,omitempty"`
	Fingerprint       string      `json:"fingerprint,omitempty"`
	Quantity          string      `json:"quantity,omitempty"`
	InitialMintTxHash string      `json:"initial_mint_tx_hash,omitempty"`
	MintOrBurnCount   int         `json:"mint_or_burn_count,omitempty"`
	OnchainMetadata   interface{} `json:"onchain_metadata,omitempty"`
	Metadata          interface{} `json:"metadata,omitempty"`
}

type AssetHistory struct {
	TxHash string `json:"tx_hash,omitempty"`
	Action string `json:"action,omitempty"`
	Amount string `json:"amount,omitempty"`
}

type AssetTransaction struct {
	TxHash      string `json:"tx_hash,omitempty"`
	TxIndex     int    `json:"tx_index,omitempty"`
	BlockHeight int    `json:"block_height,omitempty"`
}

type AssetAddress struct {
	Address  string `json:"address,omitempty"`
	Quantity string `json:"quantity,omitempty"`
}

func (c *apiClient) Assets(ctx context.Context, query APIQueryParams) (a []Asset, err error) {
	return
}

func (c *apiClient) Asset(ctx context.Context, asset string) (a Asset, err error) {
	return
}

func (c *apiClient) AssetHistory(ctx context.Context, asset string) (hist []AssetHistory, err error) {
	return
}

func (c *apiClient) AssetTransactions(ctx context.Context, asset string) (trs []AssetTransaction, err error) {
	return
}

func (c *apiClient) AssetAddresses(ctx context.Context, asset string) (addrs []AssetAddress, err error) {
	return
}

func (c *apiClient) PolicyAssets(ctx context.Context, asset string) (a []Asset, err error) {
	return
}
