package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Coord struct {
	x int
	y int
}

func main() {

	total := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, line)
	}

	newChanges := 1
	for newChanges != 0 {
		grid, newChanges = scanGrid(grid)
		total += newChanges
	}

	fmt.Println(total)
}

func scanGrid(grid []string) ([]string, int) {

	newGrid := make([]string, len(grid))
	listOfRemovable := []Coord{}
	totalChanges := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == '@' {
				totalAdjacent := 0
				// Look left
				if j > 0 {
					if grid[i][j-1] == '@' {
						totalAdjacent += 1
					}

				}
				// Look right
				if j < len(grid[i])-1 {
					if grid[i][j+1] == '@' {
						totalAdjacent += 1
					}
				}
				// Look up
				if i > 0 {
					if grid[i-1][j] == '@' {
						totalAdjacent += 1
					}
				}
				// Look down
				if i < len(grid)-1 {
					if grid[i+1][j] == '@' {
						totalAdjacent += 1
					}
				}
				// Look left-up
				if j > 0 && i > 0 {
					if grid[i-1][j-1] == '@' {
						totalAdjacent += 1
					}
				}
				// Look left-down
				if j > 0 && i < len(grid)-1 {
					if grid[i+1][j-1] == '@' {
						totalAdjacent += 1
					}
				}
				// Look right-up
				if j < len(grid[i])-1 && i > 0 {
					if grid[i-1][j+1] == '@' {
						totalAdjacent += 1
					}
				}
				// Look right-down
				if j < len(grid[i])-1 && i < len(grid)-1 {
					if grid[i+1][j+1] == '@' {
						totalAdjacent += 1
					}
				}
				if totalAdjacent < 4 {
					listOfRemovable = append(listOfRemovable, Coord{x: i, y: j})
					totalChanges += 1
				}
			}
		}
	}

	copy(newGrid, grid)
	for _, c := range listOfRemovable {
		row := []rune(newGrid[c.x])
		row[c.y] = '.'
		newGrid[c.x] = string(row)
	}

	return newGrid, totalChanges
}
