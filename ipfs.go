package blockfrost

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"

	"github.com/hashicorp/go-retryablehttp"
)

const (
	resourceIPFSAdd       = "ipfs/add"
	resourceIPFSPin       = "ipfs/pin/add"
	resourceIPFSPinList   = "ipfs/pin/list"
	resourceIPFSPinRemove = "ipfs/pin/remove"
	resourceIPFSGateway   = "ipfs/gateway"
)

type ipfsClient struct {
	server    string
	projectId string
	client    *retryablehttp.Client
	routines  int
}

type IPFSClientOptions struct {
	// The project_id to use from blockfrost. If not set
	// `BLOCKFROST_IPFS_PROJECT_ID` is loaded from env
	ProjectID string
	// Configures server to use. Can be toggled for test servers
	Server string
	// Interface implementing Do method such *http.Client
	Client *retryablehttp.Client
	// Max goroutines to use for *All Methods
	MaxRoutines int
}

// IPFSObject contains information on an IPFS object
type IPFSObject struct {
	Name     string `json:"name"`
	IPFSHash string `json:"ipfs_hash"`
	Size     string `json:"size"`
}

// IPFSPinnedObject contains information on a pinned object
type IPFSPinnedObject struct {
	TimeCreated int    `json:"time_created"`
	TimePinned  int    `json:"time_pinned"`
	IPFSHash    string `json:"ipfs_hash"`
	State       string `json:"state"`
	Size        string `json:"size"`
}

// NewIPFSClient creates and returns an IPFS client configured using
// IPFSClientOptions. It will initialize the client with default options
// if provided with empty options
func NewIPFSClient(options IPFSClientOptions) IPFSClient {
	if options.Server == "" {
		options.Server = IPFSNet
	}
	retryclient := retryablehttp.NewClient()
	retryclient.Logger = nil

	if options.ProjectID == "" {
		options.ProjectID = os.Getenv("BLOCKFROST_IPFS_PROJECT_ID")
	}

	if options.MaxRoutines == 0 {
		options.MaxRoutines = 10
	}

	client := &ipfsClient{
		server:    options.Server,
		client:    retryclient,
		projectId: options.ProjectID,
		routines:  options.MaxRoutines,
	}
	return client
}

type IPFSClient interface {
	Add(ctx context.Context, filePath string) (IPFSObject, error)
	Pin(ctx context.Context, path string) (IPFSPinnedObject, error)
	PinnedObject(ctx context.Context, path string) (IPFSPinnedObject, error)
	PinnedObjects(ctx context.Context, query APIQueryParams) ([]IPFSPinnedObject, error)
	Remove(ctx context.Context, path string) (IPFSObject, error)
	Gateway(ctx context.Context, path string) ([]byte, error)
	PinnedObjectsAll(ctx context.Context) <-chan PinnedObjectResult
}

// PinnedObjectResult contains response and error from an All method
type PinnedObjectResult struct {
	Res []IPFSPinnedObject
	Err error
}

// Add a file to IPFS storage
// You need to Pin an object to avoid it being garbage collected.
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

	req.Header.Add("Content-Type", wr.FormDataContentType())

	res, err := ip.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()
	if err = json.NewDecoder(res.Body).Decode(&ipo); err != nil {
		return
	}

	return ipo, nil
}

// PinnedObjectsAll gets all pinned objects. Returns a channel that can be used with range
func (ip *ipfsClient) PinnedObjectsAll(ctx context.Context) <-chan PinnedObjectResult {
	ch := make(chan PinnedObjectResult, ip.routines)
	jobs := make(chan methodOptions, ip.routines)
	quit := make(chan bool, ip.routines)

	wg := sync.WaitGroup{}

	for i := 0; i < ip.routines; i++ {
		wg.Add(1)
		go func(jobs chan methodOptions, ch chan PinnedObjectResult, wg *sync.WaitGroup) {
			defer wg.Done()
			for j := range jobs {
				objs, err := ip.PinnedObjects(j.ctx, j.query)
				if len(objs) != j.query.Count || err != nil {
					quit <- true
				}
				res := PinnedObjectResult{Res: objs, Err: err}
				ch <- res
			}

		}(jobs, ch, &wg)
	}
	go func() {
		defer close(ch)
		fetchScripts := true
		for i := 1; fetchScripts; i++ {
			select {
			case <-quit:
				fetchScripts = false
				return
			default:
				jobs <- methodOptions{ctx: ctx, query: APIQueryParams{Count: 100, Page: i}}
			}
		}

		wg.Wait()
	}()
	return ch
}

// Pin an object to avoid it being garbage collected
func (ip *ipfsClient) Pin(ctx context.Context, IPFSPath string) (ipo IPFSPinnedObject, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", ip.server, resourceIPFSPin, IPFSPath))
	if err != nil {
		return
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestUrl.String(), nil)
	if err != nil {
		return
	}
	res, err := ip.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&ipo); err != nil {
		return
	}
	return ipo, nil
}

// PinnedObject returns information about locally pinned IPFS object
func (ip *ipfsClient) PinnedObject(ctx context.Context, IPFSPath string) (ipo IPFSPinnedObject, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", ip.server, resourceIPFSPinList, IPFSPath))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}
	res, err := ip.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&ipo); err != nil {
		return
	}
	return ipo, nil
}

// PinnedObjects returns information about locally pinned IPFS objects. Returns
// a slice of IPFSPinnedObject(s) whose quantity and offset is controlled by
// query parameters
func (ip *ipfsClient) PinnedObjects(ctx context.Context, query APIQueryParams) (ipos []IPFSPinnedObject, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", ip.server, resourceIPFSPinList))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}
	v := req.URL.Query()
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()
	res, err := ip.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&ipos); err != nil {
		return
	}
	return ipos, nil
}

// Remove - removes pinned objects from local storage. Returns and IPFSObject
// containing removed object information
func (ip *ipfsClient) Remove(ctx context.Context, IPFSPath string) (ipo IPFSObject, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", ip.server, resourceIPFSPinRemove, IPFSPath))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, requestUrl.String(), nil)
	if err != nil {
		return
	}
	res, err := ip.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&ipo); err != nil {
		return
	}
	return ipo, nil
}

// Gateway retrieves an object from the IFPS gateway and returns a byte
// (useful if you do not want to rely on a public gateway, such as `ipfs.blockfrost.dev`).
func (ip *ipfsClient) Gateway(ctx context.Context, IPFSPath string) (ipo []byte, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/%s", ip.server, resourceIPFSGateway, IPFSPath))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}
	res, err := ip.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	byteObj, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return
	}

	return byteObj, nil
}
