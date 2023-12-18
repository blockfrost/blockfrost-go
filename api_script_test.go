package blockfrost_test

import (
	"context"
	"encoding/json"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestScriptUnmarshal(t *testing.T) {
	want := []blockfrost.Script{
		{ScriptHash: "13a3efd825703a352a8f71f4e2758d08c28c564e8dfcce9f77776ad1"},
		{ScriptHash: "e1457a0c47dfb7a2f6b8fbb059bdceab163c05d34f195b87b9f2b30e"},
		{ScriptHash: "a6e63c0ff05c96943d1cc30bf53112ffff0f34b45986021ca058ec54"},
	}
	fp := filepath.Join(testdata, "json", "script", "scripts.json")

	got := []blockfrost.Script{}
	testStructGotWant(t, fp, &got, &want)

}

func testStructGotWant(t *testing.T, fp string, got interface{}, want interface{}) {
	bytes, err := os.ReadFile(fp)
	if err != nil {
		t.Fatalf("failed to open json test file %s with err %v", fp, err)
	}

	if err := json.Unmarshal(bytes, got); err != nil {
		t.Fatalf("failed to unmarshal %s with err %v", fp, err)
	}
	if err := json.Unmarshal(bytes, &got); err != nil {
		t.Fatalf("failed to unmarshal %s with err %v", fp, err)
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v got %v", want, got)
	}

}

func TestUnmarshalRedeemer(t *testing.T) {
	want := []blockfrost.ScriptRedeemer{
		{
			TxHash:    "1a0570af966fb355a7160e4f82d5a80b8681b7955f5d44bec0dce628516157f0",
			TxIndex:   0,
			Purpose:   "spend",
			UnitMem:   "1700",
			UnitSteps: "476468",
			Fee:       "172033",
			DatumHash: "923918e403bf43c34b4ef6b48eb2ee04babed17320d8d1b9ff9ad086e86f44ec",
		},
	}
	fp := filepath.Join(testdata, "json", "script", "redeemer.json")
	got := []blockfrost.ScriptRedeemer{}

	testStructGotWant(t, fp, &got, &want)
}

func TestIntegrationResourceRedeemers(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.ScriptRedeemers(context.TODO(), "4f590a3d80ae0312bad0b64d540c3ff5080e77250e9dbf5011630016", blockfrost.APIQueryParams{})
	if err != nil {
		t.Fatal(err)
	}

	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	want := []blockfrost.ScriptRedeemer{}

	testIntUtil(t, fp, &got, &want)
}

func TestIntegrationResourceScripts(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	got, err := api.Scripts(context.TODO(), blockfrost.APIQueryParams{})
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(got, []blockfrost.Script{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestIntegrationResourceScript(t *testing.T) {
	api := blockfrost.NewAPIClient(blockfrost.APIClientOptions{})

	scripts, err := api.Scripts(context.TODO(), blockfrost.APIQueryParams{})
	testErrorHelper(t, err)
	if len(scripts) == 0 {
		t.Fatal("Failed to fetch scripts")
	}
	got, err := api.Script(context.TODO(), scripts[0].ScriptHash)
	if err != nil {
		t.Fatal(err)
	}

	if reflect.DeepEqual(got, blockfrost.Script{}) {
		t.Fatalf("got null %+v", got)
	}
}
