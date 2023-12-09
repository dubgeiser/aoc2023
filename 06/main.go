package main

import (
	"aoc2023/lib/file"
	"fmt"
	"math"
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
	fmt.Println("Read", lineCount, "lines")

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

type Solution2 struct {
	time     int
	distance int
}

func (s *Solution2) ProcessLine(i int, line string) {
	if n, err := strconv.Atoi(strings.ReplaceAll(strings.Split(line, ":")[1], " ", "")); err == nil {
		if i == 0 {
			s.time = n
		} else {
			s.distance = n
		}
	}
}

func part2() {
	s := &Solution2{}
	file.ReadLines("./input", s)
	answer := 0
	record := s.distance
	raceTime := s.time
	for startSpeed := 1; startSpeed < raceTime; startSpeed++ {
		if startSpeed*(raceTime-startSpeed) > record {
			answer++
		}
	}
	fmt.Println(answer)
}

//
// Given:
//   - raceTime:   The time the race takes
//   - startSpeed: The speed we're gonna start the race with (this is the same
//     as the button press time)
// We're looking for cases where startSpeed*(raceTime-startSpeed) > record
// <=> startSpeed*(raceTime-startSpeed) - record > 0
// <=> startSpeed*raceTime - startSpeed^2 - record > 0
// <=> -startSpeed^2 + raceTime*startSpeed - record > 0
//
// Which can be written as quadratic function:
//	ax^2 + bx + c = 0
// where x = startSpeed, a = -1, b = raceTime and c = record
//
// If we solve for record+1 (because we want to beat the record)
// we'll get 2 numbers representing an interval; the integers in that interval
// are the starting speeds (or button hold times) that'll beat the given record.
// substracting the smallest from the lowest yields the number of wins.
// The SolveQuadratic solution runs 10-20 times faster
//
// <3 Jelle and Yasmin for reminding me of Quadratic function and Discriminant!
//
func SolveQuadratic(raceTime, record int) int {
	a := -1.0
	b := float64(raceTime)
	c := float64(-(record + 1))
	D := math.Pow(b, 2) - 4*a*c
	x1 := int(math.Floor((-1*b - math.Sqrt(D)) / (2 * a)))
	x2 := int(math.Ceil((-1*b + math.Sqrt(D)) / (2 * a)))
	return x1 - x2 + 1
}

func part1a() {
	s := &Solution1{}
	file.ReadLines("./input", s)
	answer := 1
	for ti, raceTime := range s.times {
		record := s.distances[ti]
		answer *= SolveQuadratic(raceTime, record)
	}
	fmt.Println(answer)
}

func part2a() {
	s := &Solution2{}
	file.ReadLines("./input", s)
	fmt.Println(SolveQuadratic(s.time, s.distance))
}

func main() {
	part1()
	part2()
	part1a()
	part2a()
}
