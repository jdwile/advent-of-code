package main

import (
	"fmt"
	"os"
	"testing"
)

func silence() func() {
	null, _ := os.Open(os.DevNull)
	sout := os.Stdout
	serr := os.Stderr
	os.Stdout = null
	os.Stderr = null
	return func() {
		defer null.Close()
		os.Stdout = sout
		os.Stderr = serr
	}
}

func TestPartOne(t *testing.T) {
	defer silence()()
	t.Run("0 3 0 1 -3", testPartOneFunc([]int{0, 3, 0, 1, -3}, 5))
}

func testPartOneFunc(message []int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := SolvePartOne(message)
		if actual != expected {
			t.Error(fmt.Sprintf("Expected value of %v to be %d but got %d", message, expected, actual))
		}
	}
}

func TestPartTwo(t *testing.T) {
	defer silence()()
	t.Run("0 3 0 1 -3", testPartTwoFunc([]int{0, 3, 0, 1, -3}, 10))
}

func testPartTwoFunc(message []int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := SolvePartTwo(message)
		if actual != expected {
			t.Error(fmt.Sprintf("Expected value of %v to be %d but got %d", message, expected, actual))
		}
	}
}
