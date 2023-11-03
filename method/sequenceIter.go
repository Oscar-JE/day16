package method

import model "day16/model"

type SequenceIter struct {
	stateSequence StateSequence
	hasNext       bool
}

func InitSequence(start model.State, length int) SequenceIter {
	stateSeq := []model.State{start}
	for i := 1; i < length; i++ {
		stateSeq = append(stateSeq, model.NextPossibleStates(start)[0])
	}
	return SequenceIter{stateSequence: StateSequence{stateSeq}, hasNext: true}
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
