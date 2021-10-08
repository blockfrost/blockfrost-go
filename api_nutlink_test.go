package blockfrost_test

import (
	"path/filepath"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestNutlinkAddressUnMarshall(t *testing.T) {
	fp := filepath.Join(testdata, "json", "nutlink", "address.json")
	want := blockfrost.NutlinkAddress{
		Address:      "addr1qxqs59lphg8g6qndelq8xwqn60ag3aeyfcp33c2kdp46a09re5df3pzwwmyq946axfcejy5n4x0y99wqpgtp2gd0k09qsgy6pz",
		MetadataUrl:  "https://nut.link/metadata.json",
		MetadataHash: "6bf124f217d0e5a0a8adb1dbd8540e1334280d49ab861127868339f43b3948af",
		Metadata:     blockfrost.NutlinkAddressMeta{},
	}

	got := blockfrost.NutlinkAddress{}
	testStructGotWant(t, fp, &got, &want)
}
