package atem

import "fmt"

func (a *ATEM) InputPorts() []string {
	inputs := make([]string, ioConfigs[a.atemType].Inputs)
	for i := 0; i < len(inputs); i++ {
		inputs[i] = fmt.Sprintf("Input_%d", i+1)
	}
	return inputs
}

func (a *ATEM) OutputPorts() []string {
	return []string{"PGM", "AUX", "MV"}
}
