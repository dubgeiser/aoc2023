package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strconv"
	"strings"
)

type Solution struct {
	parsingParts bool
	answer1      int
	answer2      int
	workflows    map[string]*Workflow
}

func (s *Solution) ProcessLine(i int, line string) {
	if line == "" {
		s.parsingParts = true
	} else if s.parsingParts {
		p := parsePart(line)
		if s.accepts(p) {
			s.answer1 += p["x"] + p["m"] + p["a"] + p["s"]
		}
	} else {
		name, wf := parseWorkFlow(line)
		s.workflows[name] = wf
	}
}

func (s *Solution) accepts(part map[string]int) bool {
	wf := "in"
	for wf != "A" && wf != "R" {
		wf = s.workflows[wf].run(part)
	}
	return wf == "A"
}

func parsePart(line string) map[string]int {
	sPart := strings.Split(strings.Trim(line, "{}"), ",")
	part := make(map[string]int)
	for _, skv := range sPart {
		kv := strings.Split(skv, "=")
		k := kv[0]
		v, _ := strconv.Atoi(kv[1])
		part[k] = v
	}
	return part
}

func parseWorkFlow(line string) (string, *Workflow) {
	var p []string
	p = strings.Split(line, "{")
	name := p[0]
	p = strings.Split(strings.Trim(p[1], "}"), ",")
	defaultResult := p[len(p)-1]
	sRules := p[:len(p)-1]
	rules := []Rule{}
	for _, r := range sRules {
		p = strings.Split(r, ":")
		result := p[1]
		rating := string(p[0][0])
		op := string(p[0][1])
		value, _ := strconv.Atoi(p[0][2:])
		rules = append(rules, Rule{rating, op, value, result})
	}
	return name, &Workflow{rules, defaultResult}
}

type Workflow struct {
	rules         []Rule
	defaultResult string
}

func (wf *Workflow) run(part map[string]int) string {
	for _, r := range wf.rules {
		if r.applies(part) {
			return r.result
		}
	}
	return wf.defaultResult
}

type Rule struct {
	rating string
	op     string
	value  int
	result string
}

func (r *Rule) applies(part map[string]int) bool {
	if r.op == ">" {
		return part[r.rating] > r.value
	}
	return part[r.rating] < r.value
}

func main() {
	s := &Solution{workflows: make(map[string]*Workflow)}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.answer1)
	fmt.Println(s.answer2)
}
