package main

import (
	"aoc2023/lib/algo"
	"aoc2023/lib/file"
	"fmt"
	"slices"
	"strings"

	"golang.org/x/exp/maps"
)

const (
	ON  = true
	OFF = false
	HI  = true
	LO  = false
)

type Pulse struct {
	sender   string
	receiver string
	strength bool
}

func newPulse(sender, receiver string, strength bool) Pulse {
	return Pulse{sender, receiver, strength}
}

type Module struct {
	mType     string
	name      string
	receivers []string
	state     bool
	memory    map[string]bool
}

func (m *Module) ReceivePulse(p Pulse) []Pulse {
	if m.mType == "%" {
		return m.flipFlopReceives(p)
	}
	return m.conjunctionReceives(p)
}

func (m *Module) flipFlopReceives(p Pulse) []Pulse {
	if p.strength == HI {
		return []Pulse{}
	}
	m.state = !m.state
	pulses := []Pulse{}
	for _, r := range m.receivers {
		pulses = append(pulses, newPulse(m.name, r, m.state))
	}
	return pulses
}

func (m *Module) conjunctionReceives(p Pulse) []Pulse {
	m.memory[p.sender] = p.strength
	allHi := true
	for _, s := range m.memory {
		allHi = allHi && s
	}
	pulses := []Pulse{}
	for _, r := range m.receivers {
		pulses = append(pulses, newPulse(m.name, r, !allHi))
	}
	return pulses
}

func newFlipFlop(name string, receivers []string, state bool) *Module {
	return &Module{
		mType:     "%",
		name:      name,
		receivers: receivers,
		state:     state,
	}
}

func newConjuction(name string, receivers []string) *Module {
	return &Module{
		mType:     "&",
		name:      name,
		receivers: receivers,
		memory:    make(map[string]bool),
	}
}

type Solution struct {
	broadcasts []string
	modules    map[string]*Module
}

func (s *Solution) ProcessLine(i int, line string) {
	p := strings.Split(line, " -> ")
	moduleType := p[0][0]
	name := p[0][1:]
	receivers := strings.Split(p[1], ", ")
	switch moduleType {
	// In my input file, there are no flipflops and conjunctions that share the
	// same name.
	case '%':
		s.modules[name] = newFlipFlop(name, receivers, OFF)
	case '&':
		s.modules[name] = newConjuction(name, receivers)
	case 'b':
		s.broadcasts = receivers
	default:
		panic(fmt.Sprintf("Unknown module type [%v]", moduleType))
	}
}

func (s *Solution) initConjunctions() {
	for _, m := range s.modules {
		for _, r := range m.receivers {
			if rm, ok := s.modules[r]; ok {
				if rm.mType == "&" {
					rm.memory[m.name] = LO
				}
			}
		}
	}
}

func Solve1() int {
	s := &Solution{modules: make(map[string]*Module)}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("Read", lineCount, "lines")
	s.initConjunctions()
	hi := 0
	lo := 0
	q := []Pulse{}
	var p Pulse
	for i := 0; i < 1000; i++ {
		lo++ // button press
		for _, name := range s.broadcasts {
			q = append(q, newPulse("broadcaster", name, LO))
		}
		for len(q) > 0 {
			p = q[0]
			q = q[1:]
			if p.strength == LO {
				lo++
			} else {
				hi++
			}
			if m, ok := s.modules[p.receiver]; ok {
				for _, np := range m.ReceivePulse(p) {
					q = append(q, np)
				}
			}
		}
	}
	return hi * lo
}

func allPositive(namedCounts map[string]int) bool {
	all := true
	for _, n := range namedCounts {
		all = all && (n > 0)
	}
	return all
}

// Manually checking input file: Only 1 module sends to "rx", ie: "&mg"
// mg will only send rx a low pulse if all of its senders have been high
// 4 Modules send to mg: &jg, &rh, &jm, &hf
// Feels like 08-2:
// Determine when each of the 4 modules receives a hi pulse and lcm()?
func Solve2() int {
	s := &Solution{modules: make(map[string]*Module)}
	file.ReadLines("./input", s)
	s.initConjunctions()
	q := []Pulse{}
	var p Pulse
	count := 0
	hiCycles := map[string]int{"jg": 0, "rh": 0, "jm": 0, "hf": 0}
	for {
		count++
		if allPositive(hiCycles) {
			break
		}
		for _, name := range s.broadcasts {
			q = append(q, newPulse("broadcaster", name, LO))
		}
		for len(q) > 0 {
			p = q[0]
			q = q[1:]
			if slices.Contains(maps.Keys(hiCycles), p.sender) && p.receiver == "mg" && p.strength == HI {
				hiCycles[p.sender] = count
			}
			if m, ok := s.modules[p.receiver]; ok {
				for _, np := range m.ReceivePulse(p) {
					q = append(q, np)
				}
			}
		}
	}
	return algo.LCMSlice(maps.Values(hiCycles))
}

func main() {
	fmt.Println(Solve1())
	fmt.Println(Solve2())
}
