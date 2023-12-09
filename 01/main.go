//
// This solution is for history purposes.
// In `01-alt/` is a _far_ more preferable one.
//

package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strconv"
	"strings"
)

type Solution1 struct {
	value int
}

func (lp *Solution1) ProcessLine(lineIndex int, line string) {
	snumber := ""
	for i := 0; i < len(line); i++ {
		if n, err := strconv.Atoi(string(line[i])); err == nil {
			snumber = fmt.Sprint(n)
			break
		}
	}
	for i := len(line) - 1; i >= 0; i-- {
		if n, err := strconv.Atoi(string(line[i])); err == nil {
			snumber += fmt.Sprint(n)
			break
		}
	}
	if len(snumber) > 0 {
		n, _ := strconv.Atoi(snumber)
		lp.value += n
	}
}

func part1() int {
	s := &Solution1{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	return s.value
}

var wdigits [9]string = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

type Solution2 struct {
	value int
}

func Index(i int, line string, searcher func(haystack, needle string) int, comparator func(int, int) int) int {
	wpos := searcher(line, wdigits[i])
	dpos := searcher(line, fmt.Sprint(i+1))
	if wpos == -1 && dpos == -1 {
		return -1
	}
	if wpos == -1 {
		return dpos
	}
	if dpos == -1 {
		return wpos
	}
	return comparator(wpos, dpos)
}

func Min(a, b int) int {
	return min(a, b)
}

func Max(a, b int) int {
	return max(a, b)
}

func (lp *Solution2) ProcessLine(lineIndex int, line string) {
	pLo := len(line)
	iLo := 0
	pHi := -1
	iHi := 0
	for i := 0; i < 9; i++ {
		lo := Index(i, line, strings.Index, Min)
		if lo == -1 {
			continue
		}
		if lo < pLo {
			pLo = lo
			iLo = i
		}
		hi := Index(i, line, strings.LastIndex, Max)
		if hi > pHi {
			pHi = hi
			iHi = i
		}
	}
	n, _ := strconv.Atoi(fmt.Sprint(iLo+1) + fmt.Sprint(iHi+1))
	lp.value += n
}

func part2() int {
	s := &Solution2{}
	file.ReadLines("./input", s)
	return s.value
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
