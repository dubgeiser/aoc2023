package main

import (
	"aoc2023/lib/file"
	"aoc2023/lib/grids"
	"fmt"
	"slices"
)

const (
	N = iota
	E
	S
	W
)

// Pipes that can go North, South, ...
var D2P = map[int][]byte{
	N: {'|', 'L', 'J'},
	S: {'|', 'F', '7'},
	E: {'-', 'F', 'L'},
	W: {'-', '7', 'J'},
}

var DO = map[int]int{N: S, S: N, E: W, W: E}

// Is c1 connected to the dir-th with c2?
func IsConnectedTo(c1, c2 byte, dir int) bool {
	return (c1 == 'S' || slices.Contains(D2P[dir], c1)) && slices.Contains(D2P[DO[dir]], c2)
}

type Solution struct {
	M     [][]byte
	start grids.Position
}

func (s *Solution) ProcessLine(i int, line string) {
	row := []byte(line)
	s.M = append(s.M, row)
	j := slices.Index(row, 'S')
	if j >= 0 {
		s.start = grids.NewPosition(i, j)
	}
}

func (s *Solution) AdjacentPositions(p grids.Position) []grids.Position {
	a := []grids.Position{}
	// North
	if p.Row > 0 && IsConnectedTo(s.M[p.Row][p.Col], s.M[p.Row-1][p.Col], N) {
		a = append(a, grids.NewPosition(p.Row-1, p.Col))
	}
	// East
	if p.Col < len(s.M[p.Row])-1 && IsConnectedTo(s.M[p.Row][p.Col], s.M[p.Row][p.Col+1], E) {
		a = append(a, grids.NewPosition(p.Row, p.Col+1))
	}
	// South
	if p.Row < len(s.M)-1 && IsConnectedTo(s.M[p.Row][p.Col], s.M[p.Row+1][p.Col], S) {
		a = append(a, grids.NewPosition(p.Row+1, p.Col))
	}
	// West
	if p.Col > 0 && IsConnectedTo(s.M[p.Row][p.Col], s.M[p.Row][p.Col-1], W) {
		a = append(a, grids.NewPosition(p.Row, p.Col-1))
	}
	return a
}

// Based on BFS, Algo Book p. 556
// Not storing distances: We're dealing with a loop: farthest is len(loop) / 2
// The color system in the Algo book is replaced by actually building the path.
// (cf. "seen" set)
// Queue (fifo) is implemented with a slice
// Enqueue: q = append(q, p)
// Dequeue: p = q[0]; q = q[1:]
func (s *Solution) Solve1() int {
	q := []grids.Position{s.start}
	path := []grids.Position{s.start}
	var curr grids.Position
	for len(q) > 0 {
		curr = q[0]
		q = q[1:]
		for _, next := range s.AdjacentPositions(curr) {
			if !slices.Contains(path, next) {
				path = append(path, next)
				q = append(q, next)
			}
		}
	}
	return len(path) / 2
}

func (s *Solution) Solve2() int {
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
