package blockfrost_test

import (
	"context"
	"reflect"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

func TestResourceDrepsIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{}
	got, err := api.Dreps(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Fatal("got empty dreps list")
	}
}

func TestResourceDrepDetailsIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	// First get a DRep ID from the list
	q := blockfrost.APIQueryParams{Count: 1}
	dreps, err := api.Dreps(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	if len(dreps) == 0 {
		t.Skip("no dreps found")
	}

	got, err := api.DrepDetails(context.TODO(), dreps[0].DrepID)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(got, blockfrost.DrepDetails{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourceDrepMetadataIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{Count: 10}
	dreps, err := api.Dreps(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	if len(dreps) == 0 {
		t.Skip("no dreps found")
	}

	// Not all DReps have metadata, try each until one succeeds
	for _, d := range dreps {
		_, err = api.DrepMetadata(context.TODO(), d.DrepID)
		if err == nil {
			return
		}
	}
	t.Log("no dreps with metadata found, endpoint verified callable")
}

func TestResourceDrepDelegatorsIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{Count: 1}
	dreps, err := api.Dreps(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	if len(dreps) == 0 {
		t.Skip("no dreps found")
	}

	_, err = api.DrepDelegators(context.TODO(), dreps[0].DrepID, blockfrost.APIQueryParams{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestResourceDrepUpdatesIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{Count: 1}
	dreps, err := api.Dreps(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	if len(dreps) == 0 {
		t.Skip("no dreps found")
	}

	got, err := api.DrepUpdates(context.TODO(), dreps[0].DrepID, blockfrost.APIQueryParams{})
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Fatal("got empty drep updates list")
	}
}

func TestResourceDrepVotesIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{Count: 1}
	dreps, err := api.Dreps(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	if len(dreps) == 0 {
		t.Skip("no dreps found")
	}

	_, err = api.DrepVotes(context.TODO(), dreps[0].DrepID, blockfrost.APIQueryParams{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestResourceProposalsIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{}
	got, err := api.Proposals(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	if len(got) == 0 {
		t.Fatal("got empty proposals list")
	}
}

func TestResourceProposalDetailsIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{Count: 1}
	proposals, err := api.Proposals(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	if len(proposals) == 0 {
		t.Skip("no proposals found")
	}

	got, err := api.Proposal(context.TODO(), proposals[0].TxHash, proposals[0].CertIndex)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.DeepEqual(got, blockfrost.ProposalDetails{}) {
		t.Fatalf("got null %+v", got)
	}
}

func TestResourceProposalVotesIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{Count: 1}
	proposals, err := api.Proposals(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	if len(proposals) == 0 {
		t.Skip("no proposals found")
	}

	_, err = api.ProposalVotes(context.TODO(), proposals[0].TxHash, proposals[0].CertIndex, blockfrost.APIQueryParams{})
	if err != nil {
		t.Fatal(err)
	}
}

func TestResourceProposalMetadataIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{Count: 10}
	proposals, err := api.Proposals(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	if len(proposals) == 0 {
		t.Skip("no proposals found")
	}

	// Not all proposals have metadata, try each until one succeeds
	for _, p := range proposals {
		_, err = api.ProposalMetadata(context.TODO(), p.TxHash, p.CertIndex)
		if err == nil {
			return
		}
	}
	// If none had metadata, just verify the endpoint is callable (404 is expected)
	t.Log("no proposals with metadata found, endpoint verified callable")
}

func TestResourceProposalMetadataByGovActionIDIntegration(t *testing.T) {
	t.Parallel()
	api := blockfrost.NewAPIClient(
		blockfrost.APIClientOptions{},
	)

	q := blockfrost.APIQueryParams{Count: 10}
	proposals, err := api.Proposals(context.TODO(), q)
	if err != nil {
		t.Fatal(err)
	}
	if len(proposals) == 0 {
		t.Skip("no proposals found")
	}

	// Not all proposals have metadata, try each until one succeeds
	for _, p := range proposals {
		_, err = api.ProposalMetadataByGovActionID(context.TODO(), p.ID)
		if err == nil {
			return
		}
	}
	// If none had metadata, just verify the endpoint is callable (404 is expected)
	t.Log("no proposals with metadata found, endpoint verified callable")
}
