package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {

	var instructionStrings []string
	var results []int
	input := 1

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructionStrings = strings.Split(line, ",")
	}

	instructions := make([]int, len(instructionStrings))

	for i, v := range instructionStrings {
		instructions[i], _ = strconv.Atoi(v)
	}

	index := 0
	for {

		opcode, arg1, arg2, arg3 := getValues(instructions, index)
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
			results = append(results, instructions[arg1])
			index += 2
		case 99:
			fmt.Println(results)
			os.Exit(0)
		default:
			fmt.Println("Invalid opcode")
			os.Exit(1)
		}
	}
}

func getValues(instructions []int, index int) (int, int, int, int) {

	instruction := instructions[index]
	if instruction == 99 {
		return instruction, 0, 0, 0
	}

	opcode := instruction % 100

	if opcode == 3 || opcode == 4 {
		return opcode, instructions[index+1], 0, 0
	}

	if inSlice([]int{1, 2}, opcode) {

		var arg1, arg2, arg3 int
		arg3 = instructions[index+3]

		if math.Floor(float64((instruction%1000)/100)) == 1 {

			arg1 = instructions[index+1]
		} else {
			arg1 = instructions[instructions[index+1]]
		}
		if math.Floor(float64((instruction%10000)/1000)) == 1 {
			arg2 = instructions[index+2]
		} else {
			arg2 = instructions[instructions[index+2]]
		}
		return opcode, arg1, arg2, arg3
	}

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
