package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/jdwile/advent-of-code/2019/utils"
)

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	passphrases := readInputFile()
	fmt.Printf("Part 1 Answer: %d\n\n", SolvePartOne(passphrases))
	fmt.Printf("Part 2 Answer: %d\n\n", SolvePartTwo(passphrases))
}

func readInputFile() [][]string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	passphrases := make([][]string, 0)

	for scanner.Scan() {
		a := strings.Split(scanner.Text(), " ")
		passphrases = append(passphrases, a)
	}

	return passphrases
}

func isValidPassphrase(passphrase []string) bool {
	p := make(map[string]bool)
	valid := true
	for _, word := range passphrase {
		if p[word] {
			valid = false
			break
		}
		p[word] = true
	}

	return valid
}

func SolvePartOne(passphrases [][]string) int {
	defer utils.TimeTrack(time.Now(), "Day 4: Part 1")

	var c int

	for _, passphrase := range passphrases {
		if isValidPassphrase(passphrase) {
			c++
		}
	}

	return c
}

func SolvePartTwo(passphrases [][]string) int {
	defer utils.TimeTrack(time.Now(), "Day 4: Part 2")

	for i := range passphrases {
		for j := range passphrases[i] {
			l := strings.Split(passphrases[i][j], "")
			sort.Strings(l)
			passphrases[i][j] = strings.Join(l, "")
		}
	}

	var c int

	for _, passphrase := range passphrases {
		if isValidPassphrase(passphrase) {
			c++
		}
	}

	return c
}
