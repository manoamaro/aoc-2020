package main

import (
	"aoc-2020/internal"
	"fmt"
	"sort"
	"strconv"
)

var test = "35\n20\n15\n25\n47\n40\n62\n55\n65\n95\n102\n117\n150\n182\n127\n219\n299\n277\n309\n576"

func main() {

	//input := strings.Split(test, "\n")
	input := internal.ReadFileLines("cmd/day-09/input.txt")

	// Parse to Int
	inputNumbers := make([]int, 0)
	for _, s := range input {
		intValue, _ := strconv.Atoi(s)
		inputNumbers = append(inputNumbers, intValue)
	}

	preambleSize := 25

	invalidNumber := 0

	for i := preambleSize; i < len(inputNumbers); i++ {
		preamble := inputNumbers[i-preambleSize : i]
		if !isValid(preamble, inputNumbers[i]) {
			invalidNumber = inputNumbers[i]
			break
		}
	}

	fmt.Println(invalidNumber)

	part2Result := 0

	// Caterpillar method O(n)
	front, total := 0, 0
	for back := 0; back < len(inputNumbers); back++ {
		for front < len(inputNumbers) && (total+inputNumbers[front]) <= invalidNumber {
			total += inputNumbers[front]
			front += 1
		}
		if total == invalidNumber {
			foundSegment := inputNumbers[back:front]
			sort.Ints(foundSegment)
			part2Result = foundSegment[0] + foundSegment[len(foundSegment)-1]
			break
		}
		total -= inputNumbers[back]
	}

	fmt.Println(part2Result)
}

func isValid(preamble []int, value int) bool {
	for i := 0; i < len(preamble)-1; i++ {
		for j := i + 1; j < len(preamble); j++ {
			if preamble[i]+preamble[j] == value {
				return true
			}
		}
	}
	return false
}
