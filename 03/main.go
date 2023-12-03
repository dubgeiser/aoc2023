package main

import (
	"aoc2023/lib/collections"
	"aoc2023/lib/grids"
	"fmt"
	"strconv"
)

var directions = [8][2]int{
	{0, -1}, {0, 1},
	{1, 0}, {1, -1}, {1, 1},
	{-1, 0}, {-1, -1}, {-1, 1},
}

func CheckAdjancency(grid *grids.Grid[string], row, col int) bool {
	for _, dir := range directions {
		checkPos := grids.NewPosition(row+dir[0], col+dir[1])
		if !grid.InBoundsPosition(checkPos) {
			continue
		}
		c := grid.Get(checkPos)
		if _, err := strconv.Atoi(c); err != nil {
			if c != "." {
				return true
			}
		}
	}
	return false
}

func FindAdjacentGearPositions(grid *grids.Grid[string], row, col int) *collections.Set[grids.Pos] {
	adjacentGears := collections.NewSet[grids.Pos]()
	for _, dir := range directions {
		checkPos := grids.NewPosition(row+dir[0], col+dir[1])
		if !grid.InBoundsPosition(checkPos) {
			continue
		}
		if grid.Get(checkPos) == "*" {
			adjacentGears.Add(checkPos)
		}
	}
	return adjacentGears
}

func part1() {
	g := grids.GridFromFile("./input")
	answer := 0
	for row := 0; row < g.Height; row++ {
		currN := 0
		hasAdjacentSymbol := false
		for col := 0; col < g.Width; col++ {
			c := g.GetAt(row, col)
			isDigit := false
			if n, err := strconv.Atoi(c); err == nil {
				isDigit = true
				currN *= 10
				currN += n
				hasAdjacentSymbol = hasAdjacentSymbol || CheckAdjancency(g, row, col)
			}
			if col == g.Width-1 || !isDigit {
				if hasAdjacentSymbol {
					answer += currN
				}
				hasAdjacentSymbol = false
				currN = 0
			}
		}
	}
	fmt.Println(answer)
}

func part2() {
	answer := 0
	g := grids.GridFromFile("./input")
	gears2AdjacentNumbers := make(map[grids.Pos][]int)
	adjacentGears := collections.NewSet[grids.Pos]()
	for row := 0; row < g.Height; row++ {
		currN := 0
		for col := 0; col < g.Width; col++ {
			c := g.GetAt(row, col)
			isDigit := false
			if n, err := strconv.Atoi(c); err == nil {
				isDigit = true
				currN *= 10
				currN += n
				adjacentGears.AddSet(FindAdjacentGearPositions(g, row, col))
			}
			if col == g.Width-1 || !isDigit {
				for gp := range adjacentGears.Items() {
					gears2AdjacentNumbers[gp] = append(gears2AdjacentNumbers[gp], currN)
				}
				adjacentGears.Clear()
				currN = 0
			}
		}
	}
	for _, nums := range gears2AdjacentNumbers {
		if len(nums) == 2 {
			answer += nums[0] * nums[1]
		}
	}
	fmt.Println(answer)
}

func main() {
	part1()
	part2()
}
