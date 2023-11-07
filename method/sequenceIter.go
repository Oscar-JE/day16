package method

import model "day16/model"

type SequenceIter struct {
	stateSequence     StateSequence
	hasNext           bool
	intermediateValue int
}

func InitSequenceItr(start model.State, length int) SequenceIter {
	stateSeq := []model.State{start}
	for i := 1; i < length; i++ {
		stateSeq = append(stateSeq, model.NextPossibleStates(start)[0])
	}
	return SequenceIter{stateSequence: StateSequence{stateSeq}, hasNext: true, intermediateValue: 0}
}

func (itr SequenceIter) HasNext() bool {
	return itr.hasNext
}

func (itr *SequenceIter) GetNext() []model.State {

	retPath := SequenceCopy(itr.stateSequence.sequence)
	itr.increment()
	return retPath
}

func SequenceCopy(states []model.State) []model.State {
	copy := []model.State{}
	for i := range states {
		copy = append(copy, states[i].DeepCopy())
	}
	return copy
}

func (itr *SequenceIter) increment() {
	index := indexToIncrement(itr.stateSequence)
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

func indexToIncrement(s StateSequence) int { // är detta verkligen den här klassens ansvar? Nej det är itrs ansvar
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
