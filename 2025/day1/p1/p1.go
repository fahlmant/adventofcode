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

	currentNum := 50

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		direction := line[0]
		weight := line[1:]
		weightNum, err := strconv.Atoi(weight)
		if err != nil {
			log.Fatal(err)
		}

		if direction == 'L' {
			currentNum = (currentNum - weightNum) % 100
		} else {
			currentNum = (currentNum + weightNum) % 100
		}
		if currentNum < 0 {
			currentNum += 100
		}
		if currentNum == 0 {
			total += 1
		}

	}

	fmt.Println(total)
}
