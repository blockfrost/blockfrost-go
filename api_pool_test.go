package blockfrost_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestResourcePoolsIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{}
	got, err := api.Pools(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(got, blockfrost.Pools{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourcePoolsRetiredIntegration(t *testing.T) {
	t.Parallel()

	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{}
	got, err := api.PoolsRetired(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(got, []blockfrost.PoolRetired{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourcePoolsRetiringIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{}
	got, err := api.PoolsRetiring(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(got, []blockfrost.PoolRetiring{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourcePoolSpecificIntegration(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"

	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	got, err := api.Pool(context.TODO(), inputPoolID)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(got, blockfrost.Pool{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourcePoolHistoryIntegration(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{}
	got, err := api.PoolHistory(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(got, []blockfrost.PoolHistory{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourcePoolMetadataIntegration(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	got, err := api.PoolMetadata(context.TODO(), inputPoolID)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(got, blockfrost.PoolMetadata{}) {
		t.Fatalf("got null %+v", got)
	}
}
func TestResourcePoolRelaysIntegration(t *testing.T) {
	t.Parallel()

	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	got, err := api.PoolRelays(context.TODO(), inputPoolID)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(got, []blockfrost.PoolRelay{}) {
		t.Fatalf("got null %+v", got)
	}
}
func TestResourcePoolDelegatorsIntegration(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{}
	got, err := api.PoolDelegators(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(got, []blockfrost.PoolDelegator{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourcePoolBlocksIntegration(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{}
	got, err := api.PoolBlocks(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(got, blockfrost.PoolBlocks{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourcePoolUpdateIntegration(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{}
	got, err := api.PoolUpdates(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(got, []blockfrost.PoolUpdate{}) {
		t.Fatalf("got null %+v", got)
	}
}
