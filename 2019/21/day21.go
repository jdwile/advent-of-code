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
	memory := ReadInput()
	SolvePartOne(memory)
	SolvePartTwo(memory)
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

// .................
// @................
// #####.###########
// .ABCD
func SolvePartOne(m map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 21: Part 1")
	memory := make(map[int]int)
	for e := range m {
		memory[e] = m[e]
	}

	c := cpu.ConstructCPU(memory)

	program := []string{
		"OR A J",
		"AND B J",
		"AND C J",
		"NOT J J",
		"AND D J",
		"WALK",
	}
	for _, i := range program {
		for _, r := range i {
			c.Input = append(c.Input, int(r))
		}
		c.Input = append(c.Input, 10)
	}

	c = c.ExecuteProgram()

	// res := ""
	// for i := range c.Output {
	// 	if c.Output[i] == 10 {
	// 		fmt.Println(res)
	// 		res = ""
	// 	} else {
	// 		res += string(rune(c.Output[i]))
	// 	}
	// }
	fmt.Println(c.Output[len(c.Output)-1])
}

// .................
// @................
// #####.###########
// .ABCDEFGH
func SolvePartTwo(m map[int]int) {
	defer utils.TimeTrack(time.Now(), "Day 21: Part 2")
	memory := make(map[int]int)
	for e := range m {
		memory[e] = m[e]
	}

	c := cpu.ConstructCPU(memory)

	program := []string{
		"OR A J",
		"AND B J",
		"AND C J",
		"NOT J J",
		"AND D J",
		"OR E T",
		"OR H T",
		"AND T J",
		"RUN",
	}
	for _, i := range program {
		for _, r := range i {
			c.Input = append(c.Input, int(r))
		}
		c.Input = append(c.Input, 10)
	}

	c = c.ExecuteProgram()

	// res := ""
	// for i := range c.Output {
	// 	if c.Output[i] == 10 {
	// 		fmt.Println(res)
	// 		res = ""
	// 	} else {
	// 		res += string(rune(c.Output[i]))
	// 	}
	// }
	fmt.Println(c.Output[len(c.Output)-1])
}
