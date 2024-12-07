package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	result      int
	calibration []int
}

func main() {

	total := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var equations []Equation

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, ":")
		result, err := strconv.Atoi(lineSplit[0])
		if err != nil {
			panic(err)
		}
		var listOfNums []int
		listOfNumsAsSting := strings.Split(strings.TrimSpace(lineSplit[1]), " ")
		for i := range listOfNumsAsSting {
			num, err := strconv.Atoi(listOfNumsAsSting[i])
			if err != nil {
				panic(err)
			}
			listOfNums = append(listOfNums, num)
		}
		equations = append(equations, Equation{result: result, calibration: listOfNums})
	}

	for e := range equations {
		if isCalibrationValid(equations[e]) {
			total += equations[e].result
		}
	}

	fmt.Println(total)
}

func isCalibrationValid(e Equation) bool {
	addResult := getResult(e.calibration, 0, e.result, "+")
	mulResult := getResult(e.calibration, 0, e.result, "*")
	concatResult := getResult(e.calibration, 0, e.result, "||")

	return addResult || mulResult || concatResult
}

func getResult(equationNums []int, total int, target int, operation string) bool {

	var newTotal int
	var err error
	switch operation {
	case "+":
		newTotal = total + equationNums[0]
	case "*":
		newTotal = total * equationNums[0]
	case "||":
		// Convert total and number to string
		totalString := strconv.Itoa(total)
		numString := strconv.Itoa(equationNums[0])
		newTotalString := totalString + numString
		newTotal, err = strconv.Atoi(newTotalString)
		if err != nil {
			panic(err)
		}
	default:
		panic("Not a valid operation")
	}

	if len(equationNums) == 1 {
		return newTotal == target
	}

	addResult := getResult(equationNums[1:], newTotal, target, "+")
	mulResult := getResult(equationNums[1:], newTotal, target, "*")
	concatResult := getResult(equationNums[1:], newTotal, target, "||")

	return addResult || mulResult || concatResult
}
