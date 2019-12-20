package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/fahlmant/intcode"
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
			instructions := make([]int, len(instructionStrings))

			for i, v := range instructionStrings {
				instructions[i], _ = strconv.Atoi(v)
			}
			instructions[1] = i
			instructions[2] = j
			computer := intcode.Computer{PC: 0, Offset: 0, Input: []int{}, Output: 0, Instructions: instructions}

			computer.RunProgram()
			if computer.Instructions[0] == desiredOutput {
				fmt.Println((100 * i) + j)
			}
		}
	}

}
