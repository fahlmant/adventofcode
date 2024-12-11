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

	stones := make(map[int]int)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineData := strings.Split(line, " ")
		for _, stone := range lineData {
			stoneNum, err := strconv.Atoi(stone)
			if err != nil {
				panic(err)
			}
			stones[stoneNum] += 1
		}
	}

	// Run 25 blinks
	for range 75 {
		nextStones := make(map[int]int)
		for k, v := range stones {
			// Convert 0 to 1
			if k == 0 {
				nextStones[1] += v
				continue
			}
			// Convert to string to check length
			// If length is even, remove value from map and add first half and second half to map
			stoneNumString := strconv.Itoa(k)
			if len(stoneNumString)%2 == 0 {
				firstHalf := stoneNumString[:len(stoneNumString)/2]
				secondHalf := stoneNumString[len(stoneNumString)/2:]
				firstHalfInt, err := strconv.Atoi(firstHalf)
				if err != nil {
					panic(err)
				}
				secondHalfInt, err := strconv.Atoi(secondHalf)
				if err != nil {
					panic(err)
				}
				nextStones[firstHalfInt] += v
				nextStones[secondHalfInt] += v
				continue
			}
			// Default, multiple value by 2024
			nextStones[k*2024] += v
		}
		stones = nextStones
	}

	for _, v := range stones {
		total += v
	}
	fmt.Println(total)
}
