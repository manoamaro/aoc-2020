package main

import (
	"aoc-2020/internal"
	"fmt"
	"strconv"
	"strings"
)

var inputTest = []string{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"}

func part01(input []string) (count int) {
	count = 0
	for _, i := range input {
		fields := strings.Fields(i)
		amount := strings.Split(fields[0], "-")
		amountMin, _ := strconv.Atoi(amount[0])
		amountMax, _ := strconv.Atoi(amount[1])
		letter := strings.Trim(fields[1], ":")
		password := strings.TrimSpace(fields[2])
		letterCount := strings.Count(password, letter)
		if letterCount >= amountMin && letterCount <= amountMax {
			count++
		}
	}
	return
}

func part02(input []string) (count int) {
	count = 0
	for _, i := range input {
		fields := strings.Fields(i)
		positions := strings.Split(fields[0], "-")
		pos1, _ := strconv.Atoi(positions[0])
		pos2, _ := strconv.Atoi(positions[1])
		letter := strings.Trim(fields[1], ":")
		password := strings.TrimSpace(fields[2])
		password1, password2 := string(password[pos1-1]), string(password[pos2-1])
		if (password1 == letter || password2 == letter) && password1 != password2 {
			count++
		}
	}
	return
}

func main() {
	fmt.Println(part01(internal.ReadFileLines("cmd/day-02/input.txt")))
	fmt.Println(part02(internal.ReadFileLines("cmd/day-02/input.txt")))
}
