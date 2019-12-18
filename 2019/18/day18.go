package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"time"

	"github.com/jdwile/advent-of-code/2019/utils"
)

type Point struct {
	X int
	Y int
}

func main() {
	g, pos := ReadInput()
	SolvePartOne(g, pos)
	// SolvePartTwo(i)
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

func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func (p Point) toString() string {
	return "{" + strconv.Itoa(p.X) + ", " + strconv.Itoa(p.Y) + "}"
}

func PrintGrid(g map[Point]string) {
	maxX := 0
	maxY := 0

	for p := range g {
		maxX = Max(maxX, p.X)
		maxY = Max(maxY, p.Y)
	}

	for y := 0; y <= maxY; y++ {
		res := ""
		for x := 0; x <= maxX; x++ {
			res += g[Point{x, y}]
		}
		fmt.Println(res)
	}
}

func ReadInput() (map[Point]string, map[string]Point) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	grid := make(map[Point]string)

	scanner := bufio.NewScanner(file)

	p := make(map[string]Point)

	j := 0
	for scanner.Scan() {
		l := scanner.Text()
		for i := range l {
			if string(l[i]) != "." && string(l[i]) != "#" {
				p[string(l[i])] = Point{i, j}
			}
			grid[Point{i, j}] = string(l[i])
		}
		j += 1
	}
	return grid, p
}

func GetMin(Q []Point, dist map[Point]int) (Point, []Point) {
	minPoint := Q[0]

	o := make([]Point, 0)

	for i := 1; i < len(Q); i++ {
		if dist[Q[i]] < dist[minPoint] {
			o = append(o, minPoint)
			minPoint = Q[i]
		} else {
			o = append(o, Q[i])
		}
	}

	return minPoint, o
}

func ShortestPathIgnoringDoors(g map[Point]string, start, target Point) []Point {
	dist := make(map[Point]int)
	prev := make(map[Point]Point)
	visited := make(map[Point]bool)

	Q := make([]Point, 0)

	for c := range g {
		dist[c] = math.MaxInt64
		prev[c] = Point{math.MaxInt64, math.MaxInt64}

		if c == start {
			Q = append(Q, c)
		}
	}

	dist[start] = 0

	for len(Q) > 0 {
		c := Q[0]

		if len(Q) == 1 {
			Q = make([]Point, 0)
		} else {
			Q = Q[1:]
		}

		visited[c] = true

		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				x := c.X + dx
				y := c.Y + dy

				neighbor := Point{x, y}

				if g[neighbor] == "" {
					continue
				}

				if dx == 0 && dy == 0 || g[neighbor] == "#" {
					continue
				}

				if Abs(dx)+Abs(dy) == 2 {
					continue
				}

				d := dist[c] + 1
				if d < dist[neighbor] {
					dist[neighbor] = d
					prev[neighbor] = c
				}

				if !visited[neighbor] {
					Q = append(Q, neighbor)
				}
			}
		}
	}

	if dist[target] == math.MaxInt64 {
		return make([]Point, 0)
	}

	c := target

	path := make([]Point, 0)
	for c != start {
		newPath := make([]Point, len(path)+1)

		newPath[0] = c
		for i, p := range path {
			newPath[i+1] = p
		}
		path = newPath
		c = prev[c]
	}

	return path
}

func ShortestPath(g map[Point]string, start, target Point) []Point {
	dist := make(map[Point]int)
	prev := make(map[Point]Point)
	visited := make(map[Point]bool)

	Q := make([]Point, 0)

	for c := range g {
		dist[c] = math.MaxInt64
		prev[c] = Point{math.MaxInt64, math.MaxInt64}

		if c == start {
			Q = append(Q, c)
		}
	}

	dist[start] = 0

	for len(Q) > 0 {
		c := Q[0]

		if len(Q) == 1 {
			Q = make([]Point, 0)
		} else {
			Q = Q[1:]
		}

		visited[c] = true

		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				x := c.X + dx
				y := c.Y + dy

				neighbor := Point{x, y}

				if g[neighbor] == "" {
					continue
				}

				if dx == 0 && dy == 0 || g[neighbor] == "#" || (g[neighbor][0] >= 'A' && g[neighbor][0] <= 'Z') {
					continue
				}

				if Abs(dx)+Abs(dy) == 2 {
					continue
				}

				d := dist[c] + 1
				if d < dist[neighbor] {
					dist[neighbor] = d
					prev[neighbor] = c
				}

				if !visited[neighbor] {
					Q = append(Q, neighbor)
				}
			}
		}
	}

	if dist[target] == math.MaxInt64 {
		return make([]Point, 0)
	}

	c := target

	path := make([]Point, 0)
	for c != start {
		newPath := make([]Point, len(path)+1)

		newPath[0] = c
		for i, p := range path {
			newPath[i+1] = p
		}
		path = newPath
		c = prev[c]
	}

	return path
}

