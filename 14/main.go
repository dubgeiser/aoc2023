package main

import (
	"aoc2023/lib/file"
	"fmt"
	"slices"
	"strings"
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
	g := slices.Clone(s.G)
	tiltNorth(g)
	return totalLoad(g)
}

func tiltNorth(g [][]byte) {
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
}

func tiltSouth(g [][]byte) {
	for r := len(g) - 1; r >= 0; r-- {
		for c, col := range g[r] {
			if col != 'O' {
				continue
			}
			for rr := r + 1; rr < len(g) && g[rr][c] == '.'; rr++ {
				g[rr-1][c] = '.'
				g[rr][c] = 'O'
			}
		}
	}
}

func tiltWest(g [][]byte) {
	for c := 1; c < len(g[0]); c++ {
		for r := range g {
			if g[r][c] != 'O' {
				continue
			}
			for cc := c - 1; cc >= 0 && g[r][cc] == '.'; cc-- {
				g[r][cc+1] = '.'
				g[r][cc] = 'O'
			}
		}
	}
}

func tiltEast(g [][]byte) {
	for c := len(g[0]) - 1; c >= 0; c-- {
		for r := range g {
			if g[r][c] != 'O' {
				continue
			}
			for cc := c + 1; cc < len(g[0]) && g[r][cc] == '.'; cc++ {
				g[r][cc-1] = '.'
				g[r][cc] = 'O'
			}
		}
	}
}

func cycle(g [][]byte) {
	tiltNorth(g)
	tiltWest(g)
	tiltSouth(g)
	tiltEast(g)
}

func asString(g [][]byte) string {
	s := ""
	for _, row := range g {
		s += fmt.Sprintf("%s\n", string(row))
	}
	s = strings.Trim(s, "\n")
	return s
}

func asGrid(s string) [][]byte {
	var g [][]byte
	lines := strings.Split(s, "\n")
	g = make([][]byte, len(lines))
	for r, l := range lines {
		g[r] = []byte(l)
	}
	return g
}

// When you run the cycle enought times, fi. 1000 and you calculate the
// totalLoad() each time, you see the same numbers re-appearing after a couple
// of times.
// This seems to indicate that there is a point at which a couple of grid
// configuration keeps returning.
// So we need to determine at what point the cycle begins exactly: I've been
// going through the totalLoad() results by hand, but that way it is not so
// obvious what the _exact_ cycle is and when it begins, since totalLoad() can
// be the same for different grid configurations.
// But, we can keep track of the different grid configurations... at some point
// the grid will already have been in the same configuration: so it seemst that
// this is the exact point of the cycle.
// So, we can determine this point by keeping every grid configuration we have
// seen in a slice.  Once we detect a configuration we've already seen, we know
// we're restarting the cycle. So the cycle starts after len(seen).
func (s *Solution) Solve2() int {
	g := slices.Clone(s.G)
	seen := []string{}
	var sg string

	for {
		cycle(g)
		sg = asString(g)
		if slices.Contains(seen, sg) {
			break
		}
		seen = append(seen, sg)
	}
	preRunCount := len(seen)

	// preRun = len(seen) is the number of iterations that have already passed
	// when we start at the cycle for the _second_ time.
	// g's state is now the state of the grid when the cycle starts.
	// So we do the whole thing again to determine the length of the cycle.
	// But we cycle at the end of the loop, because we're already _in_ the cycle!
	seen = []string{}
	loads := []int{}
	for {
		sg = asString(g)
		if slices.Contains(seen, sg) {
			break
		}
		loads = append(loads, totalLoad(g))
		seen = append(seen, sg)
		cycle(g)
	}

	// We now know the length of the cycle, the totalLoad at each point in the
	// cycle.
	// To get the load at x iterations ->
	// (x - preRun) % cycleLength
	// and then subtract 1 for the cycle that we already had in `seen`
	return loads[(1_000_000_000-preRunCount)%(len(loads))-1]
}

func print(g [][]byte) {
	for _, row := range g {
		fmt.Println(string(row))
	}
	fmt.Println()
}

func main() {
	s := &Solution{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.Solve1())
	fmt.Println(s.Solve2())
}
