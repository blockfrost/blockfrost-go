package blockfrost

import (
	"context"
	"net/http"
	"time"
)

const (
	resourceIPFSAdd = "ipfs/add"
)

type ipfsClient struct {
	server    string
	projectId string
	client    HttpRequestDoer
}

type IPFSClientOptions struct {
	// The project_id to use from blockfrost. If not set
	// `BLOCKFROST_IPFS_PROJECT_ID` is loaded from env
	ProjectID string
	// Configures server to use. Can be toggled for test servers
	Server string
	// Interface implementing Do method such *http.Client
	Client HttpRequestDoer
}

type IPFSObject struct {
	Name     string `json:"string,omitempty"`
	IPFSHash string `json:"ipfs_hash,omitempty"`
	Size     string `json:"size,omitempty"`
}

func NewIPFSClient(options IPFSClientOptions) (IPFSClient, error) {
	if options.Server == "" {
		options.Server = IPFSNet
	}

	if options.Client == nil {
		options.Client = &http.Client{Timeout: time.Second * 5}
	}

	client := &ipfsClient{
		server:    options.Server,
		client:    options.Client,
		projectId: options.ProjectID,
	}
	return client, nil
}

type IPFSClient interface {
	Add(ctx context.Context, filePath string) (IPFSObject, error)
}

func (ip *ipfsClient) Add(ctx context.Context, filePath string) (ipfso IPFSObject, err error) {
	return
}
