package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var inputTest = []int{1721, 979, 366, 299, 675, 1456 }

func openInput() (r []int) {
	file, err := os.Open("cmd/day-01/input.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		r = append(r, value)
	}

	file.Close()

	return
}

func part01(input []int) int {
	for _, v1 := range input {
		for _, v2 := range input {
			if v1 + v2 == 2020 {
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
	input := openInput()
	fmt.Println(part01(input))
	fmt.Println(part02(input))
}
