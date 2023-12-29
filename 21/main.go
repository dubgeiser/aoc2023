package main

import (
	"aoc2023/lib/collections"
	"aoc2023/lib/file"
	"aoc2023/lib/grids"
	"fmt"
	"slices"
)

// local Position is weighted...
type Pos struct {
	Row    int
	Col    int
	Weight int
}

type Solution struct {
	G     [][]byte
	start grids.Position
}

func (s *Solution) ProcessLine(i int, line string) {
	row := []byte(line)
	s.G = append(s.G, row)
	j := slices.Index(row, 'S')
	if j >= 0 {
		s.start = grids.NewPosition(i, j)
	}
}

func (s *Solution) adjacentPositions(p grids.Position) []grids.Position {
	adj := []grids.Position{}
	dirs := []grids.Position{
		grids.NewPosition(1, 0),
		grids.NewPosition(0, 1),
		grids.NewPosition(-1, 0),
		grids.NewPosition(0, -1),
	}
	var r, c int
	for _, d := range dirs {
		r = p.Row + d.Row
		c = p.Col + d.Col
		if r >= 0 && r < len(s.G) && c >= 0 && c < len(s.G[0]) && s.G[r][c] != '#' {
			adj = append(adj, grids.NewPosition(r, c))
		}
	}
	return adj
}

func (s *Solution) solve1() int {
	answer := collections.NewSet[grids.Position]()
	q := []Pos{Pos{s.start.Row, s.start.Col, 0}}
	visited := collections.NewSet[grids.Position]()
	var curr grids.Position
	var currW Pos
	for len(q) > 0 {
		currW = q[0]
		q = q[1:]
		curr = grids.NewPosition(currW.Row, currW.Col)
		// %2 -> Elf could go back and forth between 2 tiles, but since exactly
		// 64 steps have to be taken, only the even tile may be counted!
		if currW.Weight <= 64 && currW.Weight%2 == 0 {
			answer.Add(curr)
		}
		for _, next := range s.adjacentPositions(curr) {
			if visited.Has(next) {
				continue
			}
			visited.Add(next)
			q = append(q, Pos{next.Row, next.Col, currW.Weight + 1})
		}
	}

	return answer.Len()
}

func main() {
	s := &Solution{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.solve1())
}
