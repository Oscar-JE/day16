package model

import (
	"day16/intopr"
)

type Room struct {
	valve          Valve
	connectIndexes []int
}

func (r Room) Eq(o Room) bool {
	vb := r.valve.Eq(o.valve)
	return vb && intopr.ArrayEqual(r.connectIndexes, o.connectIndexes)
}

func (r *Room) SetConnectIndexes(indexes []int) {
	r.connectIndexes = indexes
}

func (r Room) GetConnections() []int {
	return r.connectIndexes
}

type Valve struct {
	name string
	rate int
	open bool
}

func (v Valve) Eq(o Valve) bool {
	return v.name == o.name && v.rate == o.rate && v.open == o.open
}

func (v Valve) reward() int {
	if v.open {
		return v.rate
	}
	return 0
}

func add(a int, b int) int {
	return a + b
}
