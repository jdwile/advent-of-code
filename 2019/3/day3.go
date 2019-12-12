package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jdwile/advent-of-code/2019/utils"
)

type Point struct {
	X int
	Y int
}

type WirePoint struct {
	p Point
	s int
}

type Line struct {
	start WirePoint
	end   WirePoint
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func Abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}

func ManhattanDistance(x, y int) int {
	return Abs(x) + Abs(y)
}

func SignalDelayDistance(l1 Line, l2 Line, x int, y int) int {
	return Abs(l1.start.p.X-x) + Abs(l1.start.p.Y-y) + l1.start.s + Abs(l2.start.p.X-x) + Abs(l2.start.p.Y-y) + l2.start.s
}

func main() {
	paths := ReadInput()
	lines := ConvertPathsToLines(paths)
	SolvePartOne(lines)
	SolvePartTwo(lines)
}

func ReadInput() (paths [][]string) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; i < 2; i++ {
		scanner.Scan()
		paths = append(paths, strings.Split(scanner.Text(), ","))
	}

	return paths
}

func ConvertPathsToLines(paths [][]string) (lines [][]Line) {
	lines = make([][]Line, len(paths))

	for wire, path := range paths {
		lines[wire] = make([]Line, 0)
		start := WirePoint{Point{0, 0}, 0}
		var end WirePoint
		for _, i := range path {
			d := i[0]
			n, _ := strconv.Atoi(i[1:])
			switch d {
			case 'R':
				end = WirePoint{Point{start.p.X + n, start.p.Y}, start.s + n}
			case 'U':
				end = WirePoint{Point{start.p.X, start.p.Y + n}, start.s + n}
			case 'L':
				end = WirePoint{Point{start.p.X - n, start.p.Y}, start.s + n}
			case 'D':
				end = WirePoint{Point{start.p.X, start.p.Y - n}, start.s + n}
			}
			lines[wire] = append(lines[wire], Line{start, end})
			start = end
		}
	}

	return lines
}

func FindIntersection(a Point, b Point, c Point, d Point) Point {
	a1 := b.Y - a.Y
	b1 := a.X - b.X
	c1 := a1*a.X + b1*a.Y

	a2 := d.Y - c.Y
	b2 := c.X - d.X
	c2 := a2*c.X + b2*c.Y

	det := a1*b2 - a2*b1

	if det == 0 {
		return Point{math.MaxInt64, math.MaxInt64}
	}

	x := (b2*c1 - b1*c2) / det
	y := (a1*c2 - a2*c1) / det

	if (Min(a.X, b.X) <= x && x <= Max(a.X, b.X) && Min(a.Y, b.Y) <= y && y <= Max(a.Y, b.Y)) && (Min(c.X, d.X) <= x && x <= Max(c.X, d.X) && Min(c.Y, d.Y) <= y && y <= Max(c.Y, d.Y)) {
		return Point{x, y}
	}

	return Point{math.MaxInt64, math.MaxInt64}
}

func SolvePartOne(lines [][]Line) {
	defer utils.TimeTrack(time.Now(), "Day 3: Part 1")
	intersections := make([]int, 0)
	for _, line1 := range lines[0] {
		for _, line2 := range lines[1] {
			origin := Point{0, 0}
			if line1.start.p != origin && line1.end.p != origin && line2.start.p != origin && line2.end.p != origin {
				i := FindIntersection(line1.start.p, line1.end.p, line2.start.p, line2.end.p)
				if i.X == math.MaxInt64 {
					continue
				}

				intersections = append(intersections, ManhattanDistance(i.X, i.Y))
			}
		}
	}

	i := math.MaxInt64
	for _, intersection := range intersections {
		i = Min(i, intersection)
	}

	fmt.Println("Closest Intersection (Manhattan):", i)
}

func SolvePartTwo(lines [][]Line) {
	defer utils.TimeTrack(time.Now(), "Day 3: Part 2")
	intersections := make([]int, 0)
	for _, line1 := range lines[0] {
		for _, line2 := range lines[1] {
			origin := Point{0, 0}
			if line1.start.p != origin && line1.end.p != origin && line2.start.p != origin && line2.end.p != origin { // Parallel check
				i := FindIntersection(line1.start.p, line1.end.p, line2.start.p, line2.end.p)
				if i.X == math.MaxInt64 {
					continue
				}

				intersections = append(intersections, SignalDelayDistance(line1, line2, i.X, i.Y))
			}
		}
	}

	i := math.MaxInt64
	for _, intersection := range intersections {
		i = Min(i, intersection)
	}

	fmt.Println("Closest Intersection (Signal Delay):", i)
}
