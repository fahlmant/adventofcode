package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	orbits := make(map[string]string)
	youorbits := map[string]int{}

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), ")")
		orbits[line[1]] = line[0]
	}

	for o, ok := orbits["YOU"]; ok; o, ok = orbits[o] {
		youorbits[o] = len(youorbits)
	}

	distance := 0
	for orbit, ok := orbits["SAN"]; ok; orbit, ok = orbits[orbit] {
		if _, ok := youorbits[orbit]; ok {
			fmt.Println(distance + youorbits[orbit])
			break
		}
		distance++
	}

}
