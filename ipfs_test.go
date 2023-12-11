package blockfrost_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestIPFSClientInit(t *testing.T) {
	tests := []struct {
		name   string
		option blockfrost.IPFSClientOptions
		want   blockfrost.IPFSClientOptions
	}{
		{
			"test get env from key",
			blockfrost.IPFSClientOptions{},
			blockfrost.IPFSClientOptions{ProjectID: os.Getenv("BLOCKFROST_IPFS_PROJECT_ID")},
		},
		{
			"test get env from option",
			blockfrost.IPFSClientOptions{ProjectID: "Q."},
			blockfrost.IPFSClientOptions{ProjectID: "J."},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			s := httptest.NewServer(
				http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
					projectId := r.Header.Get("Project-Id")

					got := blockfrost.IPFSClientOptions{
						ProjectID: projectId,
					}

					if !reflect.DeepEqual(got, tt.want) {
						t.Fatalf("expected %v got %v", tt.want, got)
					}
				}),
			)
			tt.option.Server = s.URL
			c := blockfrost.NewIPFSClient(
				tt.option,
			)
			_, _ = c.Add(context.TODO(), "")
		})
	}
}

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

func TestIPFSResourceAddIntegration(t *testing.T) {
	ipfs := blockfrost.NewIPFSClient(blockfrost.IPFSClientOptions{})

	fp_up := filepath.Join(testdata, "json", "ipfs", "ipfs_object.json")
	got, err := ipfs.Add(context.TODO(), fp_up)
	testErrorHelper(t, err)
	// *update = true

	// want := blockfrost.IPFSObject{}
	// fp_golden := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(t.Name(), "Test"))+".golden")
	//testIntUtil(t, fp_golden, &got, &want)
	t.Run("TestIPFSResourcePinIntegration", func(ti *testing.T) {
		testIPFSResourcePinIntegration(ti, ipfs, got.IPFSHash)
	})
	t.Run("TestIPFSResourcePinnedIntegration", func(ti *testing.T) {
		testIPFSResourcePinnnedIntegration(ti, ipfs, got.IPFSHash)
	})
	t.Run("TestIPFSResourcePinnedObjectsIntegration", func(ti *testing.T) {
		testIPFSResourcePinnedObjectsIntegration(ti, ipfs)
	})
	t.Run("TestIPFSResourceGatewayIntegration", func(ti *testing.T) {
		testIPFSResourceGatewayIntegration(ti, ipfs, got.IPFSHash, fp_up)
	})
	t.Run("TestIPFSResourceRemoveIntegration", func(ti *testing.T) {
		testIPFSResourceRemoveIntegration(ti, ipfs, got.IPFSHash)
	})
	*update = false
}

func testIPFSResourcePinIntegration(t *testing.T, ipfs blockfrost.IPFSClient, ipfsPath string) {
	testNames := strings.Split(t.Name(), "/")
	got, err := ipfs.Pin(context.TODO(), ipfsPath)
	testErrorHelper(t, err)
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(testNames[len(testNames)-1], "Test"))+".golden")
	want := blockfrost.IPFSPinnedObject{}
	testIntUtil(t, fp, &got, &want)

}

func testIPFSResourcePinnnedIntegration(t *testing.T, ipfs blockfrost.IPFSClient, ipfsPath string) {
	testNames := strings.Split(t.Name(), "/")
	got, err := ipfs.PinnedObject(context.TODO(), ipfsPath)
	testErrorHelper(t, err)
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(testNames[len(testNames)-1], "Test"))+".golden")
	want := blockfrost.IPFSPinnedObject{}
	if *update {
		data, err := json.Marshal(got)
		if err != nil {
			t.Fatal(err)
		}
		WriteGoldenFile(t, fp, data)
	}
	bytes := ReadOrGenerateGoldenFile(t, fp, got)
	if err := json.Unmarshal(bytes, &want); err != nil {
		t.Fatal(err)
	}

	want.TimeCreated = 0
	want.TimePinned = 0

	got.TimeCreated = 0
	got.TimePinned = 0

	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v got %v", want, got)
	}
}

func testIPFSResourcePinnedObjectsIntegration(t *testing.T, ipfs blockfrost.IPFSClient) {
	testNames := strings.Split(t.Name(), "/")
	got, err := ipfs.PinnedObjects(context.TODO(), blockfrost.APIQueryParams{})
	testErrorHelper(t, err)
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(testNames[len(testNames)-1], "Test"))+".golden")
	want := []blockfrost.IPFSPinnedObject{}
	if *update {
		data, err := json.Marshal(got)
		if err != nil {
			t.Fatal(err)
		}
		WriteGoldenFile(t, fp, data)
	}
	bytes := ReadOrGenerateGoldenFile(t, fp, got)
	if err := json.Unmarshal(bytes, &want); err != nil {
		t.Fatal(err)
	}
	if len(got) != len(want) {
		t.Fatalf("want len and got len not matching")
	}

	for i, igot := range got {
		// igot := igot
		igot.TimeCreated = 0
		igot.TimePinned = 0
		igot.State = ""

		iwant := want[i]
		iwant.TimeCreated = 0
		iwant.TimePinned = 0
		iwant.State = ""

		if !reflect.DeepEqual(igot, iwant) {
			t.Fatalf("expected %v got %v", iwant, igot)
		}
	}

}

func testIPFSResourceRemoveIntegration(t *testing.T, ipfs blockfrost.IPFSClient, ipfsPath string) {
	testNames := strings.Split(t.Name(), "/")
	got, err := ipfs.Remove(context.TODO(), ipfsPath)
	testErrorHelper(t, err)
	fp := filepath.Join(testdata, strings.ToLower(strings.TrimPrefix(testNames[len(testNames)-1], "Test"))+".golden")
	want := blockfrost.IPFSObject{}
	testIntUtil(t, fp, &got, &want)
}

func testIPFSResourceGatewayIntegration(t *testing.T, ipfs blockfrost.IPFSClient, ipfsPath string, fp_up string) {
	got, err := ipfs.Gateway(context.TODO(), ipfsPath)
	testErrorHelper(t, err)
	want, err := os.ReadFile(fp_up)
	testErrorHelper(t, err)
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("expected %v got %v", want, got)
	}

}

func testErrorHelper(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
