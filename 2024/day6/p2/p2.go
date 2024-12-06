package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Location struct {
	x int
	y int
}

type DirectionLocation struct {
	location  Location
	direction string
}

func main() {

	total := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var grid [][]byte

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var lineArray []byte
		for i := range len(line) {
			lineArray = append(lineArray, line[i])
		}
		grid = append(grid, lineArray)
	}

	position := Location{x: 0, y: 0}
	// Find starting place of ^
	for j := range grid {
		found := false
		for i := range grid[j] {
			if grid[j][i] == '^' {
				position = Location{x: i, y: j}
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	// Since it's ^, we start going north
	currentDir := "N"

	trackedPositions := make(map[Location]bool)

	startPosition := position

	for {
		nextLocation := getLocationInDirection(currentDir, position)
		// If the next location if off the grid, we're done
		if nextLocation.x < 0 || nextLocation.x >= len(grid[0]) || nextLocation.y < 0 || nextLocation.y >= len(grid) {
			break
		}
		// If the next step is an obstacle, rotate 90 degrees to the right
		if grid[nextLocation.y][nextLocation.x] == '#' {
			currentDir = rotateDirection(currentDir)
			continue
		}
		// Set the current position to the next position and add it to the visited locations
		position = nextLocation
		trackedPositions[nextLocation] = true
	}

	for loc := range trackedPositions {
		copyGrid := make([][]byte, len(grid))
		for i := range grid {
			gridLine := make([]byte, len(grid[i]))

			copy(gridLine, grid[i])
			copyGrid[i] = gridLine
		}

		copyGrid[loc.y][loc.x] = '#'
		hasLoop := doesConfigHaveLoop(copyGrid, startPosition)
		if hasLoop {
			total += 1
		}
	}

	fmt.Println(total)
}

func doesConfigHaveLoop(grid [][]byte, startPos Location) bool {

	position := startPos
	currentDir := "N"
	visitedWithDir := make(map[DirectionLocation]bool)
	for {
		nextLocation := getLocationInDirection(currentDir, position)
		// If the next location if off the grid, we're done
		if nextLocation.x < 0 || nextLocation.x >= len(grid[0]) || nextLocation.y < 0 || nextLocation.y >= len(grid) {
			break
		}
		// If the next step is an obstacle, rotate 90 degrees to the right
		if grid[nextLocation.y][nextLocation.x] == '#' {
			currentDir = rotateDirection(currentDir)
			continue
		}
		// Set the current position to the next position and add it to the visited locations
		position = nextLocation
		dirWithPos := DirectionLocation{location: position, direction: currentDir}
		if visitedWithDir[dirWithPos] {
			return true
		}
		visitedWithDir[DirectionLocation{location: position, direction: currentDir}] = true
	}

	return false
}

// All rotations is 90 degrees to the right
func rotateDirection(direction string) string {

	switch direction {
	case "N":
		return "E"
	case "E":
		return "S"
	case "S":
		return "W"
	case "W":
		return "N"
	default:
		return direction
	}
}

func getLocationInDirection(direction string, currentLocation Location) Location {

	// Above one in y, y-1
	if direction == "N" {
		return Location{x: currentLocation.x, y: currentLocation.y - 1}
	}
	// Below one in y, y+1
	if direction == "S" {
		return Location{x: currentLocation.x, y: currentLocation.y + 1}
	}
	// Left in x, x - 1
	if direction == "W" {
		return Location{x: currentLocation.x - 1, y: currentLocation.y}
	}
	// Right in x, x+1
	if direction == "E" {
		return Location{x: currentLocation.x + 1, y: currentLocation.y}
	}

	return currentLocation
}
