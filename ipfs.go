package blockfrost

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
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

type IPFSPinnedObject struct {
	TimeCreated int    `json:"time_created,omitempty"`
	TimePinned  int    `json:"time_pinned,omitempty"`
	IPFSHash    string `json:"ipfs_hash,omitempty"`
	State       string `json:"state,omitempty"`
	Size        string `json:"size,omitempty"`
}

func NewIPFSClient(options IPFSClientOptions) (IPFSClient, error) {
	if options.Server == "" {
		options.Server = IPFSNet
	}

	if options.Client == nil {
		options.Client = &http.Client{}
	}

	if options.ProjectID == "" {
		options.ProjectID = os.Getenv("BLOCKFROST_IPFS_PROJECT_ID")
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
	Pin(ctx context.Context, path string) (IPFSPinnedObject, error)
	PinnedObject(ctx context.Context, path string) (IPFSPinnedObject, error)
	PinnedObjects(ctx context.Context, query APIPagingParams) ([]IPFSPinnedObject, error)
	Remove(ctx context.Context, path string) ([]IPFSObject, error)
	Gateway(ctx context.Context, path string) ([]IPFSObject, error)
}

func (ip *ipfsClient) Add(ctx context.Context, filePath string) (ipo IPFSObject, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", ip.server, resourceIPFSAdd))
	if err != nil {
		return
	}

	file, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer file.Close()
	body := &bytes.Buffer{}
	wr := multipart.NewWriter(body)
	part, err := wr.CreateFormFile("file", filepath.Base(filePath))
	if err != nil {
		return
	}
	if _, err = io.Copy(part, file); err != nil {
		return
	}
	if err = wr.Close(); err != nil {
		return
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestUrl.String(), body)
	if err != nil {
		return
	}

	req.Header.Add("project_id", ip.projectId)
	req.Header.Add("Content-Type", wr.FormDataContentType())

	res, err := ip.client.Do(req)
	if err != nil {
		return
	}
	ipo = IPFSObject{}
	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return ipo, handleAPIErrorResponse(res)
	}
	if err = json.NewDecoder(res.Body).Decode(&ipo); err != nil {
		return
	}

	return ipo, nil
}

func (ip *ipfsClient) Pin(ctx context.Context, path string) (ipo IPFSPinnedObject, err error) {
	return
}

func (ip *ipfsClient) PinnedObject(ctx context.Context, path string) (ipo IPFSPinnedObject, err error) {
	return
}

func (ip *ipfsClient) PinnedObjects(ctx context.Context, query APIPagingParams) (ipos []IPFSPinnedObject, err error) {
	return
}

func (ip *ipfsClient) Remove(ctx context.Context, path string) (ipo []IPFSObject, err error) {
	return
}

func (ip *ipfsClient) Gateway(ctx context.Context, path string) (ipo []IPFSObject, err error) {
	return
}
