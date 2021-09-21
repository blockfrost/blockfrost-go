package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	resourceMetadataTxLabels        = "metadata/txs/labels"
	resourceMetadataTxContentInJSON = "metadata/txs/labels" // and {label_parameter}
	resourceMetadataTxContentInCBOR = "metadata/txs/labels" // and {label_parameter}/cbor
)

// MetadataTxLabel return Transaction metadata labels
// List of all used transaction metadata labels.
type MetadataTxLabel struct {
	Label string `json:"label,omitempty"`
	Cip10 string `json:"cip10,omitempty"`
	Count string `json:"count,omitempty"`
}

// MetadataTxContentInJSON Transaction metadata content raw in JSON
// Transaction metadata per label.
// This struct are more flexible on JSONMetadata field
type MetadataTxContentInJSON struct {
	TxHash string `json:"tx_hash,omitempty"`
	// 	string or object or Array of any or integer or number or boolean Nullable
	// Content of the JSON metadata
	JSONMetadata interface{} `json:"json_metadata,omitempty"`
}

// MetadataTxContentInCBOR return Transaction metadata content in CBOR
// Transaction metadata per label.
type MetadataTxContentInCBOR struct {
	TxHash       string `json:"tx_hash,omitempty"`
	CborMetadata string `json:"cbor_metadata,omitempty"`
}

// MetadataTxLabels returns the List of all used transaction metadata labels.
func (c *apiClient) MetadataTxLabels(
	ctx context.Context,
	query APIPagingParams,
) ([]MetadataTxLabel, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/", c.server, resourceMetadataTxLabels))
	if err != nil {
		return []MetadataTxLabel{}, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []MetadataTxLabel{}, err
	}

	v := req.URL.Query()
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()
	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []MetadataTxLabel{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []MetadataTxLabel{}, handleAPIErrorResponse(res)
	}
	metadataTxs := []MetadataTxLabel{}
	err = json.NewDecoder(res.Body).Decode(&metadataTxs)
	if err != nil {
		return []MetadataTxLabel{}, err
	}
	return metadataTxs, nil
}

// MetadataTxContentInJSON returns the Transaction metadata content in JSON
// Transaction metadata per label.
func (c *apiClient) MetadataTxContentInJSON(
	ctx context.Context,
	label string,
	query APIPagingParams,
) ([]MetadataTxContentInJSON, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s", c.server, resourceMetadataTxContentInJSON, label),
	)
	if err != nil {
		return []MetadataTxContentInJSON{}, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []MetadataTxContentInJSON{}, err
	}

	v := req.URL.Query()
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()
	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []MetadataTxContentInJSON{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []MetadataTxContentInJSON{}, handleAPIErrorResponse(res)
	}
	metadataTxs := []MetadataTxContentInJSON{}
	err = json.NewDecoder(res.Body).Decode(&metadataTxs)
	if err != nil {
		return []MetadataTxContentInJSON{}, err
	}
	return metadataTxs, nil
}

// MetadataTxContentInCBOR returns the Transaction metadata content in CBOR
// Transaction metadata per label.
func (c *apiClient) MetadataTxContentInCBOR(
	ctx context.Context,
	label string,
	query APIPagingParams,
) ([]MetadataTxContentInCBOR, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.server, resourceMetadataTxContentInCBOR, label, "cbor"),
	)
	if err != nil {
		return []MetadataTxContentInCBOR{}, err
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []MetadataTxContentInCBOR{}, err
	}

	v := req.URL.Query()
	v = formatParams(v, query)
	req.URL.RawQuery = v.Encode()
	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []MetadataTxContentInCBOR{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []MetadataTxContentInCBOR{}, handleAPIErrorResponse(res)
	}
	metadataTxs := []MetadataTxContentInCBOR{}
	err = json.NewDecoder(res.Body).Decode(&metadataTxs)
	if err != nil {
		return []MetadataTxContentInCBOR{}, err
	}
	return metadataTxs, nil
}
