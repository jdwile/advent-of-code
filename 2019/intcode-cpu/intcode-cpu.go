package cpu

import (
	"strconv"
)

type CPU struct {
	Memory             map[int]int
	Input              []int
	Output             []int
	InstructionPointer int
	RelativeBase       int
	Halted             bool
}

func ConstructCPU(memory map[int]int) CPU {
	return CPU{memory, []int{}, []int{}, 0, 0, false}
}

func chooseValueMode(mode byte, i int, c CPU) int {
	switch mode {
	case '0':
		return c.Memory[i]
	case '1':
		return i
	case '2':
		return c.Memory[c.RelativeBase+i]
	}
	panic("NO MODE")
}

func chooseSetMode(mode byte, i int, c CPU) int {
	if mode == '2' {
		return c.RelativeBase + i
	}
	return i
}

func chooseValues(c CPU, jMode, kMode, lMode byte) (j, k, l int) {
	j = c.Memory[c.InstructionPointer+1]
	k = c.Memory[c.InstructionPointer+2]
	l = c.Memory[c.InstructionPointer+3]

	j = chooseValueMode(jMode, j, c)
	k = chooseValueMode(kMode, k, c)
	l = chooseSetMode(lMode, l, c)

	return j, k, l
}

func parseOpcode(n int) (int, byte, byte, byte) {
	jMode := "0"[0]
	kMode := "0"[0]
	lMode := "0"[0]

	if n <= 99 {
		return n, jMode, kMode, lMode
	}

	s := strconv.Itoa(n)
	if len(s) == 3 {
		s = "0" + s
	}
	if len(s) == 4 {
		s = "0" + s
	}

	n = n % 100
	jMode = s[2]
	kMode = s[1]
	lMode = s[0]

	return n, jMode, kMode, lMode
}

func (c CPU) ExecuteProgram(t ...int) CPU {
	count := -1
	if len(t) > 0 {
		count = t[0]
	}
	for count != 0 {
		n, jMode, kMode, lMode := parseOpcode(c.Memory[c.InstructionPointer])
		j, k, l := chooseValues(c, jMode, kMode, lMode)

		switch n {
		case 1: // add
			c.Memory[l] = j + k
			c.InstructionPointer += 4
		case 2: // multiply
			c.Memory[l] = j * k
			c.InstructionPointer += 4
		case 3: // input
			if len(c.Input) == 0 {
				count = 1
				break
			}

			j = chooseSetMode(jMode, c.Memory[c.InstructionPointer+1], c)
			k = c.Input[0]
			c.Memory[j] = k

			v := []int{}
			if len(c.Input) > 1 {
				v = c.Input[1:]
			}

			c.Input = v
			c.InstructionPointer += 2
		case 4: // output
			c.Output = append(c.Output, j)
			c.InstructionPointer += 2
		case 5: // jump true
			v := c.InstructionPointer + 3
			if j != 0 {
				v = k
			}

			c.InstructionPointer = v
		case 6: // jump false
			v := c.InstructionPointer + 3
			if j == 0 {
				v = k
			}

			c.InstructionPointer = v
		case 7: // less than
			v := 0
			if j < k {
				v = 1
			}

			c.Memory[l] = v
			c.InstructionPointer += 4
		case 8: // equal to
			v := 0
			if j == k {
				v = 1
			}

			c.Memory[l] = v
			c.InstructionPointer += 4
		case 9: // adjust relative base
			c.RelativeBase += j
			c.InstructionPointer += 2
		case 99: // end
			c.Halted = true
			count = 1
		}

		count--
	}

	return c
}
