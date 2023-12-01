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

func (lp *Solution1) ProcessLine(line string) {
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
var sdigits [9]string = [9]string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}

type Solution2 struct {
	value int
}

func getFirstIndex(i int, line string) int {
	wpos := strings.Index(line, wdigits[i])
	dpos := strings.Index(line, sdigits[i])
	if wpos == -1 && dpos == -1 {
		return -1
	}
	if wpos == -1 {
		return dpos
	}
	if dpos == -1 {
		return wpos
	}
	return min(wpos, dpos)
}

func getLastIndex(i int, line string) int {
	wpos := strings.LastIndex(line, wdigits[i])
	dpos := strings.LastIndex(line, sdigits[i])
	if wpos == -1 && dpos == -1 {
		return -1
	}
	if wpos == -1 {
		return dpos
	}
	if dpos == -1 {
		return wpos
	}
	return max(wpos, dpos)
}

func (lp *Solution2) ProcessLine(line string) {
	pLo := -1
	iLo := 0
	pHi := -1
	iHi := 0
	for i := 0; i < 9; i++ {
		lo := getFirstIndex(i, line)
		if lo == -1 {
			continue
		}
		if pLo == -1 {
			pLo = lo
			iLo = i
		}
		if lo < pLo {
			pLo = lo
			iLo = i
		}
		hi := getLastIndex(i, line)
		if pHi == -1 {
			pHi = hi
			iHi = i
		}
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
