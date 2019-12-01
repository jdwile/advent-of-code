package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	modules := readInput()
	solvePartOne(modules)
	solvePartTwo(modules)
}

func calculateFuel(mass int) int {
	fuelCost := int(math.Floor(float64(mass)/3.0)) - 2

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
	fmt.Println("\nPart 1")
	fuel := 0

	for _, mass := range modules {
		fuel += calculateFuel(mass)
	}

	fmt.Printf("Total Fuel: %d\n", fuel)
}

func solvePartTwo(modules []int) {
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
