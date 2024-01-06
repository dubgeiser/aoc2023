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

func move(p Point, dir string, steps int) []Point {
	d := map[string]Point{
		"U": {0, -1},
		"D": {0, 1},
		"L": {-1, 0},
		"R": {1, 0},
	}
	path := []Point{}
	for i := 1; i <= steps; i++ {
		path = append(path, Point{p.x + i*d[dir].x, p.y + i*d[dir].y})
	}
	return path
}

type Solution struct {
	current    Point
	currentDir string
	path       []Point
	corners    []Point
}

func isTurning(dir1, dir2 string) bool {
	return dir1 != dir2
}

func (s *Solution) ProcessLine(i int, line string) {
	parts := strings.Split(line, " ")
	dir := parts[0]
	steps, _ := strconv.Atoi(parts[1])
	// Part 1 doesn't need color.
	// color := strings.Replace(strings.Replace(parts[2], "(#", "", 1), ")", "", 1)
	s.path = append(s.path, move(s.current, dir, steps)...)

	// The start is always a corner
	// If not in start we should change direction in order to count a corner.
	if i == 0 || isTurning(s.currentDir, dir) {
		s.corners = append(s.corners, s.current)
	}
	s.current = s.path[len(s.path)-1]
	s.currentDir = dir
}

// "Area of a polygon from coordinates"
// https://www.themathdoctors.org/polygon-coordinates-and-areas/
// Shoelace Formula
func (s *Solution) solve1() int {
	area, s1, s2 := 0, 0, 0
	// "Note that this is the first row, repeated."
	// Don't mutate the original corners, we still need to solve part 2.
	corners := append(s.corners, s.corners[0])
	path := append(s.path, s.path[0])
	for i := 0; i < len(corners)-1; i++ {
		// "You have to multiply each number of the first column with the number
		// in the second column of the next row.  Sum the products -> s1."
		s1 += corners[i].x * corners[i+1].y
		// "Then multiply each number of the second column with the number in
		// the first column of the next row. Sum the products -> s2."
		s2 += corners[i].y * corners[i+1].x
	}
	// Shoelace
	area = int(0.5 * math.Abs(float64(s1)-float64(s2)))
	// Need to add the path to the area
	// But we need to take the corners into account.
	// I don't know the exact relation: but it seems that len(path)/2 would be
	// in the right direction. (we're dealing with squares, but only half of 'm
	// will be used for the area).
	// I just added the one, I figure it "closes" the loop. :shrug:
	area += len(path)/2 + 1
	return area
}

func main() {
	s := &Solution{current: Point{0, 0}}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.solve1())
}
