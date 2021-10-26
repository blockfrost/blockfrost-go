package blockfrost_test

import (
	"context"
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
	got := blockfrost.NetworkInfo{}

	testStructGotWant(t, fp, &got, &want)
}

func TestResourceNetworkIntegration(t *testing.T) {
	nullGot := blockfrost.NetworkInfo{}

	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	got, err := api.Network(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(got, nullGot) {
		t.Fatalf("got null %+v", got)
	}

}
