package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	x, y int
}

func area(point1, point2 Point) int {

	x := math.Abs(float64(point1.x-point2.x)) + 1
	y := math.Abs(float64(point1.y-point2.y)) + 1

	return int(x * y)
}

func main() {

	total := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	points := make([]Point, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := strings.Split(line, ",")
		x, _ := strconv.Atoi(splitLine[0])
		y, _ := strconv.Atoi(splitLine[1])
		point := Point{x, y}
		points = append(points, point)
	}

	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			area := area(points[i], points[j])
			if area > total {
				total = area
			}
		}
	}

	fmt.Println(total)
}
