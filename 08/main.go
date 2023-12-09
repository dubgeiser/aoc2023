package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strconv"
	"strings"
)

type Solution1 struct {
	lrs []int
	M   map[string][2]string
}

func (s *Solution1) ProcessLine(i int, line string) {
	if i == 0 {
		l := strings.ReplaceAll(strings.ReplaceAll(line, "L", "0"), "R", "1")
		for _, r := range l {
			if n, err := strconv.Atoi(string(r)); err == nil {
				s.lrs = append(s.lrs, n)
			}
		}
		return
	}
	if line == "" {
		return
	}
	var n, l, r string
	fmt.Sscanf(line, "%3s = (%3s, %3s)", &n, &l, &r)
	s.M[n] = [2]string{l, r}
}

func part1() {
	s := &Solution1{M: make(map[string][2]string)}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println()

	steps := 0
	var lr int
	for curr := "AAA"; curr != "ZZZ"; steps++ {
		lr = s.lrs[steps%len(s.lrs)]
		curr = s.M[curr][lr]
	}
	fmt.Println(steps)
}

func main() {
	part1()
}
