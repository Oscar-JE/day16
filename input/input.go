package input

type Input struct {
	Valve       Valve
	Connections []string
}

func InitInput(name string, rate int, connections []string) Input {
	valve := Valve{Name: name, Rate: rate}
	return Input{Valve: valve, Connections: connections}
}

func (i Input) Eq(other Input) bool {
	equalValve := i.Valve.Eq(other.Valve)
	equalConnections := false
	if len(i.Connections) == len(other.Connections) {
		equalConnections = true
		for conInd := range i.Connections {
			equalConnections = equalConnections && i.Connections[conInd] == other.Connections[conInd]
		}
	}
	return equalValve && equalConnections
}

type Valve struct {
	Name string
	Rate int
}

func (v Valve) Eq(other Valve) bool {
	return v.Name == other.Name && v.Rate == other.Rate
}
