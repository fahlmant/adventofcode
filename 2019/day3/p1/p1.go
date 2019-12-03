package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

type coord struct {
	x, y int
}

func main() {

	input, _ := ioutil.ReadFile("../input")
	wires := strings.Fields(string(input))
	var wirePoints []map[coord]int
	var commonPoints []coord

	for _, wire := range wires {
		seen := make(map[coord]int)
		currentX := 0
		currentY := 0
		moves := strings.Split(wire, ",")
		for _, move := range moves {
			direction := string(move[0])
			distance, _ := strconv.Atoi(move[1:])
			for i := 0; i < distance; i++ {
				switch direction {
				case "R":
					currentX++
				case "L":
					currentX--
				case "U":
					currentY++
				case "D":
					currentY--
				}
				seen[coord{x: currentX, y: currentY}]++
			}
		}

		wirePoints = append(wirePoints, seen)
	}

	for k1 := range wirePoints[0] {
		if _, ok := wirePoints[1][k1]; ok {
			commonPoints = append(commonPoints, k1)
		}
	}

	var min int
	for _, v := range commonPoints {
		mDist := math.Abs(float64(v.x)) + math.Abs(float64(v.y))
		if min == 0 || int(mDist) < min {
			min = int(mDist)
		}
	}

	fmt.Println(min)
}
