package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	var grid [1000][1000]bool

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")
		if split[0] == "toggle" {
			toggle(&grid, split)
		} else if split[0] == "turn" {
			turn(&grid, split)
		}
	}

	total := 0
	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if grid[i][j] == true {
				total++
			}
		}
	}

	fmt.Println(total)
}

func toggle(grid *[1000][1000]bool, line []string) {

	xy1 := strings.Split(line[1], ",")
	xy2 := strings.Split(line[3], ",")

	x1, _ := strconv.Atoi(xy1[0])
	x2, _ := strconv.Atoi(xy2[0])
	y1, _ := strconv.Atoi(xy1[1])
	y2, _ := strconv.Atoi(xy2[1])

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grid[i][j] = !grid[i][j]
		}
	}
}

func turn(grid *[1000][1000]bool, line []string) {

	var boolValue bool

	if line[1] == "on" {
		boolValue = true
	} else if line[1] == "off" {
		boolValue = false
	}

	xy1 := strings.Split(line[2], ",")
	xy2 := strings.Split(line[4], ",")

	x1, _ := strconv.Atoi(xy1[0])
	x2, _ := strconv.Atoi(xy2[0])
	y1, _ := strconv.Atoi(xy1[1])
	y2, _ := strconv.Atoi(xy2[1])

	for i := x1; i <= x2; i++ {
		for j := y1; j <= y2; j++ {
			grid[i][j] = boolValue
		}
	}

}
