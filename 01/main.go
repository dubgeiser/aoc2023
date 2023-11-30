package main

import (
	"aoc2023/lib/file"
	"fmt"
)

type Processor struct {
}

func (lp *Processor) ProcessLine(line string) {
}

func part1(input []int) int {
	return 0
}

func part2(input []int) int {
	return 0
}

func main() {
	lp := &Processor{}
	lineCount, err := file.ReadLines("./input", lp)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
}
