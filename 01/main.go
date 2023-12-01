package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strconv"
)

type Solution1 struct {
	value int
}

type Solution2 struct {
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

func (lp *Solution2) ProcessLine(line string) {
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

func part2() int {
	s := &Solution2{}
	file.ReadLines("./input", s)
	return s.value
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
