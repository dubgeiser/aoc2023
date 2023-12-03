package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strconv"
	"strings"
)

type Solution1 struct {
	grid   [][]string
	answer int
}

func (s *Solution1) ProcessLine(line string) {
	row := strings.Split(line, "")
	s.grid = append(s.grid, row)
}

var directions = [8][2]int{
	{0, -1}, {0, 1},
	{1, 0}, {1, -1}, {1, 1},
	{-1, 0}, {-1, -1}, {-1, 1},
}

func IsInGrid(grid [][]string, row, col int) bool {
	return row > 0 && col > 0 && row < len(grid) && col < len(grid[0])
}

func CheckAdjancency(grid [][]string, row, col int) bool {
	var checkRow, checkCol int
	for _, dir := range directions {
		checkRow = row + dir[0]
		checkCol = col + dir[1]
		if !IsInGrid(grid, checkRow, checkCol) {
			continue
		}
		c := grid[checkRow][checkCol]
		if _, err := strconv.Atoi(c); err != nil {
			if c != "." {
				return true
			}
		}
	}
	return false
}

func (s *Solution1) Solve() {
	width := len(s.grid[0])
	for row := range s.grid {
		currN := 0
		hasAdjacentSymbol := false
		for col := range s.grid[row] {
			c := s.grid[row][col]
			isDigit := false
			if n, err := strconv.Atoi(c); err == nil {
				isDigit = true
				currN *= 10
				currN += n
				hasAdjacentSymbol = hasAdjacentSymbol || CheckAdjancency(s.grid, row, col)
			}
			if col == width-1 || !isDigit {
				if hasAdjacentSymbol {
					s.answer += currN
				}
				hasAdjacentSymbol = false
				currN = 0
			}
		}
	}
}

func part1() {
	s := &Solution1{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	s.Solve()
	fmt.Println(s.answer)
}

func main() {
	part1()
}
