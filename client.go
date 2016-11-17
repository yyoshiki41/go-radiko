package radiko

import (
	"context"
	"errors"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"path"
	"time"
)

const (
	defaultEndpoint    = "https://radiko.jp"
	defaultHTTPTimeout = 120 * time.Second

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
	httpClient = &http.Client{
		Timeout: defaultHTTPTimeout,
		Jar:     nil,
	}
)

// Client represents a single connection to radiko API endpoint.
type Client struct {
	URL *url.URL

	httpClient      *http.Client
	authTokenHeader string
}

// New returns a new Client struct.
func New(authToken string) (*Client, error) {
	parsedURL, err := url.Parse(defaultEndpoint)
	if err != nil {
		return nil, err
	}

	if httpClient == nil {
		return nil, errors.New("A HTTP client is nil.")
	}

	return &Client{
		URL:             parsedURL,
		httpClient:      httpClient,
		authTokenHeader: authToken,
	}, nil
}

func (c *Client) setAuthTokenHeader(authToken string) {
	c.authTokenHeader = authToken
}

func (c *Client) setJar(jar *cookiejar.Jar) {
	c.httpClient.Jar = jar
}

func (c *Client) newRequest(ctx context.Context, verb, apiEndpoint string, params *Params) (*http.Request, error) {
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

	// Set the request's context
	if ctx == nil {
		return nil, errors.New("Context is nil.")
	}
	req = req.WithContext(ctx)

	// Add request headers
	for k, v := range params.header {
		req.Header.Set(k, v)
	}
	// For backwards compatibility with HTTP/1.0
	// https://tools.ietf.org/html/rfc7234#page-29
	req.Header.Set("pragma", "no-cache")
	// Add auth_token in HTTP Header
	if params.setAuthToken {
		req.Header.Set(radikoAuthTokenHeader, c.authTokenHeader)
	}

	return req, nil
}

// Do executes an API request.
func (c *Client) Do(req *http.Request) (*http.Response, error) {
	return c.httpClient.Do(req)
}

// Params is the list of options to pass to the request.
type Params struct {
	// query is a map of key-value pairs that will be added to the Request.
	query map[string]string
	// header is a map of key-value pairs that will be added to the Request.
	header map[string]string
	// setAuthToken is a boolean value. If true, set auth_token in HTTP Header.
	setAuthToken bool
}

// SetHTTPClient overrides the default HTTP client.
func SetHTTPClient(client *http.Client) {
	httpClient = client
}

func apiPath(apiVersion, pathStr string) string {
	return path.Join(apiVersion, "api", pathStr)
}
