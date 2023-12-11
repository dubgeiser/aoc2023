package main

import (
	"aoc2023/lib/collections"
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
// This is also the place where we can determine which pipe is at start.
func (s *Solution) IsConnectedTo(ch1, ch2 byte, dir int) bool {
	var isConnected bool
	isConnected = (ch1 == 'S' || slices.Contains(D2P[dir], ch1)) && slices.Contains(D2P[DO[dir]], ch2)
	if isConnected && ch1 == 'S' {
		s.possibleStarts = s.possibleStarts.Intersection(collections.NewSetFrom(D2P[dir]))
	}
	return isConnected
}

type Solution struct {
	M              [][]byte
	start          grids.Position
	Loop           []grids.Position
	possibleStarts *collections.Set[byte]
	chStart        byte
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
	if p.Row > 0 && s.IsConnectedTo(s.M[p.Row][p.Col], s.M[p.Row-1][p.Col], N) {
		a = append(a, grids.NewPosition(p.Row-1, p.Col))
	}
	// East
	if p.Col < len(s.M[p.Row])-1 && s.IsConnectedTo(s.M[p.Row][p.Col], s.M[p.Row][p.Col+1], E) {
		a = append(a, grids.NewPosition(p.Row, p.Col+1))
	}
	// South
	if p.Row < len(s.M)-1 && s.IsConnectedTo(s.M[p.Row][p.Col], s.M[p.Row+1][p.Col], S) {
		a = append(a, grids.NewPosition(p.Row+1, p.Col))
	}
	// West
	if p.Col > 0 && s.IsConnectedTo(s.M[p.Row][p.Col], s.M[p.Row][p.Col-1], W) {
		a = append(a, grids.NewPosition(p.Row, p.Col-1))
	}
	return a
}

// Apparently, this is called "Ray casting"
// Given a position, count the number of times it crosses the loop.
// If there's an odd number of ' | ', ' L ', ' J ' then it must be inside the loop
func (s *Solution) CountEdgeCrossings(p grids.Position) int {
	var borderChars = []byte{'|', 'L', 'J'}
	count := 0
	for c:= p.Col + 1; c <  len(s.M[p.Row]); c++ {
		ch := s.M[p.Row][c]
		if slices.Contains(borderChars, ch) {
			count++
		}
	}
	return count
}

// Based on BFS, Algo Book p. 556
// Not storing distances: We're dealing with a loop: farthest is len(loop) / 2
// The color system in the Algo book is replaced by actually building the path/loop
//
// Queue (fifo) is implemented with a slice
// Enqueue: q = append(q, p)
// Dequeue: p = q[0]; q = q[1:]
func (s *Solution) Solve1() int {
	q := []grids.Position{s.start}
	loop := []grids.Position{s.start}
	var curr grids.Position
	for len(q) > 0 {
		curr = q[0]
		q = q[1:]
		for _, next := range s.AdjacentPositions(curr) {
			if slices.Contains(loop, next) {
				continue
			}
			loop = append(loop, next)
			q = append(q, next)

		}
	}
	s.Loop = loop
	if s.possibleStarts.Size() > 1 {
		panic(s.possibleStarts)
	}
	s.M[s.start.Row][s.start.Col] = s.possibleStarts.Export()[0]
	return len(loop) / 2
}

func (s *Solution) Solve2() int {
	answer := 0
	// Cleanup the grid first, so we don't count too many edge crossings.
	for r, row := range s.M {
		for c := range row {
			if !slices.Contains(s.Loop, grids.NewPosition(r, c)) {
				s.M[r][c] = '.'
			}
		}
	}
	for r, row := range s.M {
		for c := range row {
			p := grids.NewPosition(r, c)
			if slices.Contains(s.Loop, p) {
				continue
			}
			crossCount := s.CountEdgeCrossings(p)
			if crossCount%2 == 1 {
				answer++
			}
		}
	}
	return answer
}

func main() {
	s := &Solution{possibleStarts: collections.NewSetFrom([]byte{'|', '-', 'L', 'J', '7', 'F'})}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.Solve1())
	fmt.Println(s.Solve2())
}
