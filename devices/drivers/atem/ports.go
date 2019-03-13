package atem

import "fmt"

func (a *ATEM) InputPorts() []string {
	inputs := make([]string, 12)
	for i := 0; i < len(inputs); i++ {
		inputs[i] = fmt.Sprintf("Input_%d", i+1)
	}
	inputs = append(inputs, "Tally")
	return inputs
}

func (a *ATEM) OutputPorts() []string {
	return []string{"PGM", "AUX", "MV"}
}
