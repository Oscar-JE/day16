package method

import "day16/model/state"

func TotalReward(stateSequnce []state.State) int {
	total := 0
	for i := range stateSequnce {
		state := stateSequnce[i]
		total += state.Reward()
	}
	return total
}

type StateSequnces struct {
	stateSequence []state.State
	hasNext       bool
}

func (itr StateSequnces) HasNext() bool {
	return itr.hasNext
}

func (itr *StateSequnces) GetNext() []state.State {

	retPath := SequenceCopy(itr.stateSequence)
	itr.increment()
	return retPath
}

func SequenceCopy(states []state.State) []state.State {
	copy := []state.State{}
	for i := range states {
		copy = append(copy, states[i].DeepCopy())
	}
	return copy
}

func (itr *StateSequnces) increment() {
	last := false
	for i := len(itr.stateSequence) - 1; i > 0; i-- {
		parent := itr.stateSequence[i-1]
		connected := state.NextPossiblePossition(parent)
		child := 
	}
}
