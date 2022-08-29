package blockfrost

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/hashicorp/go-retryablehttp"
)

type apiClient struct {
	server    string
	projectId string
	client    *retryablehttp.Client
	routines  int
}

// HttpRequestDoer defines methods for a http client.
type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

// APIClientOptions contains optios used to initialize an API Client using
// NewAPIClient
type APIClientOptions struct {
	// The project_id to use from blockfrost. If not set
	// `BLOCKFROST_PROJECT_ID` is loaded from env
	ProjectID string

	// Server url to use
	Server string

	// Max number of routines to use for *All methods
	MaxRoutines int

	RetryWaitMin time.Duration // Minimum time to wait
	RetryWaitMax time.Duration // Maximum time to wait
	RetryMax     int           // Maximum number of retries
}

// NewAPICLient creates a client from APIClientOptions. If no options are provided,
//  client with default configurations is returned.
func NewAPIClient(options APIClientOptions) APIClient {
	if options.Server == "" {
		options.Server = CardanoMainNet
	}

	retryclient := retryablehttp.NewClient()
	retryclient.Logger = nil

	if options.ProjectID == "" {
		options.ProjectID = os.Getenv("BLOCKFROST_PROJECT_ID")
	}

	if options.MaxRoutines == 0 {
		options.MaxRoutines = 10
	}

	client := &apiClient{
		server:    options.Server,
		client:    retryclient,
		projectId: options.ProjectID,
		routines:  options.MaxRoutines,
	}

	return client
}

