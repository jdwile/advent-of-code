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

func main() {
	memory := readInput()
	solvePart1(memory)
	solvePart2(memory)
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

func solvePart1(m map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 9: Part 1")

	c := cpu.ConstructCPU(m)
	c.Input = []int{1}

	c = c.ExecuteProgram()
	fmt.Println(c.Output)
}

func solvePart2(m map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 9: Part 2")

	c := cpu.ConstructCPU(m)
	c.Input = []int{2}

	c = c.ExecuteProgram()
	fmt.Println(c.Output)
}
