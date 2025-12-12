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

	total := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	problems := [][]string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		newLine := strings.Fields(strings.TrimSpace(line))
		problems = append(problems, newLine)
	}

	for j := range problems[0] {
		numList := []int{}
		// Get all the numbers in a given column
		for i := 0; i < len(problems)-1; i++ {
			numString := problems[i][j]
			num, err := strconv.Atoi(numString)
			if err != nil {
				log.Fatal(err)
			}
			numList = append(numList, num)
		}

		// Find the operation at the bottom of the column
		operation := problems[len(problems)-1][j]

		// Perform the operation (product or sum of column)
		if operation == "*" {
			product := numList[0]
			for k := 1; k < len(numList); k++ {
				product *= numList[k]
			}
			total += product
		}
		if operation == "+" {
			sum := numList[0]
			for k := 1; k < len(numList); k++ {
				sum += numList[k]
			}
			total += sum
		}
	}

	fmt.Println(total)
}
