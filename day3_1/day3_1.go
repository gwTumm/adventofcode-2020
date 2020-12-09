package main

import (
	"io/ioutil"
	"strings"
)

// Row, column indexed
var m [][]bool

func main() {
	b, err := ioutil.ReadFile("day3_input2.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	m = make([][]bool, len(lines))
	for n, l := range lines {
		runes := []rune(strings.TrimSpace(l))
		m[n] = make([]bool, len(runes))
		for nc, r := range runes {
			if r == '#' {
				m[n][nc] = true
			}
		}
	}

	step := []struct {
		Y, X int
	}{{1, 1}, {1, 3}, {1, 5}, {1, 7}, {2, 1}}

	mult := 1
	for _, s := range step {
		x := 0
		y := 0
		count := 0
		for {
			if y >= len(m) {
				break
			}
			if m[y][x%len(m[y])] {
				count++
			}
			x += s.X
			y += s.Y
		}
		mult *= count
	}
	println("Trees Multiplied: ", mult)
}
