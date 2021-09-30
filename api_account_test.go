package blockfrost_test

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/blockfrost/blockfrost-go"
	"github.com/stretchr/testify/assert"
)

func TestResourceAccount(t *testing.T) {
	t.Parallel()

	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			fp := filepath.Join(testdata, "json", "account", "account.json")
			file, err := ioutil.ReadFile(fp)
			if err != nil {
				t.Fatalf("an error ocurred while trying to read json test file %s", fp)
			}
			w.Write(file)
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	accountAddr, err := api.Account(context.TODO(), inputStakeAddr)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, inputStakeAddr, accountAddr.StakeAddress)
}

func TestResourceAccountRewardsHistory(t *testing.T) {
	t.Parallel()

	expectedElement := blockfrost.AccountRewardsHistory{
		Epoch:  215,
		Amount: "12695385",
		PoolID: "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			fp := filepath.Join(testdata, "json", "account", "account_rewards_history.json")
			file, err := ioutil.ReadFile(fp)
			if err != nil {
				t.Fatalf("an error ocurred while trying to read json test file %s", fp)
			}
			w.Write(file)
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{
		Count: 1,
	}
	accountArray, err := api.AccountRewardsHistory(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])

}

func TestResourceAccountHistory(t *testing.T) {
	t.Parallel()

	expectedElement := blockfrost.AccountHistory{
		ActiveEpoch: 210,
		Amount:      "12695385",
		PoolID:      "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			fp := filepath.Join(testdata, "json", "account", "account_history.json")
			file, err := ioutil.ReadFile(fp)
			if err != nil {
				t.Fatalf("an error ocurred while trying to read json test file %s", fp)
			}
			w.Write(file)
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{
		Count: 1,
	}
	accountArray, err := api.AccountHistory(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])

}

func TestResourceAccountDelegationHistory(t *testing.T) {
	t.Parallel()

	expectedElement := blockfrost.AccountDelegationHistory{
		ActiveEpoch: 210,
		TXHash:      "2dd15e0ef6e6a17841cb9541c27724072ce4d4b79b91e58432fbaa32d9572531",
		Amount:      "12695385",
		PoolID:      "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			fp := filepath.Join(testdata, "json", "account", "account_delegation_history.json")
			file, err := ioutil.ReadFile(fp)
			if err != nil {
				t.Fatalf("an error ocurred while trying to read json test file %s", fp)
			}
			w.Write(file)
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{
		Count: 1,
	}
	accountArray, err := api.AccountDelegationHistory(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])

}

func TestResourceAccountRegistrationHistory(t *testing.T) {
	t.Parallel()

	expectedElement := blockfrost.AccountRegistrationHistory{
		TXHash: "2dd15e0ef6e6a17841cb9541c27724072ce4d4b79b91e58432fbaa32d9572531",
		Action: "registered",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			fp := filepath.Join(testdata, "json", "account", "account_registration_history.json")
			file, err := ioutil.ReadFile(fp)
			if err != nil {
				t.Fatalf("an error ocurred while trying to read json test file %s", fp)
			}
			w.Write(file)
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{
		Count: 1,
	}
	accountArray, err := api.AccountRegistrationHistory(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])

}

func TestResourceAccountWithdrawalHistory(t *testing.T) {
	t.Parallel()

	expectedElement := blockfrost.AccountWithdrawalHistory{
		TXHash: "48a9625c841eea0dd2bb6cf551eabe6523b7290c9ce34be74eedef2dd8f7ecc5",
		Amount: "454541212442",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			fp := filepath.Join(testdata, "json", "account", "account_withdrawal_history.json")
			file, err := ioutil.ReadFile(fp)
			if err != nil {
				t.Fatalf("an error ocurred while trying to read json test file %s", fp)
			}
			w.Write(file)
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{
		Count: 1,
	}
	accountArray, err := api.AccountWithdrawalHistory(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])

}

func TestResourceAccountMIRHistory(t *testing.T) {
	t.Parallel()

	expectedElement := blockfrost.AccountMIRHistory{
		TXHash: "69705bba1d687a816ff5a04ec0c358a1f1ef075ab7f9c6cc2763e792581cec6d",
		Amount: "2193707473",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			fp := filepath.Join(testdata, "json", "account", "account_mir_history.json")
			file, err := ioutil.ReadFile(fp)
			if err != nil {
				t.Fatalf("an error ocurred while trying to read json test file %s", fp)
			}
			w.Write(file)
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{
		Count: 1,
	}
	accountArray, err := api.AccountMIRHistory(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])

}

func TestResourceAccountAssociatedAddresses(t *testing.T) {
	t.Parallel()

	expectedElement := blockfrost.AccountAssociatedAddress{
		Address: "addr1qx2kd28nq8ac5prwg32hhvudlwggpgfp8utlyqxu6wqgz62f79qsdmm5dsknt9ecr5w468r9ey0fxwkdrwh08ly3tu9sy0f4qd",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			fp := filepath.Join(testdata, "json", "account", "account_associated_addresses.json")
			file, err := ioutil.ReadFile(fp)
			if err != nil {
				t.Fatalf("an error ocurred while trying to read json test file %s", fp)
			}
			w.Write(file)
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{
		Count: 1,
	}
	accountArray, err := api.AccountAssociatedAddresses(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])

}

func TestResourceAccountAssociatedAssets(t *testing.T) {
	t.Parallel()

	expectedElement := blockfrost.AccountAssociatedAsset{
		Unit:     "d5e6bf0500378d4f0da4e8dde6becec7621cd8cbf5cbb9b87013d4cc537061636542756433343132",
		Quantity: "1",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			fp := filepath.Join(testdata, "json", "account", "account_associated_assets.json")
			file, err := ioutil.ReadFile(fp)
			if err != nil {
				t.Fatalf("an error ocurred while trying to read json test file %s", fp)
			}
			w.Write(file)
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{
		Count: 1,
	}
	accountArray, err := api.AccountAssociatedAssets(context.TODO(), inputStakeAddr, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])

}
