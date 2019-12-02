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

	instructions[1] = 12
	instructions[2] = 2

	index := 0
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
			fmt.Println(instructions)
			os.Exit(0)
		default:
			fmt.Println("Opcode error")
			os.Exit(1)
		}

		index += 4
	}
}
