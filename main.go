package main

import (
	"day16/intopr"
	"day16/method"
	model "day16/model"
	"day16/parse"
	"fmt"
	"time"
)

func main() {
	inputs := parse.ParseFile("input_short.txt")
	start := model.InitState(inputs)
	iter := method.InitSequenceItr(start, 30)
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
