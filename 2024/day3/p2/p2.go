package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {

	total := 0

	file, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	input := strings.Replace(string(file), "\n", "", -1)
	excludeRegex := regexp.MustCompile(`don't\(\).*?(?:do\(\))`)
	input = excludeRegex.ReplaceAllString(input, "")

	regex := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)

	// -1 means give us all matches
	mulStatements := regex.FindAllStringSubmatch(input, -1)

	for _, statement := range mulStatements {
		x, err := strconv.Atoi(statement[1])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(statement[2])
		if err != nil {
			panic(err)
		}

		total += (x * y)
	}

	fmt.Println(total)
}
