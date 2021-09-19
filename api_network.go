package blockfrost

type NetworkSupply struct {
	Max         string `json:"max,omitempty"`
	Total       string `json:"total,omitempty"`
	Circulating string `json:"circulating,omitempty"`
	Locked      string `json:"locked,omitempty"`
}

type NetworkStake struct {
	Live   string `json:"live,omitempty"`
	Active string `json:"active,omitempty"`
}
type NetworkInfo struct {
	Supply NetworkSupply `json:"supply,omitempty"`
	Stake  NetworkStake  `json:"stake,omitempty"`
}
