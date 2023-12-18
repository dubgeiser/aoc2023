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
}

func isPerfectReflection(g [][]byte, r1, r2 int) bool {
	top := g[:r2]
	bot := g[r2:]
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

func calcLines(g [][]byte) int {
	for r := 0; r < len(g)-1; r++ {
		if string(g[r]) == string(g[r+1]) && isPerfectReflection(g, r, r+1) {
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
