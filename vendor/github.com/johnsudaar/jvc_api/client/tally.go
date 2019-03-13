package client

import (
	"github.com/pkg/errors"
)

type TallyIndication string
type TallyPriority string

const (
	TallyOff     TallyIndication = "Off"
	TallyProgram TallyIndication = "Program"
	TallyPreview TallyIndication = "Preview"

	TallyPriorityCamera TallyPriority = "Camera"
	TallyPriorityWeb    TallyPriority = "Web"
)

type setStudioTallyParams struct {
	Indication TallyIndication `json:"Indication"`
}

func (c HTTPClient) SetStudioTally(indication TallyIndication) error {
	_, err := c.makeRequest("SetStudioTally", setStudioTallyParams{
		Indication: indication,
	})
	if err != nil {
		return errors.Wrap(err, "fail to call HTTP API")
	}

	return nil
}

func (c HTTPClient) SetTallyLampPriority(priority TallyPriority) error {
	_, err := c.makeRequest("SetTallyLampPriority", map[string]TallyPriority{
		"Priority": priority,
	})
	if err != nil {
		return errors.Wrap(err, "fail to call HTTP API")
	}
	return nil
}
