package radiko

import (
	"context"
	"testing"
	"time"
)

func TestTimeshiftPlaylistM3U8(t *testing.T) {
	if isOutsideJP() {
		t.Skip("Skipping test in limited mode.")
	}

	client, err := New("")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	_ = client
}

func TestEmptyStationIDTimeshiftPlaylistM3U8(t *testing.T) {
	client, err := New("")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	_, err = client.TimeshiftPlaylistM3U8(context.Background(), "", time.Now())
	if err == nil {
		t.Error("Should detect an error.")
	}
}
