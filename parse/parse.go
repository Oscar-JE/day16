package parse

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parseFiel(filePath string) []Input {
	inputs := []Input{}
	reader, err := os.Open(filePath)
	if err != nil {
		panic("Error opening file")
	}
	scanner := bufio.NewScanner(reader)
	line := ""
	for scanner.Scan() {
		line = scanner.Text()
		inputs = append(inputs, parseLine(line))
	}
	return inputs
}

func parseLine(line string) Input {
	valveAndConnectionStrings := strings.Split(line, ";")
	valveLine := valveAndConnectionStrings[0]
	connectionsLine := valveAndConnectionStrings[1]
	valve := parseValve(valveLine)
	connections := parseConnections(connectionsLine)
	return Input{Valve: valve, Connections: connections}
}

func parseValve(line string) Valve {
	line = strings.Trim(line, " ")
	words := strings.Split(line, " ")
	name := words[1]
	rateSection := words[len(words)-1]
	rateRep := strings.Split(rateSection, "=")[1]
	rate, parseError := strconv.Atoi(rateRep)
	if parseError != nil {
		panic("convert to int error")
	}
	return Valve{Name: name, Rate: rate}
}

func parseConnections(line string) []string {
	names := strings.Trim(strings.Split(line, "valves")[1], " ")
	splittedNames := strings.Split(names, ", ")
	return splittedNames
}
