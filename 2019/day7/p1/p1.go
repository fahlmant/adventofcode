package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {

	var instructionStrings []string
	max := 0

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
	possiblePhaseSettings := buildCombinationsList(0, 5)
	for _, setting := range possiblePhaseSettings {
		result := runAmplifiers(instructions, setting[0], setting[1], setting[2], setting[3], setting[4])

		if result > max {
			max = result
		}
	}

	fmt.Println(max)

}

func runAmplifiers(instructions []int, phaseA, phaseB, phaseC, phaseD, phaseE int) int {

	resultA := intcodeProcessor(instructions, phaseA, 0)
	resultB := intcodeProcessor(instructions, phaseB, resultA)
	resultC := intcodeProcessor(instructions, phaseC, resultB)
	resultD := intcodeProcessor(instructions, phaseD, resultC)
	resultE := intcodeProcessor(instructions, phaseE, resultD)

	return resultE

}

func intcodeProcessor(instructions []int, input1 int, input2 int) int {

	var results []int

	results = append(results, input1)
	results = append(results, input2)

	instructions[instructions[1]] = input1

	index := 2
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
			instructions[arg1] = results[len(results)-1]
			index += 2
		case 4:
			results = append(results, instructions[arg1])
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
		case 99:
			return results[len(results)-1]
		default:
			fmt.Println("Invalid opcode")
			fmt.Println(opcode)
			os.Exit(1)
		}
	}
}

func getValues(instructions []int, index int) (int, int, int, int) {

	instruction := instructions[index]
	opcode := instruction % 100
	if opcode == 99 {
		return instruction, 0, 0, 0
	}

	if opcode == 3 || opcode == 4 {
		return opcode, instructions[index+1], 0, 0
	}

	if inSlice([]int{1, 2, 5, 6, 7, 8}, opcode) {

		var arg1, arg2, arg3 int

		if !inSlice([]int{5, 6}, opcode) {
			arg3 = instructions[index+3]
		}

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

func buildCombinationsList(low, high int) [][]int {

	rand.Seed(time.Now().Unix())

	var combinationList [][]int

	for len(combinationList) < (factorial(high) - factorial(low)) {
		var combination string
		for len(combination) < 5 {

			if len(combination) < 1 {
				combination = strconv.Itoa(rand.Intn(high-low) + low)
			}

			nextNum := rand.Intn(high-low) + low
			if !strings.Contains(combination, strconv.Itoa(nextNum)) {
				combination = strings.Join([]string{combination, strconv.Itoa(nextNum)}, "")
			}
		}

		for _, item := range combinationList {
			if equal(item, stringToIntArray(combination)) {

			}
		}
		if !sliceInSlice(combinationList, stringToIntArray(combination)) {
			combinationList = append(combinationList, stringToIntArray(combination))
		}

	}

	return combinationList
}

func stringToIntArray(value string) []int {

	var final []int
	for _, v := range value {
		num, _ := strconv.Atoi(string(v))
		final = append(final, num)
	}
	return final
}

func sliceInSlice(slice [][]int, value []int) bool {
	for _, v := range slice {
		if equal(value, v) {
			return true
		}
	}
	return false
}

func inSlice(slice []int, value int) bool {

	for _, v := range slice {
		if value == v {
			return true
		}
	}
	return false
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func factorial(n int) (result int) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}
