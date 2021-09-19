package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// MetadataTxLabels returns the List of all used transaction metadata labels.
func (c *apiClient) MetadataTxLabels(ctx context.Context, query APIPagingParams) ([]MetadataTxLabel, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s/", c.server, resourceMetadataTxLabels))
	if err != nil {
		return []MetadataTxLabel{}, err
	}

	v := url.Values{}
	if query.Count > 0 {
		v.Set("count", fmt.Sprintf("%d", query.Count))
		requestUrl.RawQuery = v.Encode()
	}
	if query.Page > 0 {
		v.Set("page", fmt.Sprintf("%d", query.Page))
		requestUrl.RawQuery = v.Encode()
	}
	if query.Order != "" {
		v.Set("order", query.Order)
		requestUrl.RawQuery = v.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []MetadataTxLabel{}, err
	}

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

	v := url.Values{}
	if query.Count > 0 {
		v.Set("count", fmt.Sprintf("%d", query.Count))
		requestUrl.RawQuery = v.Encode()
	}
	if query.Page > 0 {
		v.Set("page", fmt.Sprintf("%d", query.Page))
		requestUrl.RawQuery = v.Encode()
	}
	if query.Order != "" {
		v.Set("order", query.Order)
		requestUrl.RawQuery = v.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []MetadataTxContentInJSON{}, err
	}

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

// MetadataTxContentInJSON returns the Transaction metadata content in JSON
// Transaction metadata per label.
func (c *apiClient) MetadataTxContentInJSONRaw(
	ctx context.Context,
	label string,
	query APIPagingParams,
) ([]MetadataTxContentInJSONRaw, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s", c.server, resourceMetadataTxContentInJSON, label),
	)
	if err != nil {
		return []MetadataTxContentInJSONRaw{}, err
	}

	v := url.Values{}
	if query.Count > 0 {
		v.Set("count", fmt.Sprintf("%d", query.Count))
		requestUrl.RawQuery = v.Encode()
	}
	if query.Page > 0 {
		v.Set("page", fmt.Sprintf("%d", query.Page))
		requestUrl.RawQuery = v.Encode()
	}
	if query.Order != "" {
		v.Set("order", query.Order)
		requestUrl.RawQuery = v.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []MetadataTxContentInJSONRaw{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []MetadataTxContentInJSONRaw{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []MetadataTxContentInJSONRaw{}, handleAPIErrorResponse(res)
	}
	metadataTxs := []MetadataTxContentInJSONRaw{}
	err = json.NewDecoder(res.Body).Decode(&metadataTxs)
	if err != nil {
		return []MetadataTxContentInJSONRaw{}, err
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

	v := url.Values{}
	if query.Count > 0 {
		v.Set("count", fmt.Sprintf("%d", query.Count))
		requestUrl.RawQuery = v.Encode()
	}
	if query.Page > 0 {
		v.Set("page", fmt.Sprintf("%d", query.Page))
		requestUrl.RawQuery = v.Encode()
	}
	if query.Order != "" {
		v.Set("order", query.Order)
		requestUrl.RawQuery = v.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return []MetadataTxContentInCBOR{}, err
	}

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
