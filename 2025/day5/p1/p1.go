package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func main() {

	total := 0

	file, err := os.ReadFile("../input")
	if err != nil {
		log.Fatal(err)
	}

	input := strings.Split(string(file), "\n\n")

	rangeStrings := strings.Split(input[0], "\n")
	ingredientStrings := strings.Split(input[1], "\n")

	ingredientNumbers := make([]int, len(ingredientStrings))

	for i, v := range ingredientStrings {
		var err error
		ingredientNumbers[i], err = strconv.Atoi(v)
		if err != nil {
			log.Fatal(err)
		}
	}

	ranges := make([]Range, len(rangeStrings))

	for i, v := range rangeStrings {
		rangeSplit := strings.Split(v, "-")
		startOfRange, err := strconv.Atoi(rangeSplit[0])
		if err != nil {
			log.Fatal(err)
		}
		endOfRange, err := strconv.Atoi(rangeSplit[1])
		if err != nil {
			log.Fatal(err)
		}
		ranges[i] = Range{start: startOfRange, end: endOfRange}
	}

	for _, v := range ingredientNumbers {
		for _, r := range ranges {
			if v >= r.start && v <= r.end {
				total += 1
				break
			}
		}
	}

	fmt.Println(total)
}
