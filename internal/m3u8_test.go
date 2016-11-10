package internal

import (
	"strings"
	"testing"
)

const inputM3U8 = `
#EXTM3U
#EXT-X-VERSION:3
#EXT-X-STREAM-INF:PROGRAM-ID=1,BANDWIDTH=52973,CODECS="mp4a.40.5"
https://radiko.jp/v2/api/ts/chunklist/NejwTOkX.m3u8
`

func TestGetURIFromM3U8(t *testing.T) {
	expected := "https://radiko.jp/v2/api/ts/chunklist/NejwTOkX.m3u8"

	s := strings.NewReader(inputM3U8)
	u, err := GetURIFromM3U8(s)
	if err != nil {
		t.Error(err)
	}
	if u != expected {
		t.Errorf("expected %d, but %d", expected, u)
	}
}
