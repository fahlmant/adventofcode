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

	var grid [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var lineData []int
		for l := range line {
			value, err := strconv.Atoi(string(line[l]))
			if err != nil {
				panic(err)
			}
			lineData = append(lineData, value)
		}
		grid = append(grid, lineData)
	}

	for j := range grid {
		for i := range grid[j] {
			if grid[j][i] == 0 {

				total += findUniqueReachableNines(grid, Location{x: i + 1, y: j}, 1)
				total += findUniqueReachableNines(grid, Location{x: i - 1, y: j}, 1)
				total += findUniqueReachableNines(grid, Location{x: i, y: j + 1}, 1)
				total += findUniqueReachableNines(grid, Location{x: i, y: j - 1}, 1)
			}
		}
	}
	fmt.Println(total)
}

func findUniqueReachableNines(grid [][]int, l Location, currentValue int) int {

	total := 0
	if !isWithinGrid(grid, l) {
		return 0
	}

	if grid[l.y][l.x] == currentValue {
		if currentValue == 9 {
			return 1
		}
		currentValue++
		total += findUniqueReachableNines(grid, Location{x: l.x + 1, y: l.y}, currentValue)
		total += findUniqueReachableNines(grid, Location{x: l.x - 1, y: l.y}, currentValue)
		total += findUniqueReachableNines(grid, Location{x: l.x, y: l.y + 1}, currentValue)
		total += findUniqueReachableNines(grid, Location{x: l.x, y: l.y - 1}, currentValue)
	}

	return total
}

func isWithinGrid(grid [][]int, l Location) bool {

	maxX := len(grid[0])
	maxY := len(grid)

	if l.x < 0 || l.x >= maxX || l.y < 0 || l.y >= maxY {
		return false
	}

	return true
}
