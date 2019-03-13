package client

import (
	"github.com/pkg/errors"
)

type HTTPClient struct {
	sessionID string
	user      string
	password  string
	ip        string
}

// New build a new HTTPClient and tries to authenticate against the camera API
func New(ip, user, password string) (*HTTPClient, error) {
	client := &HTTPClient{
		ip:       ip,
		user:     user,
		password: password,
	}

	err := client.authenticate()
	if err != nil {
		return nil, errors.Wrap(err, "fail to authenticate")
	}

	return client, nil
}
