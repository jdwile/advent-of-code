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
	t.Run("1", testPartOneFunc(1, 0))
	t.Run("12", testPartOneFunc(12, 3))
	t.Run("23", testPartOneFunc(23, 2))
	t.Run("1024", testPartOneFunc(1024, 31))
}

func testPartOneFunc(input, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := SolvePartOne(input)
		if actual != expected {
			t.Error(fmt.Sprintf("Expected distance of %v to be %d but got %d", input, expected, actual))
		}
	}
}

func TestPartTwo(t *testing.T) {
	defer silence()()
	t.Run("1", testPartTwoFunc(1, 2))
	t.Run("22", testPartTwoFunc(22, 23))
	t.Run("350", testPartTwoFunc(350, 351))
	t.Run("800", testPartTwoFunc(800, 806))
}

func testPartTwoFunc(input, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := SolvePartTwo(input)
		if actual != expected {
			t.Error(fmt.Sprintf("Expected distance of %v to be %d but got %d", input, expected, actual))
		}
	}
}
