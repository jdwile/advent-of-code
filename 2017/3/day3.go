package main

import (
	"fmt"
	"time"

	"github.com/jdwile/advent-of-code/2019/utils"
)

func Abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

func main() {
	input := 265149
	fmt.Printf("Part 1 Answer: %d\n\n", SolvePartOne(input))
	fmt.Printf("Part 2 Answer: %d\n\n", SolvePartTwo(input))
}

func SolvePartOne(input int) int {
	defer utils.TimeTrack(time.Now(), "Day 3: Part 1")

	x, y := 0, 0
	n, s := 1, 1

	for n < input {
		x++
		y--
		s += 2
		n = s * s
	}

	DXS := []int{-1, 0, 1, 0}
	DYS := []int{0, 1, 0, -1}
	dir := 0

	for n-(s-1) > input {
		n -= (s - 1)
		x += DXS[dir] * (s - 1)
		y += DYS[dir] * (s - 1)
		dir++
	}

	for n > input {
		n--
		x += DXS[dir]
		y += DYS[dir]
	}

	return Abs(x) + Abs(y)
}

func SolvePartTwo(input int) int {
	defer utils.TimeTrack(time.Now(), "Day 3: Part 2")

	return 0
}
