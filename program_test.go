package radiko

import (
	"context"
	"testing"
	"time"
)

func TestGetStationsByAreaID(t *testing.T) {
	if isOutsideJP() {
		t.Skip("Skipping test in limited mode.")
	}

	client, err := New()
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	stations, err := client.GetStationsByAreaID(ctx, areaIDTokyo, time.Now())
	if err != nil {
		t.Error(err)
	}
	if stations == nil {
		t.Error("Stations is nil.")
	}
}

func TestGetStations(t *testing.T) {
	if isOutsideJP() {
		t.Skip("Skipping test in limited mode.")
	}

	client, err := New()
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	stations, err := client.GetStations(ctx, time.Now())
	if err != nil {
		t.Error(err)
	}
	if stations == nil {
		t.Error("Stations is nil.")
	}
}

func TestGetNowProgramsByAreaID(t *testing.T) {
	if isOutsideJP() {
		t.Skip("Skipping test in limited mode.")
	}

	client, err := New()
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	programs, err := client.GetNowProgramsByAreaID(ctx, areaIDTokyo)
	if err != nil {
		t.Error(err)
	}
	if programs == nil {
		t.Errorf("Programs is nil.")
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
	stations, err := client.GetNowPrograms(ctx)
	if err != nil {
		t.Error(err)
	}
	if stations == nil {
		t.Error("Stations is nil.")
	}
}
