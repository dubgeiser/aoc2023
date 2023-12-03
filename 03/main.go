package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strconv"
	"strings"
)

type Pos struct {
	row, col int
}

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

func FindAdjacentGearPositions(grid [][]string, row, col int) map[Pos]bool {
	adjacentGears := make(map[Pos]bool)
	for _, dir := range directions {
		checkPos := Pos{row + dir[0], col + dir[1]}
		if !IsInGrid(grid, checkPos.row, checkPos.col) {
			continue
		}
		if grid[checkPos.row][checkPos.col] == "*" {
			adjacentGears[checkPos] = true
		}
	}
	return adjacentGears
}

func (s *Solution1) Solve2() {
	width := len(s.grid[0])
	gears2AdjacentNumbers := make(map[Pos][]int)
	for row := range s.grid {
		currN := 0
		adjacentGears := make(map[Pos]bool)
		for col := range s.grid[row] {
			c := s.grid[row][col]
			isDigit := false
			if n, err := strconv.Atoi(c); err == nil {
				isDigit = true
				currN *= 10
				currN += n
				for g := range FindAdjacentGearPositions(s.grid, row, col) {
					adjacentGears[g] = true
				}
			}
			if col == width-1 || !isDigit {
				for gp := range adjacentGears {
					gears2AdjacentNumbers[gp] = append(gears2AdjacentNumbers[gp], currN)
				}
				adjacentGears = nil
				adjacentGears = make(map[Pos]bool)
				currN = 0
			}
		}
	}
	for _, nums := range gears2AdjacentNumbers {
		if len(nums) == 2 {
			s.answer += nums[0] * nums[1]
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

func part2() {
	s := &Solution1{}
	file.ReadLines("./input", s)
	s.Solve2()
	fmt.Println(s.answer)
}

func main() {
	part1()
	part2()
}
