package main

import (
	"bufio"
	"fmt"
	"github.com/jdwile/advent-of-code/2019/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

type Image struct {
	W      int
	H      int
	Pixels [][]int
}

func main() {
	i := readInput()
	j := parseImage(i, 25, 6)
	solvePartOne(j)
	solvePartTwo(j)
}

func readInput() (nums []int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	arr := strings.Split(scanner.Text(), "")

	for _, n := range arr {
		i, _ := strconv.Atoi(n)
		nums = append(nums, i)
	}

	return nums
}

func parseImage(a []int, w, h int) (layers []Image) {
	var c [][]int
	i := 0

	for i < len(a) {
		c = make([][]int, h)
		for j := 0; j < h; j++ {
			c[j] = make([]int, w)
			for k := 0; k < w; k++ {
				c[j][k] = a[i]
				i++
			}
		}
		layers = append(layers, Image{w, h, c})
	}

	return layers
}

func countDigits(a [][]int, d int) (c int) {
	for i := range a {
		for j := range a[i] {
			if a[i][j] == d {
				c++
			}
		}
	}
	return c
}

func solvePartOne(l []Image) {
	defer utils.TimeTrack(time.Now(), "Day 8: Part 1")

	m := -1
	i := 0

	for j, p := range l {
		c := countDigits(p.Pixels, 0)
		if m < 0 || c < m {
			m = c
			i = j
		}
	}

	fmt.Println(countDigits(l[i].Pixels, 1) * countDigits(l[i].Pixels, 2))
}

func solvePartTwo(l []Image) {
	defer utils.TimeTrack(time.Now(), "Day 8: Part 2")

	o := l[0].Pixels

	for _, i := range l {
		for j := 0; j < i.H; j++ {
			for k := 0; k < i.W; k++ {
				if o[j][k] == 2 {
					o[j][k] = i.Pixels[j][k]
				}
			}
		}
	}

	for j := range o {
		var a string
		for k := range o[j] {
			if o[j][k] == 0 {
				a += " "
			}
			if o[j][k] == 1 {
				a += "."
			}
		}
		fmt.Println(a)
	}
}
