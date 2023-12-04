package main

import (
	"aoc2023/lib/collections"
	"aoc2023/lib/file"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Solution1 struct {
	answer int
}

func WinCount(line string) int {
	winning := collections.NewSet[int]()
	played := collections.NewSet[int]()
	line = strings.Split(line, ": ")[1]
	p := strings.Split(line, " | ")
	sw := strings.Split(p[0], " ")
	sp := strings.Split(p[1], " ")
	for _, sn := range sw {
		if n, err := strconv.Atoi(sn); err == nil {
			winning.Add(n)
		}
	}
	for _, sn := range sp {
		if n, err := strconv.Atoi(sn); err == nil {
			played.Add(n)
		}
	}
	return len(winning.Intersection(played).Items())
}

func (s *Solution1) ProcessLine(lineIndex int, line string) {
	s.answer += int(math.Pow(2, float64(WinCount(line)-1)))
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
	iProduced []int
	numCopies int
}

func (s *Solution2) ProcessLine(i int, line string) {
	winCount := WinCount(line)
	iProduced := []int{}
	for j := i + 1; j < i+1+winCount; j++ {
		iProduced = append(iProduced, j)
	}
	s.cards = append(s.cards, &Card{iProduced, 1})
}

func part2() {
	s := &Solution2{}
	file.ReadLines("./input", s)

	// We can be sure that the last card does not produce any further cards.
	// I'm still wondering if "copy" is the right name, because we count the
	// _real_, _original_ card also... which isn't a copy, right?
	// Damn you, Jelle!
	answer := 0
	for i := len(s.cards) - 1; i > -1; i-- {
		for _, j := range s.cards[i].iProduced {
			s.cards[i].numCopies += s.cards[j].numCopies
		}
		answer += s.cards[i].numCopies
	}
	fmt.Println(answer)
}

func main() {
	part1()
	part2()
}
