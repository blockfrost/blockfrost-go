package blockfrostgo_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	blockfrostgo "github.com/blockfrost/blockfrost-go/pkg/api"

	"github.com/stretchr/testify/assert"
)

func Test_Account(t *testing.T) {
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.Write([]byte(
				`{
					"stake_address": "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7",
					"active": true,
					"active_epoch": 412,
					"controlled_amount": "619154618165",
					"rewards_sum": "319154618165",
					"withdrawals_sum": "12125369253",
					"reserves_sum": "319154618165",
					"treasury_sum": "12000000",
					"withdrawable_amount": "319154618165",
					"pool_id": "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
					}`,
			))
		}),
	)
	defer s.Close()
	client := blockfrostgo.NewBlockfrostAPI(
		ApiKey,
		s.URL,
		false,
		nil,
		nil,
	)
	accountAddr, err := client.Account(context.TODO(), inputStakeAddr)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, inputStakeAddr, accountAddr.StakeAddress)
}

func Test_Account_Rewards_History(t *testing.T) {
	expectedElement := blockfrostgo.AccountRewardsHist{
		Epoch:  215,
		Amount: "12695385",
		PoolID: "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.Write([]byte(
				`[
					{
					"epoch": 215,
					"amount": "12695385",
					"pool_id": "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
					},
					{
					"epoch": 216,
					"amount": "3586329",
					"pool_id": "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
					}
					]`,
			))
		}),
	)
	defer s.Close()
	client := blockfrostgo.NewBlockfrostAPI(
		ApiKey,
		s.URL,
		false,
		nil,
		nil,
	)
	q := blockfrostgo.QueryParmasAPI{
		Count: 1,
	}

	accountArray, err := client.AccountRewards(
		context.TODO(),
		inputStakeAddr,
		q,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])
}

func Test_Account_History(t *testing.T) {
	expectedElement := blockfrostgo.AccountHistory{
		ActiveEpoch: 210,
		Amount:      "12695385",
		PoolID:      "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.Write([]byte(
				`[
					{
					"active_epoch": 210,
					"amount": "12695385",
					"pool_id": "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
					},
					{
					"active_epoch": 211,
					"amount": "22695385",
					"pool_id": "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
					}
					]`,
			))
		}),
	)
	defer s.Close()
	client := blockfrostgo.NewBlockfrostAPI(
		ApiKey,
		s.URL,
		false,
		nil,
		nil,
	)
	q := blockfrostgo.QueryParmasAPI{
		Count: 1,
	}

	accountArray, err := client.AccountHistory(
		context.TODO(),
		inputStakeAddr,
		q,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])
}

func Test_Account_Delegations(t *testing.T) {
	expectedElement := blockfrostgo.AccountDelegation{
		ActiveEpoch: 210,
		TXHash:      "2dd15e0ef6e6a17841cb9541c27724072ce4d4b79b91e58432fbaa32d9572531",
		Amount:      "12695385",
		PoolID:      "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.Write([]byte(
				`[
					{
					"active_epoch": 210,
					"tx_hash": "2dd15e0ef6e6a17841cb9541c27724072ce4d4b79b91e58432fbaa32d9572531",
					"amount": "12695385",
					"pool_id": "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
					},
					{
					"active_epoch": 242,
					"tx_hash": "1a0570af966fb355a7160e4f82d5a80b8681b7955f5d44bec0dde628516157f0",
					"amount": "12691385",
					"pool_id": "pool1kchver88u3kygsak8wgll7htr8uxn5v35lfrsyy842nkscrzyvj"
					}
					]`,
			))
		}),
	)
	defer s.Close()
	client := blockfrostgo.NewBlockfrostAPI(
		ApiKey,
		s.URL,
		false,
		nil,
		nil,
	)
	q := blockfrostgo.QueryParmasAPI{
		Count: 1,
	}

	accountArray, err := client.AccountDelegations(
		context.TODO(),
		inputStakeAddr,
		q,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])
}

