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
			if input[j][i] == 'X' {
				total += checkForXMAS(input, i, j)
			}
		}
	}
	fmt.Println(total)
}

func checkForXMAS(input [][]byte, i, j int) int {
	XMASMatches := 0
	MASString := "MAS"

	// Horizontal to the right
	if i <= len(input[j])-1-3 {
		var line []byte
		for x := 1; x <= 3; x++ {
			line = append(line, input[j][i+x])
		}
		if string(line) == MASString {
			XMASMatches += 1
		}
	}

	// Horizontal to the left
	if i >= 3 {
		var line []byte
		for x := 1; x <= 3; x++ {
			line = append(line, input[j][i-x])
		}
		if string(line) == MASString {
			XMASMatches += 1
		}
	}

	// Vertical up
	if j >= 3 {
		var line []byte
		for x := 1; x <= 3; x++ {
			line = append(line, input[j-x][i])
		}
		if string(line) == MASString {
			XMASMatches += 1
		}
	}

	// Vertical down
	if j <= len(input)-1-3 {
		var line []byte
		for x := 1; x <= 3; x++ {
			line = append(line, input[j+x][i])
		}
		if string(line) == MASString {
			XMASMatches += 1
		}
	}

	// Diagonal NE: Up (-y) and right (+x)
	if i <= len(input[j])-1-3 && j >= 3 {
		var line []byte
		for x := 1; x <= 3; x++ {
			line = append(line, input[j-x][i+x])
		}
		if string(line) == MASString {
			XMASMatches += 1
		}
	}

	// Diagonal SE: Down (+y) and right (+x)
	if i <= len(input[j])-1-3 && j <= len(input)-1-3 {
		var line []byte
		for x := 1; x <= 3; x++ {
			line = append(line, input[j+x][i+x])
		}
		if string(line) == MASString {
			XMASMatches += 1
		}
	}

	// Diagonal NW: Up (-y) and left (-x)\
	if i >= 3 && j >= 3 {
		var line []byte
		for x := 1; x <= 3; x++ {
			line = append(line, input[j-x][i-x])
		}
		if string(line) == MASString {
			XMASMatches += 1
		}
	}

	// Diagonal SW: Down (+y) and left (-x)
	if i >= 3 && j <= len(input)-1-3 {
		var line []byte
		for x := 1; x <= 3; x++ {
			line = append(line, input[j+x][i-x])
		}
		if string(line) == MASString {
			XMASMatches += 1
		}
	}

	return XMASMatches
}
