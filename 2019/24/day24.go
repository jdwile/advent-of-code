package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
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

func main() {
	grid := ReadInput()
	// SolvePartOne(grid)
	SolvePartTwo(grid)
}

func ReadInput() map[Point]string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	grid := make(map[Point]string)

	y := 0
	for scanner.Scan() {
		row := strings.Split(scanner.Text(), "")
		for x := range row {
			grid[Point{x, y}] = row[x]
		}
		y++
	}

	return grid
}

func getAdjacentBugs(grid map[Point]string, loc Point) int {
	DXS := []int{-1, 0, 1}
	DYS := []int{-1, 0, 1}

	numAdjBugs := 0
	for _, dx := range DXS {
		for _, dy := range DYS {
			if Abs(dx)+Abs(dy) != 1 {
				continue
			}

			if grid[Point{loc.X + dx, loc.Y + dy}] == "#" {
				numAdjBugs++
			}
		}
	}

	return numAdjBugs
}

func getInitGrid() map[Point]string {
	grid := make(map[Point]string)
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			grid[Point{x, y}] = "."
		}
	}
	grid[Point{2, 2}] = "?"
	return grid
}

func getBugsOnEdge(grid map[Point]string, dx, dy int) int {
	numBugs := 0
	x, y := 0, 0

	if dx == -1 || dy == -1 {
		x = 5
		y = 5
	} else if dx == 1 || dy == 1 {
		x = 0
		y = 5
	}

	for x >= 0 && y >= 0 {
		if grid[Point{x, y}] == "#" {
			numBugs++
		}
		x -= Abs(dy)
		y -= Abs(dx)
	}

	return numBugs
}

func getAdjacentBugs3D(space map[int]map[Point]string, layer int, loc Point) int {
	DXS := []int{-1, 0, 1}
	DYS := []int{-1, 0, 1}

	numAdjBugs := 0
	for _, dx := range DXS {
		for _, dy := range DYS {
			if Abs(dx)+Abs(dy) != 1 {
				continue
			}

			cur := Point{loc.X + dx, loc.Y + dy}

			if space[layer][cur] == "#" {
				numAdjBugs++
			} else if space[layer][cur] == "?" {
				numAdjBugs += getBugsOnEdge(space[layer-1], dx, dy)
			} else if space[layer][cur] == "" {
				sub := Point{2 + dx, 2 + dy}

				// fmt.Println("Checking upper layer", layer, cur, layer+1, sub)

				if space[layer+1][sub] == "#" {
					numAdjBugs++
				}
			}
		}
	}

	return numAdjBugs
}

func getHash(grid map[Point]string) string {
	hash := ""

	for y := 0; y < 5; y++ {
		row := ""
		for x := 0; x < 5; x++ {
			row += grid[Point{x, y}]
		}
		hash += row
	}

	return hash
}

func SolvePartOne(g map[Point]string) {
	cache := make(map[string]bool)

	grid := make(map[Point]string)
	for k, v := range g {
		grid[k] = v
	}

	for true {
		// fmt.Println("\u001b[2J")
		// for y := 0; y < 5; y++ {
		// 	res := ""
		// 	for x := 0; x < 5; x++ {
		// 		res += grid[Point{x, y}]
		// 	}
		// 	fmt.Println(res)
		// }

		hash := getHash(grid)
		if cache[hash] {
			fmt.Println("FIRST REPEAT")
			break
		} else {
			cache[hash] = true
		}

		nextGrid := make(map[Point]string)
		for k, v := range grid {
			numAdjBugs := getAdjacentBugs(grid, k)
			if v == "#" {
				if numAdjBugs != 1 {
					nextGrid[k] = "."
				} else {
					nextGrid[k] = "#"
				}
			} else {
				if numAdjBugs == 1 || numAdjBugs == 2 {
					nextGrid[k] = "#"
				} else {
					nextGrid[k] = "."
				}
			}
		}

		grid = nextGrid
	}

	rating := 0

	i := 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if grid[Point{x, y}] == "#" {
				rating += int(math.Pow(2, float64(i)))
			}

			i++
		}
	}

	fmt.Println(rating)
}

func hasBugs(grid map[Point]string) bool {
	for _, v := range grid {
		if v == "#" {
			return true
		}
	}
	return false
}

func SolvePartTwo(g map[Point]string) {
	space := make(map[int]map[Point]string)

	grid := make(map[Point]string)
	for k, v := range g {
		grid[k] = v
	}
	grid[Point{2, 2}] = "?"

	space[0] = grid

	for age := 1; age <= 3; age++ {

		// for layer := -2; layer <= 2; layer++ {
		// 	fmt.Println("Depth:", layer) //"\u001b[2J")
		for y := 0; y < 5; y++ {
			res := ""
			for x := 0; x < 5; x++ {
				res += space[0][Point{x, y}]
			}
			fmt.Println(res)
		}
		// }
		fmt.Println()
		nextSpace := make(map[int]map[Point]string)

		space[age] = getInitGrid()
		space[-age] = getInitGrid()

		for layer := 0; layer < 1; layer++ {
			nextSpace[layer] = space[layer]
			if Abs(layer) < age {
				for k, v := range space[layer] {
					numAdjBugs := getAdjacentBugs3D(space, layer, k)

					if v == "#" {
						if numAdjBugs != 1 {
							nextSpace[layer][k] = "."
						} else {
							nextSpace[layer][k] = "#"
						}
					} else if v == "." {
						if numAdjBugs == 1 || numAdjBugs == 2 {
							nextSpace[layer][k] = "#"
						} else {
							nextSpace[layer][k] = "."
						}
					}
				}
			}
		}

		space = nextSpace
	}

}
