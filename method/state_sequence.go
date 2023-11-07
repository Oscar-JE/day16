package method

import model "day16/model"

func TotalReward(stateSequnce []model.State) int {
	total := 0
	for i := range stateSequnce {
		state := stateSequnce[i]
		total += state.Reward()
	}
	return total
}

type StateSequence struct {
	sequence []model.State
}

func (s StateSequence) Get(index int) model.State {
	return s.sequence[index]
}

func (s StateSequence) getLast() model.State {
	return s.sequence[len(s.sequence)-1]
}

func (s StateSequence) NextPossibleStates() []model.State {
	nextStates := model.NextPossibleStates(s.getLast())
	nextPossibleStates := []model.State{}
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

func (s StateSequence) AlreadyVisited(nextState model.State) bool {
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


func (itr *SequenceIter) setToZeroAfterIndex(index int) {
	for i := index + 1; i < len(itr.stateSequence.sequence); i++ {
		subSeq := itr.stateSequence.SubSequence(i - 1)
		itr.stateSequence.sequence[i] = subSeq.NextPossibleStates()[0]
	}
}

func indexOf(element model.State, states []model.State) int {
	for i := range states {
		if states[i].Eq(element) {
			return i
		}
	}
	return -1
}
