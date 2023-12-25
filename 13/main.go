package main

import (
	"aoc2023/lib/file"
	"aoc2023/lib/grids"
	"fmt"
	"slices"
)

func printGrid(g [][]byte) {
	for _, row := range g {
		fmt.Println(string(row))
	}
}

type Solution struct {
	G       [][]byte
	answer1 int
	answer2 int
}

func (s *Solution) ProcessLine(i int, line string) {
	if line == "" {
		s.CheckGrid()
		s.G = [][]byte{}
		return
	}
	s.G = append(s.G, []byte(line))
}

func (s *Solution) CheckGrid() {
	s.answer1 += calcLines(s.G) * 100
	s.answer1 += calcLines(grids.Transpose[byte](s.G))
	s.answer2 += calcAlmostPerfectReflection(s.G, 100)
	s.answer2 += calcAlmostPerfectReflection(grids.Transpose(s.G), 1)
}

func isPerfectReflection(g [][]byte, r int) bool {
	top := slices.Clone(g[:r])
	bot := slices.Clone(g[r:])
	var g1, g2 [][]byte = top, bot
	if len(top) > len(bot) {
		g1, g2 = bot, top
		slices.Reverse(g1)
		slices.Reverse(g2)
	}
	j := 0
	for i := len(g1) - 1; i >= 0; i-- {
		if string(g1[i]) != string(g2[j]) {
			return false
		}
		j++
	}
	return true
}

func calcAlmostPerfectReflection(g [][]byte, factor int) int {
	var diffCount int
	count := 0
	for i := 0; i < len(g)-1; i++ {
		diffCount = 0
		for j := 0; j < len(g); j++ {
			above := i - j
			below := i + 1 + j
			if above < below && above >= 0 && below < len(g) {
				for c := 0; c < len(g[0]); c++ {
					if g[above][c] != g[below][c] {
						diffCount++
					}
				}
			}
		}
		if diffCount == 1 {
			count += (i + 1) * factor
		}
	}
	return count
}

func calcLines(g [][]byte) int {
	for r := 0; r < len(g)-1; r++ {
		if string(g[r]) == string(g[r+1]) && isPerfectReflection(g, r+1) {
			return r + 1
		}
	}
	return 0
}

func main() {
	s := &Solution{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	// Still need to process the last grid.
	s.CheckGrid()
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.answer1)
	fmt.Println(s.answer2)
}
