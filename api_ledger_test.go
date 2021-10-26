package blockfrost_test

import (
	"context"
	"encoding/json"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestGenesisUnmarshall(t *testing.T) {
	fp := filepath.Join(testdata, "json", "genesis", "genesis.json")
	want := blockfrost.GenesisBlock{
		ActiveSlotsCoefficient: 0.05,
		UpdateQuorum:           5,
		MaxLovelaceSupply:      "45000000000000000",
		NetworkMagic:           764824073,
		EpochLength:            432000,
		SystemStart:            1506203091,
		SlotsPerKesPeriod:      129600,
		SlotLength:             1,
		MaxKesEvolutions:       62,
		SecurityParam:          2160,
	}
	got := blockfrost.GenesisBlock{}
	testStructGotWant(t, fp, &got, &want)
}

func TestGenesisIntegration(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.Genesis(context.TODO())
	if err != nil {
		t.Fatal(err)
	}
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimLeft(t.Name(), "Test"))+".golden")
	if *update {
		data, err := json.Marshal(got)
		if err != nil {
			t.Fatal(err)
		}
		WriteGoldenFile(t, fp, data)
	}
	bytes := ReadOrGenerateGoldenFile(t, fp, got)
	want := blockfrost.GenesisBlock{}
	if err = json.Unmarshal(bytes, &want); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v got %v", want, got)
	}
}
