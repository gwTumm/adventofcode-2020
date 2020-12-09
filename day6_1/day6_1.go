package main

import (
	"io/ioutil"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("day6_input2.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")

	groups := make([][]map[rune]bool, 0)
	answers := make([]map[rune]bool, 0)

	for _, lt := range lines {
		l := strings.TrimSpace(lt)
		if len(l) == 0 {
			groups = append(groups, answers)
			answers = make([]map[rune]bool, 0)
		} else {
			answer := make(map[rune]bool)
			for _, r := range l {
				answer[r] = true
			}
			answers = append(answers, answer)
		}
	}

	if len(answers) > 0 {
		groups = append(groups, answers)
	}

	allcount := 0
	anycount := 0
	for _, group := range groups {
		anyanswered := make(map[rune]bool)
		allanswered := make(map[rune]bool)

		if len(group) == 0 {
			continue
		}

		for a := range group[0] {
			allanswered[a] = true
		}

		for _, answer := range group {
			for a := range answer {
				anyanswered[a] = true
			}
		}

		for _, answer := range group[1:] {
			missing := make(map[rune]bool)
			for a := range allanswered {
				if !answer[a] {
					missing[a] = true
				}
			}
			for a := range missing {
				delete(allanswered, a)
			}
		}

		// println("Group answered: ", len(anyanswered))
		anycount += len(anyanswered)
		allcount += len(allanswered)
	}

	println("Any sum: ", anycount)
	println("All sum: ", allcount)
}
