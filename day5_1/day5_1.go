package main

import (
	"io/ioutil"
	"math"
	"strings"
)

func seatId(row int, column int) int {
	return row*8 + column
}

func main() {
	b, err := ioutil.ReadFile("day5_input2.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	maxseatid := 0
	seatids := make(map[int]bool)

	for _, passt := range lines {
		pass := []rune(strings.TrimSpace(passt))
		start := 0
		stop := 127
		for i := 0; i < 7; i++ {
			switch pass[i] {
			case 'F':
				stop = start + (stop-start)/2
			case 'B':
				start = start + int(math.Ceil(float64(stop-start)/float64(2)))
			default:
				panic("Uknown at " + string(pass[i]))
			}
		}

		startcol := 0
		stopcol := 7
		for i := 7; i < 10; i++ {
			switch pass[i] {
			case 'L':
				stopcol = startcol + (stopcol-startcol)/2
			case 'R':
				startcol = startcol + int(math.Ceil(float64(stopcol-startcol)/float64(2)))
			default:
				panic("Uknown at " + string(pass[i]))
			}
		}

		// println("Row ", start, " Col ", startcol)
		seatid := seatId(start, startcol)
		seatids[seatid] = true
		if seatid > maxseatid {
			maxseatid = seatid
		}
	}

	println("Max seat id ", maxseatid)

	for i := 0; i < maxseatid; i++ {
		_, existslesser := seatids[i-1]
		_, existsmore := seatids[i+1]
		_, exists := seatids[i]

		if !exists && existslesser && existsmore {
			println("My seat: ", i)
		}
	}
}
