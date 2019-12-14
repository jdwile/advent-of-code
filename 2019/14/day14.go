package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jdwile/advent-of-code/2019/utils"
)

func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}

type Reaction struct {
	N int
	C map[string]int
}

func main() {
	reactions := ReadInput()
	n := SolvePartOne(reactions)
	SolvePartTwo(n, reactions)
}

func ReadInput() (reactions map[string]Reaction) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	reactions = make(map[string]Reaction)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		o := strings.Split(scanner.Text(), " => ")

		r := strings.Split(o[1], " ")
		resN, _ := strconv.Atoi(r[0])
		res := r[1]

		ingredients := make(map[string]int)

		r = strings.Split(o[0], ", ")

		for _, c := range r {
			s := strings.Split(c, " ")
			n, _ := strconv.Atoi(s[0])
			i := s[1]

			ingredients[i] = n
		}

		reaction := Reaction{resN, ingredients}
		reactions[res] = reaction
	}

	return reactions
}

func doReactions(fuel int, table map[string]Reaction) int {
	supply := make(map[string]int)
	demand := make(map[string]int)

	supply["ORE"] = demand["ORE"]
	demand["FUEL"] = fuel

	n := fuel

	for n > 0 {
		for c := range demand {
			if c == "ORE" || demand[c] == 0 {
				continue
			}

			defecit := demand[c] - supply[c]

			if defecit > 0 {
				r := int(math.Ceil(float64(defecit) / float64(table[c].N)))

				for c2, demand2 := range table[c].C {
					demand[c2] += demand2 * r
					if c2 != "ORE" {
						n += demand2 * r
					}
				}

				sum := table[c].N * r
				supply[c] += sum
			}

			n -= demand[c]
			supply[c] -= demand[c]
			demand[c] = 0
		}
	}

	return demand["ORE"]
}

func SolvePartOne(reactions map[string]Reaction) int {
	defer utils.TimeTrack(time.Now(), "Day 14: Part 1")
	r := doReactions(1, reactions)
	fmt.Println(r)
	return r
}

func SolvePartTwo(n int, reactions map[string]Reaction) {
	defer utils.TimeTrack(time.Now(), "Day 14: Part 2")

	goal := 1000000000000
	max := goal / n

	left := max
	right := max * 2

	for left <= right {
		mid := left + (right-left)/2
		ore := doReactions(mid, reactions)

		if ore < goal {
			max = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	fmt.Println(max)
}
