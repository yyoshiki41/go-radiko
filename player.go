package radiko

import (
	"io"
	"net/http"
	"os"
)

const (
	playerURL = "http://radiko.jp/apps/js/flash/myplayer-release.swf"
)

// DownloadPlayer downloads swf player file.
func DownloadPlayer(path string) error {
	resp, err := http.Get(playerURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	_, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return err
}
