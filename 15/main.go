package main

import (
	"aoc2023/lib/file"
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

type Solution struct {
	answer1 int
	answer2 int
}

func (s *Solution) ProcessLine(i int, line string) {
	boxes := [256][]string{}
	label2f := map[string]int{}
	reOp := regexp.MustCompile(`-|=`)
	for _, step := range strings.Split(line, ",") {
		// Part 1
		s.answer1 += Hash(step)

		// Part 2
		p := reOp.Split(step, -1)
		label := p[0]
		bi := Hash(label)
		f := 1
		if len(p) == 2 {
			f, _ = strconv.Atoi(p[1])
		}
		label2f[label] = f
		if strings.Contains(step, "=") {
			if !slices.Contains(boxes[bi], label) {
				boxes[bi] = append(boxes[bi], label)
			}
		} else {
			boxes[bi] = slices.DeleteFunc(boxes[bi], func(e string) bool { return e == label })
		}
	}
	for bi, b := range boxes {
		for li, label := range b {
			s.answer2 += (bi + 1) * (li + 1) * label2f[label]
		}
	}
}

func Hash(s string) int {
	hash := 0
	for _, step := range []byte(s) {
		hash += int(step)
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
