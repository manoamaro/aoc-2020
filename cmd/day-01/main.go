package main

import (
	"aoc-2020/internal"
	"fmt"
)

var inputTest = []int{1721, 979, 366, 299, 675, 1456}

func part01(input []int) int {
	for _, v1 := range input {
		for _, v2 := range input {
			if v1+v2 == 2020 {
				return v1 * v2
			}
		}
	}
	return 0
}

func part02(input []int) int {
	for _, v1 := range input {
		for _, v2 := range input {
			for _, v3 := range input {
				if v1+v2+v3 == 2020 {
					return v1 * v2 * v3
				}
			}
		}
	}
	return 0
}

func main() {
	input := internal.ReadFileSliceInt("cmd/day-01/input.txt")
	fmt.Println(part01(input))
	fmt.Println(part02(input))
}