func getAvailableKeys(g map[Point]string, start Point) map[string]Point {
	Q := []Point{start}
	visited := make(map[Point]bool)
	keys := make(map[string]Point)

	for len(Q) > 0 {
		c := Q[0]

		if len(Q) == 1 {
			Q = make([]Point, 0)
		} else {
			Q = Q[1:]
		}

		visited[c] = true

		if g[c][0] >= 'a' && g[c][0] <= 'z' {
			keys[g[c]] = c
		}

		for dy := -1; dy <= 1; dy++ {
			for dx := -1; dx <= 1; dx++ {
				x := c.X + dx
				y := c.Y + dy

				np := Point{x, y}

				if g[np] == "" {
					continue
				}

				if dx == 0 && dy == 0 || g[np] == "#" || (g[np][0] >= 'A' && g[np][0] <= 'Z') {
					continue
				}

				if Abs(dx)+Abs(dy) == 2 {
					continue
				}

				if !visited[np] {
					Q = append(Q, np)
				}
			}
		}
	}

	return keys
}

func GetShortestPaths(g map[Point]string, p map[string]Point) map[string]map[string][]Point {
	SHORTEST_PATHS := make(map[string]map[string][]Point)

	for p1 := range p {
		p1Paths := make(map[string][]Point)
		for p2 := range p {
			if p1 == p2 {
				continue
			}

			p1Paths[p2] = ShortestPathIgnoringDoors(g, p[p1], p[p2])
		}
		SHORTEST_PATHS[p1] = p1Paths
	}

	return SHORTEST_PATHS
}

func CollectAllKeys(respond chan<- int, cur Point, g map[Point]string, p map[string]Point, SHORTEST_PATHS map[string]map[string][]Point) {
	newGrid := make(map[Point]string)
	for p, v := range g {
		newGrid[p] = v
	}
	newGrid[cur] = "."

	keys := getAvailableKeys(newGrid, cur)

	// PrintGrid(g)
	// fmt.Println(g[cur], cur.toString(), "collecting keys", keys)
	var newGrids []map[Point]string
	var newPoints []map[string]Point
	var moves []int
	var keyChans []chan int
	var locs []Point

	for key, loc := range keys {
		keyChan := make(chan int)
		keyChans = append(keyChans, keyChan)
		np := make(map[string]Point)
		ng := make(map[Point]string)

		door := string(key[0] - 32)

		for p, v := range g {
			ng[p] = v
		}
		for p, v := range p {
			if p == key || p == door {
				continue
			}
			np[p] = v
		}

		ng[p[door]] = "."
		ng[cur] = "."
		np[g[cur]] = loc

		// fmt.Println("Move from", g[cur], "to", key, ":", len(SHORTEST_PATHS[g[cur]][key]))
		moves = append(moves, len(SHORTEST_PATHS[g[cur]][key]))
		newPoints = append(newPoints, np)
		newGrids = append(newGrids, ng)
		locs = append(locs, loc)
	}

	for i := range keyChans {
		go CollectAllKeys(keyChans[i], locs[i], newGrids[i], newPoints[i], SHORTEST_PATHS)
	}

	minDist := 0
	for i := range keyChans {
		d := <-keyChans[i]
		d += moves[i]
		if minDist == 0 {
			minDist = d
		}
		minDist = Min(minDist, d)
	}

	// fmt.Println("Responding with:", minDist)

	respond <- minDist
}

func SolvePartOne(g map[Point]string, p map[string]Point) {
	defer utils.TimeTrack(time.Now(), "Day 18: Part 1")
	fmt.Println("getting shortest paths..")
	SHORTEST_PATHS := GetShortestPaths(g, p)
	fmt.Println("Got shortest paths! getting distance from key collection..")

	// PrintGrid(g)
	// fmt.Println("AVAILABLE KEYS", getAvailableKeys(g, p["@"]))
	// fmt.Println(ShortestPath(g, p["@"], p["b"]))
	respond := make(chan int)
	go CollectAllKeys(respond, p["@"], g, p, SHORTEST_PATHS)

	dist := <-respond
	fmt.Println("DISTANCE:", dist)

}
