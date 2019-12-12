package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/jdwile/advent-of-code/2019/utils"
)

func main() {
	nodes := readInput()
	solvePartOne(nodes)
	solvePartTwo(nodes)
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func readInput() map[string][]string {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	nodes := make(map[string][]string)
	for scanner.Scan() {
		a := strings.Split(scanner.Text(), ")")

		if nodes[a[0]] != nil {
			if !contains(nodes[a[0]], a[1]) {
				nodes[a[0]] = append(nodes[a[0]], a[1])
			}
		} else {
			nodes[a[0]] = []string{a[1]}
		}

		if nodes[a[1]] != nil {
			if !contains(nodes[a[1]], a[0]) {
				nodes[a[1]] = append(nodes[a[1]], a[0])
			}
		} else {
			nodes[a[1]] = []string{a[0]}
		}
	}

	return nodes
}

func CountDescendants(p string, nodes map[string][]string, vi map[string]bool) int {
	c := 0

	v := make(map[string]bool)
	for key, val := range vi {
		v[key] = val
	}
	l := []string{p}

	for len(l) > 0 {
		j := l[0]
		l = l[1:]
		v[j] = true

		for _, n := range nodes[j] {
			if !v[n] {
				c += 1
				l = append(l, n)
			}
		}
	}

	return c
}

func MinDistance(s, e string, nodes map[string][]string) int {
	v := make(map[string]bool)
	d := make(map[string]int)

	l := []string{s}
	v[s] = true
	d[s] = 0

	for len(l) > 0 {
		c := l[0]
		l = l[1:]

		for _, k := range nodes[c] {
			if v[k] {
				continue
			}

			d[k] = d[c] + 1
			l = append(l, k)
			v[k] = true
		}
	}
	return d[e]
}

func solvePartOne(nodes map[string][]string) {
	defer utils.TimeTrack(time.Now(), "Day 6: Part 1")
	numDirect := 1
	numIndirect := -1
	v := make(map[string]bool)
	l := []string{"COM"}

	for len(l) > 0 {
		j := l[0]
		l = l[1:]
		v[j] = true

		if nodes[j] != nil {
			d := len(nodes[j]) - 1
			i := CountDescendants(j, nodes, v) - d
			numDirect += d
			numIndirect += i

			for _, k := range nodes[j] {
				if !v[k] {
					l = append(l, k)
				}
			}
		}
	}

	fmt.Println("Part 1:", numDirect+numIndirect)
}

func solvePartTwo(nodes map[string][]string) {
	defer utils.TimeTrack(time.Now(), "Day 6: Part 2")
	fmt.Println("Part 2:", MinDistance("YOU", "SAN", nodes)-2)
}
