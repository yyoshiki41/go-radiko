package radiko

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"path"
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

	radikoAuthTokenHeader  = "X-Radiko-AuthToken"
	radikoKeyLentghHeader  = "X-Radiko-KeyLength"
	radikoKeyOffsetHeader  = "X-Radiko-KeyOffset"
	radikoPartialKeyHeader = "X-Radiko-Partialkey"

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
	// TODO: 再設計
	URL *url.URL

	HTTPClient *http.Client
}

// New returns new Client struct.
func New() (*Client, error) {
	parsedURL, err := url.Parse(defaultEndpoint)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse url: %s", err)
	}

	if httpClient == nil {
		return nil, errors.New("HTTP Client is nil")
	}

	return &Client{
		URL:        parsedURL,
		HTTPClient: httpClient,
	}, nil
}

func (c *Client) newRequest(verb, apiEndpoint string, params *Params) (*http.Request, error) {
	// TODO: 再設計
	u := *c.URL
	u.Path = path.Join(c.URL.Path, apiEndpoint)

	// Add query parameters
	urlQuery := u.Query()
	for k, v := range params.query {
		urlQuery.Set(k, v)
	}
	u.RawQuery = urlQuery.Encode()

	req, err := http.NewRequest(verb, u.String(), nil)
	if err != nil {
		return nil, err
	}

	// Add request headers
	for k, v := range params.header {
		req.Header.Set(k, v)
	}

	return req, nil
}

// Params is the list of options to pass to the request.
type Params struct {
	// query is a map of key-value pairs that will be added to the Request.
	query map[string]string
	// header is a map of key-value pairs that will be added to the Request.
	header map[string]string
}

// SetHTTPClient overrides the default HTTP client.
func SetHTTPClient(client *http.Client) {
	httpClient = client
}

func apiPath(apiVersion, pathStr string) string {
	return path.Join(apiVersion, "api", pathStr)
}
