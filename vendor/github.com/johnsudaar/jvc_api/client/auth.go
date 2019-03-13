package client

import (
	"fmt"
	"net/http"

	"github.com/pkg/errors"
	dac "github.com/xinsnake/go-http-digest-auth-client"
)

func (c *HTTPClient) authenticate() error {
	digestTransport := dac.NewTransport(c.user, c.password)
	client := &http.Client{
		Transport: &digestTransport,
	}

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("http://%s/api.php", c.ip), nil)
	if err != nil {
		return errors.Wrap(err, "fail to build request")
	}

	resp, err := client.Do(req)
	if err != nil {
		return errors.Wrap(err, "fail to make request")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return errors.Wrapf(err, "unexpected status code: %s", resp.Status)
	}

	sessionID := ""

	for _, cookie := range resp.Cookies() {
		if cookie.Name == "SessionID" {
			sessionID = cookie.Value
		}
	}

	if sessionID == "" {
		return errors.New("No SessionID received")
	}

	c.sessionID = sessionID

	return nil
}
