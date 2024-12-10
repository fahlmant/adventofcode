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
				// Find unique 9s that are reachable by this location
				nineLocations := make(map[Location]bool)
				findUniqueReachableNines(nineLocations, grid, Location{x: i, y: j}, 0)
				total += len(nineLocations)
			}
		}
	}
	fmt.Println(total)
}

func findUniqueReachableNines(nineLocations map[Location]bool, grid [][]int, l Location, currentValue int) {

	if !isWithinGrid(grid, l) {
		return
	}

	if grid[l.y][l.x] == currentValue {
		if currentValue == 9 {
			nineLocations[l] = true
			return
		}
		currentValue++
		findUniqueReachableNines(nineLocations, grid, Location{x: l.x + 1, y: l.y}, currentValue)
		findUniqueReachableNines(nineLocations, grid, Location{x: l.x - 1, y: l.y}, currentValue)
		findUniqueReachableNines(nineLocations, grid, Location{x: l.x, y: l.y + 1}, currentValue)
		findUniqueReachableNines(nineLocations, grid, Location{x: l.x, y: l.y - 1}, currentValue)
	}

}

func isWithinGrid(grid [][]int, l Location) bool {

	maxX := len(grid[0])
	maxY := len(grid)

	if l.x < 0 || l.x >= maxX || l.y < 0 || l.y >= maxY {
		return false
	}

	return true
}
