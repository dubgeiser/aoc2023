package main

import (
	"aoc2023/lib/file"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const (
	highcard = iota
	pair
	pair2
	kind3
	fullhouse
	kind4
	kind5
)

type Solution1 struct {
	hand2bid map[string]int
	hands    []string
}

func HandValue(hand string) int {
	card2count := make(map[rune]int)
	for _, r := range hand {
		card2count[r]++
	}
	size := len(card2count)
	if size == 1 {
		return kind5
	} else if size == 2 {
		for _, count := range card2count {
			if count == 4 {
				return kind4
			}
		}
		return fullhouse
	} else if size == 3 {
		for _, count := range card2count {
			if count == 3 {
				return kind3
			}
		}
		return pair2
	} else if size == 4 {
		return pair
	}
	return highcard
}

var strength = []rune{'2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A'}
var strengthJoker = []rune{'J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A'}

func Strength(r rune) int {
	return slices.Index(strength, r)
}

func StrengthJoker(r rune) int {
	return slices.Index(strengthJoker, r)
}

func CmpHands(a, b string) bool {
	va := HandValue(a)
	vb := HandValue(b)
	if va == vb {
		for i := 0; i < len(a); i++ {
			fa := Strength(rune(a[i]))
			fb := Strength(rune(b[i]))
			if fa == fb {
				continue
			}
			return fa < fb
		}
	}
	return va < vb
}

func (s *Solution1) ProcessLine(i int, line string) {
	l := strings.Split(line, " ")
	hand := l[0]
	bid, _ := strconv.Atoi(l[1])
	s.hand2bid[hand] = bid
	s.hands = append(s.hands, hand)
}

func part1() {
	s := &Solution1{hand2bid: make(map[string]int)}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println()

	sort.Slice(s.hands, func(i, j int) bool {
		return CmpHands(s.hands[i], s.hands[j])
	})
	answer := 0
	for i, h := range s.hands {
		answer += s.hand2bid[h] * (i + 1)
	}
	fmt.Println(answer)
}

func HandValueJoker(hand string) int {
	card2count := make(map[rune]int)
	for _, r := range hand {
		card2count[r]++
	}
	rv := HandValue(hand)
	if card2count['J'] == 0 {
		return rv
	}
	// having any joker in these cases, you can always make it 5 of a kind
	if rv == kind5 || rv == kind4 || rv == fullhouse {
		return kind5
	}
	// we have either 1 or 3 jokers
	if rv == kind3 {
		return kind4
	}
	if rv == pair2 {
		if card2count['J'] == 1 {
			return fullhouse
		}
		return kind4
	}
	if rv == pair {
		return kind3
	}
	return pair
}

func CmpHandsJoker(a, b string) bool {
	va := HandValueJoker(a)
	vb := HandValueJoker(b)
	if va == vb {
		for i := 0; i < len(a); i++ {
			fa := StrengthJoker(rune(a[i]))
			fb := StrengthJoker(rune(b[i]))
			if fa == fb {
				continue
			}
			return fa < fb
		}
	}
	return va < vb
}

func part2() {
	s := &Solution1{hand2bid: make(map[string]int)}
	file.ReadLines("./input", s)
	sort.Slice(s.hands, func(i, j int) bool {
		return CmpHandsJoker(s.hands[i], s.hands[j])
	})
	answer := 0
	for i, h := range s.hands {
		answer += s.hand2bid[h] * (i + 1)
	}
	fmt.Println(answer)
}

func main() {
	part1()
	part2()
}
