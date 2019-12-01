package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Circut struct {
	output    string
	inputs    []string
	values    []int
	operation string
}

func main() {

	circuts := make(map[string]*Circut)

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		circut := parseInstruction(line)
		circuts[circut.output] = circut
	}

	fmt.Println(circuts)
}

func parseInstruction(instruction string) *Circut {

	operationRegex := regexp.MustCompile("[A-Z]+")
	outputRegex := regexp.MustCompile("-> (.+)")
	inputRegex := regexp.MustCompile("[a-z0-9]+")
	wireRegex := regexp.MustCompile("[a-z]+")

	operation := operationRegex.FindString(instruction)
	output := outputRegex.FindString(instruction)
	inputs := inputRegex.FindAllStringSubmatch(instruction, -1)

	var wireInputs []string
	var values []int

	for _, input := range inputs {
		if wireRegex.MatchString(input[0]) {
			wireInputs = append(wireInputs, input[0])
		} else {
			value, _ := strconv.Atoi(input[0])
			values = append(values, value)
		}
	}

	outputWire := strings.Split(output, " ")[1]
	return &Circut{outputWire, wireInputs[0 : len(wireInputs)-1], values, operation}
}
