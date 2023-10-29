package state

func NextPossibleStates(sIn State) []State {
	nextStates := []State{}
	nextPositions := NextPossiblePossition(sIn)
	for _, nextPosition := range nextPositions {
		s := sIn.DeepCopy()
		nextStates = append(nextStates, NextState(s, nextPosition))
	}
	return nextStates
}

func NextPossiblePossition(state State) []int {
	possibleNextPositions := []int{}
	possibleNextPositions = append(possibleNextPositions, state.agentPosition)
	connectedRooms := state.rooms[state.agentPosition].connectIndexes
	possibleNextPositions = append(possibleNextPositions, connectedRooms...)
	return possibleNextPositions
}

func NextState(state State, nextPosition int) State {
	if state.agentPosition == nextPosition {
		state.rooms[state.agentPosition].valve.open = true
	} else {
		state.agentPosition = nextPosition
	}
	return state
}
