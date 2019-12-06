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
	total := 0

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

	for orbit := range orbits {
		for v, ok := orbits[orbit]; ok; v, ok = orbits[v] {
			total++
		}
	}

	fmt.Println(total)
}
