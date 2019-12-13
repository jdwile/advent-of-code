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

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
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

func Paint(a []int, g map[Point]int) (score int) {

	for i := 0; i < len(a)-2; i += 3 {
		x := a[i]
		y := a[i+1]
		t := a[i+2]

		if x == -1 && y == 0 {
			score = t
			continue
		}

		g[Point{x, y}] = t
	}

	// xMax := 34
	// yMax := 24
	// time.Sleep(20 * time.Millisecond)
	// fmt.Println("\u001b[H")
	// for y := 0; y <= yMax; y++ {
	// 	res := ""
	// 	for x := 0; x <= xMax; x++ {
	// 		switch g[Point{x, y}] {
	// 		case 0:
	// 			res += " "
	// 		case 1:
	// 			res += "\033[1;35m▓\033[0m"
	// 		case 2:
	// 			res += "\033[1;34m□\033[0m"
	// 		case 3:
	// 			res += "\033[1;36m=\033[0m"
	// 		case 4:
	// 			res += "\033[1;33mO\033[0m"
	// 		}
	// 	}
	// 	fmt.Println(res)
	// }

	return score
}

func CountBlocks(g map[Point]int) (blocks int) {
	for _, v := range g {
		if v == 2 {
			blocks++
		}
	}
	return blocks
}

func SolvePartOne(m map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 13: Part 1")
	memory := make(map[int]int)
	for e := range m {
		memory[e] = m[e]
	}

	g := make(map[Point]int)

	c := cpu.ConstructCPU(memory)
	c = c.ExecuteProgram()

	Paint(c.Output, g)

	fmt.Println(CountBlocks(g))
}

func SolvePartTwo(m map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 13: Part 2")
	memory := make(map[int]int)
	for e := range m {
		memory[e] = m[e]
	}
	memory[0] = 2

	var score int
	g := make(map[Point]int)

	c := cpu.ConstructCPU(memory)
	c = c.ExecuteProgram()

	// fmt.Println("\u001b[2J")
	score = Paint(c.Output, g)
	c.Output = make([]int, 0)

	for CountBlocks(g) > 0 {
		var paddle, ball Point
		for p, v := range g {
			if v == 3 {
				paddle = p
			}
			if v == 4 {
				ball = p
			}
		}

		var dir int
		if paddle.X < ball.X {
			dir = 1
		} else if paddle.X > ball.X {
			dir = -1
		}

		c.Input = []int{dir}
		c = c.ExecuteProgram()
		score = Max(score, Paint(c.Output, g))
		c.Output = make([]int, 0)
	}

	fmt.Println(score)
}
