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

	instructions := make([]int, len(instructionStrings) + 10000)

	for i, v := range instructionStrings {
		instructions[i], _ = strconv.Atoi(v)
	}

	computer := intcode.Computer{PC: 0, Offset: 0, Input: []int{}, Output: []int{}, Instructions: instructions}

	computer.RunProgram()

	fmt.Println(computer.Output)
}
