package parse

import (
	"testing"
)

func TestParseLine(t *testing.T) {
	inLine := "Valve AA has flow rate=0; tunnels lead to valves DD, II, BB"
	expectedInput := InitInput("AA", 0, []string{"DD", "II", "BB"})
	actual := parseLine(inLine)
	if !actual.Eq(expectedInput) {
		t.Errorf("Incorrect input parsing")
	}
}
