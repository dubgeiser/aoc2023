package main

import (
	"aoc2023/lib/file"
	"aoc2023/lib/grids"
	"fmt"
)

type Solution struct {
	G         [][]byte
	ExpRows   []int
	ExpCols   []int
	Positions []grids.Position
}

func (s *Solution) ProcessLine(i int, line string) {
	var empty = true
	s.G = append(s.G, []byte(line))
	for c, ch := range line {
		if ch == '#' {
			empty = false
			s.Positions = append(s.Positions, grids.NewPosition(i, c))
		}
	}
	if empty {
		s.ExpRows = append(s.ExpRows, i)
	}
}

func (s *Solution) FindEmptyCols() {
	g := grids.Transpose(s.G)
	var empty bool
	for c, col := range g {
		empty = true
		for _, ch := range col {
			if ch == '#' {
				empty = false
			}
		}
		if empty {
			s.ExpCols = append(s.ExpCols, c)
		}
	}
}

func (s *Solution) CalcExpansion(p1, p2 grids.Position, factor int) int {
	exp := 0
	for _, er := range s.ExpRows {
		if min(p1.Row, p2.Row) <= er && er <= max(p1.Row, p2.Row) {
			exp += int(factor - 1)
		}
	}
	for _, ec := range s.ExpCols {
		if min(p1.Col, p2.Col) <= ec && ec <= max(p1.Col, p2.Col) {
			exp += int(factor - 1)
		}
	}
	return exp
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func ManhattanDistance(p1, p2 grids.Position) int {
	return Abs(p1.Row-p2.Row) + Abs(p1.Col-p2.Col)
}

func (s *Solution) Solve(factor int) any {
	answer := 0
	count := 1
	for i := 0; i < len(s.Positions)-1; i++ {
		for j := i + 1; j < len(s.Positions); j++ {
			p1 := s.Positions[i]
			p2 := s.Positions[j]
			d := ManhattanDistance(p1, p2) + s.CalcExpansion(p1, p2, factor)
			count++
			answer += d
		}
	}
	return answer
}

func (s *Solution) Solve1() any {
	s.FindEmptyCols()
	return s.Solve(2)
}

func (s *Solution) Solve2() any {
	return s.Solve(1000000)
}

func main() {
	s := &Solution{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.Solve1())
	fmt.Println(s.Solve2())
}
