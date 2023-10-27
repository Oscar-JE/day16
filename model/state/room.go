package state

type Room struct {
	valve          Valve
	connectIndexes []int
}

func (r *Room) SetConnectIndexes(indexes []int) {
	r.connectIndexes = indexes
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
