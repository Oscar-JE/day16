package method

import (
	"day16/input"
	"day16/model/state"
	"testing"
)

func TestTotalReward(t *testing.T) {
	inputs := []input.Input{input.InitInput("AA", 0, []string{"BB"}), input.InitInput("BB", 1, []string{"AA"})}
	s0 := state.InitState(inputs)
	s1 := s0.DeepCopy()
	s1.SetPosition("BB")
	s2 := s1.DeepCopy()
	s2.SetValve("BB", true)
	p := StateSequence{[]state.State{s0, s1, s2}}
	reward := p.TotalReward()
	if reward != 1 {
		t.Errorf("reward not wath we expected")
	}
}
