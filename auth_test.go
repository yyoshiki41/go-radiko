package radiko

import (
	"context"
	"testing"
)

func TestAuth1Fms(t *testing.T) {
	client, err := New()
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	authToken, length, offset, err := client.Auth1Fms(ctx)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(authToken) == 0 || length < 0 || offset < 0 {
		t.Errorf("AuthToken: %s, Length: %d, Offset: %d", authToken, length, offset)
	}
}

func TestAuth2Fms(t *testing.T) {
	if isOutsideJP() {
		t.Skip("Skipping test in limited mode.")
	}

	client, err := New()
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}
	_ = client

	/* TODO: Implement mock server
	ctx := context.Background()
	res, err := client.Auth2Fms(ctx, authToken, partialKey)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(res) == 0 {
		t.Error("Empty results")
	}
	*/
}
