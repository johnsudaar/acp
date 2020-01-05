package client

import "github.com/pkg/errors"

type ButtonParams struct {
	Kind string `json:"Kind"`
	Key  string `json:"Key"`
}

func (c HTTPClient) SendWebKeyEvent(kind, key string) error {
	_, err := c.makeRequest("SetWebKeyEvent", ButtonParams{
		Kind: kind,
		Key:  key,
	})
	if err != nil {
		return errors.Wrap(err, "fail to send web button event")
	}
	return nil
}
