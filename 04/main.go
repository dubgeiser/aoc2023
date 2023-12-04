package main

import (
	"aoc2023/lib/file"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

type Solution1 struct {
	answer int
}

func parseNumbers(line string) ([]int, []int) {
	var winning, played []int
	line = strings.Split(line, ": ")[1]
	p := strings.Split(line, " | ")
	sw := strings.Split(p[0], " ")
	sp := strings.Split(p[1], " ")
	for _, sn := range sw {
		if n, err := strconv.Atoi(sn); err == nil {
			winning = append(winning, n)
		}
	}
	for _, sn := range sp {
		if n, err := strconv.Atoi(sn); err == nil {
			played = append(played, n)
		}
	}
	return winning, played
}

func (s *Solution1) ProcessLine(line string) {
	winning, played := parseNumbers(line)
	score := 0
	for _, np := range played {
		if slices.Contains(winning, np) {
			if score == 0 {
				score = 1
			} else {
				score *= 2
			}
		}
	}
	s.answer += score
}

func part1() {
	s := &Solution1{}
	file.ReadLines("./input", s)
	fmt.Println(s.answer)
}

func main() {
	part1()
}
