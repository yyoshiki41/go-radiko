package radiko

import (
	"context"
	"path"
	"testing"
)

func TestAuthorizeToken(t *testing.T) {
	client, err := New("")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	pngPath := path.Join(testdataDir, "authkey.png")
	authToken, err := client.AuthorizeToken(ctx, pngPath)
	if err != nil {
		t.Error(err)
	}
	if len(authToken) == 0 {
		t.Errorf("AuthToken is empty.")
	}
}

func TestAuth1Fms(t *testing.T) {
	client, err := New("")
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

	client, err := New("")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}
	_ = client
}
