package main

import (
	"aoc2023/lib/file"
	"fmt"
)

type Solution struct {
	answer1 int
	answer2 int
}

func (s *Solution) ProcessLine(i int, line string) {
}

func main() {
	s := &Solution{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.answer1)
	fmt.Println(s.answer2)
}
