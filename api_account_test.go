package blockfrost_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestResourceAccountIntegration(t *testing.T) {
	t.Parallel()

	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	got, err := api.Account(context.TODO(), inputStakeAddr)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if reflect.DeepEqual(got, blockfrost.Account{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourceAccountRewardsHistoryIntegration(t *testing.T) {
	t.Parallel()
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{
		Count: 1,
	}
	got, err := api.AccountRewardsHistory(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if reflect.DeepEqual(got, []blockfrost.AccountRewardsHistory{}) {
		t.Fatalf("got null %+v", got)
	}

}

func TestResourceAccountHistoryIntegration(t *testing.T) {
	t.Parallel()

	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{
		Count: 1,
	}
	got, err := api.AccountHistory(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if reflect.DeepEqual(got, []blockfrost.AccountHistory{}) {
		t.Fatalf("got null %+v", got)
	}

}

func TestResourceAccountDelegationHistoryIntregration(t *testing.T) {
	t.Parallel()

	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)
	q := blockfrost.APIQueryParams{
		Count: 1,
	}
	got, err := api.AccountDelegationHistory(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if reflect.DeepEqual(got, []blockfrost.AccountDelegationHistory{}) {
		t.Fatalf("got null %+v", got)
	}

}

func TestResourceAccountRegistrationHistoryIntegration(t *testing.T) {
	t.Parallel()
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{
		Count: 1,
	}
	got, err := api.AccountRegistrationHistory(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if reflect.DeepEqual(got, []blockfrost.AccountRegistrationHistory{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourceAccountWithdrawalHistoryIntegration(t *testing.T) {
	t.Parallel()
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{
		Count: 1,
	}
	got, err := api.AccountWithdrawalHistory(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if reflect.DeepEqual(got, []blockfrost.AccountWithdrawalHistory{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourceAccountMIRHistoryIntegration(t *testing.T) {
	t.Parallel()
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{
		Count: 1,
	}
	got, err := api.AccountMIRHistory(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if reflect.DeepEqual(got, []blockfrost.AccountMIRHistory{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourceAccountAssociatedAddressesIntegration(t *testing.T) {
	t.Parallel()
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{
		Count: 1,
	}
	got, err := api.AccountAssociatedAddresses(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if reflect.DeepEqual(got, []blockfrost.AccountAssociatedAddress{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourceAccountAssociatedAssetsIntegration(t *testing.T) {
	t.Parallel()

	inputStakeAddr := "stake1u9ylzsgxaa6xctf4juup682ar3juj85n8tx3hthnljg47zctvm3rc"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{
		Count: 1,
	}
	_, err := api.AccountAssociatedAssets(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
}
