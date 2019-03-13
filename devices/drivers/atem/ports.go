package atem

func (a *ATEM) InputPorts() []string {
	return []string{"Input_1", "Input_2", "Input_3", "Input_4"}
}

func (a *ATEM) OutputPorts() []string {
	return []string{"PGM", "AUX", "MV"}
}