// APIClient defines methods implemented by the api client.
type APIClient interface {
	Info(ctx context.Context) (Info, error)
	Health(ctx context.Context) (Health, error)
	HealthClock(ctx context.Context) (HealthClock, error)
	Metrics(ctx context.Context) ([]Metric, error)
	MetricsEndpoints(ctx context.Context) ([]MetricsEndpoint, error)
	Block(ctx context.Context, hashOrNumber string) (Block, error)
	BlockLatest(ctx context.Context) (Block, error)
	BlockLatestTransactions(ctx context.Context) ([]Transaction, error)
	BlockTransactions(ctx context.Context, hashOrNumer string) ([]Transaction, error)
	BlocksNext(ctx context.Context, hashOrNumber string) ([]Block, error)
	BlocksPrevious(ctx context.Context, hashOrNumber string) ([]Block, error)
	BlockBySlot(ctx context.Context, slotNumber int) (Block, error)
	BlocksBySlotAndEpoch(ctx context.Context, slotNumber int, epochNumber int) (Block, error)
	EpochLatest(ctx context.Context) (Epoch, error)
	LatestEpochParameters(ctx context.Context) (EpochParameters, error)
	Epoch(ctx context.Context, epochNumber int) (Epoch, error)
	EpochsNext(ctx context.Context, epochNumber int, query APIQueryParams) ([]Epoch, error)
	EpochNextAll(ctx context.Context, epochNumber int) <-chan EpochResult
	EpochsPrevious(ctx context.Context, epochNumber int, query APIQueryParams) ([]Epoch, error)
	EpochPreviousAll(ctx context.Context, epochNumber int) <-chan EpochResult
	EpochStakeDistribution(ctx context.Context, epochNumber int, query APIQueryParams) ([]EpochStake, error)
	EpochStakeDistributionAll(ctx context.Context, epochNumber int) <-chan EpochStakeResult
	EpochStakeDistributionByPool(ctx context.Context, epochNumber int, poolId string, query APIQueryParams) ([]EpochStake, error)
	EpochStakeDistributionByPoolAll(ctx context.Context, epochNumber int, poolId string) <-chan EpochStakeResult
	EpochBlockDistribution(ctx context.Context, epochNumber int, query APIQueryParams) ([]string, error)
	EpochBlockDistributionAll(ctx context.Context, epochNumber int) <-chan BlockDistributionResult
	EpochBlockDistributionByPool(ctx context.Context, epochNumber int, poolId string, query APIQueryParams) ([]string, error)
	EpochBlockDistributionByPoolAll(ctx context.Context, epochNumber int, poolId string) <-chan BlockDistributionResult
	EpochParameters(ctx context.Context, epochNumber int) (EpochParameters, error)
	Address(ctx context.Context, address string) (Address, error)
	AddressDetails(ctx context.Context, address string) (AddressDetails, error)
	AddressTransactions(ctx context.Context, address string, query APIQueryParams) ([]AddressTransactions, error)
	AddressTransactionsAll(ctx context.Context, address string) <-chan AddressTxResult
	AddressUTXOs(ctx context.Context, address string, query APIQueryParams) ([]AddressUTXO, error)
	AddressUTXOsAll(ctx context.Context, address string) <-chan AddressUTXOResult
	Account(ctx context.Context, stakeAddress string) (Account, error)
	AccountHistory(ctx context.Context, stakeAddress string, query APIQueryParams) ([]AccountHistory, error)
	AccountHistoryAll(ctx context.Context, address string) <-chan AccountHistoryResult
	AccountRewardsHistory(ctx context.Context, stakeAddress string, query APIQueryParams) ([]AccountRewardsHistory, error)
	AccountRewardsHistoryAll(ctx context.Context, stakeAddress string) <-chan AccountRewardHisResult
	AccountDelegationHistory(ctx context.Context, stakeAddress string, query APIQueryParams) ([]AccountDelegationHistory, error)
	AccountDelegationHistoryAll(ctx context.Context, stakeAddress string) <-chan AccDelegationHistoryResult
	AccountRegistrationHistory(ctx context.Context, stakeAddress string, query APIQueryParams) ([]AccountRegistrationHistory, error)
	AccountRegistrationHistoryAll(ctx context.Context, stakeAddress string) <-chan AccountRegistrationHistoryResult
	AccountWithdrawalHistory(ctx context.Context, stakeAddress string, query APIQueryParams) ([]AccountWithdrawalHistory, error)
	AccountWithdrawalHistoryAll(ctx context.Context, stakeAddress string) <-chan AccountWithdrawalHistoryResult
	AccountMIRHistory(ctx context.Context, stakeAddress string, query APIQueryParams) ([]AccountMIRHistory, error)
	AccountMIRHistoryAll(ctx context.Context, stakeAddress string) <-chan AccountMIRHistoryResult
	AccountAssociatedAddresses(ctx context.Context, stakeAddress string, query APIQueryParams) ([]AccountAssociatedAddress, error)
	AccountAssociatedAddressesAll(ctx context.Context, stakeAddress string) <-chan AccountAssociatedAddressesAll
	AccountAssociatedAssets(ctx context.Context, stakeAddress string, query APIQueryParams) ([]AccountAssociatedAsset, error)
	AccountAssociatedAssetsAll(ctx context.Context, stakeAddress string) <-chan AccountAssociatedAssetsAll
	Asset(ctx context.Context, asset string) (Asset, error)
	Assets(ctx context.Context, query APIQueryParams) ([]Asset, error)
	AssetsAll(ctx context.Context) <-chan AssetResult
	AssetHistory(ctx context.Context, asset string) ([]AssetHistory, error)
	AssetTransactions(ctx context.Context, asset string) ([]AssetTransaction, error)
	AssetAddresses(ctx context.Context, asset string, query APIQueryParams) ([]AssetAddress, error)
	AssetAddressesAll(ctx context.Context, asset string) <-chan AssetAddressesAll
	AssetsByPolicy(ctx context.Context, policyId string) ([]Asset, error)
	Genesis(ctx context.Context) (GenesisBlock, error)
	MetadataTxLabels(ctx context.Context, query APIQueryParams) ([]MetadataTxLabel, error)
	MetadataTxLabelsAll(ctx context.Context) <-chan MetadataTxLabelResult
	MetadataTxContentInJSON(ctx context.Context, label string, query APIQueryParams) ([]MetadataTxContentInJSON, error)
	MetadataTxContentInJSONAll(ctx context.Context, label string) <-chan MetadataTxContentInJSONResult
	MetadataTxContentInCBOR(ctx context.Context, label string, query APIQueryParams) ([]MetadataTxContentInCBOR, error)
	MetadataTxContentInCBORAll(ctx context.Context, label string) <-chan MetadataTxContentInCBORResult
	Network(ctx context.Context) (NetworkInfo, error)
	Nutlink(ctx context.Context, address string) (NutlinkAddress, error)
	Tickers(ctx context.Context, address string, query APIQueryParams) ([]Ticker, error)
	TickersAll(ctx context.Context, address string) <-chan TickerResult
	TickerRecords(ctx context.Context, ticker string, query APIQueryParams) ([]TickerRecord, error)
	TickerRecordsAll(ctx context.Context, ticker string) <-chan TickerRecordResult
	AddressTickerRecords(ctx context.Context, address string, ticker string, query APIQueryParams) ([]TickerRecord, error)
	AddressTickerRecordsAll(ctx context.Context, address string, ticker string) <-chan TickerRecordResult
	Script(ctx context.Context, address string) (Script, error)
	Scripts(ctx context.Context, query APIQueryParams) ([]Script, error)
	ScriptsAll(ctx context.Context) <-chan ScriptAllResult
	ScriptRedeemers(ctx context.Context, address string, query APIQueryParams) ([]ScriptRedeemer, error)
	ScriptRedeemersAll(ctx context.Context, address string) <-chan ScriptRedeemerResult
	Pool(ctx context.Context, poolID string) (Pool, error)
	Pools(ctx context.Context, query APIQueryParams) (Pools, error)
	PoolsAll(ctx context.Context) <-chan PoolsResult
	PoolsRetired(ctx context.Context, query APIQueryParams) ([]PoolRetired, error)
	PoolsRetiredAll(ctx context.Context) <-chan PoolsRetiredResult
	PoolsRetiring(ctx context.Context, query APIQueryParams) ([]PoolRetiring, error)
	PoolsRetiringAll(ctx context.Context) <-chan PoolsRetiringResult
	PoolHistory(ctx context.Context, poolID string, query APIQueryParams) ([]PoolHistory, error)
	PoolHistoryAll(ctx context.Context, poolId string) <-chan PoolHistoryResult
	PoolMetadata(ctx context.Context, poolID string) (PoolMetadata, error)
	PoolRelays(ctx context.Context, poolID string) ([]PoolRelay, error)
	PoolDelegators(ctx context.Context, poolID string, query APIQueryParams) ([]PoolDelegator, error)
	PoolDelegatorsAll(ctx context.Context, poolId string) <-chan PoolDelegatorsResult
	PoolBlocks(ctx context.Context, poolID string, query APIQueryParams) (PoolBlocks, error)
	PoolBlocksAll(ctx context.Context, poolId string) <-chan PoolBlocksResult
	PoolUpdates(ctx context.Context, poolID string, query APIQueryParams) ([]PoolUpdate, error)
	PoolUpdatesAll(ctx context.Context, poolId string) <-chan PoolUpdateResult
	Transaction(ctx context.Context, hash string) (TransactionContent, error)
	TransactionUTXOs(ctx context.Context, hash string) (TransactionUTXOs, error)
	TransactionStakeAddressCerts(ctx context.Context, hash string) ([]TransactionStakeAddressCert, error)
	TransactionWithdrawals(ctx context.Context, hash string) ([]TransactionWidthrawal, error)
	TransactionMIRs(ctx context.Context, hash string) ([]TransactionMIR, error)
	TransactionMetadata(ctx context.Context, hash string) ([]TransactionMetadata, error)
	TransactionMetadataInCBORs(ctx context.Context, hash string) ([]TransactionMetadataCbor, error)
	TransactionRedeemers(ctx context.Context, hash string) ([]TransactionRedeemer, error)
	TransactionDelegationCerts(ctx context.Context, hash string) ([]TransactionDelegation, error)
	TransactionPoolUpdates(ctx context.Context, hash string) ([]TransactionPoolCert, error)
	TransactionPoolUpdateCerts(ctx context.Context, hash string) ([]TransactionPoolCert, error)
	TransactionPoolRetirementCerts(ctx context.Context, hash string) ([]TransactionPoolCert, error)
	TransactionSubmit(ctx context.Context, cbor []byte) (string, error)
}
