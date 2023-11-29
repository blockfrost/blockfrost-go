package blockfrost_test

import (
	"context"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestEpochLatestIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.EpochLatest(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	nullDef := blockfrost.Epoch{}
	if reflect.DeepEqual(got, nullDef) {
		t.Fatal(err)
	}
}

func TestLatestEpochParametersIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.LatestEpochParameters(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	nullDef := blockfrost.EpochParameters{}
	if reflect.DeepEqual(got, nullDef) {
		t.Fatal(err)
	}
}

func TestEpochIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.Epoch(context.TODO(), 225)
	if err != nil {
		t.Fatal(err)
	}
	want := blockfrost.Epoch{}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	testIntUtil(t, fp, &got, &want)
}

func TestEpochNextIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.EpochsNext(context.TODO(), 225, blockfrost.APIQueryParams{Count: 5})
	if err != nil {
		t.Fatal(err)
	}
	want := []blockfrost.Epoch{}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	testIntUtil(t, fp, &got, &want)
}

func TestEpochsPreviousIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.EpochsPrevious(context.TODO(), 225, blockfrost.APIQueryParams{Count: 5})
	if err != nil {
		t.Fatal(err)
	}
	want := []blockfrost.Epoch{}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	testIntUtil(t, fp, &got, &want)
}

func TestEpochStakeDistributionIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.EpochStakeDistribution(context.TODO(), 225, blockfrost.APIQueryParams{Count: 5})
	if err != nil {
		t.Fatal(err)
	}
	want := []blockfrost.EpochStake{}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	testIntUtil(t, fp, &got, &want)
}

func TestEpochStakeDistributionByPoolIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	pool := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	got, err := api.EpochStakeDistributionByPool(context.TODO(), 225, pool, blockfrost.APIQueryParams{Count: 5})
	if err != nil {
		t.Fatal(err)
	}
	want := []blockfrost.EpochStakeByPool{}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	testIntUtil(t, fp, &got, &want)
}

func TestEpochBlockDistributionIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.EpochBlockDistribution(context.TODO(), 225, blockfrost.APIQueryParams{Count: 5})
	if err != nil {
		t.Fatal(err)
	}
	want := []string{}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	testIntUtil(t, fp, &got, &want)
}

func TestEpochBlockDistributionByPoolIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	pool := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	got, err := api.EpochBlockDistributionByPool(context.TODO(), 225, pool, blockfrost.APIQueryParams{Count: 5})
	if err != nil {
		t.Fatal(err)
	}
	want := []string{}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	testIntUtil(t, fp, &got, &want)
}

func TestEpochParametersIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.EpochParameters(context.TODO(), 225)
	if err != nil {
		t.Fatal(err)
	}
	want := blockfrost.EpochParameters{}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	testIntUtil(t, fp, &got, &want)
}
