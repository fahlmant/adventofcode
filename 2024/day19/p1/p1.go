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
		if ok := isDesignPossible(towels, d); ok {
			total += 1
		}
	}

	fmt.Println(total)
}

func isDesignPossible(towels []string, design string) bool {

	var towelsAsBytes [][]byte

	for _, t := range towels {
		localTowel := make([]byte, len(t))
		copy(localTowel[:], t)
		towelsAsBytes = append(towelsAsBytes, localTowel)
	}

	localDesign := make([]byte, len(design))
	copy(localDesign[:], []byte(design))

	cache := make(map[string]bool)
	for _, t := range towels {
		cache[t] = true
	}

	isPossible := designCheck(towelsAsBytes, localDesign, cache, 0)

	return isPossible
}

func designCheck(towels [][]byte, target []byte, cache map[string]bool, depth int) bool {

	// If we already know the target (all of the rest of the string) is possible, don't check
	if b, exists := cache[string(target)]; exists {
		return b
	}

	// If the target is exauhsted, we've checked everything
	if len(target) == 0 {
		return true
	}

	for _, t := range towels {
		if len(t) > len(target) {
			continue
		}
		if strings.HasPrefix(string(target), string(t)) {
			result := designCheck(towels, target[len(t):], cache, depth+1)
			if result {
				cache[string(target)] = true
				return true
			}
		}
	}

	cache[string(target)] = false
	return false
}
