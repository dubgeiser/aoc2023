package main

import (
	"aoc2023/lib/algo"
	"aoc2023/lib/file"
	"fmt"
	"strconv"
	"strings"
)

type Solution struct {
	lrs    []int
	M      map[string][2]string
	starts []string
}

func (s *Solution) ProcessLine(i int, line string) {
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

func (s *Solution) CountSteps(start string, endCondition func(s string) bool) int {
	step := 0
	for curr := start; !endCondition(curr); step++ {
		curr = s.M[curr][s.lrs[step%len(s.lrs)]]
	}
	return step
}

func (s *Solution) Solve1() int {
	return s.CountSteps("AAA", func(s string) bool { return s == "ZZZ" })
}

func (s *Solution) Solve2() int {
	// Find the number of steps that each start point needs to find its end.
	// Then calculate the least common multiple of those steps.
	// The lcm() determines at which step the end points overlap, ie. the number of
	// steps it takes for all start points to simultaneously reach their end point.
	steps := []int{}
	for _, start := range s.starts {
		steps = append(steps, s.CountSteps(start, func(s string) bool { return s[2] == 'Z' }))
	}
	return algo.LCMSlice(steps)
}

func main() {
	s := &Solution{M: make(map[string][2]string)}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.Solve1())
	fmt.Println(s.Solve2())
}
