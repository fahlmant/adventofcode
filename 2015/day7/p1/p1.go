package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type instruction struct {
	output    string
	inputs    []string
	operation string
}

type wires struct {
	wireMap map[string]int
}

func main() {

	allWires := wires{wireMap: make(map[string]int, 1)}
	var instructions []instruction
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructions = append(instructions, parseInstruction(line))
	}

	allWires.runCircuts(instructions)
	fmt.Printf("%+v\n", allWires.wireMap)
}

func parseInstruction(line string) instruction {

	ins := instruction{}

	// Split on -> operator
	splitInput := strings.Split(line, "->")

	// Output register is after ->
	ins.output = strings.TrimSpace(splitInput[1])

	splitLeft := strings.Split(splitInput[0], " ")
	// If there's more than one word on the left side, parse it out
	if splitLeft[1] == "AND" || splitLeft[1] == "OR" || splitLeft[1] == "LSHIFT" || splitLeft[1] == "RSHIFT" {
		ins.operation = strings.TrimSpace(splitLeft[1])
		ins.inputs = append(ins.inputs, strings.TrimSpace(splitLeft[0]))
		ins.inputs = append(ins.inputs, strings.TrimSpace(splitLeft[2]))
	} else if splitLeft[0] == "NOT" {
		ins.operation = "NOT"
		ins.inputs = append(ins.inputs, strings.TrimSpace(splitLeft[1]))
	} else {
		// Use PLACE for the operation with no keyword
		ins.operation = "PLACE"
		ins.inputs = append(ins.inputs, strings.TrimSpace(splitLeft[0]))
	}

	return ins
}

func (w wires) executeInstruction(ins instruction) {

	var intInputs []int
	for _, input := range ins.inputs {
		input = strings.TrimSpace(input)
		if num, err := strconv.Atoi(input); err == nil {
			intInputs = append(intInputs, num)
		} else {
			intInputs = append(intInputs, w.wireMap[input])
		}
	}

	//fmt.Printf("executing instruction %+v with inputs %+v\n", ins, intInputs)

	switch ins.operation {
	case "AND":
		w.wireMap[ins.output] = intInputs[0] & intInputs[1]
	case "OR":
		w.wireMap[ins.output] = intInputs[0] | intInputs[1]
	case "LSHIFT":
		w.wireMap[ins.output] = intInputs[0] << uint(intInputs[1])
	case "RSHIFT":
		w.wireMap[ins.output] = intInputs[0] >> uint(intInputs[1])
	case "NOT":
		w.wireMap[ins.output] = intInputs[0]
	case "PLACE":
		w.wireMap[ins.output] = intInputs[0]
	default:
		panic("Error")
	}

}

func (w wires) runCircuts(instructions []instruction) {

	var instructionDone []instruction
	for len(instructions) != len(instructionDone) {
		for _, ins := range instructions {
			//fmt.Printf("%+v\n", ins)
			if !w.validateInstruction(ins) {
				continue
			}
			w.executeInstruction(ins)
			instructionDone = append(instructionDone, ins)
		}
	}
}

func (w wires) validateInstruction(ins instruction) bool {

	for _, input := range ins.inputs {
		input = strings.TrimSpace(input)
		if _, err := strconv.Atoi(input); err != nil {
			if _, ok := w.wireMap[input]; !ok {
				fmt.Printf("%+v\n", ins.inputs)
				fmt.Printf("Not valid because %s has no value\n", input)
				return false
			}
		}
	}
	return true
}
