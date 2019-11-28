package main

import (
	"bufio"
	"fmt"

	"log"
	"os"
)

type coord struct {
	x int
	y int
}

func main() {

	var coordMap = make(map[coord]int)
	total := 0
	position := coord{0, 0}

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lineByte := []byte(line)
		for _, b := range lineByte {
			if string(b) == "<" {
				position.x--
			} else if string(b) == ">" {
				position.x++
			} else if string(b) == "^" {
				position.y--
			} else if string(b) == "v" {
				position.y++
			}

			coordMap[position]++
		}
	}

	for _, v := range coordMap {
		if v > 0 {
			total++
		}
	}

	fmt.Println(total + 1)
}
