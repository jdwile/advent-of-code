package main

import (
	"bufio"
	"fmt"
	"github.com/jdwile/advent-of-code/2019/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	memory := readInput()
	solvePart1(memory)
	solvePart2(memory)
}

func readInput() map[int]int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	m := make(map[int]int)

	scanner.Scan()
	arr := strings.Split(scanner.Text(), ",")

	for i, n := range arr {
		a, _ := strconv.Atoi(n)
		m[i] = a
	}

	return m
}

type CPU struct {
	Memory             map[int]int
	Input              []int
	Output             []int
	InstructionPointer int
	RelativeBase       int
	Halted             bool
}

func ConstructCPU(memory map[int]int) CPU {
	return CPU{memory, []int{}, []int{}, 0, 0, false}
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func chooseValueMode(mode byte, i int, c CPU) int {
	switch mode {
	case '0':
		return c.Memory[i]
	case '1':
		return i
	case '2':
		return c.Memory[c.RelativeBase+i]
	}
	panic("NO MODE")
}

func chooseSetMode(mode byte, i int, c CPU) int {
	if mode == '2' {
		return c.RelativeBase + i
	}
	return i
}

func parseOpcode(n int) (int, byte, byte, byte) {
	jMode := "0"[0]
	kMode := "0"[0]
	lMode := "0"[0]

	if n > 99 {
		s := strconv.Itoa(n)
		if len(s) == 3 {
			s = "0" + s
		}
		if len(s) == 4 {
			s = "0" + s
		}

		n = n % 100
		jMode = s[2]
		kMode = s[1]
		lMode = s[0]
	}
	return n, jMode, kMode, lMode
}

func (c CPU) executeProgram() CPU {
	loop := true
	for loop {
		n, jMode, kMode, lMode := parseOpcode(c.Memory[c.InstructionPointer])

		switch n {
		case 1: // add
			j := c.Memory[c.InstructionPointer+1]
			k := c.Memory[c.InstructionPointer+2]
			l := c.Memory[c.InstructionPointer+3]
			c.Memory[chooseSetMode(lMode, l, c)] = chooseValueMode(jMode, j, c) + chooseValueMode(kMode, k, c)
			c.InstructionPointer += 4
		case 2: // multiply
			j := c.Memory[c.InstructionPointer+1]
			k := c.Memory[c.InstructionPointer+2]
			l := c.Memory[c.InstructionPointer+3]
			c.Memory[chooseSetMode(lMode, l, c)] = chooseValueMode(jMode, j, c) * chooseValueMode(kMode, k, c)
			c.InstructionPointer += 4
		case 3: // input
			if (len(c.Input) == 0) {
				loop = false
				break
			}
			j := c.Memory[c.InstructionPointer+1]
			k := c.Input[0]
			c.Memory[chooseSetMode(jMode, j, c)] = k
			if len(c.Input) <= 1 {
				c.Input = []int{}
			} else {
				c.Input = c.Input[1:]
			}
			c.InstructionPointer += 2
		case 4: // output
			j := c.Memory[c.InstructionPointer+1]
			o := chooseValueMode(jMode, j, c)
			c.Output = append(c.Output, o)
			c.InstructionPointer += 2
		case 5: // jump true
			j := c.Memory[c.InstructionPointer+1]
			k := c.Memory[c.InstructionPointer+2]
			if chooseValueMode(jMode, j, c) != 0 {
				c.InstructionPointer = chooseValueMode(kMode, k, c)
			} else {
				c.InstructionPointer += 3
			}
		case 6: // jump false
			j := c.Memory[c.InstructionPointer+1]
			k := c.Memory[c.InstructionPointer+2]
			if chooseValueMode(jMode, j, c) == 0 {
				c.InstructionPointer = chooseValueMode(kMode, k, c)
			} else {
				c.InstructionPointer += 3
			}
		case 7: // less than
			j := c.Memory[c.InstructionPointer+1]
			k := c.Memory[c.InstructionPointer+2]
			l := c.Memory[c.InstructionPointer+3]
			if chooseValueMode(jMode, j, c) < chooseValueMode(kMode, k, c) {
				c.Memory[chooseSetMode(lMode, l, c)] = 1
			} else {
				c.Memory[chooseSetMode(lMode, l, c)] = 0
			}
			c.InstructionPointer += 4
		case 8: // equal to
			j := c.Memory[c.InstructionPointer+1]
			k := c.Memory[c.InstructionPointer+2]
			l := c.Memory[c.InstructionPointer+3]
			if chooseValueMode(jMode, j, c) == chooseValueMode(kMode, k, c) {
				c.Memory[chooseSetMode(lMode, l, c)] = 1
			} else {
				c.Memory[chooseSetMode(lMode, l, c)] = 0
			}
			c.InstructionPointer += 4
		case 9: // adjust relative base
			j := c.Memory[c.InstructionPointer+1]
			c.RelativeBase += chooseValueMode(jMode, j, c)
			c.InstructionPointer += 2
		case 99: // end
			c.Halted = true
			loop = false
			break
		}
	}

	return c
}

func solvePart1(m map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 9: Part 1")

	c := ConstructCPU(m)
	c.Input = []int{1}

	c = c.executeProgram()
	fmt.Println(c.Output)
}

func solvePart2(m map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 9: Part 2")

	c := ConstructCPU(m)
	c.Input = []int{2}

	c = c.executeProgram()
	fmt.Println(c.Output)
}
