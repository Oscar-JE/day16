package method

import "day16/model/state"

type Path struct {
	states []state.State
}

func (p Path) TotalReward() int { // denna bör redan vara testbar ; behöver denna ens vara en struct. Lite overkill
	total := 0
	for i := range p.states {
		state := p.states[i]
		total += state.Reward()
	}
	return total
}

type PathIter struct {
	p       Path
	hasNext bool
}

func (itr PathIter) HasNext() bool {
	return itr.hasNext
}

func (itr PathIter) lastPath() bool { // kan denna returnera true Ja för vi kollar inte på sista
	lastPath := true
	for i := 0; i < len(itr.p.states)-1; i++ {
		nextPossibleStates := state.NextPossibleStates(itr.p.states[i])
		lastPath = lastPath && nextPossibleStates[len(nextPossibleStates)-1].Eq(itr.p.states[i+1])
	}
	return !lastPath
}

func (itr *PathIter) GetNext() Path {
	if itr.lastPath() {
		itr.hasNext = false // vi kommer returnera den sista pathen i detta anrop
	}
	retPath := itr.p
	// skapa en deepCopy  , behöver jag det ? Nej , behövs enbart på det state vi muterar
	// hela koceptet faller förresten, Hur hanterar vi om en ändring i ett state ovanför inte längre når till state som kommer efter.
	// Då kommer vi behöve en state enumrering för att hantera det. Vilekt vi ioförsig redan har men det kan komma att bli komplicerat.
	// Om inte next states alltid returnerar det som är i ordning. Fast vänta det här spelar ingen roll pga att listan som av Next states är determinisktisk
	// vid ett skifte ska ändå alla efterföljande state sättas om till det första möjliga statet i listan
	// frågan är om vi skulle komma undan med en snabbskriven rekursion och sedan kunna gå vidare ?

	return retPath
}
