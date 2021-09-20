package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// Pools returns the List of stake pools
// List of registered stake pools.
func (c *apiClient) Pools(ctx context.Context, query APIPagingParams) (Pools, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourcePool))
	if err != nil {
		return Pools{}, err
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
		return Pools{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return Pools{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return Pools{}, handleAPIErrorResponse(res)
	}
	pools := Pools{}
	err = json.NewDecoder(res.Body).Decode(&pools)
	if err != nil {
		return Pools{}, err
	}
	return pools, nil
}

// PoolsRetired returns the List of retired stake pools
// List of already retired pools.
func (c *apiClient) PoolsRetired(ctx context.Context, query APIPagingParams) ([]PoolRetired, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourcePoolRetired))
	if err != nil {
		return []PoolRetired{}, err
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
		return []PoolRetired{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []PoolRetired{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []PoolRetired{}, handleAPIErrorResponse(res)
	}
	pools := []PoolRetired{}
	err = json.NewDecoder(res.Body).Decode(&pools)
	if err != nil {
		return []PoolRetired{}, err
	}
	return pools, nil
}

// PoolsRetiring returns the List of retiring stake pools
// List of stake pools retiring in the upcoming epochs
func (c *apiClient) PoolsRetiring(
	ctx context.Context,
	query APIPagingParams,
) ([]PoolRetiring, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourcePoolRetiring))
	if err != nil {
		return []PoolRetiring{}, err
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
		return []PoolRetiring{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []PoolRetiring{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []PoolRetiring{}, handleAPIErrorResponse(res)
	}
	pools := []PoolRetiring{}
	err = json.NewDecoder(res.Body).Decode(&pools)
	if err != nil {
		return []PoolRetiring{}, err
	}
	return pools, nil
}

// PoolsRetiring returns the Specific stake pool
// Pool information.
func (c *apiClient) PoolSpecific(
	ctx context.Context,
	poolID string,
	query APIPagingParams,
) (PoolSpecific, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s", c.server, resourcePool, poolID),
	)
	if err != nil {
		return PoolSpecific{}, err
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
		return PoolSpecific{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return PoolSpecific{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return PoolSpecific{}, handleAPIErrorResponse(res)
	}
	pool := PoolSpecific{}
	err = json.NewDecoder(res.Body).Decode(&pool)
	if err != nil {
		return PoolSpecific{}, err
	}
	return pool, nil
}

// PoolHistory returns the Stake pool history
// History of stake pool parameters over epochs.
func (c *apiClient) PoolHistory(
	ctx context.Context,
	poolID string,
	query APIPagingParams,
) ([]PoolHistory, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.server, resourcePool, poolID, resourcePoolHistory),
	)
	if err != nil {
		return []PoolHistory{}, err
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
		return []PoolHistory{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []PoolHistory{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []PoolHistory{}, handleAPIErrorResponse(res)
	}
	pool := []PoolHistory{}
	err = json.NewDecoder(res.Body).Decode(&pool)
	if err != nil {
		return []PoolHistory{}, err
	}
	return pool, nil
}

// PoolMetadata returns the Stake pool metadata
// Stake pool registration metadata.
func (c *apiClient) PoolMetadata(
	ctx context.Context,
	poolID string,
	query APIPagingParams,
) (PoolMetadata, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.server, resourcePool, poolID, resourcePoolMetadata),
	)
	if err != nil {
		return PoolMetadata{}, err
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
		return PoolMetadata{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return PoolMetadata{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return PoolMetadata{}, handleAPIErrorResponse(res)
	}
	pool := PoolMetadata{}
	err = json.NewDecoder(res.Body).Decode(&pool)
	if err != nil {
		return PoolMetadata{}, err
	}
	return pool, nil
}

// PoolRelay returns the Stake pool relays
// Relays of a stake pool.
func (c *apiClient) PoolRelay(
	ctx context.Context,
	poolID string,
	query APIPagingParams,
) ([]PoolRelay, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.server, resourcePool, poolID, resourcePoolRelay),
	)
	if err != nil {
		return []PoolRelay{}, err
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
		return []PoolRelay{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []PoolRelay{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []PoolRelay{}, handleAPIErrorResponse(res)
	}
	pool := []PoolRelay{}
	err = json.NewDecoder(res.Body).Decode(&pool)
	if err != nil {
		return []PoolRelay{}, err
	}
	return pool, nil
}

// PoolDelegator returns the Stake pool delegators
// List of current stake pools delegators.
func (c *apiClient) PoolDelegator(
	ctx context.Context,
	poolID string,
	query APIPagingParams,
) ([]PoolDelegator, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.server, resourcePool, poolID, resourcePoolDelegator),
	)
	if err != nil {
		return []PoolDelegator{}, err
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
		return []PoolDelegator{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []PoolDelegator{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []PoolDelegator{}, handleAPIErrorResponse(res)
	}
	pool := []PoolDelegator{}
	err = json.NewDecoder(res.Body).Decode(&pool)
	if err != nil {
		return []PoolDelegator{}, err
	}
	return pool, nil
}

// PoolBlocks returns the Stake pool blocks
// List of stake pools blocks.
func (c *apiClient) PoolBlocks(
	ctx context.Context,
	poolID string,
	query APIPagingParams,
) (PoolBlocks, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.server, resourcePool, poolID, resourcePoolBlocks),
	)
	if err != nil {
		return PoolBlocks{}, err
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
		return PoolBlocks{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return PoolBlocks{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return PoolBlocks{}, handleAPIErrorResponse(res)
	}
	pool := PoolBlocks{}
	err = json.NewDecoder(res.Body).Decode(&pool)
	if err != nil {
		return PoolBlocks{}, err
	}
	return pool, nil
}

// PoolUpdate returns the Stake pool updates
// List of certificate updates to the stake pool.
func (c *apiClient) PoolUpdate(
	ctx context.Context,
	poolID string,
	query APIPagingParams,
) ([]PoolUpdate, error) {
	requestUrl, err := url.Parse(
		fmt.Sprintf("%s/%s/%s/%s", c.server, resourcePool, poolID, resourcePoolUpdate),
	)
	if err != nil {
		return []PoolUpdate{}, err
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
		return []PoolUpdate{}, err
	}

	req.Header.Add("project_id", c.projectId)
	req = req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return []PoolUpdate{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return []PoolUpdate{}, handleAPIErrorResponse(res)
	}
	pool := []PoolUpdate{}
	err = json.NewDecoder(res.Body).Decode(&pool)
	if err != nil {
		return []PoolUpdate{}, err
	}
	return pool, nil
}
