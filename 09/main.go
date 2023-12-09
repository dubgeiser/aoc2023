package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strconv"
	"strings"
)

type Solution1 struct {
	answer int
}

func (s *Solution1) ProcessLine(i int, line string) {
	nums := []int{}
	for _, sNum := range strings.Split(line, " ") {
		if num, err := strconv.Atoi(sNum); err == nil {
			nums = append(nums, num)
		}
	}
	s.answer += nums[len(nums) - 1] + FindNext(nums)
}

func AllZeroes(nums []int) bool {
	for _, n := range nums {
		if n != 0 {
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

func FindNext(n []int) int {
	d := Distances(n)
	if AllZeroes(d) {
		return 0
	}
	return d[len(d)-1] + FindNext(d)
}

func part1() {
	s := &Solution1{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read ", lineCount, " lines")
	fmt.Println(s.answer)
}

func main() {
	part1()
}
