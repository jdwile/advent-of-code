package main

import (
	"bufio"
	"fmt"
	"github.com/jdwile/advent-of-code/2019/utils"
	"github.com/ernestosuarez/itertools"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	instructions := readInput()
	// solvePart1(instructions, []int{1,0,4,3,2})
	solvePart2(instructions, []int{9,8,7,6,5})
}

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func chooseMode(mode bool, i int, arr []int) int {
	if mode {
		return i
	}
	return arr[i]
}

func executeProgram(instructions []int, input []int, def int) ([]int, []int) {
	// fmt.Println("Executing Program", instructions, input, def)
	output := []int{}
	i := 0
	loop := true
	for loop {
		n := instructions[i]
		ji := false
		ki := false

		if n > 99 {
			s := strconv.Itoa(n)
			if len(s) == 3 {
				s = "0" + s
			}

			n = n % 100
			ji = s[1] == '1'
			ki = s[0] == '1'
		}

		switch n {
		case 1: // add
			j := instructions[i+1]
			k := instructions[i+2]
			l := instructions[i+3]
			instructions[l] = chooseMode(ji, j, instructions) + chooseMode(ki, k, instructions)
			i += 4
		case 2: // multiply
			j := instructions[i+1]
			k := instructions[i+2]
			l := instructions[i+3]
			instructions[l] = chooseMode(ji, j, instructions) * chooseMode(ki, k, instructions)
			i += 4
		case 3: // input
			j := instructions[i+1]
			var k int;
			if len(input) > 0 {
				k = input[0]
			} else {
				k = def
			}
			instructions[j] = k
			if len(input) <= 1 {
				input = []int{}
			} else {
				input = input[1:]
			}
			i += 2
		case 4: // output
			j := instructions[i+1]
			o := chooseMode(ji, j, instructions)
			output = append(output, o);
			i += 2
			loop = false
			break
		case 5: // jump true
			j := instructions[i+1]
			k := instructions[i+2]
			if chooseMode(ji, j, instructions) != 0 {
				i = chooseMode(ki, k, instructions)
			} else {
				i += 3
			}
		case 6: // jump false
			j := instructions[i+1]
			k := instructions[i+2]
			if chooseMode(ji, j, instructions) == 0 {
				i = chooseMode(ki, k, instructions)
			} else {
				i += 3
			}
		case 7: // less than
			j := instructions[i+1]
			k := instructions[i+2]
			l := instructions[i+3]
			if chooseMode(ji, j, instructions) < chooseMode(ki, k, instructions) {
				instructions[l] = 1
			} else {
				instructions[l] = 0
			}
			i += 4 // equal to
		case 8:
			j := instructions[i+1]
			k := instructions[i+2]
			l := instructions[i+3]
			if chooseMode(ji, j, instructions) == chooseMode(ki, k, instructions) {
				instructions[l] = 1
			} else {
				instructions[l] = 0
			}
			i += 4
		case 99: // end
			fmt.Println("HALT")
			loop = false
			break
		}

		if i >= len(instructions) {
			loop = false
		}
	}

	fmt.Println("OUTPUT LENGTH", len(output))
	return output, input;
}

func readInput() []int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var nums []int
	scanner := bufio.NewScanner(file)

	scanner.Scan()
	arr := strings.Split(scanner.Text(), ",")

	for _, n := range arr {
		i, _ := strconv.Atoi(n)
		nums = append(nums, i)
	}

	return nums
}

func solvePart1(s []int, input []int) {
	defer utils.TimeTrack(time.Now(), "Day 7: Part 1")
	instructions := make([]int, len(s))

	answer := 0
	for v := range itertools.PermutationsInt(input, len(input)) {
		d := 0
		for i,_ := range []string{"A", "B", "C", "D", "E"} {
			copy(instructions, s)
			// fmt.Println(instructions);
			res,_ := executeProgram(instructions, []int{v[i], d}, d)
			d = res[0]
		}
		answer = Max(answer, d);
	}
	fmt.Println(answer)
}

func solvePart2(s []int, input []int) {
	defer utils.TimeTrack(time.Now(), "Day 7: Part 1")
	instructions := make([][]int, 5)
	for i,_ := range instructions {
		instructions[i] = make([]int, len(s));
		copy(instructions[i], s)
	}

	amps := []string{"A", "B", "C", "D", "E"}
	var signals [5][]int
	lastE := 0
	halt := false

	count := 0

	for v := range itertools.PermutationsInt(input, len(input)) {
		for i,j := range v {
			if i == 0 {
				signals[i] = []int{0, j}
			} else {
				signals[i] = []int{j}
			}
		}
		fmt.Println(v)
		for !halt {
			for i,_ := range amps {
				if (count > 20) {
					halt = true
					break
				}
				count += 1
				fmt.Println(count, amps[i], signals[i], instructions[i])
				// if i == 0 {
				// 	fmt.Println("\n", amps[i], instructions[i], "\n");
				// }
				curInstruction := make([]int, len(instructions[i]));
				copy(curInstruction, instructions[i])
				c := signals[i]
				res, c := executeProgram(instructions[i], c, c[len(c) - 1])
				copy(instructions[i], curInstruction)
				signals[i] = c
				fmt.Println(amps[i], res, c)
				if (len(res) < 0) {
					halt = true
					break
				}
				if (i == 4) {
					lastE = res[len(res) - 1]
				}
				for _,j := range res {
					signals[(i + 1) % len(signals)] = append(signals[(i + 1) % len(signals)], j)
				}
				fmt.Println("Post", count, amps[i], signals[i], instructions[i])
			}
		}
	}
	fmt.Println(lastE)
}
