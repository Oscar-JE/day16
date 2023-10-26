package state

type Room struct {
	valve          Valve
	connectIndexes []int
}

type Valve struct {
	name string
	rate int
	open bool
}

func (v Valve) reward() int {
	if v.open {
		return v.rate
	}
	return v.rate
}
