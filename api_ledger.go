package blockfrost

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

var (
	resourceGenesis = "genesis"
)

type GenesisBlock struct {
	ActiveSlotsCoefficient float32 `json:"active_slots_coefficient,omitempty"`
	UpdateQuorum           float32 `json:"update_quorum,omitempty"`
	MaxLovelaceSupply      string  `json:"max_lovelace_supply,omitempty"`
	NetworkMagic           int     `json:"network_magic,omitempty"`
	EpochLength            int     `json:"epoch_length,omitempty"`
	SystemStart            int     `json:"system_start,omitempty"`
	SlotsPerKesPeriod      int     `json:"slots_per_kes_period,omitempty"`
	SlotLength             int     `json:"slot_length,omitempty"`
	MaxKesEvolutions       int     `json:"max_kes_evolutions,omitempty"`
	SecurityParam          int     `json:"security_param,omitempty"`
}

func (c *apiClient) Genesis(ctx context.Context) (GenesisBlock, error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceGenesis))
	if err != nil {
		return GenesisBlock{}, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return GenesisBlock{}, err
	}
	req.Header.Add("project_id", c.projectId)

	res, err := c.client.Do(req)
	if err != nil {
		return GenesisBlock{}, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return GenesisBlock{}, handleAPIErrorResponse(res)
	}

	gen := GenesisBlock{}
	if err := json.NewDecoder(res.Body).Decode(&gen); err != nil {
		return gen, err
	}
	return gen, nil
}
