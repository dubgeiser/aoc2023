package main

import (
	"aoc2023/lib/collections"
	"aoc2023/lib/file"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Solution struct {
	cards   []*Card
	answer1 int
	answer2 int
}

func (s *Solution) ProcessLine(i int, line string) {
	s.answer1 += int(math.Pow(2, float64(WinCount(line)-1)))
	winCount := WinCount(line)
	iProduced := []int{}
	for j := i + 1; j < i+1+winCount; j++ {
		iProduced = append(iProduced, j)
	}
	s.cards = append(s.cards, &Card{iProduced, 1})
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

type Card struct {
	iProduced []int
	numCopies int
}

func main() {
	s := &Solution{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.answer1)

	// We can be sure that the last card does not produce any further cards.
	// I'm still wondering if "copy" is the right name, because we count the
	// _real_, _original_ card also... which isn't a copy, right?
	// Damn you, Jelle!
	for i := len(s.cards) - 1; i > -1; i-- {
		for _, j := range s.cards[i].iProduced {
			s.cards[i].numCopies += s.cards[j].numCopies
		}
		s.answer2 += s.cards[i].numCopies
	}
	fmt.Println(s.answer2)
}
