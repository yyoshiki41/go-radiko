package radiko

import "testing"

func TestDownloadBinary(t *testing.T) {
	_, err := downloadBinary()
	if err != nil {
		t.Errorf("Failed to download binary: %s", err)
	}
}
