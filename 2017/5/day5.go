package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jdwile/advent-of-code/2019/utils"
)

func main() {
	message := readInputFile()
	fmt.Printf("Part 1 Answer: %d\n\n", SolvePartOne(message))
	fmt.Printf("Part 2 Answer: %d\n\n", SolvePartTwo(message))
}

func readInputFile() []int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	message := make([]int, 0)

	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		message = append(message, n)
	}

	return message
}

func SolvePartOne(m []int) int {
	defer utils.TimeTrack(time.Now(), "Day 5: Part 1")

	message := make([]int, len(m))
	for j, n := range m {
		message[j] = n
	}

	var s, i int

	for i < len(message) {
		c := message[i]

		message[i]++
		i += c
		s++
	}

	return s
}

func SolvePartTwo(m []int) int {
	defer utils.TimeTrack(time.Now(), "Day 5: Part 2")

	message := make([]int, len(m))
	for j, n := range m {
		message[j] = n
	}

	var s, i int

	for i < len(message) {
		c := message[i]

		if c >= 3 {
			message[i]--
		} else {
			message[i]++
		}
		i += c
		s++
	}

	return s
}
