package main

import (
	"bufio"
	"fmt"
	"github.com/jdwile/advent-of-code/2019/utils"
	"math"
	"os"
	// "strconv"
	"strings"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func Abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func SortPoints(p []Point, base Point) []Point {
	for i := 0; i < len(p)-1; i++ {
		for j := i + 1; j < len(p); j++ {
			iVal := math.Atan2(p[i].Y-base.Y, p[i].X-base.X)
			jVal := math.Atan2(p[j].Y-base.Y, p[j].X-base.X)

			if jVal < iVal {
				temp := p[i]
				p[i] = p[j]
				p[j] = temp
			} else if jVal == iVal {
				iDist := math.Pow(p[i].X-base.X, 2) + math.Pow(p[i].Y-base.Y, 2)
				jDist := math.Pow(p[j].X-base.X, 1) + math.Pow(p[j].Y-base.Y, 2)

				if jDist < iDist {
					temp := p[i]
					p[i] = p[j]
					p[j] = temp
				}
			}
		}
	}
	return p
}

func main() {
	p := ReadInput()

	station := SolvePartOne(p)
	SolvePartTwo(p, station)
}

func ReadInput() (points []Point) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	y := 0
	for scanner.Scan() {
		x := 0
		for _, c := range strings.Split(scanner.Text(), "") {
			if c == "#" {
				points = append(points, Point{float64(x), float64(y)})
			}
			x += 1
		}
		y += 1
	}

	return points
}

func SolvePartOne(p []Point) Point {
	defer utils.TimeTrack(time.Now(), "Day 10: Part 1")
	m := 0
	var mP Point

	for _, p1 := range p {
		slopes := make(map[float64]bool)
		c := 0
		for _, p2 := range p {
			if p1 == p2 {
				continue
			}

			s := math.Atan2(p2.Y-p1.Y, p2.X-p1.X)
			if !slopes[s] {
				slopes[s] = true
				c += 1
			}
		}
		if c > m {
			m = c
			mP = p1
		}
	}

	fmt.Println(m, mP)
	return mP
}

func PointMeasure(a Point, b Point) float64 {
	return math.Atan2(a.Y-b.Y, a.X-b.X)
}

func SolvePartTwo(p []Point, s Point) {
	defer utils.TimeTrack(time.Now(), "Day 10: Part 2")

	sortedPoints := SortPoints(p, s)

	start := 0
	for math.Abs(PointMeasure(sortedPoints[start], s)-(-math.Pi/2)) > 0.00001 {
		start += 1
	}

	i := start
	removedCount := 0
	lastRemoved := s

	for removedCount < 200 {
		for sortedPoints[i] == lastRemoved || (PointMeasure(sortedPoints[i], s) == PointMeasure(lastRemoved, s)) {
			i = (i + 1) % len(sortedPoints)
		}

		lastRemoved = sortedPoints[i]
		removedCount++

		copy(sortedPoints[i:], sortedPoints[i+1:])
		sortedPoints[len(sortedPoints)-1] = Point{}
		sortedPoints = sortedPoints[:len(sortedPoints)-1]
	}
	fmt.Println(removedCount, lastRemoved)
}
