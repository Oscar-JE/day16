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

// logiken om att undvika loopar bör ligga i state sequence
// är detta en bra ide ?
// tar den imorgon

type StateSequence struct {
	sequence []state.State
}

func (s StateSequence) Get(index int) state.State { // todo flytta innehålet i state paketet till model
	return s.sequence[index]
}

func (s StateSequence) getLast() state.State {
	return s.sequence[len(s.sequence)-1]
}

func (s StateSequence) NextPossibleStates() []state.State {
	nextStates := state.NextPossibleStates(s.getLast())
	nextPossibleStates := []state.State{}
	for i := range nextStates {
		if !s.AlreadyVisited(nextStates[i]) {
			nextPossibleStates = append(nextPossibleStates, nextStates[i])
		}
	}
	if len(nextPossibleStates) == 0 {
		nextPossibleStates = append(nextPossibleStates, s.getLast())
	}
	return nextPossibleStates
}

func (s StateSequence) AlreadyVisited(nextState state.State) bool {
	for i := range s.sequence {
		if s.sequence[i].Eq(nextState) {
			return true
		}
	}
	return false
}
func (s StateSequence) SubSequence(index int) StateSequence {
	return StateSequence{s.sequence[0 : index+1]}
}

func (s StateSequence) indexToIncrement() int {
	for i := len(s.sequence) - 1; i > 0; i-- {
		subSeq := s.SubSequence(i - 1)
		connected := subSeq.NextPossibleStates()
		child := s.sequence[i]
		indexOfchild := indexOf(child, connected)
		if indexOfchild != len(connected)-1 {
			return i
		}
	}
	return 0
}

func (itr *StateSequences) setToZeroAfterIndex(index int) {
	for i := index + 1; i < len(itr.stateSequence.sequence); i++ {
		subSeq := itr.stateSequence.SubSequence(i - 1)
		itr.stateSequence.sequence[i] = subSeq.NextPossibleStates()[0]
	}
}

/// ______________________

type StateSequences struct {
	stateSequence StateSequence
	hasNext       bool
}

func InitSequence(start state.State, length int) StateSequences {
	stateSeq := []state.State{start}
	for i := 1; i < length; i++ {
		stateSeq = append(stateSeq, state.NextPossibleStates(start)[0])
	}
	return StateSequences{stateSequence: StateSequence{stateSeq}, hasNext: true}
}

func (itr StateSequences) HasNext() bool {
	return itr.hasNext
}

func (itr *StateSequences) GetNext() []state.State {

	retPath := SequenceCopy(itr.stateSequence.sequence)
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
	index := itr.stateSequence.indexToIncrement()
	if index == 0 {
		itr.hasNext = false
		return
	}
	subSeq := itr.stateSequence.SubSequence(index - 1) // off by one ?
	possibleNext := subSeq.NextPossibleStates()
	child := itr.stateSequence.Get(index)
	itr.stateSequence.sequence[index] = possibleNext[indexOf(child, possibleNext)+1]
	itr.setToZeroAfterIndex(index)
}

func indexOf(element state.State, states []state.State) int {
	for i := range states {
		if states[i].Eq(element) {
			return i
		}
	}
	return -1
}
