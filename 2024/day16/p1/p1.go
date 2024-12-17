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

	for n := 2; n < len(path); n++ {
		currentPos := path[n]
		prevPos := path[n-1]
		prevPrevPos := path[n-2]

		dx1 := currentPos.x - prevPos.x
		dx2 := prevPos.x - prevPrevPos.x

		dy1 := currentPos.y - prevPos.y
		dy2 := prevPos.y - prevPrevPos.y

		// If the change in x is 0, then two positions on the same row,
		// if the change in y is 0, then two positions are on the same column
		// So if between three consecutive nodes in the path, there is a change from row to column, we've moved 90 degrees
		if (dx1 == 0 && dy2 == 0) || (dx2 == 0 && dy1 == 0) {
			score += 1000
		}
	}

	// If the second node in the path is not on the same X, we've turned from east
	if path[1].y != path[0].y {
		score += 1000
	}

	return score
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
