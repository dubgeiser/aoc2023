package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strconv"
	"strings"
)

type Solution1 struct {
	times     []int
	distances []int
}

func (s *Solution1) ProcessLine(i int, line string) {
	sNrs := strings.Split(strings.Split(line, ":")[1], " ")
	for _, sn := range sNrs {
		if n, err := strconv.Atoi(sn); err == nil {
			if i == 0 {
				s.times = append(s.times, n)
			} else {
				s.distances = append(s.distances, n)
			}
		}
	}
}

func part1() {
	s := &Solution1{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read ", lineCount, " lines")

	answer := 0
	for ti, raceTime := range s.times {
		numWins := 0
		record := s.distances[ti]
		for startSpeed := 1; startSpeed < raceTime; startSpeed++ {
			if startSpeed*(raceTime-startSpeed) > record {
				numWins++
			}
		}
		if answer == 0 {
			answer = numWins
		} else if numWins > 0 {
			answer *= numWins
		}
	}
	fmt.Println(answer)
}

func main() {
	part1()
}
