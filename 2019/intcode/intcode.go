package intcode

import (
	"fmt"
	"math"
	"os"
)

func RunIntcode(instructions []int, input int) []int {

	index := 0
	offset := 0
	var results []int
	for {

		opcode, arg1, arg2, arg3 := getValues(instructions, index, offset)
		switch opcode {
		case 1:
			instructions[arg3] = arg1 + arg2
			index += 4
		case 2:
			instructions[arg3] = arg1 * arg2
			index += 4
		case 3:
			instructions[arg1] = input
			index += 2
		case 4:
			results = append(results, arg1)
			index += 2
		case 5:
			if arg1 != 0 {
				index = arg2
			} else {
				index += 3
			}
		case 6:
			if arg1 == 0 {
				index = arg2
			} else {
				index += 3
			}
		case 7:
			if arg1 < arg2 {
				instructions[arg3] = 1
			} else {
				instructions[arg3] = 0
			}
			index += 4
		case 8:
			if arg1 == arg2 {
				instructions[arg3] = 1
			} else {
				instructions[arg3] = 0
			}
			index += 4
		case 9:
			offset += arg1
			index += 2
		case 99:
			fmt.Println(results)
			os.Exit(0)
		default:
			fmt.Println(opcode)
			fmt.Println("Invalid opcode")
			os.Exit(1)
		}
	}

}

func getValues(instructions []int, index, offset int) (int, int, int, int) {

	//Get the current instruction
	instruction := instructions[index]

	//Get the opcode out of the instruction
	opcode := instruction % 100

	//99 has no arguments
	if opcode == 99 {
		//Just return the instruction as 99 ends the program
		return instruction, 0, 0, 0
	}

	//Opcodes 3 can't support immediate mode since it needs to read and store at an index
	if opcode == 3 {
		if instruction > 200 {
			//Relative mode
			return opcode, instructions[index+1] + offset, 0, 0
		}
		//Default to position mode
		return opcode, instructions[index+1], 0, 0
	}
	if opcode == 4 {
		if instruction > 200 {
			//Relative mode

			return opcode, instructions[instructions[index+1]+offset], 0, 0
		} else if instruction > 100 {
			return opcode, instructions[index+1] + offset, 0, 0
		} else {
			return opcode, instructions[instructions[index+1]], 0, 0
		}
	}

	//9 supports all three modes but only takes 1 argument
	if opcode == 9 {
		if instruction > 200 {
			//Relative mode
			return opcode, instructions[instructions[index+1]+offset], 0, 0
		} else if instruction-opcode == 100 {
			//Immediate mode
			return opcode, instructions[index+1], 0, 0
		}
		//Default to position mode
		return opcode, instructions[instructions[index+1]], 0, 0
	}

	//1,2,5,6,7,and 8 support all three modes and multiple arguments
	if inSlice([]int{1, 2, 5, 6, 7, 8}, opcode) {

		var arg1, arg2, arg3 int

		// 5 and 6 Don't take a third argument
		if !inSlice([]int{5, 6}, opcode) {
			//The third argument is always an index, cannot be a values
			if instruction > 20000 {
				//Relative mode
				arg3 = instructions[index+3] + offset
			} else {
				//Default to position mode
				arg3 = instructions[index+3]
			}
		}

		//Handle first argument's mode
		if math.Floor(float64((instruction%1000)/100)) == 1 {
			//Immediate mode
			arg1 = instructions[index+1]
		} else if math.Floor(float64((instruction%1000)/100)) == 2 {
			//Relative mode
			arg1 = instructions[instructions[index+1]+offset]
		} else {
			//Default to position mode
			arg1 = instructions[instructions[index+1]]
		}

		//Handle second argument's mode
		if math.Floor(float64((instruction%10000)/1000)) == 1 {
			//Parameter mode
			arg2 = instructions[index+2]
		} else if math.Floor(float64((instruction%10000)/1000)) == 2 {
			//Relative mode
			arg2 = instructions[instructions[index+2]+offset]
		} else {
			//Default to position mode
			arg2 = instructions[instructions[index+2]]
		}
		return opcode, arg1, arg2, arg3
	}

	//Default to return just the opcode
	fmt.Println("Dead men tell no")
	return opcode, 0, 0, 0
}

func inSlice(slice []int, value int) bool {

	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}
