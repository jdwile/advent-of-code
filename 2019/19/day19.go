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

type Point struct {
	X int
	Y int
}

func main() {
	memory := ReadInput()
	SolvePartOne(memory)
	SolvePartTwo(memory)
}

func ReadInput() map[int]int {
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

func GetTractorBeam(m map[int]int, size int) map[Point]string {
	grid := make(map[Point]string)

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			memory := make(map[int]int)
			for e := range m {
				memory[e] = m[e]
			}
			c := cpu.ConstructCPU(memory)

			c.Input = []int{x, y}

			c = c.ExecuteProgram()

			r := c.Output[0]
			c.Output = []int{}

			loc := Point{x, y}

			if r == 0 {
				grid[loc] = "."
			} else {
				grid[loc] = "#"
			}
		}
	}

	return grid
}

func SolvePartOne(m map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 19: Part 1")

	grid := GetTractorBeam(m, 50)
	count := 0

	for y := 0; y < 50; y++ {
		o := ""
		for x := 0; x < 50; x++ {
			o += grid[Point{x, y}]
			if grid[Point{x, y}] == "#" {
				count += 1
			}
		}
		fmt.Println(o)
	}
	fmt.Println(count)
}

func SolvePartTwo(m map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 19: Part 1")

	size := 1000
	grid := GetTractorBeam(m, size)

	rightCorner := size - 1
	for y := size - 1; y > 0; y-- {
		leftCorner := 0
		fmt.Println(y)
		for grid[Point{rightCorner, y}] == "." {
			rightCorner--
		}

		for grid[Point{leftCorner, y + 100}] == "." {
			leftCorner++
		}

		if rightCorner-leftCorner == 100 {
			fmt.Println(leftCorner*10000 + y)
			break
		}
	}
}
