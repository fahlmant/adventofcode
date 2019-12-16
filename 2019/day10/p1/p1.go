package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

type slope struct {
	x, y int
}

func main() {

	var lines []string
	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())

	}

	fmt.Println(findMaxSeenAsteroids(lines))

}

func findMaxSeenAsteroids(lines []string) int {

	max := 0
	grid := make([][]int, len(lines))
	for i, line := range lines {
		grid[i] = make([]int, len(line))
		for j := 0; j < len(line); j++ {
			if string(line[j]) == "#" {
				grid[i][j] = 1
			}
		}
	}

	for i := 0; i < len(lines); i++ {
		for j := 0; j < len(lines[i]); j++ {
			if grid[i][j] == 1 {
				count := countAsteroids(grid, i, j)
				if count > max {
					max = count
				}
			}
		}
	}

	return max
}

func countAsteroids(grid [][]int, i, j int) int {

	uniqueSlopes := make(map[slope]bool)

	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 1 && !(i == x && j == y) {
				gcds := math.Abs(float64(gcd(x-i, y-j)))
				uniqueSlopes[slope{(x - i) / int(gcds), (y - j) / int(gcds)}] = true

			}
		}
	}

	return len(uniqueSlopes)
}

func gcd(x, y int) int {
	for y != 0 {

		x, y = y, x%y
	}
	return x
}