func Test_Account_Registration(t *testing.T) {
	expectedElement := blockfrostgo.AccountRegistration{
		TXHash: "2dd15e0ef6e6a17841cb9541c27724072ce4d4b79b91e58432fbaa32d9572531",
		Action: "registered",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.Write([]byte(
				`[
					{
					"tx_hash": "2dd15e0ef6e6a17841cb9541c27724072ce4d4b79b91e58432fbaa32d9572531",
					"action": "registered"
					},
					{
					"tx_hash": "1a0570af966fb355a7160e4f82d5a80b8681b7955f5d44bec0dde628516157f0",
					"action": "deregistered"
					}
					]`,
			))
		}),
	)
	defer s.Close()
	client := blockfrostgo.NewBlockfrostAPI(
		ApiKey,
		s.URL,
		false,
		nil,
		nil,
	)
	q := blockfrostgo.QueryParmasAPI{
		Count: 1,
	}

	accountArray, err := client.AccountRegistrations(
		context.TODO(),
		inputStakeAddr,
		q,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])
}

func Test_Account_Withdrawals(t *testing.T) {
	expectedElement := blockfrostgo.AccountWithdrawal{
		TXHash: "48a9625c841eea0dd2bb6cf551eabe6523b7290c9ce34be74eedef2dd8f7ecc5",
		Amount: "454541212442",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.Write([]byte(
				`[
					{
					"tx_hash": "48a9625c841eea0dd2bb6cf551eabe6523b7290c9ce34be74eedef2dd8f7ecc5",
					"amount": "454541212442"
					},
					{
					"tx_hash": "4230b0cbccf6f449f0847d8ad1d634a7a49df60d8c142bb8cc2dbc8ca03d9e34",
					"amount": "97846969"
					}
					]`,
			))
		}),
	)
	defer s.Close()
	client := blockfrostgo.NewBlockfrostAPI(
		ApiKey,
		s.URL,
		false,
		nil,
		nil,
	)
	q := blockfrostgo.QueryParmasAPI{
		Count: 1,
	}

	accountArray, err := client.AccountWithdrawals(
		context.TODO(),
		inputStakeAddr,
		q,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])
}

func Test_Account_MIR_History(t *testing.T) {
	expectedElement := blockfrostgo.AccountMIR{
		TXHash: "69705bba1d687a816ff5a04ec0c358a1f1ef075ab7f9c6cc2763e792581cec6d",
		Amount: "2193707473",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.Write([]byte(
				`[
					{
					"tx_hash": "69705bba1d687a816ff5a04ec0c358a1f1ef075ab7f9c6cc2763e792581cec6d",
					"amount": "2193707473"
					},
					{
					"tx_hash": "baaa77b63d4d7d2bb3ab02c9b85978c2092c336dede7f59e31ad65452d510c13",
					"amount": "14520198574"
					}
					]`,
			))
		}),
	)
	defer s.Close()
	client := blockfrostgo.NewBlockfrostAPI(
		ApiKey,
		s.URL,
		false,
		nil,
		nil,
	)
	q := blockfrostgo.QueryParmasAPI{
		Count: 1,
	}

	accountArray, err := client.AccountMIRHistory(
		context.TODO(),
		inputStakeAddr,
		q,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])
}

func Test_Account_Associated_Address(t *testing.T) {
	expectedElement := blockfrostgo.AccountAssets{
		Unit:     "d5e6bf0500378d4f0da4e8dde6becec7621cd8cbf5cbb9b87013d4cc537061636542756433343132",
		Quantity: "1",
	}
	inputStakeAddr := "stake1ux3g2c9dx2nhhehyrezyxpkstartcqmu9hk63qgfkccw5rqttygt7"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(50 * time.Millisecond)
			w.Write([]byte(
				`[
					{
					"unit": "d5e6bf0500378d4f0da4e8dde6becec7621cd8cbf5cbb9b87013d4cc537061636542756433343132",
					"quantity": "1"
					},
					{
					"unit": "b0d07d45fe9514f80213f4020e5a61241458be626841cde717cb38a76e7574636f696e",
					"quantity": "125"
					}
					]`,
			))
		}),
	)
	defer s.Close()
	client := blockfrostgo.NewBlockfrostAPI(
		ApiKey,
		s.URL,
		false,
		nil,
		nil,
	)
	q := blockfrostgo.QueryParmasAPI{
		Count: 1,
	}

	accountArray, err := client.AccountAssetsWithAddress(
		context.TODO(),
		inputStakeAddr,
		q,
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, accountArray[0])
}
