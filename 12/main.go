package main

import (
	"aoc2023/lib/file"
	"fmt"
	"math"
	"regexp"
	"strings"
)

var regexD = regexp.MustCompile(`\.+`)
var regexQ = regexp.MustCompile(`\?+`)

type Solution struct {
	answer1 int
}

func (s *Solution) ProcessLine(i int, line string) {
	parts := strings.Split(line, " ")
	springs := parts[0]
	sGroups := strings.Split(parts[1], ",")
	var groups []int
	var d int
	for _, sc := range sGroups {
		fmt.Sscanf(sc, "%d", &d)
		groups = append(groups, d)
	}
	s.answer1 += CountValidConfigurations(springs, groups)
}

func IsValidConfiguration(springs string, groups []int) bool {
	confs := regexD.Split(strings.Trim(springs, "."), -1)
	if len(confs) != len(groups) {
		return false
	}
	for i := 0; i < len(confs); i++ {
		if len(confs[i]) != groups[i] {
			return false
		}
	}
	return true
}

var P []byte = []byte{'.', '#'}

func isZero(number []int) bool {
	for _, n := range number {
		if n > 0 {
			return false
		}
	}
	return true
}

// Decrement given number, following the overflow defined by m,
// which is the max index each corresponding number may be.
// Note that this will not count subzero
func decrement(number []int, m []int) []int {
	for i := 0; i < len(number); i++ {
		if number[i] > 0 {
			number[i]--
			return number
		}
		number[i] = m[i]
	}
	return number
}

func CountValidConfigurations(springs string, groups []int) int {
	unknowns := regexQ.FindAllString(springs, -1)
	replaces := make([][]string, len(unknowns))
	total := 0
	// Keep track of the indices of the replaces, so that we can iterate
	// through all of them...
	rIndices := []int{}
	for i, u := range unknowns {
		replaces[i] = Permutations(P, len(u))
		rIndices = append(rIndices, len(replaces[i])-1)
	}

	var check string
	rMax := make([]int, len(rIndices))
	copy(rMax, rIndices)

	// Get around the fact that {0 0 0} must be processed too
	rIndices[0]++
	for !isZero(rIndices) {
		rIndices = decrement(rIndices, rMax)
		check = springs
		for i, unknown := range unknowns {
			check = strings.Replace(check, unknown, replaces[i][rIndices[i]], 1)
		}
		if IsValidConfiguration(check, groups) {
			total++
		}
	}
	return total
}

// Convert a given integer to an arbitrary 'number' system.
// Pool: all the characters of the number system
// n: number of positions that we're calculating in
func BaseConvert(number int, pool []byte, n int) []byte {
	var converted []byte
	l := len(pool)
	for i := 0; i < n; i++ {
		converted = append(converted, pool[number%l])
		number = number / l
	}
	return converted
}

// Note that this cache always assumes the same pool.
// This reduces make()'ing maps on the fly.
var permCache = make(map[int][]string)

// Take n number of pool and generate all permutations.
func Permutations(pool []byte, n int) []string {
	if v, ok := permCache[n]; ok {
		return v
	}
	var perms []string
	base := float64(len(pool))
	exp := float64(n)
	for i := 0; i < int(math.Pow(base, exp)); i++ {
		perms = append(perms, string(BaseConvert(i, pool, n)))
	}
	permCache[n] = perms
	return perms
}

func main() {
	s := &Solution{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.answer1)
}
