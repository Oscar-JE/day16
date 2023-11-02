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

type StateSequences struct {
	stateSequence []state.State
	hasNext       bool
}

func InitSequence(start state.State, length int) StateSequences {
	stateSeq := []state.State{start}
	for i := 1 ; i < length ; i++ {
			
	}
}

func (itr StateSequences) HasNext() bool {
	return itr.hasNext
}

func (itr *StateSequences) GetNext() []state.State {

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

func (itr *StateSequences) increment() {
	index := itr.indexToIncrement()
	if index == 0 {
		itr.hasNext = false
		return
	}
	parent := itr.stateSequence[index-1]
	possibleNext := state.NextPossibleStates(parent)
	child := itr.stateSequence[index]
	itr.stateSequence[index] = possibleNext[indexOf(child, possibleNext)+1]
	itr.setToZeroAfterIndex(index)
}

func (itr StateSequences) indexToIncrement() int {
	for i := len(itr.stateSequence) - 1; i > 0; i-- {
		parent := itr.stateSequence[i-1]
		connected := state.NextPossibleStates(parent)
		child := itr.stateSequence[i]
		indexOfchild := indexOf(child, connected)
		if indexOfchild != len(connected)-1 {
			return i
		}
	}
	return 0
}

func (itr *StateSequences) setToZeroAfterIndex(index int) {
	for i := index + 1; i < len(itr.stateSequence); i++ {
		parent := itr.stateSequence[i-1]
		itr.stateSequence[i] = state.NextPossibleStates(parent)[0]
	}
}

func indexOf(element state.State, states []state.State) int {
	for i := range states {
		if states[i].Eq(element) {
			return i
		}
	}
	return -1
}
