package model

import (
	"day16/input"
	"day16/intopr"
	"strconv"
)

type State struct {
	agentPosition int
	rooms         []Room
}

func InitState(inputs []input.Input) State {
	rooms := []Room{}
	for _, inputEl := range inputs {
		valve := Valve{name: inputEl.Valve.Name, rate: inputEl.Valve.Rate, open: false}
		connections := input.FindIndexes(inputEl.Connections, inputs)
		room := Room{valve: valve, connectIndexes: connections}
		rooms = append(rooms, room)
	}
	agentPosition := input.FindIndex("AA", inputs)
	return State{agentPosition: agentPosition, rooms: rooms}
}

func (s *State) SetPosition(roomRep string) {
	s.agentPosition = s.findIndexOfRoom(roomRep)
}

func (s State) findIndexOfRoom(roomRep string) int {
	index := -1
	for i := range s.rooms {
		if s.rooms[i].valve.name == roomRep {
			index = i
			break
		}
	}
	return index
}

func (s *State) SetValve(roomRep string, open bool) {
	index := s.DeepCopy().findIndexOfRoom(roomRep)
	s.rooms[index].valve.open = open
}

func (s State) DeepCopy() State {
	copy := s
	newRoomArray := []Room{}
	for i := range s.rooms {
		newRoomArray = append(newRoomArray, s.rooms[i]) // Det bör leda till en djup tycker jag
	}
	copy.rooms = newRoomArray
	return copy
}

func (s State) Indexed() int {
	// hur gör vi detta så att det inte blir några krockar ?
	// Det blir ett väldigt stort antal positioner. Kan bli här som value iteration fallerar
	// varje position ska kombineras med varje möjlig konfiguration av valven
	// # positioner = #valv , # valvConfigurationer = 2^{#valv}
	// om vi tänker oss det som en matris med formen (#positioner, #vavlvConfigurationer)
	// då skulle index med rad först bli följande index = position * #valvConfigurationer + vavlvConfiguration
	return s.agentPosition*s.maximumValveConfigurationNumber() + s.valveCongigurationNumber()
}

func (s State) valveCongigurationNumber() int {
	valveConfigNumber := 0
	for i := range s.rooms {
		if s.rooms[i].valve.open {
			valveConfigNumber += intopr.PowerOfTwo(i)
		}
	}
	return valveConfigNumber
}
func (s State) maximumValveConfigurationNumber() int {
	return intopr.PowerOfTwo(len(s.rooms)) - 1
}

func (s State) Eq(other State) bool {
	if len(s.rooms) != len(other.rooms) { // borde dock vara omöjligt att tappa ett antal rum
		return false
	}
	rommsEq := true
	for i := range s.rooms {
		rommsEq = rommsEq && s.rooms[i].Eq(other.rooms[i])
	}
	agentb := s.agentPosition == other.agentPosition
	return agentb && rommsEq
}

func (s State) Reward() int {
	r := 0
	for _, room := range s.rooms {
		r += room.valve.reward()
	}
	return r
}

func (s State) String() string {
	rep := ""
	rep += "Agent Position : " + s.rooms[s.agentPosition].valve.name + "\n"
	for ir := range s.rooms {
		room := s.rooms[ir]
		roomName := room.valve.name
		connectionNames := ""
		on := ""
		if room.valve.open {
			on = " open "
		} else {
			on = " closed "
		}
		for _, irr := range room.connectIndexes {
			connectionNames += s.rooms[irr].valve.name + " "
		}
		rep += roomName + on + "\t Rate: " + strconv.Itoa(room.valve.rate) + " \t Connections: " + connectionNames + "\n"
	}
	return rep
}
