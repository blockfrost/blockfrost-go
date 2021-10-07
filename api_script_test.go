package blockfrost_test

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
		},
	}
	fp := filepath.Join(testdata, "json", "script", "redeemer.json")
	got := []blockfrost.ScriptRedeemer{}

	testStructGotWant(t, fp, &got, &want)
}

func TestIntegrationResourceRedeemers(t *testing.T) {
	fp := filepath.Join(testdata, "json", "script", "redeemer.json")
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatal(err)
	}
	want := []blockfrost.ScriptRedeemer{}
	if err = json.Unmarshal(data, &want); err != nil {
		t.Fatal(err)
	}
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(data)
	}))

	api, err := blockfrost.NewAPIClient(blockfrost.APIClientOptions{Server: s.URL})
	if err != nil {
		t.Fatal(err)
	}

	addr := "e1457a0c47dfb7a2f6b8fbb059bdceab163c05d34f195b87b9f2b30e"
	got, err := api.ScriptRedeemers(context.TODO(), addr, blockfrost.APIQueryParams{})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected  %v got %v", want, got)
	}
}

func TestIntegrationResourceScripts(t *testing.T) {
	fp := filepath.Join(testdata, "json", "script", "scripts.json")
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatal(err)
	}
	want := []blockfrost.Script{}
	if err = json.Unmarshal(data, &want); err != nil {
		t.Fatal(err)
	}
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(data)
	}))
	api, err := blockfrost.NewAPIClient(blockfrost.APIClientOptions{Server: s.URL})
	if err != nil {
		t.Fatal(err)
	}

	got, err := api.Scripts(context.TODO(), blockfrost.APIQueryParams{})
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected  %v got %v", want, got)
	}
}

func TestIntegrationResourceScript(t *testing.T) {
	fp := filepath.Join(testdata, "json", "script", "script.json")
	data, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatal(err)
	}
	want := blockfrost.Script{}
	if err = json.Unmarshal(data, &want); err != nil {
		t.Fatal(err)
	}
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(data)
	}))
	api, err := blockfrost.NewAPIClient(blockfrost.APIClientOptions{Server: s.URL})
	if err != nil {
		t.Fatal(err)
	}

	got, err := api.Script(context.TODO(), "")
	if err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected  %v got %v", want, got)
	}
}
