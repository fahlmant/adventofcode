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

const SECONDS = 100
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

	for range SECONDS {
		grid = runStep(grid)
	}

	// Count up the number of Robots in each quadrant
	result := calculateQuadrants(grid)

	total = (result[0] * result[1] * result[2] * result[3])

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

func calculateQuadrants(grid map[Vector][]Robot) []int {

	q1Total := 0
	q2Total := 0
	q3Total := 0
	q4Total := 0

	for j := 0; j < height/2; j++ {
		for i := 0; i < width/2; i++ {
			pos := Vector{x: i, y: j}
			q1Total += len(grid[pos])
		}
		for k := width/2 + 1; k < width; k++ {
			pos := Vector{x: k, y: j}
			q2Total += len(grid[pos])
		}
	}

	for j := height/2 + 1; j < height; j++ {
		for i := 0; i < width/2; i++ {
			pos := Vector{x: i, y: j}
			q3Total += len(grid[pos])
		}
		for k := width/2 + 1; k < width; k++ {
			pos := Vector{x: k, y: j}
			q4Total += len(grid[pos])
		}
	}

	return []int{q1Total, q2Total, q3Total, q4Total}

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
