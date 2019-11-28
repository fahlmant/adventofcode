package main

import (
	"bufio"
	"fmt"

	"log"
	"os"
	"strconv"
)

func main() {

	var total int
	var freqChanges []int
	var numseem []int

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		num, _ := strconv.Atoi(line[1:])
		if string(line[0]) == "+" {
			freqChanges = append(freqChanges, num)
		} else if string(line[0]) == "-" {
			freqChanges = append(freqChanges, -num)
		}
	}

	for {
		for _, item := range freqChanges {
			total += item
			if find(numseem, total) {
				fmt.Println(total)
				os.Exit(0)
			} else {
				numseem = append(numseem, total)
			}
		}
	}
}

func find(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}
