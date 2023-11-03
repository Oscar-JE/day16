package model

import (
	"day16/input"
	"testing"
)

func TestConstructorDept(t *testing.T) {
	inputs := []input.Input{input.InitInput("AA", 0, []string{"BB"}), input.InitInput("BB", 1, []string{"AA"})}
	s := InitState(inputs)
	s.rooms[1].valve.open = true
	if s.rooms[0].valve.open {
		t.Errorf("state have shallow copies of room")
	}
}
