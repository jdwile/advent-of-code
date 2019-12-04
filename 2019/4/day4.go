package main

import (
	"fmt"
	"github.com/jdwile/advent-of-code/2019/utils"
	"strconv"
	"time"
)

func main() {
	solvePartOne()
	solvePartTwo()
}

func AlwaysIncreases(n int) bool {
	s := strconv.Itoa(n)

	for i := 0; i < len(s)-1; i++ {
		if s[i] > s[i+1] {
			return false
		}
	}

	return true
}

func DoubleExists(n int) bool {
	s := strconv.Itoa(n)

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}

	return false
}

func ExplicitDoubleExists(n int) bool {
	s := strconv.Itoa(n)
	r := 1

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			r += 1
		} else if r == 2 {
			return true
		} else {
			r = 1
		}
	}

	if r == 2 {
		return true
	}

	return false
}

func solvePartOne() {
	defer utils.TimeTrack(time.Now(), "Day 4: Part 1")
	n := 0
	for i := 138307; i <= 654504; i++ {
		if AlwaysIncreases(i) && DoubleExists(i) {
			n += 1
		}
	}

	fmt.Println(n)
}

func solvePartTwo() {
	defer utils.TimeTrack(time.Now(), "Day 4: Part 2")
	n := 0
	for i := 138307; i <= 654504; i++ {
		if AlwaysIncreases(i) && ExplicitDoubleExists(i) {
			n += 1
		}
	}

	fmt.Println(n)
}
