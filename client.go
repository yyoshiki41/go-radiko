package radiko

import (
	"net/http"
	"net/url"
)

const (
	defaultEndpoint = "https://radiko.jp"
	// envEndpoint is the environment variable that overrrides the defaultEndpoint.
	envEndpoint = "RADIKO_ENDPOINT"

	apiV2 = "v2"
	apiV3 = "v3"

	// HTTP Headers
	radikoAppHeader        = "X-Radiko-App"
	radikoAppVersionHeader = "X-Radiko-App-Version"
	radikoUserHeader       = "X-Radiko-User"
	radikoDeviceHeader     = "X-Radiko-Device"

	radikoAuthTokenHeader = "X-Radiko-AuthToken"
	radikoKeyLentghHeader = "X-Radiko-KeyLength"
	radikoKeyOffsetHeader = "X-Radiko-KeyOffset"

	radikoApp        = "pc_ts"
	radikoAppVersion = "4.0.0"
	radikoUser       = "test-stream"
	radikoDevice     = "pc"
)

// Client represents a single connection to radiko API endpoint.
type Client struct {
	URL *url.URL

	HTTPClient *http.Client
	HTTPHeader http.Header
}
