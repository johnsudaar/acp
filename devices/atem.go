package devices

import (
	"encoding/json"
	"fmt"

	"github.com/johnsudaar/acp/events"
	"github.com/pkg/errors"
)

type ATEM struct {
	IP     string `json:"ip"`
	Port   string `json:"port"`
	MyPort string `json:"my_port"`
}

func init() {
	Register("ATEM", NewATEM)
}

func NewATEM(message json.RawMessage) (Device, error) {
	var atem ATEM
	err := json.Unmarshal(message, &atem)
	if err != nil {
		return nil, errors.Wrap(err, "invalid payload")
	}
	return &atem, nil

}

func (a *ATEM) WriteEvent(outputName string, event events.Event) error {
	if !ValidOutput(outputName, a) {
		return fmt.Errorf("invalid output name: %s", outputName)
	}
	return nil
}

func (a *ATEM) Connect(inputName string, port OutputPort) error {
	if !ValidInput(inputName, a) {
		return fmt.Errorf("invalid input name: %s", inputName)
	}

	if !port.Valid() {
		return fmt.Errorf("invalid output name: %s", port.Name)
	}

	return nil
}

func (a ATEM) AvailableInputs() []string {
	return []string{
		"Input1",
		"Input2",
		"Input3",
		"Input4",
	}
}

func (a ATEM) AvailableOutputs() []string {
	return []string{}
}
