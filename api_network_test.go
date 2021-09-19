package blockfrost_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestNetworkUnmarshal(t *testing.T) {
	want := blockfrost.NetworkInfo{
		Supply: blockfrost.NetworkSupply{
			Max:         "45000000000000000",
			Total:       "32890715183299160",
			Circulating: "32412601976210393",
			Locked:      "125006953355",
		},
		Stake: blockfrost.NetworkStake{
			Live:   "23204950463991654",
			Active: "22210233523456321",
		},
	}

	fp := filepath.Join(testdata, "json", "network", "network.json")
	bytes, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatalf("an error ocurred while trying to read json test file %s", fp)
	}

	got := blockfrost.NetworkInfo{}
	if err := json.Unmarshal(bytes, &got); err != nil {
		t.Fatalf("failed to unmarshal %s with err %v", fp, err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v got %v", want, got)
	}
}
