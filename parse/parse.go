package parse

import (
	"bufio"
	"day16/input"
	"os"
	"strconv"
	"strings"
)

func ParseFile(filePath string) []input.Input {
	inputs := []input.Input{}
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

func parseLine(line string) input.Input {
	valveAndConnectionStrings := strings.Split(line, ";")
	valveLine := valveAndConnectionStrings[0]
	connectionsLine := valveAndConnectionStrings[1]
	valve := parseValve(valveLine)
	connections := parseConnections(connectionsLine)
	return input.Input{Valve: valve, Connections: connections}
}

func parseValve(line string) input.Valve {
	line = strings.Trim(line, " ")
	words := strings.Split(line, " ")
	name := words[1]
	rateSection := words[len(words)-1]
	rateRep := strings.Split(rateSection, "=")[1]
	rate, parseError := strconv.Atoi(rateRep)
	if parseError != nil {
		panic("convert to int error")
	}
	return input.Valve{Name: name, Rate: rate}
}

func parseConnections(line string) []string {
	spliter := "valve"
	if strings.Contains(line, "valves") {
		spliter = "valves"
	}
	names := strings.Split(line, spliter)[1]
	names = strings.Trim(names, " ")
	splittedNames := []string{}
	if strings.Contains(names, ", ") {
		splittedNames = strings.Split(names, ", ")
	} else {
		splittedNames = append(splittedNames, strings.Trim(names, " "))
	}

	return splittedNames
}
