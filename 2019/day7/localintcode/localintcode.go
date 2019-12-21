package localintcode

import (
	"fmt"
	"strconv"
)

type Computer struct {
	PC, Offset, Output  int
	Instructions, Input []int
}

func (c *Computer) RunProgram() {

	for {
		instruction := c.Instructions[c.PC]
		opcode, arg1, arg2, arg3 := c.parseOpcode(instruction)

		switch opcode {
		case 1:
			c.Instructions[arg3] = c.Instructions[arg1] + c.Instructions[arg2]
			c.PC += 4
		case 2:
			c.Instructions[arg3] = c.Instructions[arg1] * c.Instructions[arg2]
			c.PC += 4
		case 3:
			if len(c.Input) >= 1{
				input := c.Input[0]
				c.Input = c.Input[1:]
				c.Instructions[arg1] = input
			}
			c.PC += 2
		case 4:
			c.Output = c.Instructions[arg1]
			c.PC +=2
		case 5:
			if c.Instructions[arg1] != 0 {
				c.PC = c.Instructions[arg2]
			} else {
				c.PC += 3
			}
		case 6:
			if c.Instructions[arg1] == 0 {
				c.PC = c.Instructions[arg2]
			} else {
				c.PC += 3
			}
		case 7:
			if c.Instructions[arg1] < c.Instructions[arg2] {
				c.Instructions[arg3] = 1
			} else {
				c.Instructions[arg3] = 0
			}
			c.PC += 4
		case 8:
			if c.Instructions[arg1] == c.Instructions[arg2] {
				c.Instructions[arg3] = 1
			} else {
				c.Instructions[arg3] = 0
			}
			c.PC += 4
		case 99:
			return
		default:
			fmt.Printf("Invalid Opcode: %d\n", opcode)
			return
		}
	}
}

func (c *Computer) parseOpcode(instruction int) (int, int, int, int) {

	var opcode, arg1, arg2, arg3 int

	//Cast instruction to string for easier parsing
	instructionString := strconv.Itoa(instruction)
	//Get the opcode out of the instruction
	if len(instructionString) < 2 {
		//Handle case that opcode is only a single digit
		opcode = instruction
		instructionString = ""
	} else {
		//Extract opcode from instruction and remove it from string
		opcode, _ = strconv.Atoi(instructionString[len(instructionString) - 2:])
		instructionString = instructionString[:len(instructionString) - 2]
	}

	//Set the arguments to the defaultopode

	//99 Doens't have any arguments
	if opcode != 99 {
		arg1 = c.Instructions[c.PC+1]
		//9 only has 1 argument
		if opcode != 9 {
			arg2 = c.Instructions[c.PC+2]
			//3 and 4 only have 2 arguments
			if !inSlice([]int{3,4}, opcode) {
				arg3 = c.Instructions[c.PC+3]
			}
		}
	}

	//Get modes for each argument and assign index accordingly
	//Arg 3 mode
	if len(instructionString) == 3 {
		mode, _ := strconv.Atoi(string(instructionString[0]))
		if mode == 1 {
			arg3 = c.PC + 3
		} else if mode == 2 {

		}
		instructionString = instructionString[1:]
	}
	//Arg 2 mode
	if len(instructionString) == 2 {
		mode, _ := strconv.Atoi(string(instructionString[0]))
		if mode == 1 {
			arg2 = c.PC + 2
		} else if mode == 2 {

		}
		instructionString = instructionString[1:]
	}
	//Arg 1 mode
	if len(instructionString) == 1 {
		mode, _ := strconv.Atoi(string(instructionString[0]))
		if mode == 1 {
			arg1 = c.PC + 1
		} else if mode == 2 {

		}
		instructionString = instructionString[1:]
	}

	return opcode, arg1, arg2, arg3
}

func inSlice(slice []int, value int) bool {

	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}