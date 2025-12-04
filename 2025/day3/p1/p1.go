package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	total := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Find the first number by finding the maximum number, sans the final number
		firstNumIndex := 0
		firstMaxNum := 0
		for i := 0; i < len(line)-1; i++ {
			newNum, err := strconv.Atoi(string(line[i]))
			if err != nil {
				log.Fatal(err)
			}
			if newNum > firstMaxNum {
				firstMaxNum = newNum
				firstNumIndex = i
			}
		}

		// Find the second  number, which is the max between the first number and the end
		secondNumIndex := firstNumIndex + 1
		secondMaxNum := 0
		for i := firstNumIndex + 1; i < len(line); i++ {
			newNum, err := strconv.Atoi(string(line[i]))
			if err != nil {
				log.Fatal(err)
			}
			if newNum > secondMaxNum {
				secondMaxNum = newNum
				secondNumIndex = i
			}
		}

		battery := fmt.Sprintf("%c%c", line[firstNumIndex], line[secondNumIndex])
		batteryPower, err := strconv.Atoi(battery)
		if err != nil {
			log.Fatal(err)
		}
		total += batteryPower
	}

	fmt.Println(total)
}
