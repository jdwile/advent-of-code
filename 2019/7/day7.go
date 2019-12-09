package main

import (
	"bufio"
	"fmt"
	"github.com/ernestosuarez/itertools"
	"github.com/jdwile/advent-of-code/2019/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	instructions := readInput()
	solvePart1(instructions, []int{1, 0, 4, 3, 2})
	solvePart2(instructions, []int{9, 8, 7, 6, 5})
}

type CPU struct {
	Instructions []int
	Input        []int
	Output       []int
	Counter      int
	Halted       bool
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func chooseMode(mode bool, i int, arr []int) int {
	if mode {
		return i
	}
	return arr[i]
}

func executeProgram(c CPU) CPU {
	loop := true
	for loop {
		n := c.Instructions[c.Counter]
		ji := false
		ki := false

		if n > 99 {
			s := strconv.Itoa(n)
			if len(s) == 3 {
				s = "0" + s
			}

			n = n % 100
			ji = s[1] == '1'
			ki = s[0] == '1'
		}

		switch n {
		case 1: // add
			j := c.Instructions[c.Counter+1]
			k := c.Instructions[c.Counter+2]
			l := c.Instructions[c.Counter+3]
			c.Instructions[l] = chooseMode(ji, j, c.Instructions) + chooseMode(ki, k, c.Instructions)
			c.Counter += 4
		case 2: // multiply
			j := c.Instructions[c.Counter+1]
			k := c.Instructions[c.Counter+2]
			l := c.Instructions[c.Counter+3]
			c.Instructions[l] = chooseMode(ji, j, c.Instructions) * chooseMode(ki, k, c.Instructions)
			c.Counter += 4
		case 3: // input
			j := c.Instructions[c.Counter+1]
			var k int
			if len(c.Input) > 0 {
				k = c.Input[0]
			} else {
				loop = false
				break
			}
			c.Instructions[j] = k
			if len(c.Input) <= 1 {
				c.Input = []int{}
			} else {
				c.Input = c.Input[1:]
			}
			c.Counter += 2
		case 4: // output
			j := c.Instructions[c.Counter+1]
			o := chooseMode(ji, j, c.Instructions)
			c.Output = append(c.Output, o)
			c.Counter += 2
		case 5: // jump true
			j := c.Instructions[c.Counter+1]
			k := c.Instructions[c.Counter+2]
			if chooseMode(ji, j, c.Instructions) != 0 {
				c.Counter = chooseMode(ki, k, c.Instructions)
			} else {
				c.Counter += 3
			}
		case 6: // jump false
			j := c.Instructions[c.Counter+1]
			k := c.Instructions[c.Counter+2]
			if chooseMode(ji, j, c.Instructions) == 0 {
				c.Counter = chooseMode(ki, k, c.Instructions)
			} else {
				c.Counter += 3
			}
		case 7: // less than
			j := c.Instructions[c.Counter+1]
			k := c.Instructions[c.Counter+2]
			l := c.Instructions[c.Counter+3]
			if chooseMode(ji, j, c.Instructions) < chooseMode(ki, k, c.Instructions) {
				c.Instructions[l] = 1
			} else {
				c.Instructions[l] = 0
			}
			c.Counter += 4
		case 8: // equal to
			j := c.Instructions[c.Counter+1]
			k := c.Instructions[c.Counter+2]
			l := c.Instructions[c.Counter+3]
			if chooseMode(ji, j, c.Instructions) == chooseMode(ki, k, c.Instructions) {
				c.Instructions[l] = 1
			} else {
				c.Instructions[l] = 0
			}
			c.Counter += 4
		case 99: // end
			c.Halted = true
			loop = false
			break
		}

		if c.Counter >= len(c.Instructions) {
			loop = false
		}
	}

	return c
}

func readInput() []int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	arr := strings.Split(scanner.Text(), ",")

	for _, n := range arr {
		i, _ := strconv.Atoi(n)
		nums = append(nums, i)
	}

	return nums
}

func solvePart1(s []int, input []int) {
	defer utils.TimeTrack(time.Now(), "Day 7: Part 1")
	instructions := make([]int, len(s))

	answer := 0
	for v := range itertools.PermutationsInt(input, len(input)) {
		d := 0
		for i, _ := range []string{"A", "B", "C", "D", "E"} {
			copy(instructions, s)
			c := CPU{instructions, []int{v[i], d}, []int{}, 0, false}
			c = executeProgram(c)
			d = c.Output[0]
		}
		answer = Max(answer, d)
	}
	fmt.Println(answer)
}

func solvePart2(s []int, input []int) {
	defer utils.TimeTrack(time.Now(), "Day 7: Part 2")

	amps := []string{"A", "B", "C", "D", "E"}
	m := 0

	for v := range itertools.PermutationsInt(input, len(input)) {
		var lastE int
		halt := false
		cpus := make([]CPU, 5)
		for i, j := range v {
			instructions := make([]int, len(s))
			copy(instructions, s)
			cpus[i] = CPU{instructions, []int{}, []int{}, 0, false}
			if i == 0 {
				cpus[i].Input = []int{j, 0}
			} else {
				cpus[i].Input = []int{j}
			}
		}
		for !halt {
			for i, _ := range amps {
				if !cpus[i].Halted {
					cpus[i] = executeProgram(cpus[i])
					if i == len(amps)-1 && len(cpus[i].Output) > 0 {
						lastE = cpus[i].Output[len(cpus[i].Output)-1]
					}
					for _, j := range cpus[i].Output {
						cpus[(i+1)%len(cpus)].Input = append(cpus[(i+1)%len(cpus)].Input, j)
					}
					cpus[i].Output = []int{}
				}
			}
			halt = true
			for _, cpu := range cpus {
				halt = halt && cpu.Halted
			}
		}
		m = Max(m, lastE)
	}
	fmt.Println(m)
}
