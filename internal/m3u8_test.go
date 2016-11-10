package internal

import (
	"bufio"
	"os"
	"path"
	"testing"
)

func readTestData(fileName string) *os.File {
	const testDir = "github.com/yyoshiki41/go-radiko/testdata"

	GOPATH := os.Getenv("GOPATH")
	if GOPATH == "" {
		panic("$GOPATH is empty.")
	}
	f, err := os.Open(path.Join(GOPATH, "src", testDir, fileName))
	if err != nil {
		panic(err)
	}
	return f
}

func TestGetURIFromM3U8(t *testing.T) {
	expected := "https://radiko.jp/v2/api/ts/chunklist/NejwTOkX.m3u8"

	input := bufio.NewReader(readTestData("uri.m3u8"))
	u, err := GetURIFromM3U8(input)
	if err != nil {
		t.Error(err)
	}
	if u != expected {
		t.Errorf("expected %d, but %d", expected, u)
	}
}

func TestGetChunklistFromM3U8(t *testing.T) {
	input := bufio.NewReader(readTestData("chunklist.m3u8"))
	chunklist, err := GetChunklistFromM3U8(input)
	if err != nil {
		t.Error(err)
	}
	if len(chunklist) == 0 {
		t.Error("chunklist is empty.")
	}
}
