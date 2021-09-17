package blockfrost

import "fmt"

const (
	CardanoMainNet = "https://cardano-mainnet.blockfrost.io/api/v0"
	CardanoTestNet = "https://cardano-testnet.blockfrost.io/api/v0"
	IPFS           = "https://ipfs.blockfrost.io/api/v0"
)

// resource paths
const (
	resourceHealth                             = "health"
	resourceHealthClock                        = "health/clock"
	resourceMetrics                            = "metrics"
	resourceMetricsEndpoint                    = "metrics/endpoints"
	resourceBlock                              = "blocks"
	resourceBlocksLatest                       = "blocks/latest"
	resourceBlocksLatestTransactions           = "blocks/latest/txs"
	resourceBlocksSlot                         = "blocks/slot"
	resourceBlocksEpoch                        = "blocks/epoch"
	resourceAccount                            = "accounts"
	resourceAccountHistory                     = "history"
	resourceAccountRewardsHistory              = "rewards"
	resourceAccountDelegationHistory           = "delegations"
	resourceAccountRegistrationHistory         = "registrations"
	resourceAccountWithdrawalHistory           = "withdrawals"
	resourceAccountMIRHistory                  = "mirs"
	resourceAccountAssociatedAddress           = "addresses"
	resourceAccountAddressWithAssetsAssociated = "addresses/assets"
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
}

// Account return Specific account address
// Obtain information about a specific stake account.
type Account struct {
	StakeAddress       string `json:"stake_address,omitempty"`
	Active             bool   `json:"active,omitempty"`
	ActiveEpoch        int64  `json:"active_epoch,omitempty"`
	ControlledAmount   string `json:"controlled_amount,omitempty"`
	RewardsSum         string `json:"rewards_sum,omitempty"`
	WithdrawalsSum     string `json:"withdrawals_sum,omitempty"`
	ReservesSum        string `json:"reserves_sum,omitempty"`
	TreasurySum        string `json:"treasury_sum,omitempty"`
	WithdrawableAmount string `json:"withdrawable_amount,omitempty"`
	PoolID             string `json:"pool_id,omitempty"`
}

// AccountRewardsHist return Account reward history
// Obtain information about the reward history of a specific account.
type AccountRewardsHistory struct {
	Epoch  int32  `json:"epoch,omitempty"`
	Amount string `json:"amount,omitempty"`
	PoolID string `json:"pool_id,omitempty"`
}

// AccountHistory return Account history
// Obtain information about the history of a specific account.
type AccountHistory struct {
	ActiveEpoch int32  `json:"active_epoch,omitempty"`
	Amount      string `json:"amount,omitempty"`
	PoolID      string `json:"pool_id,omitempty"`
}

// AccountDelegationHistory return Account delegation history
// Obtain information about the delegation of a specific account.
type AccountDelegationHistory struct {
	ActiveEpoch int32  `json:"active_epoch,omitempty"`
	TXHash      string `json:"tx_hash,omitempty"`
	Amount      string `json:"amount,omitempty"`
	PoolID      string `json:"pool_id,omitempty"`
}

// AccountRegistrationHistory return Account registration history
// Obtain information about the registrations and deregistrations of a specific account.
type AccountRegistrationHistory struct {
	TXHash string `json:"tx_hash,omitempty"`
	Action string `json:"action,omitempty"`
}

// AccountWithdrawalHistory return Account withdrawal history
// Obtain information about the withdrawals of a specific account.
type AccountWithdrawalHistory struct {
	TXHash string `json:"tx_hash,omitempty"`
	Amount string `json:"amount,omitempty"`
}

// AccountMIRHistory return Account MIR history
// Obtain information about the MIRs of a specific account.
type AccountMIRHistory struct {
	TXHash string `json:"tx_hash,omitempty"`
	Amount string `json:"amount,omitempty"`
}

// AccountAssociatedAddress return Account associated addresses
// Obtain information about the addresses of a specific account.
type AccountAssociatedAddress struct {
	Address string `json:"address,omitempty"`
}

// AccountAssetsWithAddress return Assets associated with the account addresses
// Obtain information about assets associated with addresses of a specific account.
// Be careful, as an account could be part of a mangled address and does not necessarily mean the addresses are owned by user as the account.
type AccountAssetsWithAddress struct {
	Unit     string `json:"unit,omitempty"`
	Quantity string `json:"quantity,omitempty"`
}
