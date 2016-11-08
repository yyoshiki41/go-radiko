package radiko

import (
	"io/ioutil"
	"os"
	"testing"
)

const (
	areaIDTokyo = "JP13"
)

func createTestTempDir(t *testing.T) (string, func()) {
	dir, err := ioutil.TempDir("", "test-go-radiko")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %s", err)
	}

	return dir, func() { os.RemoveAll(dir) }
}
