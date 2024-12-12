package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Location struct {
	x, y int
}

type Shape struct {
	locations []Location
	perimeter int
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

	shapesMap := findContigousShapes(grid)

	for _, shapeList := range shapesMap {
		for _, shape := range shapeList {
			total += calculateFenceCost(shape)
		}
	}
	fmt.Println(total)
}

func findContigousShapes(grid [][]byte) map[byte][]Shape {

	shapes := make(map[byte][]Shape)

	visited := make(map[Location]bool)

	for j := range grid {
		for i := range grid[j] {
			l := Location{x: i, y: j}
			letter := grid[l.y][l.x]
			if !visited[l] {
				shapes[letter] = append(shapes[letter], findShape(grid, visited, letter, l))
			}
		}
	}
	return shapes
}

func findShape(grid [][]byte, visited map[Location]bool, letter byte, location Location) Shape {

	var shape Shape

	// If the position is outside the grid (the previous letter is a border) or the position contains another letter
	// we know the edge between this position and the previous is a border of the shape
	if !isWithinGrid(grid, location) || grid[location.y][location.x] != letter {
		shape.perimeter = 1
		return shape
	}

	if visited[location] {
		return shape
	}

	visited[location] = true

	shape.locations = []Location{location}

	nextLocations := []Location{
		{x: location.x + 1, y: location.y},
		{x: location.x - 1, y: location.y},
		{x: location.x, y: location.y + 1},
		{x: location.x, y: location.y - 1},
	}

	for _, nl := range nextLocations {
		nextShape := findShape(grid, visited, letter, nl)
		shape.locations = append(shape.locations, nextShape.locations...)
		shape.perimeter += nextShape.perimeter
	}

	return shape
}

func calculateFenceCost(s Shape) int {

	total := 0

	area := len(s.locations)

	total = area * s.perimeter

	return total
}

func isWithinGrid(grid [][]byte, l Location) bool {

	maxX := len(grid[0])
	maxY := len(grid)

	if l.x < 0 || l.x >= maxX || l.y < 0 || l.y >= maxY {
		return false
	}

	return true
}
