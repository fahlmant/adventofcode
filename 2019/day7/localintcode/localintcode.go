package localintcode

import (
	"fmt"
	"strconv"
)

type Computer struct {
	PC, Offset                  int
	Instructions, Input, Output []int
}

func (c *Computer) RunProgram() {

	for {
		fmt.Println(c.PC)
		instruction := c.Instructions[c.PC]
		opcode, arg1, arg2, arg3 := c.parseOpcode(instruction)
		fmt.Printf("PC: %d\n", c.PC)
		switch opcode {
		case 1:
			fmt.Println("Opcode 1")
			fmt.Printf("Adding memaddr %d value %d with memaddr %d value %d at address %d, currently holding %d\n", arg1, c.Instructions[arg1], arg2, c.Instructions[arg2], arg3, c.Instructions[arg3])
			c.Instructions[arg3] = c.Instructions[arg1] + c.Instructions[arg2]
			fmt.Printf("Address %d now is %d\n", arg3, c.Instructions[arg3])
			c.PC += 4
		case 2:
			fmt.Println("Opcode 2")
			c.Instructions[arg3] = c.Instructions[arg1] * c.Instructions[arg2]
			c.PC += 4
		case 3:
			fmt.Println("Opcode 3")
			if len(c.Input) >= 1 {
				input := c.Input[0]
				c.Input = c.Input[1:]
				fmt.Printf("Placing input value %d at address %d, currently holding %d\n", input, arg1, c.Instructions[arg1])
				c.Instructions[arg1] = input
				fmt.Printf("Address %d now is %d\n", arg1, c.Instructions[arg1])
			} else {
				return
			}
			c.PC += 2
		case 4:
			fmt.Println("Opcode 4")
			c.Output = append(c.Output, c.Instructions[arg1])
			c.PC += 2
			return
		case 5:
			fmt.Println("Opcode 5")
			fmt.Printf("Jump if non zero, addr %d value %d\n", arg1, c.Instructions[arg1])
			if c.Instructions[arg1] != 0 {
				fmt.Printf("Setting PC to %d from address %d\n", c.Instructions[arg2], arg2)
				c.PC = c.Instructions[arg2]
			} else {
				c.PC += 3
			}
		case 6:
			fmt.Println("Opcode 6")
			if c.Instructions[arg1] == 0 {
				c.PC = c.Instructions[arg2]
			} else {
				c.PC += 3
			}
		case 7:
			fmt.Println("Opcode 7")
			if c.Instructions[arg1] < c.Instructions[arg2] {
				c.Instructions[arg3] = 1
			} else {
				c.Instructions[arg3] = 0
			}
			c.PC += 4
		case 8:
			fmt.Println("Opcode 8")
			if c.Instructions[arg1] == c.Instructions[arg2] {
				c.Instructions[arg3] = 1
			} else {
				c.Instructions[arg3] = 0
			}
			c.PC += 4
		case 9:
			fmt.Println("Opcode 9")
			c.Offset += c.Instructions[arg1]
			c.PC += 2
		case 99:
			return
		default:
			fmt.Printf("Invalid Opcode: %d\n", opcode)
			return
		}
		fmt.Println(c.Instructions)
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
		opcode, _ = strconv.Atoi(instructionString[len(instructionString)-2:])
		instructionString = instructionString[:len(instructionString)-2]
	}

	//Set the arguments to the defaultopode

	//99 Doens't have any arguments
	if opcode != 99 {
		arg1 = c.Instructions[c.PC+1]
		//9 only has 1 argument
		if !inSlice([]int{3, 4, 9}, opcode) {
			arg2 = c.Instructions[c.PC+2]
			//3 and 4 only have 2 arguments

			arg3 = c.Instructions[c.PC+3]
		}
	}

	//Get modes for each argument and assign index accordingly
	//Arg 3 mode
	if len(instructionString) == 3 {
		mode, _ := strconv.Atoi(string(instructionString[0]))
		if mode == 1 {
			arg3 = c.PC + 3
		} else if mode == 2 {
			arg3 = c.Instructions[c.PC+3] + c.Offset
		}
		instructionString = instructionString[1:]
	}
	//Arg 2 mode
	if len(instructionString) == 2 {
		mode, _ := strconv.Atoi(string(instructionString[0]))
		if mode == 1 {
			arg2 = c.PC + 2
		} else if mode == 2 {
			arg2 = c.Instructions[c.PC+2] + c.Offset
		}
		instructionString = instructionString[1:]
	}
	//Arg 1 mode
	if len(instructionString) == 1 {
		mode, _ := strconv.Atoi(string(instructionString[0]))
		if mode == 1 {
			arg1 = c.PC + 1
		} else if mode == 2 {
			arg1 = c.Instructions[c.PC+1] + c.Offset
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
