package state

import "day16/input"

type State struct {
	agentPosition int
	rooms         []Room
}

func InitState(inputs []input.Input) State {
	rooms := []Room{}
	for _, input := range inputs {
		valve := Valve{name: input.Valve.Name, rate: input.Valve.Rate}
		room := Room{valve: valve, connectIndexes: []int{}}
		rooms = append(rooms, room)
	}
	for i := range inputs {

	}
	agentPosition := 0
	return State{agentPosition: agentPosition, rooms: rooms}
}

func (s State) Reward() int {
	r := 0
	for _, room := range s.rooms {
		r += room.valve.reward()
	}
	return r
}

func (s State) nextPossiblePositions() []int {
	return []int{}
}
