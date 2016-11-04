package radiko

import (
	"io/ioutil"
	"os"
	"testing"
)

func createTestTempDir(t *testing.T) (string, func()) {
	dir, err := ioutil.TempDir("", "test-go-radiko")
	if err != nil {
		t.Fatalf("Failed to create temp dir: ", err)
	}

	return dir, func() { os.RemoveAll(dir) }
}
