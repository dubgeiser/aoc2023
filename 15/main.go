package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strings"
)

type Solution struct {
	answer1 int
	answer2 int
}

func (s *Solution) ProcessLine(i int, line string) {
	sequence := strings.Split(line, ",")
	for _, step := range sequence {
		s.answer1 += Hash(step)
	}
}

func Hash(s string) int {
	hash := 0
	for _, step := range []byte(s) {
		hash +=int(step)
		hash = (hash * 17) % 256
	}
	return hash
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
