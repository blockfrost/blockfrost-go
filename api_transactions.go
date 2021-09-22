package blockfrost

type TransactionContent struct {
	// Count of asset mints and burns within the transaction
	AssetMintOrBurnCount int `json:"asset_mint_or_burn_count"`

	// Block hash
	Block string `json:"block"`

	// Block number
	BlockHeight int `json:"block_height"`

	// Count of the delegations within the transaction
	DelegationCount int `json:"delegation_count"`

	// Deposit within the transaction in Lovelaces
	Deposit string `json:"deposit"`

	// Fees of the transaction in Lovelaces
	Fees string `json:"fees"`

	// Transaction hash
	Hash string `json:"hash"`

	// Transaction index within the block
	Index int `json:"index"`

	// Left (included) endpoint of the timelock validity intervals
	InvalidBefore *string `json:"invalid_before"`

	// Right (excluded) endpoint of the timelock validity intervals
	InvalidHereafter *string `json:"invalid_hereafter"`

	// Count of the MIR certificates within the transaction
	MirCertCount int `json:"mir_cert_count"`
	OutputAmount []struct {
		// The quantity of the unit
		Quantity string `json:"quantity"`

		// The unit of the value
		Unit string `json:"unit"`
	} `json:"output_amount"`

	// Count of the stake pool retirement certificates within the transaction
	PoolRetireCount int `json:"pool_retire_count"`

	// Count of the stake pool registration and update certificates within the transaction
	PoolUpdateCount int `json:"pool_update_count"`

	// Size of the transaction in Bytes
	Size int `json:"size"`

	// Slot number
	Slot int `json:"slot"`

	// Count of the stake keys (de)registration and delegation certificates within the transaction
	StakeCertCount int `json:"stake_cert_count"`

	// Count of UTXOs within the transaction
	UtxoCount int `json:"utxo_count"`

	// Count of the withdrawals within the transaction
	WithdrawalCount int `json:"withdrawal_count"`
}

type TransactionUTXOs struct {
	// Transaction hash
	Hash   string `json:"hash"`
	Inputs []struct {
		// Input address
		Address string `json:"address"`
		Amount  []struct {
			// The quantity of the unit
			Quantity string `json:"quantity"`

			// The unit of the value
			Unit string `json:"unit"`
		} `json:"amount"`

		// UTXO index in the transaction
		OutputIndex float32 `json:"output_index"`

		// Hash of the UTXO transaction
		TxHash string `json:"tx_hash"`
	} `json:"inputs"`
	Outputs []struct {
		// Output address
		Address string `json:"address"`
		Amount  []struct {
			// The quantity of the unit
			Quantity string `json:"quantity"`

			// The unit of the value
			Unit string `json:"unit"`
		} `json:"amount"`
	} `json:"outputs"`
}

type TransactionStakeAddressCert struct {
	// Delegation stake address
	Address string `json:"address"`

	// Index of the certificate within the transaction
	CertIndex int `json:"cert_index"`

	// Registration boolean, false if deregistration
	Registration bool `json:"registration"`
}

type TransactionDelegations struct {
	// Epoch in which the delegation becomes active
	ActiveEpoch int `json:"active_epoch"`

	// Bech32 delegation stake address
	Address string `json:"address"`

	// Index of the certificate within the transaction
	CertIndex int `json:"cert_index"`

	// Index of the certificate within the transaction
	Index int `json:"index"`

	// Bech32 ID of delegated stake pool
	PoolId string `json:"pool_id"`
}

type TransactionWidrawal struct {
	// Bech32 withdrawal address
	Address string `json:"address"`

	// Withdrawal amount in Lovelaces
	Amount string `json:"amount"`
}

type TransactionContentMirs struct {
	// Bech32 stake address
	Address string `json:"address"`

	// MIR amount in Lovelaces
	Amount string `json:"amount"`

	// Index of the certificate within the transaction
	CertIndex int `json:"cert_index"`

	// Source of MIR funds
	Pot string `json:"pot"`
}

type TransactionContentPoolCerts []struct {
	// Epoch that the delegation becomes active
	ActiveEpoch int `json:"active_epoch"`

	// Index of the certificate within the transaction
	CertIndex int `json:"cert_index"`

	// Fixed tax cost of the stake pool in Lovelaces
	FixedCost string `json:"fixed_cost"`

	// Margin tax cost of the stake pool
	MarginCost float32 `json:"margin_cost"`
	Metadata   struct {
		// Description of the stake pool
		Description string `json:"description"`

		// Hash of the metadata file
		Hash string `json:"hash"`

		// Home page of the stake pool
		Homepage string `json:"homepage"`

		// Name of the stake pool
		Name string `json:"name"`

		// Ticker of the stake pool
		Ticker string `json:"ticker"`

		// URL to the stake pool metadata
		Url string `json:"url"`
	} `json:"metadata"`
	Owners []string `json:"owners"`

	// Stake pool certificate pledge in Lovelaces
	Pledge string `json:"pledge"`

	// Bech32 encoded pool ID
	PoolId string `json:"pool_id"`
	Relays []struct {
		// DNS name of the relay
		Dns string `json:"dns"`

		// DNS SRV entry of the relay
		DnsSrv string `json:"dns_srv"`

		// IPv4 address of the relay
		Ipv4 string `json:"ipv4"`

		// IPv6 address of the relay
		Ipv6 string `json:"ipv6"`

		// Network port of the relay
		Port int `json:"port"`
	} `json:"relays"`

	// Bech32 reward account of the stake pool
	RewardAccount string `json:"reward_account"`

	// VRF key hash
	VrfKey string `json:"vrf_key"`
}

type TransactionContentPoolRetires struct {
	// Index of the certificate within the transaction
	CertIndex int `json:"cert_index"`

	// Bech32 stake pool ID
	PoolId string `json:"pool_id"`

	// Retiring epoch
	RetiringEpoch int `json:"retiring_epoch"`
}

type TransactionContentMetadata struct {
	// Content of the metadata
	JsonMetadata interface{} `json:"json_metadata"`

	// Metadata label
	Label string `json:"label"`
}

type TransactionContentMetadataCbor struct {
	// Content of the CBOR metadata
	CborMetadata *string `json:"cbor_metadata"`

	// Metadata label
	Label string `json:"label"`
}

type TransactionRedeemer struct {
	TxIndex   int    `json:"tx_index"`
	Purpose   string `json:"purpose"`
	UnitMem   string `json:"unit_mem"`
	UnitSteps string `json:"unit_steps"`
	Fee       string `json:"fee"`
}
