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

func Abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func main() {
	memory := ReadInput()
	c := SolvePartOne(memory)
	SolvePartTwo(c)
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

func getDir(dx, dy int) int {
	if dx == -1 {
		return 3
	} else if dx == 1 {
		return 4
	} else if dy == -1 {
		return 2
	}
	return 1
}

func DFSSearch(depth int, c cpu.CPU, pos Point, visited map[Point]bool, marks map[Point]string) (cpu.CPU, bool, map[Point]bool, map[Point]string) {

	// time.Sleep(250 * time.Millisecond)
	dxys := []Point{Point{}, Point{0, 1}, Point{0, -1}, Point{-1, 0}, Point{1, 0}}

	if visited[pos] {
		return c, false, visited, marks
	}

	marks[pos] = " "
	visited[pos] = true

	// Test all directions for valid nodes
	for d := 1; d <= 4; d++ {
		dxy := dxys[d]
		newPos := Point{pos.X + dxy.X, pos.Y + dxy.Y}

		if visited[newPos] {
			continue
		}

		// Navigate to new point
		c.Input = []int{d}
		c = c.ExecuteProgram()
		res := c.Output[0]
		c.Output = make([]int, 0)

		if res == 2 {
			fmt.Println("\u001b[2J")
			fmt.Println("FOUND OXYGEN TANK")
			for i := -25; i < 25; i++ {
				res := ""
				for j := -25; j < 25; j++ {
					c := Point{i, j}
					if (c == Point{0, 0}) {
						res += "X"
					} else if c == pos {
						res += "O"
					} else if len(marks[c]) > 0 {
						res += marks[c]
					} else {
						res += "□"
					}
				}
				fmt.Println(res)
			}
			fmt.Println(depth + 1)
			return c, true, visited, marks
		}

		if res == 0 {
			marks[newPos] = "▓"
			visited[newPos] = true
			continue
		}

		// Recurse from point
		var found bool
		_, found, visited, marks = DFSSearch(depth+1, c, newPos, visited, marks)

		if found {
			return c, true, visited, marks
		}

		// Navigate back to current position
		dxy.X = dxy.X * -1
		dxy.Y = dxy.Y * -1
		c.Input = []int{getDir(dxy.X, dxy.Y)}
		c = c.ExecuteProgram()
		c.Output = make([]int, 0)
	}
	return c, false, visited, marks
}

func DFSDepthCount(depth int, c cpu.CPU, pos Point, visited map[Point]bool) (cpu.CPU, int, map[Point]bool) {
	maxDepth := depth

	dxys := []Point{Point{}, Point{0, 1}, Point{0, -1}, Point{-1, 0}, Point{1, 0}}

	if visited[pos] {
		return c, depth, visited
	}

	visited[pos] = true

	// Test all directions for valid nodes
	for d := 1; d <= 4; d++ {
		dxy := dxys[d]
		newPos := Point{pos.X + dxy.X, pos.Y + dxy.Y}

		if visited[newPos] {
			continue
		}

		// Navigate to new point
		c.Input = []int{d}
		c = c.ExecuteProgram()
		res := c.Output[0]
		c.Output = make([]int, 0)

		if res == 0 || res == 2 {
			continue
		}

		// Recurse from point
		var d int
		c, d, visited = DFSDepthCount(depth+1, c, newPos, visited)

		maxDepth = Max(maxDepth, d)

		// Navigate back to current position
		dxy.X = dxy.X * -1
		dxy.Y = dxy.Y * -1
		c.Input = []int{getDir(dxy.X, dxy.Y)}
		c = c.ExecuteProgram()
		c.Output = make([]int, 0)
	}

	if depth == 0 {

	}

	return c, maxDepth, visited
}

func SolvePartOne(m map[int]int) cpu.CPU {
	defer utils.TimeTrack(time.Now(), "Day 15: Part 1")
	memory := make(map[int]int)
	for e := range m {
		memory[e] = m[e]
	}

	c := cpu.ConstructCPU(memory)

	marks := make(map[Point]string)
	visited := make(map[Point]bool)

	c, _, _, _ = DFSSearch(0, c, Point{0, 0}, visited, marks)
	return c
}

func SolvePartTwo(c cpu.CPU) {
	defer utils.TimeTrack(time.Now(), "Day 15: Part 2")

	visited := make(map[Point]bool)

	var depth int
	_, depth, _ = DFSDepthCount(0, c, Point{0, 0}, visited)
	fmt.Println(depth)
}
