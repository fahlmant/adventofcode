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

		batteryIndicies := []int{}

		previousIndex := 0
		for i := 0; i < 12; i++ {

			max := line[previousIndex]
			currentIndex := previousIndex
			// Find the max value that the next battery can be
			for j := previousIndex; j <= len(line)-(12-i); j++ {
				if line[j] > max {
					max = line[j]
					currentIndex = j
				}
			}
			batteryIndicies = append(batteryIndicies, currentIndex)
			previousIndex = currentIndex + 1
		}

		var battery string
		for _, in := range batteryIndicies {
			battery += string(line[in])
		}

		batteryPower, err := strconv.Atoi(battery)
		if err != nil {
			log.Fatal(err)
		}
		total += batteryPower
	}

	fmt.Println(total)
}
