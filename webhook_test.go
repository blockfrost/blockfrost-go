package blockfrost_test

import (
	"encoding/json"
	"testing"

	"github.com/blockfrost/blockfrost-go"
)

var validPayload string = `{"id":"47668401-c3a4-42d4-bac1-ad46515924a3","webhook_id":"cf68eb9c-635f-415e-a5a8-6233638f28d7","created":1650013856,"type":"block","payload":{"time":1650013853,"height":7126256,"hash":"f49521b67b440e5030adf124aee8f88881b7682ba07acf06c2781405b0f806a4","slot":58447562,"epoch":332,"epoch_slot":386762,"slot_leader":"pool1njjr0zn7uvydjy8067nprgwlyxqnznp9wgllfnag24nycgkda25","size":34617,"tx_count":13,"output":"13403118309871","fees":"4986390","block_vrf":"vrf_vk197w95j9alkwt8l4g7xkccknhn4pqwx65c5saxnn5ej3cpmps72msgpw69d","previous_block":"9e3f5bfc9f0be44cf6e14db9ed5f1efb6b637baff0ea1740bb6711786c724915","next_block":null,"confirmations":0}}`

func TestVerifyWebhookSignature(t *testing.T) {
	event, _ := blockfrost.VerifyWebhookSignatureIgnoringTolerance([]byte(validPayload),
		// 2 signatures - first one is invalid, second one is valid
		"t=1650013856,v1=abc,t=1650013856,v1=f4c3bb2a8b0c8e21fa7d5fdada2ee87c9c6f6b0b159cc22e483146917e195c3e",
		"59a1eb46-96f4-4f0b-8a03-b4d26e70593a")

	_, ok := event.(*blockfrost.WebhookEventBlock)
	if !ok {
		jsonData, _ := json.Marshal(event)
		t.Fatalf("Invalid webhook type %s", jsonData)
	}

	jsonData, err := json.Marshal(event)
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
