package main

import (
	"day16/model/state"
	"day16/parse"
	"fmt"
)

func main() {
	inputs := parse.ParseFile("input_short.txt")
	s := state.InitState(inputs)
	fmt.Println(s)
	println("next states")
	nextStates := state.NextPossibleStates(s)
	for _, el := range nextStates {
		fmt.Println(el)
	}

}
