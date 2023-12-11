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
	g := transpose(s.G)
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

func (s *Solution) CalcExpansion(p1, p2 grids.Position) int {
	exp := 0
	for _, er := range s.ExpRows {
		if min(p1.Row, p2.Row) <= er && er <= max(p1.Row, p2.Row) {
			exp++
		}
	}
	for _, ec := range s.ExpCols {
		if min(p1.Col, p2.Col) <= ec && ec <= max(p1.Col, p2.Col) {
			exp++
		}
	}
	return exp
}

func transpose(m [][]byte) [][]byte {
	t := make([][]byte, len(m))
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[0]); c++ {
			t[c] = append(t[c], m[r][c])
		}
	}
	return t
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

func (s *Solution) Solve1() any {
	s.FindEmptyCols()
	answer := 0
	count := 1
	for i := 0; i < len(s.Positions)-1; i++ {
		for j := i+1; j < len(s.Positions); j++ {
			p1 := s.Positions[i]
			p2 := s.Positions[j]
			d := ManhattanDistance(p1, p2) + s.CalcExpansion(p1, p2)
			count++
			answer += d
		}
	}
	return answer
}

func (s *Solution) Solve2() any {
	return 0
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
