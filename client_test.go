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

func TestEmptyHTTPClient(t *testing.T) {
	var c *http.Client
	SetHTTPClient(c)

	client, err := New()
	if err == nil {
		t.Errorf(
			"Should detect HTTPClient is nil.\nclient: %v", client)
	}
}

func TestSetHTTPClient(t *testing.T) {
	const expected = 1 * time.Second
	SetHTTPClient(&http.Client{Timeout: expected})

	client, err := New()
	if err != nil {
		t.Errorf("Failed to construct client: %s", err)
	}
	if client.HTTPClient.Timeout != expected {
		t.Errorf("expected %d, but %d", expected, client.HTTPClient.Timeout)
	}
}
