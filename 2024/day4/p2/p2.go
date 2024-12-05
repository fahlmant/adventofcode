package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	total := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var input [][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var lineArray []byte
		for i := range len(line) {
			lineArray = append(lineArray, line[i])
		}
		input = append(input, lineArray)
	}

	for j := range input {
		for i := range input[j] {
			if input[j][i] == 'A' {
				total += checkForXMAS(input, i, j)
			}
		}
	}
	fmt.Println(total)
}

func checkForXMAS(input [][]byte, i, j int) int {
	XMASMatches := 0

	if i > len(input[j])-2 || i < 1 {
		return 0
	}
	if j > len(input)-2 || j < 1 {
		return 0
	}

	// M M
	//  A
	// S S
	if input[j-1][i-1] == 'M' && input[j-1][i+1] == 'M' && input[j+1][i-1] == 'S' && input[j+1][i+1] == 'S' {
		XMASMatches += 1
	}

	// S S
	//  A
	// M M
	if input[j-1][i-1] == 'S' && input[j-1][i+1] == 'S' && input[j+1][i-1] == 'M' && input[j+1][i+1] == 'M' {
		XMASMatches += 1
	}

	// M S
	//  A
	// M S
	if input[j-1][i-1] == 'M' && input[j-1][i+1] == 'S' && input[j+1][i-1] == 'M' && input[j+1][i+1] == 'S' {
		XMASMatches += 1
	}

	// S M
	//  A
	// S M
	if input[j-1][i-1] == 'S' && input[j-1][i+1] == 'M' && input[j+1][i-1] == 'S' && input[j+1][i+1] == 'M' {
		XMASMatches += 1
	}

	return XMASMatches
}
