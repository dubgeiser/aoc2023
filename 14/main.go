package main

import (
	"aoc2023/lib/file"
	"fmt"
)

type Solution struct {
	G [][]byte
}

func (s *Solution) ProcessLine(i int, line string) {
	s.G = append(s.G, []byte(line))
}

func totalLoad(g [][]byte) int {
	answer := 0
	L := len(g)
	for r, row := range g {
		for _, col := range row {
			if col == 'O' {
				answer += L - r
			}
		}
	}
	return answer
}

func (s *Solution) Solve1() int {
	g := copyGrid(s.G)
	for r := 1; r < len(g); r++ {
		for c, col := range g[r] {
			if col != 'O' {
				continue
			}
			for rr := r - 1; rr >= 0 && g[rr][c] == '.'; rr-- {
				g[rr+1][c] = '.'
				g[rr][c] = 'O'
			}
		}
	}
	return totalLoad(g)
}

func copyGrid(src [][]byte) [][]byte {
	g := make([][]byte, len(src))
	for r, row := range src {
		g[r] = make([]byte, len(src[0]))
		for c, col := range row {
			g[r][c] = col
		}
	}
	return g
}

func print(g [][]byte) {
	for _, row := range g {
		fmt.Println(string(row))
	}
}

func main() {
	s := &Solution{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.Solve1())
}
