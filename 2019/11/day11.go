package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	. "github.com/jdwile/advent-of-code/2019/intcode-cpu"
	"github.com/jdwile/advent-of-code/2019/utils"
)

type Point struct {
	X int
	Y int
}

func main() {
	memory := ReadInput()
	SolvePart1(memory)
	SolvePart2(memory)
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

func Paint(g map[Point]rune, c CPU, robot Point, startingColor int) map[Point]rune {
	dirs := []string{"UP", "RIGHT", "DOWN", "LEFT"}
	dxys := []Point{Point{0, 1}, Point{1, 0}, Point{0, -1}, Point{-1, 0}}

	dir := 0
	startFlag := true

	for true {
		if startFlag {
			c.Input = []int{startingColor}
			startFlag = false
		} else {
			if g[robot] == 'W' {
				c.Input = []int{1}
			} else {
				c.Input = []int{0}
			}
		}

		c = c.ExecuteProgram()

		if len(c.Output) == 0 {
			break
		}

		color := c.Output[0]
		turn := c.Output[1]
		c.Output = make([]int, 0)

		if color == 0 {
			g[robot] = 'B'
		} else {
			g[robot] = 'W'
		}

		if turn == 0 {
			dir = (dir - 1)
			if dir < 0 {
				dir += len(dirs)
			}
		} else {
			dir = (dir + 1) % len(dirs)
		}

		robot = Point{robot.X + dxys[dir].X, robot.Y + dxys[dir].Y}
	}

	return g
}

func SolvePart1(m map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 11: Part 1")
	memory := make(map[int]int)
	for e := range m {
		memory[e] = m[e]
	}

	g := make(map[Point]rune)
	c := ConstructCPU(memory)
	robot := Point{0, 0}

	g = Paint(g, c, robot, 0)

	fmt.Println(len(g))
}

func SolvePart2(m map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 11: Part 2")
	memory := make(map[int]int)
	for e := range m {
		memory[e] = m[e]
	}

	g := make(map[Point]rune)
	c := ConstructCPU(memory)
	robot := Point{0, 0}

	g = Paint(g, c, robot, 1)

	res := make([][]string, 11)
	for i := range res {
		res[i] = make([]string, 40)
		for j := range res[i] {
			res[i][j] = " "
		}
	}

	for i := range g {
		if g[i] == 'W' {
			res[i.Y+10][i.X] = "▓"
		}
	}

	for i := range res {
		o := ""
		for j := range res[i] {
			o += res[len(res)-i-1][j]
		}
		fmt.Println(o)
	}
}
