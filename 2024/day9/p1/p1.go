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

	// Build out the filesystem using the diskmap
	currentValue := 0
	for i := range diskMap {
		if i%2 == 0 || i == 0 {
			for range diskMap[i] {
				fileSystem = append(fileSystem, currentValue)
			}
			currentValue += 1
		} else {
			for range diskMap[i] {
				fileSystem = append(fileSystem, -1)
			}
		}
	}

	lastNumberIndex := getLastValidDiskId(fileSystem)
	for i := 0; i < lastNumberIndex; i++ {
		if fileSystem[i] == -1 {
			fileSystem[i] = fileSystem[lastNumberIndex]
			fileSystem[lastNumberIndex] = -1
			lastNumberIndex = getLastValidDiskId(fileSystem)
		}
	}

	index := 0
	for {
		if fileSystem[index] == -1 {
			break
		}
		total += index * fileSystem[index]
		index += 1
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
