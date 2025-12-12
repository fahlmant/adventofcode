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

	beams := make(map[int]int, 0)

	for i := range grid[0] {
		if grid[0][i] == 'S' {
			beams[i] = 1
		}
	}

	for _, j := range grid[1:] {
		for b := range beams {
			if j[b] == '.' {
				continue
			}
			if j[b] == '^' {
				beams[b-1] += beams[b]
				beams[b+1] += beams[b]
				delete(beams, b)
			}
		}
	}

	for _, num := range beams {
		total += num
	}
	fmt.Println(total)
}
