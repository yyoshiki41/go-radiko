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
	if programs != nil {
		t.Errorf("Programs is nil.")
	}
}

func TestGetStationMaps(t *testing.T) {
	client, err := New()
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	m, err := client.GetStationMaps(ctx, areaIDTokyo)
	if err != nil {
		t.Errorf("Error: %s", err)
	}
	if len(m) == 0 {
		t.Errorf("Program map is empty.")
	}
}
