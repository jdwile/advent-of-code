package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/jdwile/advent-of-code/2019/utils"
)

func main() {
	instructions := ReadInput()
	SolvePartOne(instructions)
	// SolvePartTwo(memory)
}

func ReadInput() (instructions []string) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		instructions = append(instructions, scanner.Text())
	}

	return instructions
}

func SolvePartOne(instructions []string) {
	defer utils.TimeTrack(time.Now(), "Day 22: Part 1")
	SIZE := 10007

	deck := make([]int, SIZE)
	for i := range deck {
		deck[i] = i
	}

	for _, instruction := range instructions {
		if instruction == "deal into new stack" { // Reverse
			newDeck := make([]int, SIZE)
			for i := range deck {
				newDeck[SIZE-i-1] = deck[i]
			}
			deck = newDeck
		} else if strings.Contains(instruction, "cut") { // Cut
			arr := strings.Split(instruction, " ")
			n, _ := strconv.Atoi(arr[len(arr)-1])

			if n < 0 {
				n += SIZE
			}

			temp := deck[n:]
			for i := 0; i < n; i++ {
				temp = append(temp, deck[i])
			}
			deck = temp
		} else { // Deal
			arr := strings.Split(instruction, " ")
			n, _ := strconv.Atoi(arr[len(arr)-1])
			// fmt.Println("Dealing with increment", n)

			newDeck := make([]int, SIZE)

			j := 0
			for i := range deck {
				newDeck[j] = 0 + deck[i]
				// fmt.Println(newDeck, deck[i], j)

				j += n
				j = j % SIZE

			}
			// fmt.Println(deck, "dealt into", newDeck)
			deck = newDeck
		}
	}

	for i, v := range deck {
		if v == 2019 {
			fmt.Println(i)
		}
	}
}
