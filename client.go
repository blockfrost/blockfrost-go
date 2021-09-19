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
}

func NewAPIClient(options APIClientOptions) (APIClient, error) {
	if options.Server == "" {
		options.Server = CardanoMainNet
	}

	if options.Client == nil {
		// TODO: Make configurable. Timeout, retries ...
		options.Client = &http.Client{Timeout: time.Second * 5}
	}

	if options.ProjectID == "" {
		options.ProjectID = os.Getenv("BLOCKFROST_PROJECT_ID")
	}

	client := &apiClient{
		server:    options.Server,
		client:    options.Client,
		projectId: options.ProjectID,
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
}
