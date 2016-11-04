package radiko

import (
	"net/http"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	_, err := New()
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}
}

func TestSetHTTPClient(t *testing.T) {
	const timeout = 1 * time.Second
	SetHTTPClient(&http.Client{Timeout: timeout})

	client, err := New()
	if err != nil {
		t.Errorf("Failed to construct client: %s", err)
	}
	if client.HTTPClient.Timeout != timeout {
		t.Errorf("Failed to set http client: %s", err)
	}
}
