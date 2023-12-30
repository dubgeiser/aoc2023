package main

import (
	"aoc2023/lib/collections"
	"aoc2023/lib/file"
	"container/heap"
	"fmt"
	"strconv"
)

type Grid [][]int

type CityBlock struct {
	row, col int
	heatLoss int
	dirRow   int
	dirCol   int
	dirCount int
	index    int
}

// String func so we can use a CityBlock in a visited set without taking
// heat loss or index in de priority queue into account.
func (cb *CityBlock) String() string {
	return fmt.Sprintf("(%d,%d,%d,%d,%d)", cb.row, cb.col, cb.dirRow, cb.dirCol, cb.dirCount)
}

type PriorityQueue []*CityBlock

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heatLoss < pq[j].heatLoss
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	cb := x.(*CityBlock)
	cb.index = n
	*pq = append(*pq, cb)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	cb := old[n-1]
	old[n-1] = nil // memory!
	cb.index = -1  // just in case
	*pq = old[0 : n-1]
	return cb
}

type Solution struct {
	G Grid
}

func (s *Solution) ProcessLine(i int, line string) {
	s.G = append(s.G, toInt(line))
}

func toInt(line string) []int {
	l := []int{}
	for _, s := range line {
		n, _ := strconv.Atoi(string(s))
		l = append(l, n)
	}
	return l
}

func (g Grid) adjacents(cb *CityBlock) []*CityBlock {
	adj := []*CityBlock{}
	for _, dir := range [4][2]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}} {
		rr, cc := cb.row+dir[0], cb.col+dir[1]
		if rr < 0 || cc < 0 || rr >= len(g) || cc >= len(g[0]) {
			continue
		}
		if cb.dirRow == -dir[0] && cb.dirCol == -dir[1] {
			continue
		}
		dirCount := 1
		if cb.dirRow == dir[0] && cb.dirCol == dir[1] {
			dirCount = cb.dirCount + 1
		}
		if dirCount > 3 {
			continue
		}
		adj = append(adj, &CityBlock{
			row:      rr,
			col:      cc,
			dirCount: dirCount,
			heatLoss: cb.heatLoss + g[rr][cc],
			dirRow:   dir[0],
			dirCol:   dir[1],
		})
	}
	return adj
}

func (s *Solution) solve1() int {
	visited := collections.NewSet[string]()
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &CityBlock{row: 0, col: 0, heatLoss: 0, dirRow: 0, dirCol: 0, dirCount: 0})
	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*CityBlock)
		if curr.row == len(s.G)-1 && curr.col == len(s.G[0])-1 {
			return curr.heatLoss
		}
		if visited.Has(curr.String()) {
			continue
		}
		for _, adj := range s.G.adjacents(curr) {
			heap.Push(pq, adj)
		}
		visited.Add(curr.String())
	}
	return 0
}

func main() {
	s := &Solution{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.solve1())
}
