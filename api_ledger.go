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

// GenesisBlock contains the information of the genesis block of the network.
type GenesisBlock struct {
	ActiveSlotsCoefficient float32 `json:"active_slots_coefficient"`
	UpdateQuorum           float32 `json:"update_quorum"`
	MaxLovelaceSupply      string  `json:"max_lovelace_supply"`
	NetworkMagic           int     `json:"network_magic"`
	EpochLength            int     `json:"epoch_length"`
	SystemStart            int     `json:"system_start"`
	SlotsPerKesPeriod      int     `json:"slots_per_kes_period"`
	SlotLength             int     `json:"slot_length"`
	MaxKesEvolutions       int     `json:"max_kes_evolutions"`
	SecurityParam          int     `json:"security_param"`
}

// Genesis returns the information about blockchain genesis.
func (c *apiClient) Genesis(ctx context.Context) (gen GenesisBlock, err error) {
	requestUrl, err := url.Parse(fmt.Sprintf("%s/%s", c.server, resourceGenesis))
	if err != nil {
		return
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, requestUrl.String(), nil)
	if err != nil {
		return
	}

	res, err := c.handleRequest(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err = json.NewDecoder(res.Body).Decode(&gen); err != nil {
		return
	}
	return gen, nil
}
