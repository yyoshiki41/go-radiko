package radiko

import (
	"net/http"

	"github.com/yyoshiki41/go-radiko/internal"
)

// GetChunklistFromM3U8 returns a slice of url.
func GetChunklistFromM3U8(uri string) ([]string, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return internal.GetChunklistFromM3U8(resp.Body)
}
