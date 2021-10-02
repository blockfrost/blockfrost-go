package blockfrost_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
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
			fp := filepath.Join(testdata, "json", "pool", "pools.json")
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
		t.Fatal(err)
	}
	q := blockfrost.APIPagingParams{}
	pools, err := api.Pools(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
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
			fp := filepath.Join(testdata, "json", "pool", "pools_retired.json")
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
		t.Fatal(err)
	}
	q := blockfrost.APIPagingParams{}
	pools, err := api.PoolsRetired(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
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
			fp := filepath.Join(testdata, "json", "pool", "pools_retiring.json")
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
		t.Fatal(err)
	}
	q := blockfrost.APIPagingParams{}
	pools, err := api.PoolsRetiring(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedElement, pools[0])
}

func TestResourcePoolSpecific(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	expectedElement := blockfrost.Pool{
		PoolID:         inputPoolID,
		Hex:            "0f292fcaa02b8b2f9b3c8f9fd8e0bb21abedb692a6d5058df3ef2735",
		VrfKey:         "0b5245f9934ec2151116fb8ec00f35fd00e0aa3b075c4ed12cce440f999d8233",
		BlocksMinted:   69,
		LiveStake:      "6900000000",
		LiveSize:       0.42,
		LiveSaturation: 0.93,
		LiveDelegators: 127,
		ActiveStake:    "4200000000",
		ActiveSize:     0.43,
		DeclaredPledge: "5000000000",
		LivePledge:     "5000000001",
		MarginCost:     0.05,
		FixedCost:      "340000000",
		RewardAccount:  "stake1uxkptsa4lkr55jleztw43t37vgdn88l6ghclfwuxld2eykgpgvg3f",
		Owners:         []string{"stake1u98nnlkvkk23vtvf9273uq7cph5ww6u2yq2389psuqet90sv4xv9v"},
		Registration: []string{
			"9f83e5484f543e05b52e99988272a31da373f3aab4c064c76db96643a355d9dc",
			"7ce3b8c433bf401a190d58c8c483d8e3564dfd29ae8633c8b1b3e6c814403e95",
			"3e6e1200ce92977c3fe5996bd4d7d7e192bcb7e231bc762f9f240c76766535b9",
		},
		Retirement: []string{"252f622976d39e646815db75a77289cf16df4ad2b287dd8e3a889ce14c13d1a8"},
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			fp := filepath.Join(testdata, "json", "pool", "pools_specific.json")
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
		t.Fatal(err)
	}
	pool, err := api.Pool(context.TODO(), inputPoolID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedElement, pool)
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
			fp := filepath.Join(testdata, "json", "pool", "pools_history.json")
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
		t.Fatal(err)
	}
	q := blockfrost.APIPagingParams{}
	got, err := api.PoolHistory(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedElement, got[0])
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
			fp := filepath.Join(testdata, "json", "pool", "pools_metadata.json")
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
		t.Fatal(err)
	}
	pool, err := api.PoolMetadata(context.TODO(), inputPoolID)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedElement, pool)
}
func TestResourcePoolRelays(t *testing.T) {
	t.Parallel()
	want := []blockfrost.PoolRelay{{
		Ipv4:   "4.4.4.4",
		Ipv6:   "https://stakenuts.com/mainnet.json",
		DNS:    "relay1.stakenuts.com",
		DNSSrv: "_relays._tcp.relays.stakenuts.com",
		Port:   3001,
	},
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			json.NewEncoder(w).Encode(want)
		}),
	)
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatal(err)
	}
	got, err := api.PoolRelays(context.TODO(), "")
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, want, got)
}
func TestResourcePoolDelegators(t *testing.T) {
	t.Parallel()
	inputPoolID := "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2q3lkdy"
	expectedElement := blockfrost.PoolDelegator{
		Address:   "stake1ux4vspfvwuus9uwyp5p3f0ky7a30jq5j80jxse0fr7pa56sgn8kha",
		LiveStake: "1137959159981411",
	}
	s := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
			time.Sleep(20 * time.Millisecond)
			fp := filepath.Join(testdata, "json", "pool", "pools_delegator.json")
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
		t.Fatal(err)
	}
	q := blockfrost.APIPagingParams{}
	pool, err := api.PoolDelegators(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatal(err)
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
			fp := filepath.Join(testdata, "json", "pool", "pools_block.json")
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
		t.Fatal(err)
	}
	q := blockfrost.APIPagingParams{}
	pool, err := api.PoolBlocks(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatal(err)
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
			fp := filepath.Join(testdata, "json", "pool", "pools_update.json")
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
		t.Fatal(err)
	}
	q := blockfrost.APIPagingParams{}
	pool, err := api.PoolUpdates(context.TODO(), inputPoolID, q)
	if err != nil {
		t.Fatal(err)
	}
	assert.Equal(t, expectedElement, pool[0])
}
