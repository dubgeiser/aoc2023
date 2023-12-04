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

type Solution2 struct {
	cards []*Card
}

type Card struct {
	winning   []int
	played    []int
	iProduced []int
	numCopies int
}

func (s *Solution2) ProcessLine(line string) {
	winning, played := parseNumbers(line)
	s.cards = append(s.cards, &Card{winning, played, []int{}, 1})
}

func part2() {
	s := &Solution2{}
	file.ReadLines("./input", s)
	for i, c := range s.cards {
		numWins := 0
		for _, np := range c.played {
			if slices.Contains(c.winning, np) {
				numWins++
			}
		}
		for j := i + 1; j < i+1+numWins; j++ {
			c.iProduced = append(c.iProduced, j)
		}
	}

	// We can be sure that the last card does not produce any further cards.
	// I'm still wondering if "copy" is the right name, because we count the
	// _real_, _original_ card also... which isn't a copy, right?
	// Damn you, Jelle!
	for i := len(s.cards) - 1; i > -1; i-- {
		for _, j := range s.cards[i].iProduced {
			s.cards[i].numCopies += s.cards[j].numCopies
		}
	}

	answer := 0
	for _, c := range s.cards {
		answer += c.numCopies
	}
	fmt.Println(answer)
}

func main() {
	part1()
	part2()
}
