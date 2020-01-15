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
	t.Run("Given Example", testPartOneFunc([][]int{{5, 1, 9, 5}, {7, 5, 3}, {2, 4, 6, 8}}, 18))
}

func testPartOneFunc(spreadsheet [][]int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := SolvePartOne(spreadsheet)
		if actual != expected {
			t.Error(fmt.Sprintf("Expected captcha value of %v to be %d but got %d", spreadsheet, expected, actual))
		}
	}
}

func TestPartTwo(t *testing.T) {
	defer silence()()
	t.Run("Given Example", testPartTwoFunc([][]int{{5, 9, 2, 8}, {9, 4, 7, 3}, {3, 8, 6, 5}}, 9))
}

func testPartTwoFunc(spreadsheet [][]int, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := SolvePartTwo(spreadsheet)
		if actual != expected {
			t.Error(fmt.Sprintf("Expected captcha value of %v to be %d but got %d", spreadsheet, expected, actual))
		}
	}
}
