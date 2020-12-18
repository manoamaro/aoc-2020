package main

import (
	"aoc-2020/internal"
	"fmt"
	"strings"
)

func main() {

	rawInput := internal.ReadFileLines("cmd/day-06/input.txt")
	input := strings.Split(strings.Join(rawInput, "\n"), "\n\n")

	countAnyone := 0
	countEveryone := 0
	for _, group := range input {
		eachPeople := strings.Split(group, "\n")
		allTogether := strings.Join(eachPeople, "")

		answers := make(map[string]int)

		for _, answer := range allTogether {
			answers[string(answer)]++
		}

		countAnyone += len(answers)

		for _, count := range answers {
			if count == len(eachPeople) {
				countEveryone++
			}
		}
	}
	fmt.Println(countAnyone)
	fmt.Println(countEveryone)
}
