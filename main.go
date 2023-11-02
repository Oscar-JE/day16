package main

import (
	"day16/intopr"
	"day16/method"
	"day16/model/state"
	"day16/parse"
	"fmt"
	"time"
)

func main() {
	inputs := parse.ParseFile("input_short.txt")
	start := state.InitState(inputs)
	iter := method.InitSequence(start, 15)
	startTime := time.Now()
	maximum := 0
	for iter.HasNext() {
		sequence := iter.GetNext()
		score := method.TotalReward(sequence)
		maximum = intopr.Max(maximum, score)
	}
	fmt.Println(maximum)
	fmt.Println(time.Since(startTime).Seconds())
}
