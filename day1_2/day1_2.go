package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("day1_input2.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(b), "\n")
	list := make([]int, len(lines))
	for n, l := range lines {
		list[n], _ = strconv.Atoi(strings.TrimSpace(l))
	}
	for l1, n1 := range list {
		for l2, n2 := range list {
			for l3, n3 := range list {
				if l1 != l2 && l1 != l3 {
					if n1+n2+n3 == 2020 {
						println("Answer: ", n1*n2*n3)
						return
					}
				}
			}
		}
	}
}
