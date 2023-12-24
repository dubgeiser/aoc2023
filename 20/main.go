package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strings"
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

func (s *Solution) Solve1() int {
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

func main() {
	s := &Solution{modules: make(map[string]*Module)}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println()
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.Solve1())
}
