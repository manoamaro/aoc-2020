package internal

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func OpenFile(path string) *os.File {
	file, err := os.Open("cmd/day-01/input.txt")
	if err != nil {
		log.Fatalf("failed to open")
	}
	return file
}

func ReadFileSliceInt(path string) (r []int) {
	file := OpenFile(path)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		value, _ := strconv.Atoi(scanner.Text())
		r = append(r, value)
	}

	file.Close()
	return
}
