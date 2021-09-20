package blockfrost_test

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"reflect"
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
	bytes, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatalf("an error ocurred while trying to read json test file %s", fp)
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
		},
	}
	fp := filepath.Join(testdata, "json", "script", "redeemer.json")
	got := []blockfrost.ScriptRedeemer{}

	testStructGotWant(t, fp, &got, &want)
}
