package radiko

import (
	"context"
	"testing"
)

func TestGetStations(t *testing.T) {
	if isOutsideJP() {
		t.Skip("Skipping test in limited mode.")
	}

	client, err := New()
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	stations, err := client.GetStations(ctx, areaIDTokyo, "20161109")
	if err != nil {
		t.Error(err)
	}
	if stations == nil {
		t.Error("Stations is nil.")
	}
}

func TestGetNowPrograms(t *testing.T) {
	if isOutsideJP() {
		t.Skip("Skipping test in limited mode.")
	}

	client, err := New()
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	programs, err := client.GetNowPrograms(ctx, areaIDTokyo)
	if err != nil {
		t.Error(err)
	}
	if programs == nil {
		t.Errorf("Programs is nil.")
	}
}
