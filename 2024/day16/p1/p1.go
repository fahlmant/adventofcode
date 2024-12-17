package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Vector struct {
	x, y int
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
		var lineData []byte
		for l := range line {
			lineData = append(lineData, line[l])
		}
		grid = append(grid, lineData)
	}

	var startingPosition Vector
	var endingPosition Vector
	for j := range grid {
		for i := range grid[j] {
			if grid[j][i] == 'S' {
				startingPosition = Vector{x: i, y: j}
			} else if grid[j][i] == 'E' {
				endingPosition = Vector{x: i, y: j}
			}
		}
	}

	visited := make(map[Vector]bool)
	path := []Vector{}
	allPaths := [][]Vector{}
	allPaths = findAllPaths(grid, startingPosition, endingPosition, visited, path, allPaths)
	minScore := calculateScore(allPaths[0])
	for i := 1; i < len(allPaths); i++ {
		newScore := calculateScore(allPaths[i])
		if newScore < minScore {
			minScore = newScore
		}
	}

	total = minScore

	fmt.Println(total)
}

func findAllPaths(grid [][]byte, currentPosition, endingPosition Vector, visited map[Vector]bool, path []Vector, allPaths [][]Vector) [][]Vector {

	if currentPosition == endingPosition {
		allPaths = append(allPaths, path)
		return allPaths
	}

	visited[currentPosition] = true

	for _, dir := range []Vector{{x: 1, y: 0}, {x: -1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1}} {
		newPos := Vector{x: currentPosition.x + dir.x, y: currentPosition.y + dir.y}
		if isValidPosition(grid, newPos) && !visited[newPos] {
			newPath := append(path, newPos)
			allPaths = findAllPaths(grid, newPos, endingPosition, visited, newPath, allPaths)
		}
	}

	visited[currentPosition] = false

	return allPaths
}

func calculateScore(path []Vector) int {

	score := 0

	score += len(path)

	turns := countTurns(path)

	// Check if the second node is a turn from east
	if path[1].y != path[0].y {
		turns += 1
	}

	score += turns * 1000
	return score

}

func countTurns(path []Vector) int {
	total := 0
	if len(path) < 3 {
		return 0
	}

	for i := 2; i < len(path); i++ {
		prev := path[i-1]
		beforePrev := path[i-2]
		current := path[i]

		// Calculate movement directions
		dx1, dy1 := prev.x-beforePrev.x, prev.y-beforePrev.y
		dx2, dy2 := current.x-prev.x, current.y-prev.y

		// A 90-degree turn occurs if the movement switches axes
		if (dx1 == 0 && dy2 == 0) || (dy1 == 0 && dx2 == 0) {
			total += 1
		}
	}
	return total
}

func isVectorIn(list []Vector, target Vector) bool {

	for _, v := range list {
		if v == target {
			return true
		}
	}
	return false
}

func isValidPosition(grid [][]byte, v Vector) bool {

	if v.x < 0 || v.y < 0 || v.y >= len(grid) || v.x >= len(grid[v.y]) || grid[v.y][v.x] == '#' {
		return false
	}

	return true
}

func printGrid(grid [][]byte, path []Vector) {

	for j := range grid {
		for i := range grid[j] {
			pos := Vector{x: i, y: j}
			if isVectorIn(path, pos) {
				fmt.Printf("x")
			} else {
				fmt.Printf(string(grid[j][i]))
			}
		}
		fmt.Printf("\n")
	}
}
