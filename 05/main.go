package main

import (
	"aoc2023/lib/file"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Solution struct {
	currMaps MapCollection
	seeds    []int
	maps     []MapCollection
}

type Map struct {
	dstStart int
	srcStart int
	length   int
}

type MapCollection struct {
	maps []Map
}

func (c *MapCollection) Add(m Map) {
	c.maps = append(c.maps, m)
}

func (c *MapCollection) FindDestination(src int) int {
	dst := src
	for _, m := range c.maps {
		if m.HasDestination(src) {
			dst = m.FindDestination(src)
			break
		}
	}
	return dst
}

func (m *Map) HasDestination(src int) bool {
	return src >= m.srcStart && src < m.srcStart+m.length
}

func (m *Map) FindDestination(src int) int {
	if !m.HasDestination(src) {
		panic("Only call when m.HasDestion()")
	}
	return m.dstStart + (src - m.srcStart)
}

func parseSeeds(line string) []int {
	seeds := []int{}
	for _, s := range strings.Split(line, " ")[1:] {
		if snr, err := strconv.Atoi(s); err == nil {
			seeds = append(seeds, snr)
		}
	}
	return seeds
}

func (s *Solution) ProcessLine(i int, line string) {
	if i == 0 {
		s.seeds = parseSeeds(line)
	} else if line == "" {
		return
	} else if strings.HasSuffix(line, "map:") {
		if len(s.currMaps.maps) > 0 {
			s.maps = append(s.maps, s.currMaps)
		}
		s.currMaps = MapCollection{}
	} else {
		sMap := strings.Split(line, " ")
		dstStart, _ := strconv.Atoi(sMap[0])
		srcStart, _ := strconv.Atoi(sMap[1])
		length, _ := strconv.Atoi(sMap[2])
		m := Map{dstStart: dstStart, srcStart: srcStart, length: length}
		s.currMaps.Add(m)
	}
}

func part1() {
	s := &Solution{}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println()

	// Store the last maps!!!
	// In ProcessLine() we add this _before_ the actual parsing of the maps.
	s.maps = append(s.maps, s.currMaps)

	minLocation := math.MaxInt
	var res int
	for _, seed := range s.seeds {
		res = seed
		for _, c := range s.maps {
			res = c.FindDestination(res)
		}
		minLocation = min(minLocation, res)
	}

	fmt.Println(minLocation)
}

func part2() {
	s := &Solution{}
	file.ReadLines("./input", s)
	s.maps = append(s.maps, s.currMaps)
	minLocation := math.MaxInt
	fmt.Println()
	var res int
	for i := 0; i < len(s.seeds); i += 2 {
		seedStart := s.seeds[i]
		seedLength := s.seeds[i] + s.seeds[i+1]
		for seed := seedStart; seed < seedLength; seed++ {
			res = seed
			for _, c := range s.maps {
				res = c.FindDestination(res)
			}
			minLocation = min(minLocation, res)
		}
	}
	fmt.Println(minLocation)
}

func main() {
	part1()
	part2()
}
