package main

import (
	"aoc-2020/internal"
	"fmt"
	"strconv"
	"strings"
)

var test = "nop +0\nacc +1\njmp +4\nacc +3\njmp -3\nacc -99\nacc +1\njmp -4\nacc +6"

func main() {

	//input := strings.Split(test, "\n")
	input := internal.ReadFileLines("cmd/day-08/input.txt")
	accumulator, _ := testInstructions(input)
	fmt.Println(accumulator)

	for i := 0; i < len(input); i++ {
		instruction := input[i]
		if strings.Contains(instruction, "acc") {
			continue
		} else if strings.Contains(instruction, "jmp") {
			// switch
			input[i] = strings.ReplaceAll(input[i], "jmp", "nop")
		} else if strings.Contains(instruction, "nop") {
			// switch
			input[i] = strings.ReplaceAll(input[i], "nop", "jmp")
		}

		// test
		accumulatorP2, finished := testInstructions(input)
		if finished {
			fmt.Println(accumulatorP2)
			break
		}

		// switch back to old instruction
		input[i] = instruction
	}

}

func testInstructions(instructions []string) (acc int, finished bool) {
	visitedInstructions := make(map[int]int, 0)

	for i := 0; i < len(instructions); {
		_, exists := visitedInstructions[i]
		if exists {
			return acc, false
		}
		instructionLine := instructions[i]
		rawInstruction := strings.Split(instructionLine, " ")
		visitedInstructions[i] = i
		switch rawInstruction[0] {
		case "nop":
			i++
		case "acc":
			value, _ := strconv.Atoi(rawInstruction[1])
			acc += value
			i++
		case "jmp":
			value, _ := strconv.Atoi(rawInstruction[1])
			i += value
		}
	}

	return acc, true
}
