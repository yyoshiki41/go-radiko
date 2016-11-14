package radiko

import (
	"context"
	"encoding/base64"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

// AuthorizeToken returns an enables auth_token and error,
// and sets auth_token in Client.
// Is is a alias function that wraps Auth1Fms and Auth2Fms.
func (c *Client) AuthorizeToken(ctx context.Context, pngFilePath string) (string, error) {
	authToken, length, offset, err := c.Auth1Fms(ctx)
	if err != nil {
		return "", err
	}

	f, err := os.Open(pngFilePath)
	if err != nil {
		return "", err
	}

	b := make([]byte, length)
	if _, err = f.ReadAt(b, offset); err != nil {
		return "", err
	}
	partialKey := base64.StdEncoding.EncodeToString(b)

	if _, err := c.Auth2Fms(ctx, authToken, partialKey); err != nil {
		return "", err
	}

	c.setAuthTokenHeader(authToken)
	return authToken, nil
}

// Auth1Fms returns authToken, keyLength, keyOffset and error.
func (c *Client) Auth1Fms(ctx context.Context) (string, int64, int64, error) {
	apiEndpoint := apiPath(apiV2, "auth1_fms")

	req, err := c.newRequest(ctx, "POST", apiEndpoint, &Params{
		header: map[string]string{
			radikoAppHeader:        radikoApp,
			radikoAppVersionHeader: radikoAppVersion,
			radikoUserHeader:       radikoUser,
			radikoDeviceHeader:     radikoDevice,
		},
	})
	if err != nil {
		return "", 0, 0, err
	}

	resp, err := c.Call(req)
	defer resp.Body.Close()

	authToken := resp.Header.Get(radikoAuthTokenHeader)
	keyLength := resp.Header.Get(radikoKeyLentghHeader)
	keyOffset := resp.Header.Get(radikoKeyOffsetHeader)

	length, err := strconv.ParseInt(keyLength, 10, 64)
	if err != nil {
		return "", 0, 0, err
	}
	offset, err := strconv.ParseInt(keyOffset, 10, 64)
	if err != nil {
		return "", 0, 0, err
	}

	return authToken, length, offset, err
}

// Auth2Fms enables the given authToken.
func (c *Client) Auth2Fms(ctx context.Context, authToken, partialKey string) ([]string, error) {
	apiEndpoint := apiPath(apiV2, "auth2_fms")

	req, err := c.newRequest(ctx, "POST", apiEndpoint, &Params{
		header: map[string]string{
			radikoAppHeader:        radikoApp,
			radikoAppVersionHeader: radikoAppVersion,
			radikoUserHeader:       radikoUser,
			radikoDeviceHeader:     radikoDevice,
			radikoAuthTokenHeader:  authToken,
			radikoPartialKeyHeader: partialKey,
		},
	})
	if err != nil {
		return nil, err
	}

	resp, err := c.Call(req)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	s := strings.Split(string(b), ",")
	return s, nil
}
