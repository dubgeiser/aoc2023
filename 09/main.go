package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strconv"
	"strings"
)

const (
	previous = iota
	next
)

type Solution struct {
	answer1 int
	answer2 int
}

func (s *Solution) ProcessLine(i int, line string) {
	nums := []int{}
	for _, sNum := range strings.Split(line, " ") {
		if num, err := strconv.Atoi(sNum); err == nil {
			nums = append(nums, num)
		}
	}
	s.answer1 += nums[len(nums)-1] + Find(nums, next)
	s.answer2 += nums[0] - Find(nums, previous)
}

func AllSame(nums []int) bool {
	for _, n := range nums {
		if n != nums[0] {
			return false
		}
	}
	return true
}

func Distances(n []int) []int {
	var j int
	var d []int
	for i := 0; i < len(n)-1; i++ {
		j = i + 1
		d = append(d, n[j]-n[i])
	}
	return d
}

func Find(n []int, where int) int {
	d := Distances(n)
	if AllSame(d) {
		return d[0]
	}
	if where == next {
		return d[len(d)-1] + Find(d, next)
	}
	return d[0] - Find(d, previous)
}

func main() {
	s := &Solution{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.answer1)
	fmt.Println(s.answer2)
}
