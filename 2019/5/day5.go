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
	instructions := readInput()
	solve(instructions, []int{1})
	solve(instructions, []int{5})
}

func chooseMode(mode bool, i int, arr []int) int {
	if mode {
		return i
	}
	return arr[i]
}

func executeProgram(instructions []int, input []int) {
	i := 0
	loop := true
	for loop {
		n := instructions[i]
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
			j := instructions[i+1]
			k := instructions[i+2]
			l := instructions[i+3]
			instructions[l] = chooseMode(ji, j, instructions) + chooseMode(ki, k, instructions)
			i += 4
		case 2: // multiply
			j := instructions[i+1]
			k := instructions[i+2]
			l := instructions[i+3]
			instructions[l] = chooseMode(ji, j, instructions) * chooseMode(ki, k, instructions)
			i += 4
		case 3: // input
			j := instructions[i+1]
			k := input[0]
			instructions[j] = k
			input = input[1:]
			i += 2
		case 4: // output
			j := instructions[i+1]
			output := chooseMode(ji, j, instructions)
			fmt.Println("PROGRAM OUTPUT:", output)
			i += 2
		case 5: // jump true
			j := instructions[i+1]
			k := instructions[i+2]
			if chooseMode(ji, j, instructions) != 0 {
				i = chooseMode(ki, k, instructions)
			} else {
				i += 3
			}
		case 6: // jump false
			j := instructions[i+1]
			k := instructions[i+2]
			if chooseMode(ji, j, instructions) == 0 {
				i = chooseMode(ki, k, instructions)
			} else {
				i += 3
			}
		case 7: // less than
			j := instructions[i+1]
			k := instructions[i+2]
			l := instructions[i+3]
			if chooseMode(ji, j, instructions) < chooseMode(ki, k, instructions) {
				instructions[l] = 1
			} else {
				instructions[l] = 0
			}
			i += 4 // equal to
		case 8:
			j := instructions[i+1]
			k := instructions[i+2]
			l := instructions[i+3]
			if chooseMode(ji, j, instructions) == chooseMode(ki, k, instructions) {
				instructions[l] = 1
			} else {
				instructions[l] = 0
			}
			i += 4
		case 99: // end
			loop = false
			break
		}

		if i >= len(instructions) {
			loop = false
		}
	}
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

func solve(s []int, input []int) {
	defer utils.TimeTrack(time.Now(), "Day 5")
	instructions := make([]int, len(s))
	copy(instructions, s)

	executeProgram(instructions, input)
}
