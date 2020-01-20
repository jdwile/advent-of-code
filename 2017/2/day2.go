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
	spreadsheet := readInputFile()
	fmt.Printf("Part 1 Answer: %d\n\n", SolvePartOne(spreadsheet))
	fmt.Printf("Part 2 Answer: %d\n\n", SolvePartTwo(spreadsheet))
}

func readInputFile() [][]int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	spreadsheet := make([][]int, 0)

	for scanner.Scan() {
		a := strings.Split(scanner.Text(), "\t")
		numArr := make([]int, len(a))
		for i, v := range a {
			numArr[i], _ = strconv.Atoi(v)
		}

		spreadsheet = append(spreadsheet, numArr)
	}

	return spreadsheet
}

func SolvePartOne(spreadsheet [][]int) int {
	defer utils.TimeTrack(time.Now(), "Day 2: Part 1")

	var res int

	for _, row := range spreadsheet {
		l, s := row[0], row[0]

		for _, n := range row {
			l = utils.Max(l, n)
			s = utils.Min(s, n)
		}

		res += l - s
	}

	return res
}

func SolvePartTwo(spreadsheet [][]int) int {
	defer utils.TimeTrack(time.Now(), "Day 2: Part 2")
	var res int

	for _, row := range spreadsheet {
		found := false

		for i := 0; i < len(row)-1; i++ {
			for j := i + 1; j < len(row); j++ {
				if row[i] < row[j] {
					if row[j]%row[i] == 0 {
						res += row[j] / row[i]
						found = true
						break
					}
				} else {
					if row[i]%row[j] == 0 {
						res += row[i] / row[j]
						found = true
						break
					}
				}
			}
			if found {
				break
			}
		}
	}

	return res
}
