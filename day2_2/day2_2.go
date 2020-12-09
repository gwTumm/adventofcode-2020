package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

type Pw struct {
	Min  int
	Max  int
	Char rune
	Pw   string
}

var reg = regexp.MustCompile(`^(\d+)-(\d+) (.): (.*)`)

func ReadPw(s string) Pw {
	m := reg.FindStringSubmatch(s)
	if len(m) == 0 {
		panic("Not matched for " + s)
	}

	min, _ := strconv.Atoi(m[1])
	max, _ := strconv.Atoi(m[2])
	var char rune
	for _, r := range m[3] {
		char = r
		break
	}
	pw := m[4]
	return Pw{Min: min, Max: max, Char: char, Pw: pw}
}

func (p Pw) Check() bool {
	runes := []rune(p.Pw)

	b1 := runes[p.Min-1] == p.Char
	b2 := runes[p.Max-1] == p.Char

	return (b1 || b2) && !(b1 && b2)
}

func main() {
	b, err := ioutil.ReadFile("day2_input2.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")

	valid := 0
	for _, l := range lines {
		pw := ReadPw(l)
		if pw.Check() {
			valid++
		}
	}

	println("Valid: ", valid)
}
