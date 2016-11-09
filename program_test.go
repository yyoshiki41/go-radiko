package radiko

import (
	"context"
	"testing"
)

func TestGetNowPrograms(t *testing.T) {
	client, err := New()
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	programs, err := client.GetNowPrograms(ctx, areaIDTokyo)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if programs == nil {
		t.Errorf("Programs is nil.")
	}
}

func TestGetStationList(t *testing.T) {
	client, err := New()
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	s, err := client.GetStationList(ctx, areaIDTokyo)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if s == nil {
		t.Error("Stations is empty.")
	}
}
