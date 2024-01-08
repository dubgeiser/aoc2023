package main

import (
	"aoc2023/lib/file"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

var D map[string]Point = map[string]Point{
	"U": {0, -1},
	"D": {0, 1},
	"L": {-1, 0},
	"R": {1, 0},
}

type Solution struct {
	isPart2 bool
	path    []Point
	corners []Point
	lenPath int
}

func (s *Solution) ProcessLine(i int, line string) {
	if i == 0 {
		s.corners = []Point{{0, 0}}
	}
	parts := strings.Split(line, " ")
	dir := D[parts[0]]
	steps, _ := strconv.Atoi(parts[1])
	color := strings.Replace(strings.Replace(parts[2], "(#", "", 1), ")", "", 1)
	if s.isPart2 {
		ss, _ := strconv.ParseInt(color[:len(color)-1], 16, 64)
		steps = int(ss)
		dir = D[map[byte]string{
			'0': "R",
			'1': "D",
			'2': "L",
			'3': "U",
		}[color[len(color)-1]]]
	}
	s.lenPath += steps
	p := s.corners[i]
	s.corners = append(s.corners, Point{p.x + steps*dir.x, p.y + steps*dir.y})
}

// Shoelace via formula:
// https://en.wikipedia.org/wiki/Shoelace_formula#Other_formulas
// The one in solve1() had the wrong answer for part2-input (1 short!)
func (s *Solution) solve2() int {
	area := 0
	var p, pp, pn Point
	for i := 0; i < len(s.corners); i++ {
		p = s.corners[i]
		if i == 0 {
			pp = s.corners[len(s.corners)-1]
		} else {
			pp = s.corners[i-1]
		}
		if i == len(s.corners)-1 {
			pn = s.corners[0]
		} else {
			pn = s.corners[i+1]
		}
		area += p.x * (pp.y - pn.y)
	}
	area = int(0.5*math.Abs(float64(area)))
	// Make sure we have the _complete_ inner area, see "Pick's theorem"
	pick := area - s.lenPath/2 + 1
	// and add the path around it.
	return pick + s.lenPath
}

// "Area of a polygon from coordinates"
// https://www.themathdoctors.org/polygon-coordinates-and-areas/
// Shoelace Formula + Pick's theorem
func (s *Solution) solve1() int {
	s1, s2 := 0, 0
	corners := append(s.corners, s.corners[0])
	for i := 0; i < len(corners)-1; i++ {
		s1 += corners[i].x * corners[i+1].y
		s2 += corners[i].y * corners[i+1].x
	}
	// sample:   952408144115 (ok)
	// input: 122103860427464 (too low)
	// Unfortunately; this came 1 short in part2, which, in the end made me
	// code the shoelace via https://en.wikipedia.org/wiki/Shoelace_formula#Other_formulas
	return int(0.5*math.Abs(float64(s1)-float64(s2))) + s.lenPath/2 + 1
}

func main() {
	s := &Solution{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.solve2())

	s2 := &Solution{}
	s2.isPart2 = true
	file.ReadLines("./input", s2)
	fmt.Println(s2.solve2())
}
