package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strconv"
	"strings"
)

type Solution1 struct {
	lrs []int
	M   map[string][2]string
}

func (s *Solution1) ProcessLine(i int, line string) {
	if i == 0 {
		l := strings.ReplaceAll(line, "L", "0")
		l = strings.ReplaceAll(l, "R", "1")
		for _, r := range l {
			if n, err := strconv.Atoi(string(r)); err == nil {
				s.lrs = append(s.lrs, n)
			}
		}
		return
	}
	if line == "" { return }
	a := strings.Split(line, " = ")
	key := a[0]
	lr := strings.Split(a[1], ", ")
	l := strings.TrimLeft(lr[0], "(")
	r := strings.TrimRight(lr[1], ")")
	s.M[key] = [2]string{l, r}
}

func part1() {
	s := &Solution1{M: make(map[string][2]string)}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println()

	steps := 0
	curr := "AAA"
	var lr int
	var next string
	for i := 0; i <= len(s.lrs); i++ {
		if i == len(s.lrs) {
			i = 0
		}
		lr = s.lrs[i]
		next = s.M[curr][lr]
		steps++
		if next == "ZZZ" {
			break
		}
		curr = next
	}
	fmt.Println(steps)
}

func main() {
	part1()
}
