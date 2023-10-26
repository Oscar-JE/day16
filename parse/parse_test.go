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

func TestParseLineGG(t *testing.T) {
	inLine := "Valve HH has flow rate=22; tunnel leads to valve GG"
	expectedInput := InitInput("HH", 22, []string{"GG"})
	actual := parseLine(inLine)
	if !actual.Eq(expectedInput) {
		t.Errorf("Incorrect input parsing")
	}
}
