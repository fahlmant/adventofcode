package main

import (
	"fmt"
	"strconv"
)

func main() {

	lowerBound := 256310
	upperBound := 732736
	count := 0

	for i := lowerBound; i <= upperBound; i++ {
		if isValid(i) {
			count++
		}

	}

	fmt.Println(count)
}

func isValid(num int) bool {

	if !hasSameValuesAdjacent(num) {
		return false
	}

	if !increasingLeftToRight(num) {
		return false
	}
	return true
}

func hasSameValuesAdjacent(num int) bool {

	numString := strconv.Itoa(num)
	counts := make(map[byte]int)

	for i := 0; i < len(numString); i++ {
		counts[numString[i]]++
	}

	for _, v := range counts {
		if v == 2 {
			return true
		}
	}

	return false
}

func increasingLeftToRight(num int) bool {

	numString := strconv.Itoa(num)
	for i := 0; i < len(numString)-1; i++ {
		num1, _ := strconv.Atoi(string(numString[i]))
		num2, _ := strconv.Atoi(string(numString[i+1]))
		if !(num1 <= num2) {
			return false
		}
	}
	return true
}
