package grids

import (
	"aoc2023/lib/file"
	"strings"
)

var allDirections = [8]Position{
	{0, -1}, {0, 1},
	{1, 0}, {1, -1}, {1, 1},
	{-1, 0}, {-1, -1}, {-1, 1},
}

type Position struct {
	Row int
	Col int
}

func NewPosition(row, col int) Position {
	return Position{Row: row, Col: col}
}

type Grid[T any] struct {
	items  [][]T
	Height int
	Width  int
}

// rows <-> cols
func Transpose[T any](m [][]T) [][]T {
	t := make([][]T, len(m))
	for r := 0; r < len(m); r++ {
		for c := 0; c < len(m[0]); c++ {
			t[c] = append(t[c], m[r][c])
		}
	}
	return t
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

func (g *Grid[T]) GetAt(row int, col int) T {
	return g.items[row][col]
}

func (g *Grid[T]) Get(p Position) T {
	return g.GetAt(p.Row, p.Col)
}

func (g *Grid[T]) Set(row int, col int, v T) *Grid[T] {
	g.items[row][col] = v
	return g
}

func (g *Grid[T]) InBounds(row, col int) bool {
	return row > 0 && col > 0 && row < g.Height && col < g.Width
}

func (g *Grid[T]) InBoundsPosition(p Position) bool {
	return g.InBounds(p.Row, p.Col)
}

func (g *Grid[T]) AdjacentPositions(row, col int) []Position {
	a := []Position{}
	var check Position
	for _, p := range allDirections {
		check = Position{row + p.Row, col + p.Col}
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
