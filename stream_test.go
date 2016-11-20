package radiko

import (
	"context"
	"testing"
)

func TestGetStreamMultiURL(t *testing.T) {
	client, err := New("")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	items, err := client.GetStreamMultiURL(ctx, "LFR")
	if err != nil {
		t.Error(err)
	}

	const expected = 4
	if actual := len(items); expected != actual {
		t.Errorf("expected %d, but %d", expected, actual)
	}
}
