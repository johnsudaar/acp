package atem

import (
	"fmt"

	"github.com/johnsudaar/acp/devices/types/switcher"
	"github.com/johnsudaar/atem"
	"github.com/pkg/errors"
)

var _ switcher.Switchable = &ATEM{}

func (a *ATEM) Switch(output, input string) error {
	if output != "PGM" {
		return fmt.Errorf("Cannot switch %s output", output)
	}
	var inputValue atem.VideoSource
	switch input {
	case "Input_1":
		inputValue = atem.Input_1
	case "Input_2":
		inputValue = atem.Input_2
	case "Input_3":
		inputValue = atem.Input_3
	case "Input_4":
		inputValue = atem.Input_4
	case "Input_5":
		inputValue = atem.Input_5
	case "Input_6":
		inputValue = atem.Input_6
	case "Input_7":
		inputValue = atem.Input_7
	case "Input_8":
		inputValue = atem.Input_8
	case "Input_9":
		inputValue = atem.Input_9
	case "Input_10":
		inputValue = atem.Input_10
	}

	err := a.client.SetProgram(atem.MESource0, inputValue)
	if err != nil {
		return errors.Wrap(err, "fail to switch")
	}
	return nil
}
