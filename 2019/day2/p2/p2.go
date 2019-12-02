package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var instructionStrings []string
	desiredOutput := 19690720

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

	for i := 0; i <= 99; i++ {
		for j := 0; j <= 99; j++ {
			result := runIntcode(instructionStrings, i, j)
			if result == nil {
				continue
			}
			if result[0] == desiredOutput {
				fmt.Printf("I: %d, J: %d, Answer: %d\n", i, j, (100*i)+j)
				os.Exit(0)
			}
		}
	}
}

func runIntcode(instructionString []string, x int, y int) []int {

	instructions := make([]int, len(instructionString))

	for i, v := range instructionString {
		instructions[i], _ = strconv.Atoi(v)
	}

	instructions[1] = x
	instructions[2] = y
	index := 0
	j := 0
	for {

		switch instructions[index] {
		case 1:
			index1 := instructions[index+1]
			index2 := instructions[index+2]
			outputIndex := instructions[index+3]
			instructions[outputIndex] = instructions[index1] + instructions[index2]
		case 2:
			index1 := instructions[index+1]
			index2 := instructions[index+2]
			outputIndex := instructions[index+3]
			instructions[outputIndex] = instructions[index1] * instructions[index2]
		case 99:
			return instructions
		default:
			fmt.Printf("Opcode error: %d, %d, %d\n", x, y, index)
			return nil
		}
		j++
		index += 4
	}
}
