package state

import "day16/input"

type State struct {
	agentPosition int
	rooms         []Room
}

func InitState(inputs []input.Input) State {
	rooms := []Room{}
	for _, inputEl := range inputs {
		valve := Valve{name: inputEl.Valve.Name, rate: inputEl.Valve.Rate}
		connections := input.FindIndexes(inputEl.Connections, inputs)
		room := Room{valve: valve, connectIndexes: connections}
		rooms = append(rooms, room)
	}
	agentPosition := input.FindIndex("AA", inputs)
	return State{agentPosition: agentPosition, rooms: rooms}
}

func (s State) Reward() int {
	r := 0
	for _, room := range s.rooms {
		r += room.valve.reward()
	}
	return r
}
