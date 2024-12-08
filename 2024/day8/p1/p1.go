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

				distanceX := antennaLocations[j].x - antennaLocations[i].x
				distanceY := antennaLocations[j].y - antennaLocations[i].y

				antinodeLocationA := Location{x: antennaLocations[i].x - distanceX, y: antennaLocations[i].y - distanceY}
				antinodeLocationB := Location{x: antennaLocations[j].x + distanceX, y: antennaLocations[j].y + distanceY}

				for _, location := range []Location{antinodeLocationA, antinodeLocationB} {
					if location.x < len(grid[0]) && location.x >= 0 && location.y < len(grid) && location.y >= 0 {
						antinodeLocations[location] = true
					}
				}
			}
		}
	}

	total = len(antinodeLocations)

	fmt.Println(total)
}
