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
		case 1:
			j := instructions[i+1]
			k := instructions[i+2]
			l := instructions[i+3]
			instructions[l] = chooseMode(ji, j, instructions) + chooseMode(ki, k, instructions)
			i += 4
		case 2:
			j := instructions[i+1]
			k := instructions[i+2]
			l := instructions[i+3]
			instructions[l] = chooseMode(ji, j, instructions) * chooseMode(ki, k, instructions)
			i += 4
		case 3:
			j := instructions[i+1]
			k := input[0]
			instructions[j] = k
			input = input[1:]
			i += 2
		case 4:
			j := instructions[i+1]
			output := chooseMode(ji, j, instructions)
			fmt.Println("OUTPUT:", output)
			i += 2
		case 5:
			j := instructions[i+1]
			k := instructions[i+2]
			if chooseMode(ji, j, instructions) != 0 {
				i = chooseMode(ki, k, instructions)
			} else {
				i += 3
			}
		case 6:
			j := instructions[i+1]
			k := instructions[i+2]
			if chooseMode(ji, j, instructions) == 0 {
				i = chooseMode(ki, k, instructions)
			} else {
				i += 3
			}
		case 7:
			j := instructions[i+1]
			k := instructions[i+2]
			l := instructions[i+3]
			if chooseMode(ji, j, instructions) < chooseMode(ki, k, instructions) {
				instructions[l] = 1
			} else {
				instructions[l] = 0
			}
			i += 4
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
		case 99:
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

func solve(s []int, i []int) {
	defer utils.TimeTrack(time.Now(), "Day 2: Part 1")
	instructions := make([]int, len(s))
	copy(instructions, s)
	input := make([]int, len(i))
	copy(input, i)

	executeProgram(instructions, input)
}

// func solvePartTwo(s []int) {
// 	defer utils.TimeTrack(time.Now(), "Day 2: Part 2")
// 	target := 19690720

// 	instructions := make([]int, len(s))
// 	found := false

// 	for noun := 0; noun <= 99; noun++ {
// 		for verb := 0; verb <= 99; verb++ {
// 			copy(instructions, s)
// 			instructions[1] = noun
// 			instructions[2] = verb
// 			res := executeProgram(instructions)

// 			if res == target {
// 				found = true
// 				fmt.Println(100*noun+verb, res)
// 			}
// 		}

// 		if found {
// 			break
// 		}
// 	}
// }
