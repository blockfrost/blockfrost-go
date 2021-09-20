package blockfrost

import "fmt"

const (
	CardanoMainNet = "https://cardano-mainnet.blockfrost.io/api/v0"
	CardanoTestNet = "https://cardano-testnet.blockfrost.io/api/v0"
	IPFS           = "https://ipfs.blockfrost.io/api/v0"
)

// resource paths
const (
	resourceHealth                   = "health"
	resourceHealthClock              = "health/clock"
	resourceMetrics                  = "metrics"
	resourceMetricsEndpoint          = "metrics/endpoints"
	resourceBlock                    = "blocks"
	resourceBlocksLatest             = "blocks/latest"
	resourceBlocksLatestTransactions = "blocks/latest/txs"
	resourceBlocksSlot               = "blocks/slot"
	resourceBlocksEpoch              = "blocks/epoch"
	resourceMetadataTxLabels         = "metadata/txs/labels"
	resourceMetadataTxContentInJSON  = "metadata/txs/labels" // and {label_parameter}
	resourceMetadataTxContentInCBOR  = "metadata/txs/labels" // and {label_parameter}/cbor
	resourcePool                     = "pools"
	resourcePoolRetired              = "pools/retired"
	resourcePoolRetiring             = "pools/retiring"
	resourcePoolHistory              = "history"
	resourcePoolMetadata             = "metadata"
	resourcePoolRelay                = "relays"
	resourcePoolDelegator            = "delegators"
	resourcePoolBlocks               = "blocks"
	resourcePoolUpdate               = "updates"
)

// APIError is used to describe errors from the API.
// See https://docs.blockfrost.io/#section/Errors
type APIError struct {
	Response interface{}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error, %+v", e.Response)
}

// Autobanned defines model for HTTP `418` (Auto Banned).
type AutoBanned struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// BadRequest defines model for HTTP `400` (Bad Request)
type BadRequest struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// InternalServerError defines model for HTTP `500` (Internal Server Error)
type InternalServerError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// NotFound defines model for HTTP `404` (Resource Not Found).
type NotFound struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// OverusageLimit defines model for HTTP `429` (Over Usage).
type OverusageLimit struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// UnauthorizedEror defines model for HTTP `403` (Unauthorized).
type UnauthorizedError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// Info defines model for root endpoint `/`
type Info struct {
	Url     string `json:"url,omitempty"`
	Version string `json:"version,omitempty"`
}

// Health describes boolean for backend server health status.
type Health struct {
	IsHealthy bool `json:"is_healthy,omitempty"`
}

// HealthClock describes current UNIX time
type HealthClock struct {
	ServerTime int64 `json:"server_time,omitempty"`
}

// Metric describes the Blockfrost usage metrics
type Metric struct {
	// Sum of all calls for a particular day
	Calls int `json:"calls,omitempty"`

	// Starting time of the call count interval (ends midnight UTC) in UNIX time
	Time int `json:"time,omitempty"`
}

// MetricsEndpoint
type MetricsEndpoint struct {
	// Sum of all calls for a particular day and endpoint
	Calls int `json:"calls,omitempty"`

	// Endpoint parent name
	Endpoint string `json:"endpoint,omitempty"`

	// Starting time of the call count interval (ends midnight UTC) in UNIX time
	Time int `json:"time,omitempty"`
}

// Block defines content of a block
type Block struct {
	Time          int    `json:"time,omitempty"`
	Height        int    `json:"height,omitempty"`
	Hash          string `json:"hash,omitempty"`
	Slot          int    `json:"slot,omitempty"`
	Epoch         int    `json:"epoch,omitempty"`
	EpochSlot     int    `json:"epoch_slot,omitempty"`
	SlotLeader    string `json:"slot_leader"`
	Size          int    `json:"size,omitempty"`
	TxCount       int    `json:"tx_count,omitempty"`
	Output        string `json:"output,omitempty"`
	Fees          string `json:"fees,omitempty"`
	BlockVRF      string `json:"block_vrf,omitempty"`
	PreviousBlock string `json:"previous_block,omitempty"`
	NextBlock     string `json:"next_block,omitempty"`
	Confirmations int    `json:"confirmations,omitempty"`
}

type Transaction string

type APIPagingParams struct {
	Count int
	Page  int
	Order string
	From  string
	To    string
}

type APIQueryParams APIPagingParams

// MetadataTxLabel return Transaction metadata labels
// List of all used transaction metadata labels.
type MetadataTxLabel struct {
	Label string `json:"label,omitempty"`
	Cip10 string `json:"cip10,omitempty"`
	Count string `json:"count,omitempty"`
}

