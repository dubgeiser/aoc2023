package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strconv"
	"strings"
)

var start = [3]int{12, 13, 14}

type Solution1 struct {
	value int
}

func isValidGrab(grab [3]int) bool {
	return grab[0] <= start[0] && grab[1] <= start[1] && grab[2] <= start[2]
}

var color2index = map[string]int{"red": 0, "green": 1, "blue": 2}

func parseGrab(sGrab string) [3]int {
	grab := [3]int{0}
	cc := strings.Split(sGrab, " ")
	count, _ := strconv.Atoi(cc[0])
	color := cc[1]
	grab[color2index[color]] = count
	return grab
}

func parseGrabs(sGrabs []string) [][3]int {
	var grabs [][3]int
	for _, g := range sGrabs {
		grabs = append(grabs, parseGrab(g))
	}
	return grabs
}

func parseSets(setInput string) [][3]int {
	var grabs [][3]int
	sets := strings.Split(setInput, "; ")
	for _, s := range sets {
		for _, g := range parseGrabs(strings.Split(s, ", ")) {
			grabs = append(grabs, g)
		}
	}
	return grabs
}

func parseGame(line string) (int, [][3]int) {
	var id int
	var grabs [][3]int
	fmt.Sscanf(line, "Game %d", &id)
	grabs = parseSets(strings.Replace(line, fmt.Sprintf("Game %d: ", id), "", 1))
	return id, grabs
}

func (lp *Solution1) ProcessLine(line string) {
	id, grabs := parseGame(line)
	isPossible := true
	for _, grab := range grabs {
		isPossible = isPossible && isValidGrab(grab)
	}
	if isPossible {
		lp.value += id
	}
}

func part1() int {
	s := &Solution1{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	return s.value
}

func part2() int {
	s := &Solution2{}
	file.ReadLines("./input", s)
	return s.answer
}

type Solution2 struct {
	answer int
}

func (lp *Solution2) ProcessLine(line string) {
	_, grabs := parseGame(line)
	mvg := findMinimumViableGrab(grabs)
	lp.answer += mvg[0] * mvg[1] * mvg[2]
}

func findMinimumViableGrab(grabs [][3]int) [3]int {
	mvg := [3]int{0}
	for _, g := range grabs {
		for i, cube := range g {
			mvg[i] = max(mvg[i], cube)
		}
	}
	return mvg
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
