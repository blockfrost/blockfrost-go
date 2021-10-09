package blockfrost_test

import (
	"path/filepath"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestIPFSObjectUnmarshal(t *testing.T) {
	fp := filepath.Join(testdata, "json", "ipfs", "ipfs_object.json")
	want := blockfrost.IPFSObject{
		Name:     "README.md",
		IPFSHash: "QmZbHqiCxKEVX7QfijzJTkZiSi3WEVTcvANgNAWzDYgZDr",
		Size:     "125297",
	}
	got := blockfrost.IPFSObject{}
	testStructGotWant(t, fp, &got, &want)
}
