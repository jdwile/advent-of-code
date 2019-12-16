package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jdwile/advent-of-code/2019/utils"
)

func Abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func main() {
	i := ReadInput()
	SolvePartOne(i)
	SolvePartTwo(i)
}

func ReadInput() (input []int) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	s := strings.Split(scanner.Text(), "")
	for _, a := range s {
		n, _ := strconv.Atoi(a)
		input = append(input, n)
	}
	return input
}

func SolvePartOne(signal []int) {
	defer utils.TimeTrack(time.Now(), "Day 16: Part 1")
	FFT := []int{0, 1, 0, -1}

	for phase := 0; phase < 100; phase++ {
		var output []int
		for i := range signal {
			d := 0

			for j, c := range signal {
				d += c * FFT[((j+1)/(i+1))%4]
			}

			output = append(output, Abs(d)%10)
		}
		signal = output
	}
	o := ""
	for i := 0; i < 8; i++ {
		o += strconv.Itoa(signal[i])
	}
	fmt.Println(o)
}

func SolvePartTwo(a []int) {
	defer utils.TimeTrack(time.Now(), "Day 16: Part 2")

	// fmt.Println(a)
	signal := ""
	for _, n := range a {
		signal += strconv.Itoa(n)
	}
	offset, _ := strconv.Atoi(signal[:7])
	signal = strings.Repeat(signal, 10000)

	output := make([]int, 0)

	for i := offset; i < len(signal); i++ {
		output = append(output, int(signal[i]-'0'))
	}

	for phase := 0; phase < 100; phase++ {
		d := 0

		for i := range output {
			d += output[len(output)-1-i]
			output[len(output)-1-i] = d % 10
		}
	}
	o := ""
	for i := 0; i < 8; i++ {
		o += strconv.Itoa(output[i])
	}
	fmt.Println(o)
}
