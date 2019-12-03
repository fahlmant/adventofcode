package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Particle struct {
	px, py, vx, vy int
}

func main() {

	var particleList []*Particle
	var secondDist = make(map[int]int)

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		px, _ := strconv.Atoi(strings.TrimSpace(line[10:16]))
		py, _ := strconv.Atoi(strings.TrimSpace(line[18:24]))
		vx, _ := strconv.Atoi(strings.TrimSpace(line[36:38]))
		vy, _ := strconv.Atoi(strings.TrimSpace(line[40:42]))

		particleList = append(particleList, &Particle{px: px, py: py, vx: vx, vy: vy})
	}

	var minYDist int
	var keySecond int
	for seconds := 0; seconds < 100000; seconds++ {
		maxY := 0
		minY := particleList[0].py
		//maxX := 0
		//minX := particleList[0].px
		for _, particle := range particleList {
			if particle.py > maxY {
				maxY = particle.py
			}
			if particle.py < minY {
				minY = particle.py
			}
			particle.px += particle.vx
			particle.py += particle.vy
		}
		dist := maxY - minY
		if seconds == 0 {
			minYDist = dist
		} else {
			if minYDist > dist {
				minYDist = dist
				keySecond = seconds
			}
		}

		secondDist[seconds] = dist

		if seconds == 10558 {
			printGrid(particleList, 1000, 1000)
			os.Exit(0)
		}
	}

	fmt.Println(keySecond)
}

func printGrid(p []*Particle, x int, y int) {

	grid := make([][]string, x)
	for i := 0; i < x; i++ {
		grid[i] = make([]string, y)
		for j := 0; j < y; j++ {
			grid[i][j] = "-"
		}
	}

	for _, particle := range p {
		if particle.px < x && particle.py < y {
			grid[particle.px][particle.py] = "x"
		}

	}

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
			fmt.Printf("%s", strings.Trim(grid[i][j], "\n"))
		}
		fmt.Println()
	}

}
