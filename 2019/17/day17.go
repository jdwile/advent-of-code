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

func getGrid(a []int) [][]string {
	grid := make([][]string, 51)

	row := 0
	grid[row] = make([]string, 0)
	for i := range a {
		if a[i] == 10 {
			if len(grid[row]) > 0 {
				row += 1
				grid[row] = make([]string, 0)
				continue
			}
		}

		grid[row] = append(grid[row], string(rune(a[i])))
	}

	return grid
}

func SolvePartOne(m map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 17: Part 1")
	memory := make(map[int]int)
	for e := range m {
		memory[e] = m[e]
	}

	c := cpu.ConstructCPU(memory)

	c = c.ExecuteProgram()

	o := c.Output
	c.Output = make([]int, 0)

	g := getGrid(o)

	sum := 0

	for row := 1; row < len(g)-1; row++ {
		for col := 1; col < len(g[row])-1; col++ {
			if g[row][col] != "#" {
				continue
			}

			if len(g[row+1]) > col && g[row+1][col] != "#" || len(g[row-1]) > col && g[row-1][col] != "#" {
				continue
			}

			if g[row][col+1] != "#" || g[row][col-1] != "#" {
				continue
			}

			g[row][col] = "O"
			fmt.Println(col, row)
			sum += row * col
		}
	}

	for _, row := range g {
		fmt.Println(strings.Join(row, ""))
	}

	fmt.Println(sum)
}

func SolvePartTwo(m map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 17: Part 2")
	memory := make(map[int]int)
	for e := range m {
		memory[e] = m[e]
	}

	memory[0] = 2

	c := cpu.ConstructCPU(memory)
	var program [4]string
	program[0] = "A,B,A,C,B,A,C,A,C,B\n"
	program[1] = "L,12,L,8,L,8\n"
	program[2] = "L,12,R,4,L,12,R,6\n"
	program[3] = "R,4,L,12,L,12,R,6\n"

	for _, l := range program {
		for _, r := range l {
			c.Input = append(c.Input, int(r))
		}
	}

	c.Input = append(c.Input, int('n'))
	c.Input = append(c.Input, int('\n'))

	c = c.ExecuteProgram()
	fmt.Println(c.Output[len(c.Output)-1])
}
