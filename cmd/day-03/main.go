package main

import "aoc-2020/internal"

func countTrees(input []string, jumpRight int, jumpDown int) (result int) {
	for i := 0; i < len(input); i += jumpDown {
		line := input[i]
		lineIdx := (i / jumpDown) * jumpRight
		realLineIdx := lineIdx % len(line)
		pos := string(line[realLineIdx])
		if pos == "#" {
			result++
		}
	}
	return
}

func main() {
	lines := internal.ReadFileLines("cmd/day-03/input.txt")
	println(countTrees(lines, 3, 1))

	slope1Trees := countTrees(lines, 1, 1)
	slope2Trees := countTrees(lines, 3, 1)
	slope3Trees := countTrees(lines, 5, 1)
	slope4Trees := countTrees(lines, 7, 1)
	slope5Trees := countTrees(lines, 1, 2)

	println(slope1Trees * slope2Trees * slope3Trees * slope4Trees * slope5Trees)
}
