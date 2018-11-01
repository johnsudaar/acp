package devices

import (
	"encoding/json"
	"fmt"

	"github.com/johnsudaar/acp/events"
	"github.com/pkg/errors"
)

type JVCHM660 struct {
	Name     string `json:"name,omitempty"`
	IP       string `json:"ip"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func init() {
	Register("JVC_HM_660", NewJVCHM660)
}

func NewJVCHM660(message json.RawMessage) (Device, error) {
	var cam JVCHM660
	err := json.Unmarshal(message, &cam)
	if err != nil {
		return nil, errors.Wrap(err, "invalid payload")
	}
	return &cam, nil
}

func (j *JVCHM660) WriteEvent(outputName string, event events.Event) error {
	if !ValidOutput(outputName, j) {
		return fmt.Errorf("invalid output name: %s", outputName)
	}

	return nil
}

func (j *JVCHM660) Connect(inputName string, port OutputPort) error {
	return fmt.Errorf("invalid input name: %s", inputName)
}

func (j JVCHM660) AvailableInputs() []string {
	return []string{}
}

func (j JVCHM660) AvailableOutputs() []string {
	return []string{
		"VideoOut",
	}
}
