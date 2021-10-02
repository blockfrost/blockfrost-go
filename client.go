package blockfrost

import (
	"context"
	"net/http"
	"os"
	"time"
)

type apiClient struct {
	server    string
	projectId string
	client    HttpRequestDoer
	routines  int
}

type HttpRequestDoer interface {
	Do(req *http.Request) (*http.Response, error)
}

type APIClientOptions struct {
	// The project_id to use from blockfrost. If not set
	// `BLOCKFROST_PROJECT_ID` is loaded from env
	ProjectID string
	// Configures server to use. Can be toggled for test servers
	Server string
	// Interface implementing Do method such *http.Client
	Client HttpRequestDoer
	// Number of goroutines to use for *All methods
	Routines int
}

func NewAPIClient(options APIClientOptions) (APIClient, error) {
	if options.Server == "" {
		options.Server = CardanoMainNet
	}

	if options.Client == nil {
		// TODO: Make configurable. Timeout, retries ...
		options.Client = &http.Client{Timeout: time.Second * 10}
	}

	if options.ProjectID == "" {
		options.ProjectID = os.Getenv("BLOCKFROST_PROJECT_ID")
	}

	if options.Routines == 0 {
		options.Routines = 10
	}

	client := &apiClient{
		server:    options.Server,
		client:    options.Client,
		projectId: options.ProjectID,
		routines:  options.Routines,
	}

	return client, nil
}

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
	Address(ctx context.Context, address string) (Address, error)
	AddressDetails(ctx context.Context, address string) (AddressDetails, error)
	AddressTransactions(ctx context.Context, address string, query APIPagingParams) ([]AddressTransactions, error)
	AddressTransactionsAll(ctx context.Context, address string) <-chan AddressTxResult
	AddressUTXOs(ctx context.Context, address string, query APIPagingParams) ([]AddressUTXO, error)
	AddressUTXOsAll(ctx context.Context, address string) <-chan AddressUTXOResult
	Account(ctx context.Context, stakeAddress string) (Account, error)
	AccountHistory(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountHistory, error)
	AccountHistoryAll(ctx context.Context, address string) <-chan AccountHistoryResult
	AccountRewardsHistory(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountRewardsHistory, error)
	AccountRewardsHistoryAll(ctx context.Context, stakeAddress string) <-chan AccountRewardHisResult
	AccountDelegationHistory(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountDelegationHistory, error)
	AccountDelegationHistoryAll(ctx context.Context, stakeAddress string) <-chan AccDelegationHistoryResult
	AccountRegistrationHistory(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountRegistrationHistory, error)
	AccountRegistrationHistoryAll(ctx context.Context, stakeAddress string) <-chan AccountRegistrationHistoryResult
	AccountWithdrawalHistory(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountWithdrawalHistory, error)
	AccountWithdrawalHistoryAll(ctx context.Context, stakeAddress string) <-chan AccountWithdrawalHistoryResult
	AccountMIRHistory(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountMIRHistory, error)
	AccountMIRHistoryAll(ctx context.Context, stakeAddress string) <-chan AccountMIRHistoryResult
	AccountAssociatedAddresses(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountAssociatedAddress, error)
	AccountAssociatedAddressesAll(ctx context.Context, stakeAddress string) <-chan AccountAssociatedAddressesAll
	AccountAssociatedAssets(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountAssociatedAsset, error)
	AccountAssociatedAssetsAll(ctx context.Context, stakeAddress string) <-chan AccountAssociatedAssetsAll
	Asset(ctx context.Context, asset string) (Asset, error)
	Assets(ctx context.Context, query APIPagingParams) ([]Asset, error)
	AssetsAll(ctx context.Context, poolId string) <-chan AssetResult
	AssetHistory(ctx context.Context, asset string) ([]AssetHistory, error)
	AssetTransactions(ctx context.Context, asset string) ([]AssetTransaction, error)
	AssetAddresses(ctx context.Context, asset string) ([]AssetAddress, error)
	AssetsByPolicy(ctx context.Context, policyId string) ([]Asset, error)
	Genesis(ctx context.Context) (GenesisBlock, error)
	MetadataTxLabels(ctx context.Context, query APIPagingParams) ([]MetadataTxLabel, error)
	MetadataTxLabelsAll(ctx context.Context) <-chan MetadataTxLabelResult
	MetadataTxContentInJSON(ctx context.Context, label string, query APIPagingParams) ([]MetadataTxContentInJSON, error)
	MetadataTxContentInJSONAll(ctx context.Context, label string) <-chan MetadataTxContentInJSONResult
	MetadataTxContentInCBOR(ctx context.Context, label string, query APIPagingParams) ([]MetadataTxContentInCBOR, error)
	MetadataTxContentInCBORAll(ctx context.Context, label string) <-chan MetadataTxContentInCBORResult
	Network(ctx context.Context) (NetworkInfo, error)
	Script(ctx context.Context, address string) (Script, error)
	Scripts(ctx context.Context, query APIPagingParams) ([]Script, error)
	ScriptsAll(ctx context.Context) <-chan ScriptAllResult
	ScriptRedeemers(ctx context.Context, address string, query APIPagingParams) ([]ScriptRedeemer, error)
	ScriptRedeemersAll(ctx context.Context, address string) <-chan ScriptRedeemerResult
	Pool(ctx context.Context, poolID string) (Pool, error)
	Pools(ctx context.Context, query APIPagingParams) (Pools, error)
	PoolsAll(ctx context.Context) <-chan PoolsResult
	PoolsRetired(ctx context.Context, query APIPagingParams) ([]PoolRetired, error)
	PoolsRetiredAll(ctx context.Context) <-chan PoolsRetiredResult
	PoolsRetiring(ctx context.Context, query APIPagingParams) ([]PoolRetiring, error)
	PoolsRetiringAll(ctx context.Context) <-chan PoolsRetiringResult
	PoolHistory(ctx context.Context, poolID string, query APIPagingParams) ([]PoolHistory, error)
	PoolHistoryAll(ctx context.Context, poolId string) <-chan PoolHistoryResult
	PoolMetadata(ctx context.Context, poolID string) (PoolMetadata, error)
	PoolRelays(ctx context.Context, poolID string) ([]PoolRelay, error)
	PoolDelegators(ctx context.Context, poolID string, query APIPagingParams) ([]PoolDelegator, error)
	PoolDelegatorsAll(ctx context.Context, poolId string) <-chan PoolDelegatorsResult
	PoolBlocks(ctx context.Context, poolID string, query APIPagingParams) (PoolBlocks, error)
	PoolBlocksAll(ctx context.Context, poolId string) <-chan PoolBlocksResult
	PoolUpdates(ctx context.Context, poolID string, query APIPagingParams) ([]PoolUpdate, error)
	PoolUpdatesAll(ctx context.Context, poolId string) <-chan PoolUpdateResult
	Transaction(ctx context.Context, hash string) (TransactionContent, error)
	TransactionUTXOs(ctx context.Context, hash string) (TransactionUTXOs, error)
	TransactionStakeAddressCerts(ctx context.Context, hash string) ([]TransactionStakeAddressCert, error)
	TransactionWithdrawals(ctx context.Context, hash string) ([]TransactionWidthrawal, error)
	TransactionMIRs(ctx context.Context, hash string) ([]TransactionMIR, error)
	TransactionMetadata(ctx context.Context, hash string) ([]TransactionMetadata, error)
	TransactionMetadataInCBORs(ctx context.Context, hash string) ([]TransactionMetadataCbor, error)
	TransactionRedeemers(ctx context.Context, hash string) ([]TransactionMetadata, error)
	TransactionDelegationCerts(ctx context.Context, hash string) ([]TransactionDelegation, error)
	TransactionPoolUpdates(ctx context.Context, hash string) ([]TransactionPoolCert, error)
	TransactionPoolUpdateCerts(ctx context.Context, hash string) ([]TransactionPoolCert, error)
	TransactionPoolRetirementCerts(ctx context.Context, hash string) ([]TransactionPoolCert, error)
}
