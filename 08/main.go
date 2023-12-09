package main

import (
	"aoc2023/lib/algo"
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

type Solution2 struct {
	lrs    []int
	M      map[string][2]string
	starts []string
}

func (s *Solution2) ProcessLine(i int, line string) {
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
	if n[2] == 'A' {
		s.starts = append(s.starts, n)
	}
}

// Find the number of steps that each start point needs to find its end.
// Then calculate the least common multiple of those steps.
// The lcm() determines at which step the end points overlap, ie. the number of
// steps it takes for all start points to simultaneously reach their end point.
func part2() {
	s := &Solution2{M: make(map[string][2]string)}
	file.ReadLines("./input", s)

	steps := []int{}
	var lr int
	for _, start := range s.starts {
		step := 0
		for curr := start; curr[2] != 'Z'; step++ {
			lr = s.lrs[step%len(s.lrs)]
			curr = s.M[curr][lr]
		}
		steps = append(steps, step)
	}
	fmt.Println(algo.LCMSlice(steps))
}

func main() {
	part1()
	part2()
}
