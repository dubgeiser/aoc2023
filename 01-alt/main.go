package main

import (
	"aoc2023/lib/file"
	"bytes"
	"fmt"
	"regexp"
	"strings"
)

type Solution struct {
	answer1 int
	answer2 int
}

func WrittenDigit(c rune, line string) (int, bool) {
	for i, wd := range writtenDigits {
		if strings.HasPrefix(line, wd) {
			return i + 1, true
		}
	}
	return 0, false
}

var re = regexp.MustCompile(`[1-9]`)
var writtenDigits = [9]string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func (s *Solution) ProcessLine(i int, line string) {
	digits := re.FindAll([]byte(line), -1)
	s.answer1 += 10*int(bytes.Runes(digits[0])[0]-'0') + int(bytes.Runes(digits[len(digits)-1])[0]-'0')

	mixDigits := []int{}
	for i, c := range line {
		if c > '0' && c <= '9' {
			mixDigits = append(mixDigits, int(c-'0'))
		} else if d, ok := WrittenDigit(c, line[i:]); ok {
			mixDigits = append(mixDigits, d)
		}
	}
	s.answer2 += 10*mixDigits[0] + mixDigits[len(mixDigits)-1]
}

func main() {
	s := &Solution{}
	file.ReadLines("../01/input", s)
	fmt.Println(s.answer1)
	fmt.Println(s.answer2)
}