// MetadataTxContentInJSONRaw Transaction metadata content raw in JSON
// Transaction metadata per label.
// This struct are more flexible on JSONMetadata field
type MetadataTxContentInJSONRaw struct {
	TxHash string `json:"tx_hash,omitempty"`
	// 	string or object or Array of any or integer or number or boolean Nullable
	// Content of the JSON metadata
	JSONMetadata map[string]interface{} `json:"json_metadata,omitempty"`
}

// MetadataTxContentInJSON return Transaction metadata content in JSON
// Transaction metadata per label.
// This struct are stronger typed on JSONMetadata field
type MetadataTxContentInJSON struct {
	TxHash       string       `json:"tx_hash,omitempty"`
	JSONMetadata JSONMetadata `json:"json_metadata,omitempty"`
}

type JSONMetadata struct {
	Adausd []Adausd `json:"ADAUSD,omitempty"`
}

type Adausd struct {
	Value  string `json:"value,omitempty"`
	Source string `json:"source,omitempty"`
}

// MetadataTxContentInCBOR return Transaction metadata content in CBOR
// Transaction metadata per label.
type MetadataTxContentInCBOR struct {
	TxHash       string `json:"tx_hash,omitempty"`
	CborMetadata string `json:"cbor_metadata,omitempty"`
}

// List of stake pools
// List of registered stake pools.
type Pools []string

// List of retired stake pools
// List of already retired pools.
type PoolRetired struct {
	PoolID string `json:"pool_id,omitempty"`
	Epoch  int    `json:"epoch,omitempty"`
}

type PoolRetiring struct {
	PoolID string `json:"pool_id,omitempty"`
	Epoch  int    `json:"epoch,omitempty"`
}

// PoolSpecific return Specific stake pool
// Pool information.
type PoolSpecific struct {
	PoolID         string   `json:"pool_id,omitempty"`
	Hex            string   `json:"hex,omitempty"`
	VrfKey         string   `json:"vrf_key,omitempty"`
	BlocksMinted   int64    `json:"blocks_minted,omitempty"`
	LiveStake      string   `json:"live_stake,omitempty"`
	LiveSize       float64  `json:"live_size,omitempty"`
	LiveSaturation float64  `json:"live_saturation,omitempty"`
	LiveDelegators int64    `json:"live_delegators,omitempty"`
	ActiveStake    string   `json:"active_stake,omitempty"`
	ActiveSize     float64  `json:"active_size,omitempty"`
	DeclaredPledge string   `json:"declared_pledge,omitempty"`
	LivePledge     string   `json:"live_pledge,omitempty"`
	MarginCost     float64  `json:"margin_cost,omitempty"`
	FixedCost      string   `json:"fixed_cost,omitempty"`
	RewardAccount  string   `json:"reward_account,omitempty"`
	Owners         []string `json:"owners,omitempty"`
	Registration   []string `json:"registration,omitempty"`
	Retirement     []string `json:"retirement,omitempty"`
}

// PoolHistory return Stake pool history
// History of stake pool parameters over epochs.
type PoolHistory struct {
	Epoch           int64   `json:"epoch,omitempty"`
	Blocks          int64   `json:"blocks,omitempty"`
	ActiveStake     string  `json:"active_stake,omitempty"`
	ActiveSize      float64 `json:"active_size,omitempty"`
	DelegatorsCount int64   `json:"delegators_count,omitempty"`
	Rewards         string  `json:"rewards,omitempty"`
	Fees            string  `json:"fees,omitempty"`
}

// PoolMetadata return Stake pool metadata
// Stake pool registration metadata.
type PoolMetadata struct {
	PoolID      string `json:"pool_id,omitempty"`
	Hex         string `json:"hex,omitempty"`
	URL         string `json:"url,omitempty"`
	Hash        string `json:"hash,omitempty"`
	Ticker      string `json:"ticker,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Homepage    string `json:"homepage,omitempty"`
}

// PoolRelay return Stake pool relays
// Relays of a stake pool.
type PoolRelay struct {
	Ipv4   string `json:"ipv4,omitempty"`
	Ipv6   string `json:"ipv6,omitempty"`
	DNS    string `json:"dns,omitempty"`
	DNSSrv string `json:"dns_srv,omitempty"`
	Port   int64  `json:"port,omitempty"`
}

// PoolDelegator return Stake pool delegators
// List of current stake pools delegators.
type PoolDelegator struct {
	Address   string `json:"address"`
	LiveStake string `json:"live_stake"`
}

// PoolBlocks return Stake pool blocks
// List of stake pools blocks.
type PoolBlocks []string

// PoolUpdate return Stake pool updates
// List of certificate updates to the stake pool.
type PoolUpdate struct {
	TxHash    string `json:"tx_hash,omitempty"`
	CERTIndex int64  `json:"cert_index,omitempty"`
	Action    string `json:"action,omitempty"`
}
