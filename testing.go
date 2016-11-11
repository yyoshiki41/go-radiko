package radiko

import (
	"flag"
	"io/ioutil"
	"os"
	"path"
	"testing"
)

const (
	areaIDTokyo = "JP13"
)

var (
	outsideJP bool

	goPath      string
	testdataDir string
)

func init() {
	flag.BoolVar(&outsideJP, "outjp", false, "Skip tests if outside Japan.")
	flag.Parse()

	goPath := os.Getenv("GOPATH")
	if goPath == "" {
		panic("$GOPATH is empty.")
	}
	testdataDir = path.Join(goPath, "src", "github.com/yyoshiki41/go-radiko", "testdata")
}

// For skipping tests.
// radiko.jp restricts use from outside Japan.
func isOutsideJP() bool {
	return outsideJP
}

func createTestTempDir(t *testing.T) (string, func()) {
	dir, err := ioutil.TempDir("", "test-go-radiko")
	if err != nil {
		t.Fatalf("Failed to create temp dir: %s", err)
	}

	return dir, func() { os.RemoveAll(dir) }
}
