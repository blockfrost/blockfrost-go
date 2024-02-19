package blockfrost_test

import (
	"encoding/json"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

var validPayload string = `{"id":"47668401-c3a4-42d4-bac1-ad46515924a3","webhook_id":"cf68eb9c-635f-415e-a5a8-6233638f28d7","created":1650013856,"type":"block","payload":{"time":1650013853,"height":7126256,"hash":"f49521b67b440e5030adf124aee8f88881b7682ba07acf06c2781405b0f806a4","slot":58447562,"epoch":332,"epoch_slot":386762,"slot_leader":"pool1njjr0zn7uvydjy8067nprgwlyxqnznp9wgllfnag24nycgkda25","size":34617,"tx_count":13,"output":"13403118309871","fees":"4986390","block_vrf":"vrf_vk197w95j9alkwt8l4g7xkccknhn4pqwx65c5saxnn5ej3cpmps72msgpw69d","previous_block":"9e3f5bfc9f0be44cf6e14db9ed5f1efb6b637baff0ea1740bb6711786c724915","next_block":null,"confirmations":0}}`

func TestWebhookEpochUnmarshal(t *testing.T) {
	const epochWebhookPayload string = `{"id":"5ffcaf65-7961-4377-9741-fa0c76176a4b","webhook_id":"b592db93-ec26-4ecc-8800-8a14b3a2806f","created":1654811689,"api_version":1,"type":"epoch","payload":{"previous_epoch":{"epoch":343,"start_time":1654379091,"end_time":1654811091,"first_block_time":1654379116,"last_block_time":1654811087,"block_count":20994,"tx_count":463239,"output":"106038169691018243","fees":"162340782180","active_stake":"24538091587045780"},"current_epoch":{"epoch":344,"start_time":1654811091,"end_time":1655243091}}}`
	var txEvent blockfrost.WebhookEventEpoch

	err := json.Unmarshal([]byte(epochWebhookPayload), &txEvent)

	if err != nil {
		t.Fatalf("Error parsing webhook JSON: %v\n", err)
		return
	}
}

func TestWebhookDelegationUnmarshal(t *testing.T) {
	const delegationWebhookPayload string = `{"id":"1e42b77f-80cf-4860-b0e8-1f49f9cf39ae","webhook_id":"59acf943-73cf-403c-9f03-cae3d654ce65","created":1708341715,"api_version":1,"type":"delegation","payload":[{"tx":{"hash":"15947383c54e1412fcb93cabc5c290cfd3d19d4a85fc3857f509d80fad926cf3","block":"1856de7062e94fae14922d8592ace4e6755d2b401c2ca950329a06334f6681f8","block_height":9952712,"block_time":1708341706,"slot":116775415,"index":19,"output_amount":[{"unit":"lovelace","quantity":"4267389351"}],"fees":"174125","deposit":"2000000","size":425,"invalid_before":null,"invalid_hereafter":"116782570","utxo_count":2,"withdrawal_count":0,"mir_cert_count":0,"delegation_count":1,"stake_cert_count":1,"pool_update_count":0,"pool_retire_count":0,"asset_mint_or_burn_count":0,"redeemer_count":0,"valid_contract":true},"delegations":[{"index":1,"cert_index":1,"address":"stake1u8tjqpexqt9jpj88z6syk5u3cfkmm7647ghm4pslufrr4lst2p42s","pool_id":"pool1f2wfjqkf2wx6jq93pdck6hgmy9zgw32lmvrq9zejl7scqxjqfze","active_epoch":469,"pool":{"pool_id":"pool1f2wfjqkf2wx6jq93pdck6hgmy9zgw32lmvrq9zejl7scqxjqfze","hex":"4a9c9902c9538da900b10b716d5d1b214487455fdb06028b32ffa180","vrf_key":"18dd29025d4dbf5899ff54da2ab4a0f25f49a0b04d339d64108aa55538a1d484","blocks_minted":73,"blocks_epoch":19,"live_stake":"25984523381075","live_size":0.001142245872813442,"live_saturation":0.3664818900732455,"live_delegators":1379,"active_stake":"21504417226880","active_size":0.000939885072940621,"declared_pledge":"2000000","live_pledge":"4000000","margin_cost":0.06,"fixed_cost":"340000000","reward_account":"stake1u8ah6l0g6unccujkufkwufjqstj8tv3639ftkykvk69a65s8h62ap","owners":["stake1uxgm4h9atd0s457acgx85vnc0kamg68h8j2wtrzrgp7t5eqm9vydu"],"registration":["7c3103489d13e34cc0897c35b825284032603e8b931db589087ffbbe5b768f17"],"retirement":[]}}]}]}`
	var txEvent blockfrost.WebhookEventDelegation

	err := json.Unmarshal([]byte(delegationWebhookPayload), &txEvent)

	if err != nil {
		t.Fatalf("Error parsing webhook JSON: %v\n", err)
		return
	}
}

func TestWebhookTransactionUnmarshal(t *testing.T) {
	const txWebhookPayload string = `{"id":"22f5af45-5292-49f4-bbe3-20e30af7faa9","webhook_id":"59acf943-73cf-403c-9f03-cae3d654ce65","created":1708341161,"api_version":1,"type":"transaction","payload":[{"tx":{"hash":"b3aedfba6472ecb080f6ea80a1c38072c18b3be864bdf4d7f6f3b7c419cb0e6a","block":"31092d73df46de327ffbdd20094006443e7b7c195511732d238cc5aeda2136be","block_height":1949042,"block_time":1708341158,"slot":52657958,"index":0,"output_amount":[{"unit":"lovelace","quantity":"19434547"},{"unit":"2b556df9f37c04ef31b8f7f581c4e48174adcf5041e8e52497d815564e6f646546656564","quantity":"1"}],"fees":"282339","deposit":"0","size":666,"invalid_before":"52657924","invalid_hereafter":"52658044","utxo_count":6,"withdrawal_count":0,"mir_cert_count":0,"delegation_count":0,"stake_cert_count":0,"pool_update_count":0,"pool_retire_count":0,"asset_mint_or_burn_count":0,"redeemer_count":1,"valid_contract":true},"inputs":[{"address":"addr_test1wqlcn3pks3xdptxjw9pqrqtcx6ev694sstsruw3phd57ttg0lh0zq","amount":[{"unit":"lovelace","quantity":"64000000"}],"tx_hash":"1a28375b01a8928f21a87598e24165ab5806e7c3e77a5eb4d78c56e3cf07894c","output_index":0,"data_hash":null,"inline_datum":null,"reference_script_hash":"3f89c436844cd0acd2714201817836b2cd16b082e03e3a21bb69e5ad","collateral":false,"reference":true},{"address":"addr_test1vz7st3e4f5tqzyldkwdr9gkwvpzlfr6364egl7ha4ck7emctt2gnq","amount":[{"unit":"lovelace","quantity":"9000000"}],"tx_hash":"1a34755269f26e60f22a1f6bceb65228b9636e6e095bb63bab0db9d46c6b9050","output_index":1,"data_hash":null,"inline_datum":null,"reference_script_hash":null,"collateral":false,"reference":false},{"address":"addr_test1wqlcn3pks3xdptxjw9pqrqtcx6ev694sstsruw3phd57ttg0lh0zq","amount":[{"unit":"lovelace","quantity":"2000000"},{"unit":"2b556df9f37c04ef31b8f7f581c4e48174adcf5041e8e52497d815564e6f646546656564","quantity":"1"}],"tx_hash":"739752466614abac5c25a839dbfef31fbcaa10034d3ac12b4714767c5f722b7b","output_index":0,"data_hash":"d1cf4baf9990dc6298cf0c9819ba04a5d9471ac72664bf74a3eeff437e8eb7c2","inline_datum":"d87a9fd8799f581cbd05c7354d160113edb39a32a2ce6045f48f51d5728ffafdae2decefd8799fd8799f1a000b1d231b0000018dc0f5c159ffffffff","reference_script_hash":null,"collateral":false,"reference":false},{"address":"addr_test1vz7st3e4f5tqzyldkwdr9gkwvpzlfr6364egl7ha4ck7emctt2gnq","amount":[{"unit":"lovelace","quantity":"8716886"}],"tx_hash":"b2640f54d48b4402dab521934ff410cadd4c1c715f47f2aaffccad771cecd887","output_index":2,"data_hash":null,"inline_datum":null,"reference_script_hash":null,"collateral":false,"reference":false},{"address":"addr_test1vz7st3e4f5tqzyldkwdr9gkwvpzlfr6364egl7ha4ck7emctt2gnq","amount":[{"unit":"lovelace","quantity":"9000000"}],"tx_hash":"739752466614abac5c25a839dbfef31fbcaa10034d3ac12b4714767c5f722b7b","output_index":1,"data_hash":null,"inline_datum":null,"reference_script_hash":null,"collateral":true,"reference":false}],"outputs":[{"address":"addr_test1wqlcn3pks3xdptxjw9pqrqtcx6ev694sstsruw3phd57ttg0lh0zq","amount":[{"unit":"lovelace","quantity":"2000000"},{"unit":"2b556df9f37c04ef31b8f7f581c4e48174adcf5041e8e52497d815564e6f646546656564","quantity":"1"}],"output_index":0,"data_hash":"ef6fe82cd8a45498a3fc35fc45a6158bb92654571805c96b1f7810b675e354d7","inline_datum":"d87a9fd8799f581cbd05c7354d160113edb39a32a2ce6045f48f51d5728ffafdae2decefd8799fd8799f1a000b1d231b0000018dc1113c83ffffffff","collateral":false,"reference_script_hash":null},{"address":"addr_test1vz7st3e4f5tqzyldkwdr9gkwvpzlfr6364egl7ha4ck7emctt2gnq","amount":[{"unit":"lovelace","quantity":"9000000"}],"output_index":1,"data_hash":null,"inline_datum":null,"collateral":false,"reference_script_hash":null},{"address":"addr_test1vz7st3e4f5tqzyldkwdr9gkwvpzlfr6364egl7ha4ck7emctt2gnq","amount":[{"unit":"lovelace","quantity":"8434547"}],"output_index":2,"data_hash":null,"inline_datum":null,"collateral":false,"reference_script_hash":null},{"address":"addr_test1vz7st3e4f5tqzyldkwdr9gkwvpzlfr6364egl7ha4ck7emctt2gnq","amount":[{"unit":"lovelace","quantity":"5392385"}],"output_index":3,"data_hash":null,"inline_datum":null,"collateral":true,"reference_script_hash":null}]}]}`
	var txEvent blockfrost.WebhookEventTransaction

	err := json.Unmarshal([]byte(txWebhookPayload), &txEvent)

	if err != nil {
		t.Fatalf("Error parsing webhook JSON: %v\n", err)
		return
	}
}

func TestWebhookBlockUnmarshal(t *testing.T) {
	var txEvent blockfrost.WebhookEventBlock

	err := json.Unmarshal([]byte(validPayload), &txEvent)

	if err != nil {
		t.Fatalf("Error parsing webhook JSON: %v\n", err)
		return
	}

}

func TestVerifyWebhookSignature(t *testing.T) {
	event, _ := blockfrost.VerifyWebhookSignatureIgnoringTolerance([]byte(validPayload),
		// 2 signatures - first one is invalid, second one is valid
		"t=1650013856,v1=abc,t=1650013856,v1=f4c3bb2a8b0c8e21fa7d5fdada2ee87c9c6f6b0b159cc22e483146917e195c3e",
		"59a1eb46-96f4-4f0b-8a03-b4d26e70593a")

	// Unmarshal the event data into an appropriate struct depending on its Type
	var blockEvent blockfrost.WebhookEventBlock
	switch event.Type {
	case "block":
		err := json.Unmarshal([]byte(validPayload), &blockEvent)

		if err != nil {
			t.Fatalf("Error parsing webhook JSON: %v\n", err)
			return
		}
	default:
		t.Fatalf("Invalid webhook type: %s\n", event.Type)
	}

	jsonData, err := json.Marshal(blockEvent)
	if err != nil {
		t.Fatalf("Error marshaling to JSON: %s", err)
	}

	jsonString := string(jsonData)

	if jsonString != validPayload {
		t.Fatalf("\nexpected %v \ngot %v", validPayload, jsonString)

	}
}

func TestVerifyWebhookSignatureOutOfTolerance(t *testing.T) {
	_, err := blockfrost.VerifyWebhookSignature([]byte(validPayload),
		"t=1650013856,v1=f4c3bb2a8b0c8e21fa7d5fdada2ee87c9c6f6b0b159cc22e483146917e195c3e",
		"59a1eb46-96f4-4f0b-8a03-b4d26e70593a")

	if err != blockfrost.ErrTooOld {
		t.Fatalf("\nsuccess expected %v \ngot %v", blockfrost.ErrTooOld, err)
	}

}
func TestVerifyWebhookSignatureInvalidHeader(t *testing.T) {
	_, err := blockfrost.VerifyWebhookSignature([]byte(validPayload),
		"v1=f4c3bb2a8b0c8e21fa7d5fdada2ee87c9c6f6b0b159cc22e483146917e195c3e",
		"59a1eb46-96f4-4f0b-8a03-b4d26e70593a")

	if err != blockfrost.ErrInvalidHeader {
		t.Fatalf("\nsuccess expected %v \ngot %v", blockfrost.ErrInvalidHeader, err)
	}

}
func TestVerifyWebhookSignatureNoSupportedSchema(t *testing.T) {
	_, err := blockfrost.VerifyWebhookSignature([]byte(validPayload),
		"v42=f4c3bb2a8b0c8e21fa7d5fdada2ee87c9c6f6b0b159cc22e483146917e195c3e",
		"59a1eb46-96f4-4f0b-8a03-b4d26e70593a")

	if err != blockfrost.ErrNoValidSignature {
		t.Fatalf("\nsuccess expected %v \ngot %v", blockfrost.ErrNoValidSignature, err)
	}

}
func TestVerifyWebhookSignatureNoHeader(t *testing.T) {
	_, err := blockfrost.VerifyWebhookSignature([]byte(validPayload),
		"",
		"59a1eb46-96f4-4f0b-8a03-b4d26e70593a")

	if err != blockfrost.ErrNotSigned {
		t.Fatalf("\nsuccess expected %v \ngot %v", blockfrost.ErrNotSigned, err)
	}

}
