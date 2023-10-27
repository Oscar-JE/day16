package input

func FindIndexes(connections []string, inputs []Input) []int {
	connectionsIndexes := []int{}
	for _, connection := range connections {
		connectionsIndexes = append(connectionsIndexes, FindIndex(connection, inputs))
	}
	return connectionsIndexes
}

func FindIndex(name string, inputs []Input) int {
	for i := range inputs {
		if name == inputs[i].Valve.Name {
			return i
		}
	}
	panic("no input index was found for name")
}
