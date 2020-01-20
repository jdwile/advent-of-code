package main

import (
	"fmt"
	"testing"

	"github.com/jdwile/advent-of-code/2017/utils"
)

func TestPartOne(t *testing.T) {
	defer utils.Mute()()
	t.Run("1122", testPartOneFunc("1122", 3))
	t.Run("1111", testPartOneFunc("1111", 4))
	t.Run("1234", testPartOneFunc("1234", 0))
	t.Run("91212129", testPartOneFunc("91212129", 9))
}

func testPartOneFunc(captcha string, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := SolvePartOne(captcha)
		if actual != expected {
			t.Error(fmt.Sprintf("Expected captcha value of %s to be %d but got %d", captcha, expected, actual))
		}
	}
}

func TestPartTwo(t *testing.T) {
	defer utils.Mute()()
	t.Run("1212", testPartTwoFunc("1212", 6))
	t.Run("1221", testPartTwoFunc("1221", 0))
	t.Run("123425", testPartTwoFunc("123425", 4))
	t.Run("123123", testPartTwoFunc("123123", 12))
	t.Run("12131415", testPartTwoFunc("12131415", 4))
}

func testPartTwoFunc(captcha string, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := SolvePartTwo(captcha)
		if actual != expected {
			t.Error(fmt.Sprintf("Expected captcha value of %s to be %d but got %d", captcha, expected, actual))
		}
	}
}
