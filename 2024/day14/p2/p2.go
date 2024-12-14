package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	x, y int
}

type Robot struct {
	pos Vector
	vel Vector
}

const SECONDS = 10000
const width = 101
const height = 103

func main() {

	total := 0

	file, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(file), "\n")

	grid := make(map[Vector][]Robot)

	for _, s := range input {
		r := buildRobot(s)
		grid[r.pos] = append(grid[r.pos], r)
	}

	for i := range SECONDS {
		grid = runStep(grid)
		// Just assuming the tree will have all values of 1
		allLensOne := true
		for _, v := range grid {
			if len(v) > 1 {
				allLensOne = false
			}
		}
		if allLensOne {
			// Adding 1 second since we start at the 0th seconds
			total = i + 1
			break
		}
	}

	printGrid(grid)

	fmt.Println(total)
}

func buildRobot(s string) Robot {

	parts := strings.Fields(s)

	var pX, pY, vX, vY int

	// Extract and parse the values
	fmt.Sscanf(parts[0], "p=%d,%d", &pX, &pY)
	fmt.Sscanf(parts[1], "v=%d,%d", &vX, &vY)

	return Robot{pos: Vector{x: pX, y: pY}, vel: Vector{x: vX, y: vY}}
}

func runStep(grid map[Vector][]Robot) map[Vector][]Robot {

	newGrid := make(map[Vector][]Robot)

	for _, v := range grid {
		for _, robot := range v {
			newX := robot.pos.x + robot.vel.x
			newY := robot.pos.y + robot.vel.y
			if newX >= width {
				newX -= width
			}
			if newY >= height {
				newY -= height
			}
			if newX < 0 {
				newX += width
			}
			if newY < 0 {
				newY += height
			}
			newPos := Vector{x: newX, y: newY}
			newRobot := Robot{pos: newPos, vel: robot.vel}
			newGrid[newPos] = append(newGrid[newPos], newRobot)
		}
	}

	return newGrid
}

// Nice helper
func printGrid(grid map[Vector][]Robot) {
	for j := range height {
		for i := range width {
			len := len(grid[Vector{x: i, y: j}])
			if len == 0 {
				fmt.Printf(".")
			} else {
				numString := strconv.Itoa(len)
				fmt.Printf("%s", numString)
			}
		}
		fmt.Printf("\n")
	}
}
