package blockfrost

import "fmt"

const (
	CardanoMainNet = "https://cardano-mainnet.blockfrost.io/api/v0"
	CardanoTestNet = "https://cardano-testnet.blockfrost.io/api/v0"
	IPFS           = "https://ipfs.blockfrost.io/api/v0"
)

const (
	resourceHealth = "health"
)

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
