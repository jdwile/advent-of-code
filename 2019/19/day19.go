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

func GetSquare(resp chan<- string, m map[int]int, loc Point) {
	memory := make(map[int]int)
	for e := range m {
		memory[e] = m[e]
	}
	c := cpu.ConstructCPU(memory)

	c.Input = []int{loc.X, loc.Y}

	c = c.ExecuteProgram()

	r := c.Output[0]
	c.Output = []int{}

	if r == 0 {
		resp <- "."
	} else {
		resp <- "#"
	}
}

func GetTractorBeam(m map[int]int, size int) map[Point]string {
	defer utils.TimeTrack(time.Now(), "Get tractor beam")
	grid := make(map[Point]string)

	workerChans := make([][]chan string, 0)

	for y := 0; y < size; y++ {
		chans := make([]chan string, 0)
		for x := 0; x < size; x++ {
			c := make(chan string)
			chans = append(chans, c)
		}

		workerChans = append(workerChans, chans)
	}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			p := Point{x, y}
			go GetSquare(workerChans[y][x], m, p)
		}
	}

	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			p := Point{x, y}
			r := <-workerChans[y][x]
			grid[p] = r
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

	GOAL_SIZE := 100

	rightCorner := 0
	y := 5
	for true {
		cR := make(chan string)
		go GetSquare(cR, m, Point{rightCorner, y})
		rightCornerTile := <-cR

		for rightCornerTile != "#" {
			rightCorner++
			cR = make(chan string)
			go GetSquare(cR, m, Point{rightCorner, y})
			rightCornerTile = <-cR
		}
		for rightCornerTile == "#" {
			rightCorner++
			cR = make(chan string)
			go GetSquare(cR, m, Point{rightCorner, y})
			rightCornerTile = <-cR
		}
		rightCorner--

		leftCorner := rightCorner - (GOAL_SIZE - 1)
		cL := make(chan string)
		go GetSquare(cL, m, Point{leftCorner, y + (GOAL_SIZE - 1)})
		leftCornerTile := <-cL

		if leftCornerTile == "#" {
			fmt.Println(leftCorner*10000 + y)
			break
		}
		y++
	}
}
