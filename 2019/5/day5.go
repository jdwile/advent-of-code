package main

import (
	"bufio"
	"fmt"
	. "github.com/jdwile/advent-of-code/2019/intcode-cpu"
	"github.com/jdwile/advent-of-code/2019/utils"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	solvePartOne()
	solvePartTwo()
}

func readInput() map[int]int {
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

func solvePartOne() {
	m := readInput()
	defer utils.TimeTrack(time.Now(), "Day 5: Part 1")

	c := ConstructCPU(m)
	c.Input = []int{1}
	c = c.ExecuteProgram()
	fmt.Println(c.Output[len(c.Output)-1])
}

func solvePartTwo() {
	m := readInput()
	defer utils.TimeTrack(time.Now(), "Day 5: Part 2")

	c := ConstructCPU(m)
	c.Input = []int{5}
	c = c.ExecuteProgram()
	fmt.Println(c.Output[len(c.Output)-1])
}
