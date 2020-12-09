package main

import (
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

var nums []int

const preamble = 25

func checknum(n int) bool {
	if len(nums) < preamble {
		return true
	}

	for i := len(nums) - preamble; i < len(nums); i++ {
		for j := len(nums) - preamble; j < len(nums); j++ {
			if i != j && nums[i]+nums[j] == n {
				return true
			}
		}
	}

	return false
}

func main() {
	b, err := ioutil.ReadFile("day9_input2.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	invalid := 0

	for i, lt := range lines {
		l := strings.TrimSpace(lt)
		num, err := strconv.Atoi(l)

		if err != nil {
			panic(err)
		}

		if checknum(num) {
			nums = append(nums, num)
		} else {
			println("Not valid: line ", i, " ", num)
			invalid = num
			break
		}
	}

testiter:
	for i := range nums {
		sum := 0
		smallest := math.MaxInt64
		biggest := 0
		for j, n2 := range nums[i:] {
			sum += n2
			if n2 < smallest {
				smallest = n2
			}
			if n2 > biggest {
				biggest = n2
			}
			if sum == invalid {
				println("Contagious from ", i, " to ", i+j, ": ", invalid)
				println("Biggest ", biggest, " Smallest ", smallest, " sum ", smallest+biggest)
				break testiter
			}
			if sum > invalid {
				break
			}
		}
	}
}
