package main

import (
	"bufio"
	"fmt"
	"github.com/jdwile/advent-of-code/2019/utils"
	"os"
	"strconv"
	"time"
)

func main() {
	modules := readInput()
	solvePartOne(modules)
	solvePartTwo(modules)
}

func calculateFuel(mass int) int {
	fuelCost := (mass / 3) - 2

	if fuelCost < 0 {
		return 0
	}

	return fuelCost
}

func readInput() []int {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	var numbers []int
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		numbers = append(numbers, num)
	}

	return numbers
}

func solvePartOne(modules []int) {
	defer utils.TimeTrack(time.Now(), "Day 1: Part 1")
	fmt.Println("\nPart 1")
	fuel := 0

	for _, mass := range modules {
		fuel += calculateFuel(mass)
	}

	fmt.Printf("Total Fuel: %d\n", fuel)
}

func solvePartTwo(modules []int) {
	defer utils.TimeTrack(time.Now(), "Day 1: Part 2")
	fmt.Println("\nPart 2")
	fuel := 0

	for _, mass := range modules {
		currentFuelCost := calculateFuel(mass)
		for currentFuelCost > 0 {
			fuel += currentFuelCost
			currentFuelCost = calculateFuel(currentFuelCost)
		}
	}

	fmt.Printf("Total Fuel: %d\n", fuel)
}
