package blockfrost

import "fmt"

const (
	CardanoMainNet = "https://cardano-mainnet.blockfrost.io/api/v0"
	CardanoTestNet = "https://cardano-testnet.blockfrost.io/api/v0"
	IPFS           = "https://ipfs.blockfrost.io/api/v0"
)

const (
	resourceHealth          = "health"
	resourceHealthClock     = "health/clock"
	resourceMetrics         = "metrics"
	resourceMetricsEndpoint = "metrics/endpoints"
	resourceBlocksLatest    = "blocks/latest"
)

// Errors returned by API
type APIError struct {
	Response interface{}
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API Error, %+v", e.Response)
}

// Autobanned defines model for autobanned.
type AutoBanned struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// BadRequest defines model for bad_request.
type BadRequest struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// InternalServerError defines model for internal_server_error.
type InternalServerError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// NotFound defines model for not_found.
type NotFound struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// OverusageLimit defines model for overusage_limit.
type OverusageLimit struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// Unauthorized defines model for unauthorized_error.
type UnauthorizedError struct {
	Error      string `json:"error"`
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}

// Health Group types
type Info struct {
	Url     string `json:"url,omitempty"`
	Version string `json:"version,omitempty"`
}

type Health struct {
	IsHealthy bool `json:"is_healthy,omitempty"`
}

type HealthClock struct {
	ServerTime int64 `json:"server_time,omitempty"`
}

// Metric Group types
type Metric struct {
	// Sum of all calls for a particular day
	Calls int `json:"calls,omitempty"`

	// Starting time of the call count interval (ends midnight UTC) in UNIX time
	Time int `json:"time,omitempty"`
}

type MetricsEndpoint struct {
	// Sum of all calls for a particular day and endpoint
	Calls int `json:"calls,omitempty"`

	// Endpoint parent name
	Endpoint string `json:"endpoint,omitempty"`

	// Starting time of the call count interval (ends midnight UTC) in UNIX time
	Time int `json:"time,omitempty"`
}

// Blocks Group types
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
