package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strings"
)

type Solution1 struct {
	value int
}

func (s *Solution1) ProcessLine(line string) {
	digits := []int{}
	for _, c := range line {
		if c > '0' && c <= '9' {
			digits = append(digits, int(c-'0'))
		}
	}
	s.value += 10*digits[0] + digits[len(digits)-1]
}

type Solution2 struct {
	value int
}

func (s *Solution2) ProcessLine(line string) {
	digits := []int{}
	for i, c := range line {
		if c > '0' && c <= '9' {
			digits = append(digits, int(c-'0'))
		} else {
			if d, ok := WrittenDigit(c, line[i:]); ok {
				digits = append(digits, d)
			}
		}
	}
	s.value += 10*digits[0] + digits[len(digits)-1]
}

var writtenDigits = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func WrittenDigit(c rune, line string) (int, bool) {
	for i, wd := range writtenDigits {
		if strings.HasPrefix(line, wd) {
			return i + 1, true
		}
	}
	return 0, false
}

func part1() int {
	s := &Solution1{}
	file.ReadLines("../01/input", s)
	return s.value
}

func part2() int {
	s := &Solution2{}
	file.ReadLines("../01/input", s)
	return s.value
}

func main() {
	fmt.Println(part1())
	fmt.Println(part2())
}
