package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	cpu "github.com/jdwile/advent-of-code/2019/intcode-cpu"
)

func main() {
	memory := ReadInput()
	SolvePartOne(memory)
	// SolvePartTwo(memory)
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

func SolvePartOne(m map[int]int) {
	memory := make(map[int]int)
	for k, v := range m {
		memory[k] = v
	}

	c := cpu.ConstructCPU(memory)
	c = c.ExecuteProgram()

	fmt.Print("\u001b[2J")
	res := ""
	for i := range c.Output {
		res += string(rune(c.Output[i]))
	}
	fmt.Println(res + "\n")
	c.Output = []int{}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)

		if text == "exit" {
			break
		}

		for i := range text {
			c.Input = append(c.Input, int(text[i]))
		}
		c.Input = append(c.Input, '\n')

		c = c.ExecuteProgram()

		res := ""
		for i := range c.Output {
			res += string(rune(c.Output[i]))
		}
		fmt.Println(res + "\n")
		c.Output = []int{}
	}
}
