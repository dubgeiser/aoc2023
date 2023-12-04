package grids

import (
	"aoc2023/lib/file"
	"fmt"
	"strings"
)

var allDirections = [8]Pos{
	{0, -1}, {0, 1},
	{1, 0}, {1, -1}, {1, 1},
	{-1, 0}, {-1, -1}, {-1, 1},
}

var xyDirections = [4][2]int{}

type Pos struct {
	Row int
	Col int
}

func NewPosition(row, col int) Pos {
	pos := Pos{Row: row, Col: col}
	return pos
}

type Grid[T any] struct {
	items  [][]T
	Height int
	Width  int
}

func NewGrid[T any](height int, width int, v T) *Grid[T] {
	grid := &Grid[T]{Width: width, Height: height}
	grid.items = make([][]T, height)
	for iRow := range grid.items {
		grid.items[iRow] = make([]T, width)
		for iCol := range grid.items[iRow] {
			grid.Set(iRow, iCol, v)
		}
	}
	return grid
}

func (g *Grid[T]) String() string {
	return fmt.Sprint(g.items)
}

func (g *Grid[T]) GetAt(row int, col int) T {
	return g.items[row][col]
}

func (g *Grid[T]) Get(pos Pos) T {
	return g.GetAt(pos.Row, pos.Col)
}

func (g *Grid[T]) Set(row int, col int, v T) *Grid[T] {
	g.items[row][col] = v
	return g
}

func (g *Grid[T]) InBounds(row, col int) bool {
	return row > 0 && col > 0 && row < g.Height && col < g.Width
}

func (g *Grid[T]) InBoundsPosition(pos Pos) bool {
	return g.InBounds(pos.Row, pos.Col)
}

func (g *Grid[T]) AdjacentPositions(row, col int) []Pos {
	a := []Pos{}
	var check Pos
	for _, p := range allDirections {
		check = Pos{row + p.Row, col + p.Col}
		if g.InBoundsPosition(check) {
			a = append(a, check)
		}
	}
	return a
}

type stringGridBuilder struct {
	grid  [][]string
	width int
}

func (b *stringGridBuilder) ProcessLine(i int, line string) {
	b.width = len(line)
	b.grid = append(b.grid, strings.Split(line, ""))
}

func GridFromFile(fn string) *Grid[string] {
	g := &Grid[string]{}
	gb := &stringGridBuilder{}
	file.ReadLines(fn, gb)
	g.items = gb.grid
	g.Width = gb.width
	g.Height = len(gb.grid)
	return g
}
