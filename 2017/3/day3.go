package main

import (
	"fmt"
	"time"

	"github.com/jdwile/advent-of-code/2017/utils"
)

func Abs(a int) int {
	if a < 0 {
		return -1 * a
	}
	return a
}

type Point struct {
	X int
	Y int
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

func getSumNeighbors(grid map[Point]int, pos Point) int {
	DXS := []int{-1, 0, 1}
	DYS := []int{-1, 0, 1}

	n := 0

	for _, dx := range DXS {
		for _, dy := range DYS {
			if dx == 0 && dy == 0 {
				continue
			}

			n += grid[Point{pos.X + dx, pos.Y + dy}]
		}
	}

	if n == 0 {
		return 1
	}

	return n
}

func SolvePartTwo(input int) int {
	defer utils.TimeTrack(time.Now(), "Day 3: Part 2")

	var x, y int
	grid := make(map[Point]int)

	var m int
	loop := true

	for loop {
		grid[Point{x, y}] = getSumNeighbors(grid, Point{x, y})
		if grid[Point{x, y}] > input {
			loop = false
			break
		}

		m++
		x++

		for y+1 != m {
			grid[Point{x, y}] = getSumNeighbors(grid, Point{x, y})
			if grid[Point{x, y}] > input {
				loop = false
				break
			}
			y++
		}

		if !loop {
			break
		}

		for Abs(x-1) != m {
			grid[Point{x, y}] = getSumNeighbors(grid, Point{x, y})
			if grid[Point{x, y}] > input {
				loop = false
				break
			}
			x--
		}

		if !loop {
			break
		}

		for Abs(y-1) != m {
			grid[Point{x, y}] = getSumNeighbors(grid, Point{x, y})
			if grid[Point{x, y}] > input {
				loop = false
				break
			}
			y--
		}

		if !loop {
			break
		}

		for x+1 != m {
			grid[Point{x, y}] = getSumNeighbors(grid, Point{x, y})
			if grid[Point{x, y}] > input {
				loop = false
				break
			}
			x++
		}

		if !loop {
			break
		}

		grid[Point{x, y}] = getSumNeighbors(grid, Point{x, y})
		if grid[Point{x, y}] > input {
			loop = false
			break
		}
	}

	return grid[Point{x, y}]
}
