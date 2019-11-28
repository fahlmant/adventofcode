package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type coord struct {
	x int
	y int
}

func main() {

	var coordMap = make(map[coord]int)
	total := 0

	positionSanta := coord{0, 0}
	positionRobot := coord{0, 0}
	var position *coord

	line, err := ioutil.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	for index, b := range strings.Split(string(line), "") {
		if index%2 == 0 {
			position = &positionRobot
		} else {
			position = &positionSanta
		}
		if string(b) == "<" {
			position.x--
		} else if string(b) == ">" {
			position.x++
		} else if string(b) == "^" {
			position.y--
		} else if string(b) == "v" {
			position.y++
		}

		coordMap[*position]++
	}

	for _, v := range coordMap {
		if v > 0 {
			total++
		}
	}

	fmt.Println(len(coordMap))
}
