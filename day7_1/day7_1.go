package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

var nobag = regexp.MustCompile(`^(\w+ \w+) bags contain no other bags\.$`)
var submatch = regexp.MustCompile(`^(\d+) (\w+ \w+) bags?$`)
var match = regexp.MustCompile(`^(\w+ \w+) bags contain (.*)\.$`)
var bags = make(map[string]map[string]int)
var canContain = make(map[string]map[string]bool)

func subCount(c string) int {
	count := 0
	for sub, subcount := range bags[c] {
		count += subcount * (1 + subCount(sub))
	}
	return count
}

func main() {
	b, err := ioutil.ReadFile("day7_input2.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	for _, lt := range lines {
		l := strings.TrimSpace(lt)
		if a := nobag.FindStringSubmatch(l); len(a) > 0 {
			color := a[1]
			bags[color] = nil
		} else if a := match.FindStringSubmatch(l); len(a) > 0 {
			color := a[1]
			subcolors := make(map[string]int)
			substring := a[2]
			substrings := strings.Split(substring, ",")
			for _, st := range substrings {
				s := strings.TrimSpace(st)
				subs := submatch.FindStringSubmatch(s)
				for i := 1; i < len(subs); i += 2 {
					count, err := strconv.Atoi(subs[i])
					if err != nil {
						panic(err)
					}
					color := subs[i+1]
					subcolors[color] = count
				}
			}
			bags[color] = subcolors
			// println(color, len(subcolors))
		} else {
			panic("Unknown line " + lt)
		}
	}

	// Build canContain map
	for c, subs := range bags {
		if _, exists := canContain[c]; !exists {
			canContain[c] = make(map[string]bool)
		}
		for sub := range subs {
			canContain[c][sub] = true
		}
	}

	changed := false
	for {
		changed = false
		for c, subs := range canContain {
			for s := range subs {
				for transc := range canContain[s] {
					if !subs[transc] {
						changed = true
					}
					canContain[c][transc] = true
				}
			}
		}

		if !changed {
			break
		}
	}

	countShinyGold := 0
	for _, subs := range canContain {
		if subs["shiny gold"] {
			countShinyGold++
		}
	}

	// println("Shiny gold: ", countShinyGold)
	// println("Shiny gold subs: ", len(bags["shiny gold"]))
	println("Shiny gold contain: ", subCount("shiny gold"))
}
