package main

import (
	"aoc2023/lib/file"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/exp/maps"
)

type Solution struct {
	parsingParts bool
	answer1      int
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

const (
	MIN = 0
	MAX = 1
)

// Go through the workflows, starting with "in" and count the combination of
// accepted ranges for the given workflow to a running total and return that
// running total.
//
// If a workflow leads to acceptance (A); all the xmas ranges are ok, thus we
// can count all combinations, which will be literally the product of the
// distances of the ranges).
//
// If a workflow leads to rejection (R) We don't need to count anything (0).
//
// To actually count a workflow: go through its rules and for each rule a range
// can be determined for which the rules applies and for which it doesn't.
// For the applicable ranges: add their count (wf: Rule.result) to the running
// total.
// For the unapplicable: adjust the original range to reduce it and try again.
// If we did this for all rules: add the count for the ranges for the default
// result of the workflow to the running total.
func (s *Solution) countCombinations(ranges map[string][2]int, wf string) int {
	if wf == "R" {
		return 0
	}
	if wf == "A" {
		combinations := 1
		for _, r := range maps.Values(ranges) {
			combinations *= r[MAX] - r[MIN] + 1 // range is inclusive.
		}
		return combinations
	}

	count := 0
	workflow := s.workflows[wf]
	for _, r := range workflow.rules {
		rmin := ranges[r.rating][MIN]
		rmax := ranges[r.rating][MAX]
		var rApplicable [2]int
		var rUnapplicable [2]int
		if r.op == "<" {
			rApplicable = [2]int{rmin, min(r.value-1, rmax)}
			rUnapplicable = [2]int{max(r.value, rmin), rmax}
		} else {
			rApplicable = [2]int{max(r.value+1, rmin), rmax}
			rUnapplicable = [2]int{rmin, min(r.value, rmax)}
		}
		if rApplicable[MIN] <= rApplicable[MAX] {
			// Range for which the rule applies is a valid range.
			// -> Add the count of that range to the current count.
			// Do not adjust the originally given range, there are still
			// operations to be done on it!
			rr := map[string][2]int{}
			maps.Copy(rr, ranges)
			rr[r.rating] = [2]int{rApplicable[MIN], rApplicable[MAX]}
			count += s.countCombinations(rr, r.result)
		}
		if rUnapplicable[MIN] <= rUnapplicable[MAX] {
			// Range for which the rule does not apply is a valid range:
			// -> Adjust the original range for this rules' rating
			ranges[r.rating] = [2]int{rUnapplicable[MIN], rUnapplicable[MAX]}
		}
	}
	// The ranges have been adjusted for all rules that don't apply, add the
	// count for the workflow's default result to the running total.
	return count + s.countCombinations(ranges, workflow.defaultResult)
}

func main() {
	s := &Solution{workflows: make(map[string]*Workflow)}
	lineCount, err := file.ReadLines("./input", s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Read", lineCount, "lines")
	fmt.Println(s.answer1)

	xmasRanges := map[string][2]int{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000},
	}
	fmt.Println(s.countCombinations(xmasRanges, "in"))
}
