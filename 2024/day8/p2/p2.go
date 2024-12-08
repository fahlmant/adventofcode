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

	// Keep list of locations per unique antenna character
	antennaLocationMap := make(map[byte][]Location)
	for j := range grid {
		for i := range grid[j] {
			if grid[j][i] != '.' {
				antennaLocationMap[grid[j][i]] = append(antennaLocationMap[grid[j][i]], Location{x: i, y: j})
			}
		}
	}

	antinodeLocations := make(map[Location]bool)
	for _, antennaLocations := range antennaLocationMap {
		// For each pair of antennas of the same character,
		// calculate antinode positions and add them to the locations
		for i := 0; i < len(antennaLocations)-1; i++ {
			// Offset j to be one more than i
			for j := i + 1; j < len(antennaLocations); j++ {

				// Each antenna is now an antinode
				antinodeLocations[Location{x: antennaLocations[i].x, y: antennaLocations[i].y}] = true
				antinodeLocations[Location{x: antennaLocations[j].x, y: antennaLocations[j].y}] = true

				distanceX := antennaLocations[j].x - antennaLocations[i].x
				distanceY := antennaLocations[j].y - antennaLocations[i].y

				// Move from the antenna to the edge of the grid in increments of the distance
				// First backwards, then forewards
				newLocationA := Location{x: antennaLocations[i].x - distanceX, y: antennaLocations[i].y - distanceY}
				for {
					if newLocationA.x >= len(grid[0]) || newLocationA.x < 0 || newLocationA.y >= len(grid) || newLocationA.y < 0 {
						break
					}
					antinodeLocations[newLocationA] = true
					newLocationA = Location{x: newLocationA.x - distanceX, y: newLocationA.y - distanceY}
				}

				newLocationB := Location{x: antennaLocations[i].x + distanceX, y: antennaLocations[i].y + distanceY}
				for {
					if newLocationB.x >= len(grid[0]) || newLocationB.x < 0 || newLocationB.y >= len(grid) || newLocationB.y < 0 {
						break
					}
					antinodeLocations[newLocationB] = true
					newLocationB = Location{x: newLocationB.x + distanceX, y: newLocationB.y + distanceY}
				}
			}
		}
	}

	total = len(antinodeLocations)

	fmt.Println(total)
}
