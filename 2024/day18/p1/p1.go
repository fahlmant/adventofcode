package main

import (
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	x, y int
}

type Node struct {
	pos  Vector
	dist int
}

func main() {

	total := 0

	file, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(file), "\n")

	var corruptedSpace []Vector

	for i := range 1024 {
		coords := strings.Split(string(input[i]), ",")
		x, err := strconv.Atoi(coords[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(coords[1])
		if err != nil {
			panic(err)
		}
		corruptedSpace = append(corruptedSpace, Vector{x: x, y: y})
	}

	total = bfs(71, 71, corruptedSpace)

	fmt.Println(total)
}

func bfs(col, row int, corruptedSpaces []Vector) int {

	shortestPath := 0

	// Build grid
	var grid [][]int
	for range row {
		newRow := make([]int, col)
		grid = append(grid, newRow)
	}

	// Fill in grid with corrupted spaces
	for _, cs := range corruptedSpaces {
		grid[cs.y][cs.x] = 1
	}

	// Perform bfs
	queue := list.New()

	startingNode := Node{pos: Vector{x: 0, y: 0}, dist: 0}
	// Push the starting node (0,0) onto the queue
	queue.PushBack(startingNode)

	visited := make(map[Vector]bool)
	visited[startingNode.pos] = true

	for queue.Len() > 0 {

		currentNode := queue.Remove(queue.Front()).(Node)

		// At the bottom, done
		if currentNode.pos.y == len(grid)-1 && currentNode.pos.x == len(grid[0])-1 {
			return currentNode.dist
		}

		for _, dir := range []Vector{{x: 1, y: 0}, {x: -1, y: 0}, {x: 0, y: 1}, {x: 0, y: -1}} {
			nextPosition := Vector{x: currentNode.pos.x + dir.x, y: currentNode.pos.y + dir.y}
			if isValidPosition(grid, nextPosition) && !visited[nextPosition] {
				queue.PushBack(Node{pos: nextPosition, dist: currentNode.dist + 1})
				visited[nextPosition] = true
			}
		}
	}

	return shortestPath
}

func isValidPosition(grid [][]int, v Vector) bool {

	if v.x < 0 || v.y < 0 || v.y >= len(grid) || v.x >= len(grid[v.y]) || grid[v.y][v.x] == 1 {
		return false
	}

	return true
}
