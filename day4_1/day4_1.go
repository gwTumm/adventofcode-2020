package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Validator func(v string) bool

type Passport map[string]string

var hclRegex = regexp.MustCompile(`^#[0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f][0-9a-f]$`)
var eclAllowed = map[string]bool{"amb": true, "blu": true, "brn": true, "gry": true, "grn": true, "hzl": true, "oth": true}
var pidRegex = regexp.MustCompile(`^\d\d\d\d\d\d\d\d\d$`)
var fourRegex = regexp.MustCompile(`^\d\d\d\d$`)

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var rules = map[string]Validator{
	"byr": func(v string) bool {
		year, err := strconv.Atoi(v)
		return fourRegex.MatchString(v) && err == nil && year >= 1920 && year <= 2002
	},
	"iyr": func(v string) bool {
		year, err := strconv.Atoi(v)
		return fourRegex.MatchString(v) && err == nil && year >= 2010 && year <= 2020
	},
	"eyr": func(v string) bool {
		year, err := strconv.Atoi(v)
		return fourRegex.MatchString(v) && err == nil && year >= 2020 && year <= 2030
	},
	"hgt": func(v string) bool {
		cm := strings.HasSuffix(v, "cm")
		in := strings.HasSuffix(v, "in")
		height, err := strconv.Atoi(v[:len(v)-2])
		return err == nil && ((cm && height >= 150 && height <= 193) || (in && height >= 59 && height <= 76))
	},
	"hcl": func(v string) bool {
		return hclRegex.MatchString(v)
	},
	"ecl": func(v string) bool {
		return eclAllowed[v]
	},
	"pid": func(v string) bool {
		return pidRegex.MatchString(v)
	},
}

func (p Passport) Check() bool {
	for _, k := range requiredFields {
		v, found := p[k]
		if !found {
			return false
		}

		rule, found := rules[k]
		if !found {
			panic("Rule missing: " + k)
		}

		if !rule(v) {
			return false
		}
	}
	return true
}

func main() {
	b, err := ioutil.ReadFile("day4_input2.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	var passport Passport = make(map[string]string)
	count := 0
	totalcount := 0

	for _, l := range lines {
		trimmed := strings.TrimSpace(l)
		if len(trimmed) == 0 {
			totalcount++
			if passport.Check() {
				count++
			}
			passport = make(map[string]string)
		} else {
			data := strings.Split(trimmed, " ")
			for _, pair := range data {
				index := strings.Index(pair, ":")
				key := pair[:index]
				value := pair[index+1:]
				passport[key] = value
			}
		}
	}

	if len(passport) > 0 {
		totalcount++
	}
	if passport.Check() {
		count++
	}
	println("Count: ", count, " Total: ", totalcount)
}
