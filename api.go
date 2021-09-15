package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *apiClient) Health(ctx context.Context) (Health, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceHealth))
	if err != nil {
		return Health{}, err
	}
	req, err := http.NewRequest(http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return Health{}, err
	}
	req.Header.Add("project_id", c.projectId)
	req.WithContext(ctx)

	res, err := c.client.Do(req)
	if err != nil {
		return Health{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return Health{}, handleAPIErrorResponse(res)
	}

	health := Health{}
	json.NewDecoder(res.Body).Decode(&health)
	return health, nil
}

func (c *apiClient) Info(ctx context.Context) (Info, error) {
	req, err := http.NewRequest(http.MethodGet, c.server, nil)
	if err != nil {
		return Info{}, err
	}
	req.Header.Add("project_id", c.projectId)
	req.WithContext(ctx)

	res, err := c.client.Do(req)

	if err != nil {
		return Info{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return Info{}, handleAPIErrorResponse(res)
	}

	info := Info{}
	json.NewDecoder(res.Body).Decode(&info)
	return info, nil

}

func handleAPIErrorResponse(res *http.Response) error {
	var err error
	switch res.StatusCode {
	case 400:
		br := BadRequest{}
		if err = json.NewDecoder(res.Body).Decode(&br); err != nil {
			return err
		}
		return &APIError{
			Response: br,
		}
	case 403:
		ua := UnauthorizedError{}
		if err = json.NewDecoder(res.Body).Decode(&ua); err != nil {
			return err
		}
		return &APIError{
			Response: ua,
		}
	case 404:
		nf := NotFound{}
		if err = json.NewDecoder(res.Body).Decode(&nf); err != nil {
			return err
		}
		return &APIError{
			Response: nf,
		}
	case 429:
		ol := OverusageLimit{}
		if err = json.NewDecoder(res.Body).Decode(&ol); err != nil {
			return err
		}
		return &APIError{
			Response: ol,
		}
	case 418:
		ab := AutoBanned{}
		if err = json.NewDecoder(res.Body).Decode(&ab); err != nil {
			return err
		}
		return &APIError{
			Response: ab,
		}
	case 500:
		ise := InternalServerError{}
		if err = json.NewDecoder(res.Body).Decode(&ise); err != nil {
			return err
		}
		return &APIError{
			Response: ise,
		}
	default:
		return &APIError{}
	}
}
