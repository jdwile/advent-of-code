package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/ernestosuarez/itertools"
	cpu "github.com/jdwile/advent-of-code/2019/intcode-cpu"
	"github.com/jdwile/advent-of-code/2019/utils"
)

func main() {
	solvePart1([]int{1, 0, 4, 3, 2})
	solvePart2([]int{9, 8, 7, 6, 5})
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
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

func solvePart1(input []int) {
	m := readInput()
	defer utils.TimeTrack(time.Now(), "Day 7: Part 1")

	answer := 0
	for v := range itertools.PermutationsInt(input, len(input)) {
		d := 0
		for i := range []string{"A", "B", "C", "D", "E"} {
			c := cpu.ConstructCPU(m)
			c.Input = []int{v[i], d}
			c = c.ExecuteProgram()
			d = c.Output[0]
		}
		answer = Max(answer, d)
	}
	fmt.Println(answer)
}

func solvePart2(input []int) {
	m := readInput()
	defer utils.TimeTrack(time.Now(), "Day 7: Part 2")

	amps := []string{"A", "B", "C", "D", "E"}
	r := 0

	for v := range itertools.PermutationsInt(input, len(input)) {
		var lastE int
		halt := false
		cpus := make([]cpu.CPU, 5)
		for i, j := range v {
			cpus[i] = cpu.ConstructCPU(m)
			if i == 0 {
				cpus[i].Input = []int{j, 0}
			} else {
				cpus[i].Input = []int{j}
			}
		}
		for !halt {
			for i, _ := range amps {
				if !cpus[i].Halted {
					cpus[i] = cpus[i].ExecuteProgram()
					if i == len(amps)-1 && len(cpus[i].Output) > 0 {
						lastE = cpus[i].Output[len(cpus[i].Output)-1]
					}
					for _, j := range cpus[i].Output {
						cpus[(i+1)%len(cpus)].Input = append(cpus[(i+1)%len(cpus)].Input, j)
					}
					cpus[i].Output = []int{}
				}
			}
			halt = true
			for _, cpu := range cpus {
				halt = halt && cpu.Halted
			}
		}
		r = Max(r, lastE)
	}
	fmt.Println(r)
}
