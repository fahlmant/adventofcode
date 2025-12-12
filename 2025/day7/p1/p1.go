package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	total := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	grid := [][]byte{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		grid = append(grid, []byte(line))
	}

	for i := range grid {
		for j := range grid[i] {
			switch grid[i][j] {
			case 'S':
				grid[i+1][j] = '|'
			case '.':
				if i > 0 {
					if grid[i-1][j] == '|' {
						grid[i][j] = '|'
					}
				}
			case '^':
				if i == 0 {
					continue
				}
				if i > 0 {
					if grid[i-1][j] != '|' {
						continue
					}
				}
				total += 1
				if j > 0 {
					grid[i][j-1] = '|'
				}
				if j < len(grid[i])-1 {
					grid[i][j+1] = '|'
				}
			default:
				continue
			}
		}
	}
	fmt.Println(total)
}
