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

func Abs(x int) int {
	if x < 0 {
		return 0 - x
	}
	return x
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

type Point struct {
	X int
	Y int
	Z int
}

type Moon struct {
	Position Point
	Velocity Point
}

func main() {
	m := ReadInput()
	// for i := range m {
	// 	fmt.Println("Pos:", m[i].Position, "Vel:", m[i].Velocity)
	// }
	SolvePartOne(m)
	SolvePartTwo(m)
}

func ReadInput() (moons []Moon) {
	file, _ := os.Open("./input.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		l := scanner.Text()
		s := strings.Split(l[1:len(l)-1], ",")
		var c [3]int
		for i, a := range s {
			c[i], _ = strconv.Atoi(strings.Split(a, "=")[1])
		}
		pos := Point{c[0], c[1], c[2]}
		vel := Point{0, 0, 0}
		moons = append(moons, Moon{pos, vel})
	}

	return moons
}

func MoveMoons(moons []Moon) {
	// Gravity
	for i := range moons {
		for j := range moons {
			if i == j {
				continue
			}

			if moons[j].Position.X > moons[i].Position.X {
				moons[i].Velocity.X += 1
			} else if moons[j].Position.X < moons[i].Position.X {
				moons[i].Velocity.X -= 1
			}

			if moons[j].Position.Y > moons[i].Position.Y {
				moons[i].Velocity.Y += 1
			} else if moons[j].Position.Y < moons[i].Position.Y {
				moons[i].Velocity.Y -= 1
			}

			if moons[j].Position.Z > moons[i].Position.Z {
				moons[i].Velocity.Z += 1
			} else if moons[j].Position.Z < moons[i].Position.Z {
				moons[i].Velocity.Z -= 1
			}
		}
	}

	// Apply Velocity
	for i := range moons {
		moons[i].Position.X += moons[i].Velocity.X
		moons[i].Position.Y += moons[i].Velocity.Y
		moons[i].Position.Z += moons[i].Velocity.Z
	}
}

func SolvePartOne(moons []Moon) {
	defer utils.TimeTrack(time.Now(), "Day 12: Part 1")
	for t := 1; t <= 1000; t++ {
		MoveMoons(moons)
	}

	// Calculate Potential Energy
	e := 0
	for _, m := range moons {
		potential := Abs(m.Position.X) + Abs(m.Position.Y) + Abs(m.Position.Z)
		kinetic := Abs(m.Velocity.X) + Abs(m.Velocity.Y) + Abs(m.Velocity.Z)
		e += potential * kinetic
	}
	fmt.Println(e)
}

func SolvePartTwo(moons []Moon) {
	defer utils.TimeTrack(time.Now(), "Day 12: Part 2")

	initialState := ReadInput()

	var cycles [3]int
	markX := true
	markY := true
	markZ := true
	t := 0
	loop := true

	for loop {
		t += 1
		MoveMoons(moons)
		xLoop := true
		yLoop := true
		zLoop := true
		for i := range moons {
			xLoop = xLoop && moons[i].Position.X == initialState[i].Position.X && moons[i].Velocity.X == initialState[i].Velocity.X
			yLoop = yLoop && moons[i].Position.Y == initialState[i].Position.Y && moons[i].Velocity.Y == initialState[i].Velocity.Y
			zLoop = zLoop && moons[i].Position.Z == initialState[i].Position.Z && moons[i].Velocity.Z == initialState[i].Velocity.Z
		}

		if xLoop {
			if cycles[0] == 0 {
				cycles[0] = t
			} else if markX {
				cycles[0] = t - cycles[0]
				markX = false
			}
		}
		if yLoop {
			if cycles[1] == 0 {
				cycles[1] = t
			} else if markY {
				cycles[1] = t - cycles[1]
				markY = false
			}
		}
		if zLoop {
			if cycles[2] == 0 {
				cycles[2] = t
			} else if markZ {
				cycles[2] = t - cycles[2]
				markZ = false
			}
		}
		loop = markX || markY || markZ
	}
	fmt.Println(LCM(cycles[0], cycles[1], cycles[2]))
}
