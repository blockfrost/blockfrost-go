package blockfrost

type Script struct {
	ScriptHash     string `json:"script_hash,omitempty"`
	Type           string `json:"type,omitempty"`
	SerialisedSize int    `json:"serialised_size,omitempty"`
}

type ScriptRedeemer struct {
	TxHash    string `json:"tx_hash,omitempty"`
	TxIndex   int    `json:"tx_index,omitempty"`
	Purpose   string `json:"purpose,omitempty"`
	UnitMem   string `json:"unit_mem,omitempty"`
	UnitSteps string `json:"unit_steps,omitempty"`
	Fee       string `json:"fee,omitempty"`
}
