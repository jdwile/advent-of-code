package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jdwile/advent-of-code/2017/utils"
)

func main() {
	captcha := readInputFile()
	fmt.Printf("Part 1 Answer: %d\n\n", SolvePartOne(captcha))
	fmt.Printf("Part 2 Answer: %d\n\n", SolvePartTwo(captcha))
}

func readInputFile() string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	return scanner.Text()
}

func SolvePartOne(captcha string) int {
	defer utils.TimeTrack(time.Now(), "Day 1: Part 1")

	var res int
	a := strings.Split(captcha, "")

	for i := range a {
		if a[i] == a[(i+1)%len(a)] {
			n, _ := strconv.Atoi(a[i])
			res += n
		}
	}

	return res
}

func SolvePartTwo(captcha string) int {
	defer utils.TimeTrack(time.Now(), "Day 1: Part 2")

	var res int
	a := strings.Split(captcha, "")

	for i := range a {
		if a[i] == a[(i+len(a)/2)%len(a)] {
			n, _ := strconv.Atoi(a[i])
			res += n
		}
	}

	return res
}
