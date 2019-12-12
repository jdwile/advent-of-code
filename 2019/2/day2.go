package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	cpu "github.com/jdwile/advent-of-code/2019/intcode-cpu"
	"github.com/jdwile/advent-of-code/2019/utils"
)

func main() {
	solvePartOne()
	solvePartTwo()
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

func solvePartOne() {
	defer utils.TimeTrack(time.Now(), "Day 2: Part 1")
	m := readInput()
	m[1] = 12
	m[2] = 2

	c := cpu.ConstructCPU(m)
	c = c.ExecuteProgram()
	fmt.Println(c.Memory[0])
}

func solvePartTwo() {
	defer utils.TimeTrack(time.Now(), "Day 2: Part 2")
	target := 19690720
	found := false

	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			m := readInput()
			m[1] = noun
			m[2] = verb

			c := cpu.ConstructCPU(m)
			c = c.ExecuteProgram()

			if c.Memory[0] == target {
				found = true
				fmt.Println(100*noun + verb)
			}
		}

		if found {
			break
		}
	}
}
