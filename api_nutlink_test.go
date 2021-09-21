package blockfrost_test

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestNutlinkAddressUnMarshall(t *testing.T) {
	want := blockfrost.NutlinkAddress{
		Address:      "addr1qxqs59lphg8g6qndelq8xwqn60ag3aeyfcp33c2kdp46a09re5df3pzwwmyq946axfcejy5n4x0y99wqpgtp2gd0k09qsgy6pz",
		MetadataUrl:  "https://nut.link/metadata.json",
		MetadataHash: "6bf124f217d0e5a0a8adb1dbd8540e1334280d49ab861127868339f43b3948af",
		Metadata:     blockfrost.NutlinkAddressMeta{},
	}
	fp := filepath.Join(testdata, "json", "nutlink", "address_meta.json")
	bytes, err := ioutil.ReadFile(fp)
	if err != nil {
		t.Fatalf("an error ocurred while trying to read json test file %s", fp)
	}
	got := blockfrost.NutlinkAddress{}
	if err = json.Unmarshal(bytes, &got); err != nil {
		t.Fatalf("failed to unmarshal %s with err %v", fp, err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v got %v", want, got)
	}

	log.Printf("Got: %+v\nWant: %+v", got, want)

}
