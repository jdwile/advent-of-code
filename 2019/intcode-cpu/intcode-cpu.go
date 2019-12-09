package cpu

import "strconv"

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

	if n > 99 {
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
	}
	return n, jMode, kMode, lMode
}

func (c CPU) ExecuteProgram() CPU {
	loop := true
	for loop {
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
				loop = false
				break
			}
			j = chooseSetMode(jMode, c.Memory[c.InstructionPointer+1], c)
			k = c.Input[0]
			c.Memory[j] = k
			if len(c.Input) <= 1 {
				c.Input = []int{}
			} else {
				c.Input = c.Input[1:]
			}
			c.InstructionPointer += 2
		case 4: // output
			c.Output = append(c.Output, j)
			c.InstructionPointer += 2
		case 5: // jump true
			if j != 0 {
				c.InstructionPointer = k
			} else {
				c.InstructionPointer += 3
			}
		case 6: // jump false
			if j == 0 {
				c.InstructionPointer = k
			} else {
				c.InstructionPointer += 3
			}
		case 7: // less than
			if j < k {
				c.Memory[l] = 1
			} else {
				c.Memory[l] = 0
			}
			c.InstructionPointer += 4
		case 8: // equal to
			if j == k {
				c.Memory[l] = 1
			} else {
				c.Memory[l] = 0
			}
			c.InstructionPointer += 4
		case 9: // adjust relative base
			c.RelativeBase += j
			c.InstructionPointer += 2
		case 99: // end
			c.Halted = true
			loop = false
			break
		}
	}

	return c
}
