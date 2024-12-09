package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Location struct {
	x, y int
}

func main() {

	total := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var diskMap []int
	var fileSystem []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		input := scanner.Text()
		for l := range input {
			lineNumber, err := strconv.Atoi(string(input[l]))
			if err != nil {
				panic(err)
			}
			diskMap = append(diskMap, lineNumber)
		}
	}

	valueToCountMap := make(map[int]int)
	// Build out the filesystem using the diskmap
	currentValue := 0
	for i := range diskMap {
		if i%2 == 0 || i == 0 {
			for range diskMap[i] {
				fileSystem = append(fileSystem, currentValue)
			}
			// Track the number of occurances per value
			valueToCountMap[currentValue] = diskMap[i]
			currentValue += 1
		} else {
			for range diskMap[i] {
				fileSystem = append(fileSystem, -1)
			}
		}
	}

	currentValue--

	for currentValue > 0 {
		// Get the number of times the current value occurs
		lenOfEntry := valueToCountMap[currentValue]

		// Find staring place of currentValue
		var startIndexOfValue int
		for x := range fileSystem {
			if fileSystem[x] == currentValue {
				startIndexOfValue = x
				break
			}
		}

		// Find sequences of -1 and see if current value can fit there
		var lenOfEmpty int
		var indexOfEmpty int
		for i := range fileSystem {
			if fileSystem[i] == -1 {
				lenOfEmpty += 1
				if lenOfEmpty == 1 {
					indexOfEmpty = i
				}
			} else if fileSystem[i] == currentValue {
				// If the current value is found, then there's no open space to the left, done
				break
			} else {
				lenOfEmpty = 0
				continue
			}

			if lenOfEmpty == lenOfEntry {
				for x := indexOfEmpty; x < indexOfEmpty+lenOfEntry; x++ {
					fileSystem[x] = currentValue
				}
				for y := startIndexOfValue; y < startIndexOfValue+lenOfEntry; y++ {
					fileSystem[y] = -1
				}
				break
			}
		}
		currentValue--

	}

	for i := range fileSystem {
		if fileSystem[i] > -1 {
			total += i * fileSystem[i]
		}

	}

	fmt.Println(total)
}

func getLastValidDiskId(fileSystem []int) int {
	for i := len(fileSystem) - 1; i >= 0; i-- {
		if fileSystem[i] != -1 {
			return i
		}
	}

	return -1
}
