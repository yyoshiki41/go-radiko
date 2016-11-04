package radiko

import (
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	defaultHTTPTimeout = 90 * time.Second
	defaultEndpoint    = "https://radiko.jp"
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

var (
	httpClient = &http.Client{Timeout: defaultHTTPTimeout}
)

// Client represents a single connection to radiko API endpoint.
type Client struct {
	URL *url.URL

	HTTPClient *http.Client
	HTTPHeader http.Header
}

// New returns new Client struct.
func New() (*Client, error) {
	parsedURL, err := url.Parse(defaultEndpoint)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse url: %s", err)
	}

	return &Client{
		URL:        parsedURL,
		HTTPClient: httpClient,
		HTTPHeader: make(http.Header),
	}, nil
}

// SetHTTPClient overrides the default HTTP client.
func SetHTTPClient(client *http.Client) {
	httpClient = client
}
