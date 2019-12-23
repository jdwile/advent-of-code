package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	cpu "github.com/jdwile/advent-of-code/2019/intcode-cpu"
	"github.com/jdwile/advent-of-code/2019/utils"
)

func main() {
	memory := ReadInput()
	SolvePartOne(memory)
	SolvePartTwo(memory)
}

func InitializeNetwork(NIC map[int]int, n int) []cpu.CPU {
	network := make([]cpu.CPU, n)

	for i := range network {
		m := make(map[int]int)
		for k, v := range NIC {
			m[k] = v
		}

		network[i] = cpu.ConstructCPU(m)
	}

	return network
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

func SolvePartOne(NIC map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 23: Part 1")
	network := InitializeNetwork(NIC, 50)
	nextInputs := make([][]int, len(network))
	outputs := make([][]int, len(network))
	for i := range nextInputs {
		nextInputs[i] = []int{i}
		outputs[i] = make([]int, 0)
	}

	loop := true

	for loop {
		for i := range nextInputs {
			if nextInputs[i][0] == -1 {
				if len(network[i].Input) == 0 {
					network[i].Input = append(network[i].Input, -1)
				}
				continue
			}
			if len(network[i].Input) > 0 && network[i].Input[0] == -1 {
				network[i].Input = make([]int, 0)
			}
			for j := range nextInputs[i] {
				network[i].Input = append(network[i].Input, nextInputs[i][j])
			}
		}

		for i := range nextInputs {
			nextInputs[i] = []int{-1}
		}

		loop = false
		for i := range network {
			loop = loop || !network[i].Halted

			network[i] = network[i].ExecuteProgram(1)

			if len(network[i].Output) > 0 {
				outputs[i] = append(outputs[i], network[i].Output[0])
				network[i].Output = make([]int, 0)
			}
		}

		for i := range outputs {
			if len(outputs[i]) == 3 {
				addr, x, y := outputs[i][0], outputs[i][1], outputs[i][2]
				outputs[i] = make([]int, 0)

				if addr >= len(network) {
					fmt.Println(y)
					loop = false
				} else {
					if nextInputs[addr][0] == -1 {
						nextInputs[addr] = []int{x, y}
					} else {
						nextInputs[addr] = append(nextInputs[addr], x)
						nextInputs[addr] = append(nextInputs[addr], y)
					}
				}
			}
		}
	}
}

type NAC struct {
	X int
	Y int
}

func SolvePartTwo(NIC map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 23: Part 1")
	network := InitializeNetwork(NIC, 50)
	nextInputs := make([][]int, len(network))
	outputs := make([][]int, len(network))

	nac := NAC{math.MaxInt64, math.MaxInt64}
	nacMap := make(map[NAC]bool)

	for i := range nextInputs {
		nextInputs[i] = []int{i}
		outputs[i] = make([]int, 0)
	}

	loop := true

	for loop {
		for i := range nextInputs {
			if nextInputs[i][0] == -1 {
				if len(network[i].Input) == 0 {
					network[i].Input = append(network[i].Input, -1)
				}
				continue
			}
			if len(network[i].Input) > 0 && network[i].Input[0] == -1 {
				network[i].Input = make([]int, 0)
			}
			for j := range nextInputs[i] {
				network[i].Input = append(network[i].Input, nextInputs[i][j])
			}
		}

		for i := range nextInputs {
			nextInputs[i] = []int{-1}
		}

		loop = false
		for i := range network {
			loop = loop || !network[i].Halted

			network[i] = network[i].ExecuteProgram(1)

			if len(network[i].Output) > 0 {
				outputs[i] = append(outputs[i], network[i].Output[0])
				network[i].Output = make([]int, 0)
			}
		}

		for i := range outputs {
			if len(outputs[i]) == 3 {
				addr, x, y := outputs[i][0], outputs[i][1], outputs[i][2]
				outputs[i] = make([]int, 0)

				if addr >= len(network) {
					nac = NAC{x, y}
				} else {
					if nextInputs[addr][0] == -1 {
						nextInputs[addr] = []int{x, y}
					} else {
						nextInputs[addr] = append(nextInputs[addr], x)
						nextInputs[addr] = append(nextInputs[addr], y)
					}
				}
			}
		}

		allIdle := true
		for i := range network {
			allIdle = allIdle && network[i].Idle
		}

		if allIdle && nac.X != math.MaxInt64 {
			if nacMap[nac] {
				fmt.Println(nac.Y)
				loop = false
				break
			}

			nacMap[nac] = true
			network[0].Input = []int{nac.X, nac.Y}
			nac = NAC{math.MaxInt64, math.MaxInt64}
		}
	}
}
