package radiko

// Login returns
func (c *Client) Login(mail, password string) (string, error) {
	apiEndpoint := "ap/member/login/login"

	req, err := c.newRequest(ctx, "POST", apiEndpoint, &Params{})

	resp, err := c.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return "", nil
}
