package radiko

import (
	"context"
	"testing"
	"time"

	"github.com/yyoshiki41/go-radiko/internal"
)

func TestGetStationsByAreaID(t *testing.T) {
	if isOutsideJP() {
		t.Skip("Skipping test in limited mode.")
	}

	client, err := New("")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	stations, err := client.GetStationsByAreaID(ctx, areaIDTokyo, time.Now())
	if err != nil {
		t.Error(err)
	}
	if len(stations) == 0 {
		t.Error("Stations is nil.")
	}
}

func TestGetStations(t *testing.T) {
	if isOutsideJP() {
		t.Skip("Skipping test in limited mode.")
	}

	client, err := New("")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	stations, err := client.GetStations(ctx, time.Now())
	if err != nil {
		t.Error(err)
	}
	if len(stations) == 0 {
		t.Error("Stations is nil.")
	}
}

func TestGetNowProgramsByAreaID(t *testing.T) {
	if isOutsideJP() {
		t.Skip("Skipping test in limited mode.")
	}

	client, err := New("")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	programs, err := client.GetNowProgramsByAreaID(ctx, areaIDTokyo)
	if err != nil {
		t.Error(err)
	}
	if len(programs) == 0 {
		t.Errorf("Programs is nil.")
	}
}

func TestGetNowPrograms(t *testing.T) {
	if isOutsideJP() {
		t.Skip("Skipping test in limited mode.")
	}

	client, err := New("")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	ctx := context.Background()
	programs, err := client.GetNowPrograms(ctx)
	if err != nil {
		t.Error(err)
	}
	if len(programs) == 0 {
		t.Error("Programs is nil.")
	}
}

func TestGetProgramByStartTime(t *testing.T) {
	if isOutsideJP() {
		t.Skip("Skipping test in limited mode.")
	}

	client, err := New("")
	if err != nil {
		t.Fatalf("Failed to construct client: %s", err)
	}

	// Tests in ANN
	stationID := "LFR"
	n := time.Now()
	if n.Weekday() == time.Sunday {
		// If it is Sunday, ANN will not be broadcasted.
		n.Add(-24 * time.Hour)
	}
	y, m, d := n.Date()
	// ANN starts at 01:00 AM on Monday to Saturday in JST.
	start := time.Date(y, m, d, 16, 0, 0, 0, time.UTC)
	end := time.Date(y, m, d, 18, 0, 0, 0, time.UTC)

	ctx := context.Background()
	prog, err := client.GetProgramByStartTime(ctx, stationID, start)
	if err != nil {
		t.Error(err)
	}
	expected := internal.Datetime(end)
	if expected != prog.To {
		t.Errorf("expected %s, but %s", expected, prog.To)
	}
}
