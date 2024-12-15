package main

import (
	"fmt"
	"os"
	"strings"
)

type Vector struct {
	x, y int
}

func main() {

	total := 0

	file, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(file), "\n\n")

	var grid [][]byte
	var moves []byte

	// Create grid
	for _, l := range strings.Split(input[0], "\n") {
		grid = append(grid, []byte(l))
	}

	// Create list of instructions
	for _, m := range strings.Split(input[1], "\n") {
		moves = append(moves, []byte(m)...)
	}

	for _, m := range moves {
		moveRobot(grid, m)
	}

	// Calculate scores
	for j := range grid {
		for i := range grid[j] {
			if grid[j][i] == 'O' {
				total += (j * 100) + (i)
			}
		}
	}

	fmt.Println(total)
}

func moveRobot(grid [][]byte, move byte) {

	// Find robot location
	robotLocation := Vector{x: 0, y: 0}
	for j := range grid {
		for i := range grid[j] {
			if grid[j][i] == '@' {
				robotLocation = Vector{x: i, y: j}
			}
		}
	}

	dir := getDirection(move)

	var moveLocations []Vector
	nextPos := Vector{x: robotLocation.x + dir.x, y: robotLocation.y + dir.y}
	for {
		// Check if we're out of bounds. This should never happen
		if !isWithinGrid(grid, nextPos) {
			fmt.Println(robotLocation)
			fmt.Println(grid[robotLocation.y][robotLocation.x])
			fmt.Println(nextPos)
			panic("Out of bounds")
		}
		// Check the value of the next location
		// If it's a #, then we can't move because there are no open spots between the robot and the wall
		// If it's a ., we're done and we can move everything
		// If it's a O, it's a box and we can move it

		nextSpace := grid[nextPos.y][nextPos.x]
		if nextSpace == '#' {
			moveLocations = []Vector{}
			return
		}
		moveLocations = append(moveLocations, nextPos)
		if nextSpace == '.' {
			break
		}
		nextPos = Vector{x: nextPos.x + dir.x, y: nextPos.y + dir.y}
	}

	for i := len(moveLocations); i > 0; i-- {
		p := Vector{x: moveLocations[i-1].x, y: moveLocations[i-1].y}
		grid[p.y][p.x] = grid[p.y-dir.y][p.x-dir.x]
	}
	grid[robotLocation.y][robotLocation.x] = '.'

}

func getDirection(move byte) Vector {

	dir := Vector{x: 0, y: 0}
	switch move {
	case '^':
		dir = Vector{x: 0, y: -1}
	case 'v':
		dir = Vector{x: 0, y: 1}
	case '<':
		dir = Vector{x: -1, y: 0}
	case '>':
		dir = Vector{x: 1, y: 0}
	default:
		panic("Unrecognized character")
	}

	return dir
}

func isWithinGrid(grid [][]byte, v Vector) bool {

	if v.x < 0 || v.y < 0 || v.y >= len(grid) || v.x >= len(grid[v.y]) {
		return false
	}

	return true
}

func printGrid(grid [][]byte) {
	for j := range grid {
		for i := range grid[j] {
			fmt.Printf("%s", string(grid[j][i]))
		}
		fmt.Printf("\n")
	}
}
