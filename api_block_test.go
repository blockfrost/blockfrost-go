package blockfrost_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestResourceBlocksLatest(t *testing.T) {
	t.Parallel()
	blockHash := "4ea1ba291e8eef538635a53e59fddba7810d1679631cc3aed7c8e6c4091a516a"
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(
			[]byte(
				`{
					"time": 1641338934,
					"height": 15243593,
					"hash": "4ea1ba291e8eef538635a53e59fddba7810d1679631cc3aed7c8e6c4091a516a",
					"slot": 412162133,
					"epoch": 425,
					"epoch_slot": 12,
					"slot_leader": "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2qnikdy",
					"size": 3,
					"tx_count": 1,
					"output": "128314491794",
					"fees": "592661",
					"block_vrf": "vrf_vk1wf2k6lhujezqcfe00l6zetxpnmh9n6mwhpmhm0dvfh3fxgmdnrfqkms8ty",
					"previous_block": "43ebccb3ac72c7cebd0d9b755a4b08412c9f5dcb81b8a0ad1e3c197d29d47b05",
					"next_block": "8367f026cf4b03e116ff8ee5daf149b55ba5a6ec6dec04803b8dc317721d15fa",
					"confirmations": 4698
				  }`,
			))
	}))

	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatal(err)
	}

	block, err := api.BlockLatest(context.TODO())
	if err != nil {
		t.Fatal(err)
	}

	if block == (blockfrost.Block{}) {
		t.Fatal("Get nil block ")
	}

	if block.Hash != blockHash {
		t.Fatalf("Expected block.Hash to be %s, got %s", blockHash, block.Hash)
	}
}

func TestResourceBlock(t *testing.T) {
	t.Parallel()
	blockHash := "4ea1ba291e8eef538635a53e59fddba7810d1679631cc3aed7c8e6c4091a516a"
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(
			[]byte(
				`{
					"time": 1641338934,
					"height": 15243593,
					"hash": "4ea1ba291e8eef538635a53e59fddba7810d1679631cc3aed7c8e6c4091a516a",
					"slot": 412162133,
					"epoch": 425,
					"epoch_slot": 12,
					"slot_leader": "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2qnikdy",
					"size": 3,
					"tx_count": 1,
					"output": "128314491794",
					"fees": "592661",
					"block_vrf": "vrf_vk1wf2k6lhujezqcfe00l6zetxpnmh9n6mwhpmhm0dvfh3fxgmdnrfqkms8ty",
					"previous_block": "43ebccb3ac72c7cebd0d9b755a4b08412c9f5dcb81b8a0ad1e3c197d29d47b05",
					"next_block": "8367f026cf4b03e116ff8ee5daf149b55ba5a6ec6dec04803b8dc317721d15fa",
					"confirmations": 4698
					}`,
			))
	}))
	defer s.Close()
	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatal(err)
	}
	block, err := api.Block(context.TODO(), blockHash)
	if err != nil {
		t.Fatal(err)
	}

	if block.Hash != blockHash {
		t.Fatalf("Expected block.Hash to be %s, got %s", blockHash, block.Hash)
	}
}

func TestResourceBlockBySlot(t *testing.T) {
	t.Parallel()
	slot := 412162133
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(
			[]byte(
				`{
					"time": 1641338934,
					"height": 15243593,
					"hash": "4ea1ba291e8eef538635a53e59fddba7810d1679631cc3aed7c8e6c4091a516a",
					"slot": 412162133,
					"epoch": 425,
					"epoch_slot": 12,
					"slot_leader": "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2qnikdy",
					"size": 3,
					"tx_count": 1,
					"output": "128314491794",
					"fees": "592661",
					"block_vrf": "vrf_vk1wf2k6lhujezqcfe00l6zetxpnmh9n6mwhpmhm0dvfh3fxgmdnrfqkms8ty",
					"previous_block": "43ebccb3ac72c7cebd0d9b755a4b08412c9f5dcb81b8a0ad1e3c197d29d47b05",
					"next_block": "8367f026cf4b03e116ff8ee5daf149b55ba5a6ec6dec04803b8dc317721d15fa",
					"confirmations": 4698
				  }`,
			))
	}))
	defer s.Close()

	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatal(err)
	}
	block, err := api.BlockBySlot(context.TODO(), slot)
	if err != nil {
		t.Fatal(err)
	}

	if block.Slot != slot {
		t.Fatalf("Expected block.Slot to be %d, got %d", slot, block.Slot)
	}
}

func TestResourceBlockBySlotAndEpoch(t *testing.T) {
	t.Parallel()
	slot, epoch := 412162133, 425
	s := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(
			[]byte(
				`{
					"time": 1641338934,
					"height": 15243593,
					"hash": "4ea1ba291e8eef538635a53e59fddba7810d1679631cc3aed7c8e6c4091a516a",
					"slot": 412162133,
					"epoch": 425,
					"epoch_slot": 12,
					"slot_leader": "pool1pu5jlj4q9w9jlxeu370a3c9myx47md5j5m2str0naunn2qnikdy",
					"size": 3,
					"tx_count": 1,
					"output": "128314491794",
					"fees": "592661",
					"block_vrf": "vrf_vk1wf2k6lhujezqcfe00l6zetxpnmh9n6mwhpmhm0dvfh3fxgmdnrfqkms8ty",
					"previous_block": "43ebccb3ac72c7cebd0d9b755a4b08412c9f5dcb81b8a0ad1e3c197d29d47b05",
					"next_block": "8367f026cf4b03e116ff8ee5daf149b55ba5a6ec6dec04803b8dc317721d15fa",
					"confirmations": 4698
				  }`,
			))
	}))
	defer s.Close()

	api, err := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{Server: s.URL},
	)
	if err != nil {
		t.Fatal(err)
	}
	block, err := api.BlocksBySlotAndEpoch(context.TODO(), slot, epoch)
	if err != nil {
		t.Fatal(err)
	}

	if block.Slot != slot {
		t.Fatalf("Expected block.Slot to be %d, got %d", slot, block.Slot)
	}
	if block.Epoch != epoch {
		t.Fatalf("Expected block.Epoch to be %d, got %d", epoch, block.Epoch)
	}
}
