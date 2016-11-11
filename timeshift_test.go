package radiko

import "testing"

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
