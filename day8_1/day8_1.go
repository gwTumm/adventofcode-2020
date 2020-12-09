package main

import (
	"io/ioutil"
	"strconv"
	"strings"
)

type Instruction struct {
	Cmd     string
	Offset  int
	Visited bool
}

func main() {
	b, err := ioutil.ReadFile("day8_input2.txt")
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(b), "\n")
	var instructions []*Instruction

	for _, lt := range lines {
		l := strings.TrimSpace(lt)
		sp := strings.Split(l, " ")

		cmd := sp[0]
		sign := sp[1][0]
		count, err := strconv.Atoi(sp[1][1:])
		if err != nil {
			panic(err)
		}

		offset := 0
		if sign == '+' {
			offset = count
		} else {
			offset = -count
		}

		instructions = append(instructions, &Instruction{
			Cmd:    cmd,
			Offset: offset,
		})
	}

	// Count nops and jmps
	var mods []*Instruction
	for _, inst := range instructions {
		if inst.Cmd == "nop" || inst.Cmd == "jmp" {
			mods = append(mods, inst)
		}
	}

	terminated := false
	acc := 0
	for _, mod := range mods {
		pos := 0
		acc = 0
		terminated = false

		// Reset insts
		for _, i := range instructions {
			i.Visited = false
		}

		for {
			if pos > len(instructions) || pos < 0 {
				panic("Illegal pos: " + strconv.Itoa(pos))
			}
			if pos == len(instructions) { // EOF
				terminated = true
				break
			}

			inst := instructions[pos]
			if inst.Visited {
				break
			}
			inst.Visited = true
			cmd := inst.Cmd

			if mod == inst {
				if inst.Cmd == "jmp" {
					cmd = "nop"
				} else if inst.Cmd == "nop" {
					cmd = "jmp"
				}
			}

			// println(inst.Cmd, inst.Offset)
			switch cmd {
			case "acc":
				acc += inst.Offset
			case "jmp":
				pos += inst.Offset
				continue
			case "nop":
			default:
				panic("Unknown instruction " + cmd)
			}
			pos++
		}

		if terminated {
			break
		}
	}

	if !terminated {
		panic("Couldn't find")
	}

	println("Acc: ", acc)
}
