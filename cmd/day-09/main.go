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

	// Prefix sums to better performance
	prefixSums := make([]int64, len(inputNumbers)+1)
	for i := 1; i < len(prefixSums); i++ {
		prefixSums[i] = prefixSums[i-1] + int64(inputNumbers[i-1])
	}

	part2Result := 0

	// scan for the segment, using an increasing window size
	for windowSize := 2; windowSize < len(inputNumbers)-1; windowSize++ {
		// test the segments
		for i := 0; i < len(inputNumbers)-windowSize; i++ {
			if prefixSums[i+windowSize]-prefixSums[i] == int64(invalidNumber) {
				foundSegment := inputNumbers[i : i+windowSize]
				sort.Ints(foundSegment)
				part2Result = foundSegment[0] + foundSegment[len(foundSegment)-1]
				break
			}
		}
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
