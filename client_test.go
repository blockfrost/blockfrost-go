package blockfrost_test

import (
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestClientInit(t *testing.T) {
	// Include WithProjectID(project_id) to add project ID. Ignore to use env BLOCKFROST_PROJECT_ID
	_ = blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)
}
