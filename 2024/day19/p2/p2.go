package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	total := 0
	file, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(file), "\n\n")

	towels := strings.Split(input[0], ", ")

	desiredDesigns := strings.Split(input[1], "\n")

	for _, d := range desiredDesigns {
		total += isDesignPossible(towels, d)
	}

	fmt.Println(total)
}

func isDesignPossible(towels []string, design string) int {

	var towelsAsBytes [][]byte

	for _, t := range towels {
		localTowel := make([]byte, len(t))
		copy(localTowel[:], t)
		towelsAsBytes = append(towelsAsBytes, localTowel)
	}

	localDesign := make([]byte, len(design))
	copy(localDesign[:], []byte(design))

	cache := make(map[string]int)

	isPossible := designCheck(towelsAsBytes, localDesign, cache)

	return isPossible
}

func designCheck(towels [][]byte, target []byte, cache map[string]int) int {

	// If we already know the target (all of the rest of the string) is possible, don't check
	if count, exists := cache[string(target)]; exists {
		return count
	}

	// If the target is exauhsted, we've checked everything
	if len(target) == 0 {
		return 1
	}

	runningCount := 0
	for _, t := range towels {
		if len(t) > len(target) {
			continue
		}
		if strings.HasPrefix(string(target), string(t)) {
			runningCount += designCheck(towels, target[len(t):], cache)
		}
	}

	cache[string(target)] = runningCount
	return runningCount
}
