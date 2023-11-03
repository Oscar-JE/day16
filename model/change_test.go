package model

import (
	"day16/input"
	"testing"
)

func TestStateChangeSimple(t *testing.T) {
	inputs := []input.Input{input.InitInput("AA", 0, []string{"BB"}), input.InitInput("BB", 1, []string{"AA"})}
	s := InitState(inputs)
	exp0 := s.DeepCopy()
	exp0.rooms[0].valve.open = true
	exp1 := s.DeepCopy()
	exp1.agentPosition = 1
	nexts := NextPossibleStates(s)
	if !exp1.Eq(nexts[0]) {
		t.Errorf("Wrong order in nextStates of faulty calculation")
	}
}

func TestDeepCopy(t *testing.T) {
	inputs := []input.Input{input.InitInput("AA", 0, []string{"BB"}), input.InitInput("BB", 1, []string{"AA"})}
	s := InitState(inputs)
	sCopy := s.DeepCopy()
	sCopy.rooms[0].valve.open = true
	if s.rooms[0].valve.open {
		t.Errorf("failed to produce a deep copy")
	}
}
