package blockfrost_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/blockfrost/blockfrost-go"
	"github.com/stretchr/testify/assert"
)

func TestResourcePools(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.Write([]byte(
				`[
					"pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy",
					"pool1hn7hlwrschqykupwwrtdfkvt2u4uaxvsgxyh6z63703p2knj288",
					"pool1ztjyjfsh432eqetadf82uwuxklh28xc85zcphpwq6mmezavzad2"
				]`,
			))
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{}
	pools, err := api.Pools(context.TODO(), q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, inputPoolID, pools[0])
}

func TestResourcePoolsRetired(t *testing.T) {
	t.Parallel()
	expectedElement := blockfrost.PoolRetired{
		PoolID: "pool19u64770wqp6s95gkajc8udheske5e6ljmpq33awxk326zjaza0q",
		Epoch:  225,
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.Write([]byte(
				`[
					{
						"pool_id": "pool19u64770wqp6s95gkajc8udheske5e6ljmpq33awxk326zjaza0q",
						"epoch": 225
					},
					{
						"pool_id": "pool1dvla4zq98hpvacv20snndupjrqhuc79zl6gjap565nku6et5zdx",
						"epoch": 215
					}
				]`,
			))
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{}
	pools, err := api.PoolsRetired(context.TODO(), q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, pools[0])
}

func TestResourcePoolsRetiring(t *testing.T) {
	t.Parallel()
	expectedElement := blockfrost.PoolRetiring{
		PoolID: "pool19u64770wqp6s95gkajc8udheske5e6ljmpq33awxk326zjaza0q",
		Epoch:  225,
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.Write([]byte(
				`[
					{
						"pool_id": "pool19u64770wqp6s95gkajc8udheske5e6ljmpq33awxk326zjaza0q",
						"epoch": 225
					},
					{
						"pool_id": "pool1dvla4zq98hpvacv20snndupjrqhuc79zl6gjap565nku6et5zdx",
						"epoch": 215
					}
				]`,
			))
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{}
	pools, err := api.PoolsRetiring(context.TODO(), q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, pools[0])
}

func TestResourcePoolSpecific(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	expectedElement := blockfrost.PoolHistory{
		Epoch:           233,
		Blocks:          22,
		ActiveStake:     "20485965693569",
		ActiveSize:      1.2345,
		DelegatorsCount: 115,
		Rewards:         "206936253674159",
		Fees:            "1290968354",
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.Write([]byte(
				`[
					{
					  "epoch": 233,
					  "blocks": 22,
					  "active_stake": "20485965693569",
					  "active_size": 1.2345,
					  "delegators_count": 115,
					  "rewards": "206936253674159",
					  "fees": "1290968354"
					}
				  ]`,
			))
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{}
	pool, err := api.PoolHistory(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, pool[0])
}

func TestResourcePoolHistory(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	expectedElement := blockfrost.PoolHistory{
		Epoch:           233,
		Blocks:          22,
		ActiveStake:     "20485965693569",
		ActiveSize:      1.2345,
		DelegatorsCount: 115,
		Rewards:         "206936253674159",
		Fees:            "1290968354",
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.Write([]byte(
				`[
					{
					  "epoch": 233,
					  "blocks": 22,
					  "active_stake": "20485965693569",
					  "active_size": 1.2345,
					  "delegators_count": 115,
					  "rewards": "206936253674159",
					  "fees": "1290968354"
					}
				  ]`,
			))
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{}
	pool, err := api.PoolHistory(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, pool[0])
}

func TestResourcePoolMetadata(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	expectedElement := blockfrost.PoolMetadata{
		PoolID:      inputPoolID,
		Hex:         "0f292fcaa02b8b2f9b3c8f9fd8e0bb21abedb692a6d5058df3ef2735",
		URL:         "https://stakenuts.com/mainnet.json",
		Hash:        "47c0c68cb57f4a5b4a87bad896fc274678e7aea98e200fa14a1cb40c0cab1d8c",
		Ticker:      "NUTS",
		Name:        "Stake Nuts",
		Description: "The best pool ever",
		Homepage:    "https://stakentus.com/",
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.Write([]byte(
				`{
					"pool_id": "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy",
					"hex": "0f292fcaa02b8b2f9b3c8f9fd8e0bb21abedb692a6d5058df3ef2735",
					"url": "https://stakenuts.com/mainnet.json",
					"hash": "47c0c68cb57f4a5b4a87bad896fc274678e7aea98e200fa14a1cb40c0cab1d8c",
					"ticker": "NUTS",
					"name": "Stake Nuts",
					"description": "The best pool ever",
					"homepage": "https://stakentus.com/"
				  }`,
			))
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{}
	pool, err := api.PoolMetadata(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, pool)
}
func TestResourcePoolRelay(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	expectedElement := blockfrost.PoolRelay{
		Ipv4:   "4.4.4.4",
		Ipv6:   "https://stakenuts.com/mainnet.json",
		DNS:    "relay1.stakenuts.com",
		DNSSrv: "_relays._tcp.relays.stakenuts.com",
		Port:   3001,
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.Write([]byte(
				`[
					{
					  "ipv4": "4.4.4.4",
					  "ipv6": "https://stakenuts.com/mainnet.json",
					  "dns": "relay1.stakenuts.com",
					  "dns_srv": "_relays._tcp.relays.stakenuts.com",
					  "port": 3001
					}
				  ]`,
			))
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{}
	pool, err := api.PoolRelay(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, pool[0])
}
func TestResourcePoolDelegator(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	expectedElement := blockfrost.PoolDelegator{
		Address:   "stake1ux4vspfvwuus9uwyp5p3f0ky7a30jq5j80jxse0fr7pa56sgn8kha",
		LiveStake: "1137959159981411",
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.Write([]byte(
				`[
					{
					  "address": "stake1ux4vspfvwuus9uwyp5p3f0ky7a30jq5j80jxse0fr7pa56sgn8kha",
					  "live_stake": "1137959159981411"
					},
					{
					  "address": "stake1uylayej7esmarzd4mk4aru37zh9yz0luj3g9fsvgpfaxulq564r5u",
					  "live_stake": "16958865648"
					},
					{
					  "address": "stake1u8lr2pnrgf8f7vrs9lt79hc3sxm8s2w4rwvgpncks3axx6q93d4ck",
					  "live_stake": "18605647"
					}
				  ]`,
			))
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{}
	pool, err := api.PoolDelegator(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, pool[0])
}

func TestResourcePoolBlocks(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	expectedElement := blockfrost.PoolBlocks{
		"d8982ca42cfe76b747cc681d35d671050a9e41e9cfe26573eb214e94fe6ff21d",
		"026436c539e2ce84c7f77ffe669f4e4bbbb3b9c53512e5857dcba8bb0b4e9a8c",
		"bcc8487f419b8c668a18ea2120822a05df6dfe1de1f0fac3feba88cf760f303c",
		"86bf7b4a274e0f8ec9816171667c1b4a0cfc661dc21563f271acea9482b62df7",
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.Write([]byte(
				`[
					"d8982ca42cfe76b747cc681d35d671050a9e41e9cfe26573eb214e94fe6ff21d",
					"026436c539e2ce84c7f77ffe669f4e4bbbb3b9c53512e5857dcba8bb0b4e9a8c",
					"bcc8487f419b8c668a18ea2120822a05df6dfe1de1f0fac3feba88cf760f303c",
					"86bf7b4a274e0f8ec9816171667c1b4a0cfc661dc21563f271acea9482b62df7"
				  ]`,
			))
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{}
	pool, err := api.PoolBlocks(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, pool)
}

func TestResourcePoolUpdate(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	expectedElement := blockfrost.PoolUpdate{
		TxHash:    "6804edf9712d2b619edb6ac86861fe93a730693183a262b165fcc1ba1bc99cad",
		CERTIndex: 0,
		Action:    "registered",
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			w.Write([]byte(
				`[
					{
					  "tx_hash": "6804edf9712d2b619edb6ac86861fe93a730693183a262b165fcc1ba1bc99cad",
					  "cert_index": 0,
					  "action": "registered"
					},
					{
					  "tx_hash": "9c190bc1ac88b2ab0c05a82d7de8b71b67a9316377e865748a89d4426c0d3005",
					  "cert_index": 0,
					  "action": "deregistered"
					},
					{
					  "tx_hash": "e14a75b0eb2625de7055f1f580d70426311b78e0d36dd695a6bdc96c7b3d80e0",
					  "cert_index": 1,
					  "action": "registered"
					}
				  ]`,
			))
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatalf(err.Error())
	}
	q := blockfrost.APIPagingParams{}
	pool, err := api.PoolUpdate(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatalf(err.Error())
	}
	assert.Equal(t, expectedElement, pool[0])
}
