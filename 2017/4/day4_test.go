package main

import (
	"fmt"
	"testing"

	"github.com/jdwile/advent-of-code/2017/utils"
)

func TestPartOne(t *testing.T) {
	defer utils.Mute()()
	t.Run("aa bb cc dd ee", testPartOneFunc([][]string{[]string{"aa", "bb", "cc", "dd", "ee"}}, 1))
	t.Run("aa aa cc dd ee", testPartOneFunc([][]string{[]string{"aa", "aa", "cc", "dd", "ee"}}, 0))
	t.Run("aa bb cc dd aaa", testPartOneFunc([][]string{[]string{"aa", "bb", "cc", "dd", "aaa"}}, 1))
}

func testPartOneFunc(passphrases [][]string, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := SolvePartOne(passphrases)
		if actual != expected {
			t.Error(fmt.Sprintf("Expected value of %v to be %d but got %d", passphrases, expected, actual))
		}
	}
}

func TestPartTwo(t *testing.T) {
	defer utils.Mute()()
	t.Run("abcd dcba", testPartTwoFunc([][]string{[]string{"abcd", "dcba"}}, 0))
	t.Run("a aa aaa", testPartTwoFunc([][]string{[]string{"a", "aa", "aaa"}}, 1))
}

func testPartTwoFunc(passphrases [][]string, expected int) func(*testing.T) {
	return func(t *testing.T) {
		actual := SolvePartTwo(passphrases)
		if actual != expected {
			t.Error(fmt.Sprintf("Expected value of %v to be %d but got %d", passphrases, expected, actual))
		}
	}
}
