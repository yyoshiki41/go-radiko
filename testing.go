package radiko

import (
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"testing"
)

const (
	areaIDTokyo = "JP13"
)

var (
	outsideJP bool

	testdataDir string
)

func init() {
	// FIXME:
	// affects the outside packages
	flag.BoolVar(&outsideJP, "outjp", false, "Skip tests if outside Japan.")
	flag.Parse()

	GOPATH := os.Getenv("GOPATH")
	testdataDir = filepath.Join(GOPATH, "src", "github.com/yyoshiki41/go-radiko", "testdata")
}

// For skipping tests.
// radiko.jp restricts use from outside Japan.
func isOutsideJP() bool {
	return outsideJP
}

// Should restore defaultHTTPClient if SetHTTPClient is called.
func teardownHTTPClient() {
	SetHTTPClient(&http.Client{Timeout: defaultHTTPTimeout})
}

func createTestTempDir(t *testing.T) (string, func()) {
	dir, err := ioutil.TempDir("", "test-go-radiko")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %s", err)
	}

	return dir, func() { os.RemoveAll(dir) }
}
