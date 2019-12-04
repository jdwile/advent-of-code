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
	solvePartOne(instructions)
	solvePartTwo(instructions)
}

func executeProgram(instructions []int) int {
	for i := 0; i < len(instructions); i += 4 {
		n := instructions[i]
		switch n {
		case 1:
			j := instructions[i+1]
			k := instructions[i+2]
			l := instructions[i+3]
			instructions[l] = instructions[j] + instructions[k]
		case 2:
			j := instructions[i+1]
			k := instructions[i+2]
			l := instructions[i+3]
			instructions[l] = instructions[j] * instructions[k]
		case 99:
			break
		}
	}
	return instructions[0]
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

func solvePartOne(s []int) {
	defer utils.TimeTrack(time.Now(), "Day 2: Part 1")
	instructions := make([]int, len(s))
	copy(instructions, s)

	instructions[1] = 12
	instructions[2] = 2
	res := executeProgram(instructions)
	fmt.Println(res)
}

func solvePartTwo(s []int) {
	defer utils.TimeTrack(time.Now(), "Day 2: Part 2")
	target := 19690720

	instructions := make([]int, len(s))
	found := false

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			copy(instructions, s)
			instructions[1] = noun
			instructions[2] = verb
			res := executeProgram(instructions)

			if res == target {
				found = true
				fmt.Println(100*noun+verb, res)
			}
		}

		if found {
			break
		}
	}
}
