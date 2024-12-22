package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"log"
	"math"
	"os"
)

type Vector struct {
	x, y int
}

// Node tracks information about each position in the grid
type Node struct {
	position          Vector
	cost              int
	previousDirection Vector
	index             int
}

var NORTH = Vector{x: 0, y: -1}
var EAST = Vector{x: 1, y: 0}
var SOUTH = Vector{x: 0, y: 1}
var WEST = Vector{x: -1, y: 0}

// Priority queue type to use with heap
type PQ []*Node

// Functions to implement heap interface
func (pq PQ) Len() int {
	return len(pq)
}

func (pq PQ) Less(a, b int) bool {
	return pq[a].cost < pq[b].cost
}

func (pq PQ) Swap(a, b int) {
	pq[a], pq[b] = pq[b], pq[a]
	pq[a].index = a
	pq[b].index = b
}

func (pq *PQ) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PQ) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil
	node.index = -1
	*pq = old[0 : n-1]
	return node
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

	total = findShortestPath(grid, startingPosition, endingPosition)

	fmt.Println(total)
}

func findShortestPath(grid [][]byte, startPosition, endPosition Vector) int {

	shortestDistances := make(map[Vector]int)

	pq := make(PQ, 0)
	heap.Init(&pq)

	startNode := Node{
		position:          startPosition,
		cost:              0,
		previousDirection: EAST,
	}

	pq.Push(&startNode)
	shortestDistances[startPosition] = 0

	for pq.Len() > 0 {
		current := heap.Pop(&pq).(*Node)

		if current.position == endPosition {
			return current.cost
		}

		if dist, ok := shortestDistances[current.position]; ok && dist < current.cost {
			continue
		}

		for _, dir := range []Vector{NORTH, SOUTH, EAST, WEST} {
			newPosition := Vector{x: current.position.x + dir.x, y: current.position.y + dir.y}

			if !isValidPosition(grid, newPosition) {
				continue
			}

			newCost := current.cost + 1
			if dir != current.previousDirection {
				newCost += 1000
			}

			if oldCost, ok := shortestDistances[newPosition]; !ok || newCost < oldCost {
				shortestDistances[newPosition] = newCost

				heap.Push(&pq, &Node{
					position:          newPosition,
					cost:              newCost,
					previousDirection: dir,
				})
			}

		}

	}

	return math.MaxInt
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
