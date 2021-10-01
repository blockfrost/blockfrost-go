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
	AddressUTXOs(ctx context.Context, address string, query APIPagingParams) ([]AddressUTXO, error)
	Account(ctx context.Context, stakeAddress string) (Account, error)
	AccountHistory(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountHistory, error)
	AccountRewardsHistory(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountRewardsHistory, error)
	AccountDelegationHistory(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountDelegationHistory, error)
	AccountRegistrationHistory(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountRegistrationHistory, error)
	AccountWithdrawalHistory(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountWithdrawalHistory, error)
	AccountMIRHistory(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountMIRHistory, error)
	AccountAssociatedAddresses(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountAssociatedAddress, error)
	AccountAssociatedAssets(ctx context.Context, stakeAddress string, query APIPagingParams) ([]AccountAssociatedAsset, error)
	Genesis(ctx context.Context) (GenesisBlock, error)
	MetadataTxLabels(ctx context.Context, query APIPagingParams) ([]MetadataTxLabel, error)
	MetadataTxContentInJSON(ctx context.Context, label string, query APIPagingParams) ([]MetadataTxContentInJSON, error)
	MetadataTxContentInCBOR(ctx context.Context, label string, query APIPagingParams) ([]MetadataTxContentInCBOR, error)
	Network(ctx context.Context) (NetworkInfo, error)
	Script(ctx context.Context, address string) (Script, error)
	Scripts(ctx context.Context, query APIPagingParams) ([]Script, error)
	ScriptsAll(ctx context.Context) <-chan ScriptAllResult
	ScriptRedeemers(ctx context.Context, address string, query APIPagingParams) ([]ScriptRedeemer, error)
	ScriptRedeemersAll(ctx context.Context, address string) <-chan ScriptRedeemerResult
	Pools(ctx context.Context, query APIPagingParams) (Pools, error)
	PoolsRetired(ctx context.Context, query APIPagingParams) ([]PoolRetired, error)
	PoolsRetiring(ctx context.Context, query APIPagingParams) ([]PoolRetiring, error)
	PoolSpecific(ctx context.Context, poolID string, query APIPagingParams) (PoolSpecific, error)
	PoolHistory(ctx context.Context, poolID string, query APIPagingParams) ([]PoolHistory, error)
	PoolMetadata(ctx context.Context, poolID string, query APIPagingParams) (PoolMetadata, error)
	PoolRelays(ctx context.Context, poolID string, query APIPagingParams) ([]PoolRelay, error)
	PoolDelegators(ctx context.Context, poolID string, query APIPagingParams) ([]PoolDelegator, error)
	PoolBlocks(ctx context.Context, poolID string, query APIPagingParams) (PoolBlocks, error)
	PoolUpdate(ctx context.Context, poolID string, query APIPagingParams) ([]PoolUpdate, error)
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
