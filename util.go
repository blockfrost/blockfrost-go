package blockfrost

import (
	"encoding/json"
	"net/http"
)

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
