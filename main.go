package main

import (
	"day16/model/state"
	"day16/parse"
	"fmt"
)

func main() {
	inputs := parse.ParseFile("input_short.txt")
	state := state.InitState(inputs)
	fmt.Println(inputs)
	fmt.Println(state)

}
