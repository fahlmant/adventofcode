package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	results := make(map[int]map[rune]int, 250)
	i := 0
	twoCount := 0
	threeCount := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		results[i] = countLetters(line)
		if valueInMap(2, results[i]) {
			twoCount++
		}
		if valueInMap(3, results[i]) {
			threeCount++
		}
		i++
	}

	fmt.Println(twoCount)
	fmt.Println(threeCount)
	fmt.Println(twoCount * threeCount)

}

func countLetters(line string) map[rune]int {

	counts := make(map[rune]int, len(line))

	for _, char := range line {
		counts[char]++
	}

	return counts
}

func valueInMap(value int, valueMap map[rune]int) bool {

	for _, v := range valueMap {
		if v == value {
			return true
		}
	}
	return false
}
