package radiko

import (
	"context"
	"io/ioutil"
	"strconv"
	"strings"
)

// Auth1Fms returns authToken, keyLength, keyOffset and error.
func (c *Client) Auth1Fms(ctx context.Context) (string, int64, int64, error) {
	apiEndpoint := apiPath(apiV2, "auth1_fms")

	req, err := c.newRequest("POST", apiEndpoint, &Params{
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

	req = req.WithContext(ctx)
	resp, err := c.HTTPClient.Do(req)
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

	req, err := c.newRequest("POST", apiEndpoint, &Params{
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

	req = req.WithContext(ctx)
	resp, err := c.HTTPClient.Do(req)
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	s := strings.Split(string(b), ",")
	return s, nil
}
