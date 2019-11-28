package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {

	niceStrings := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if containsThreeVowels(line) && containsLetterPair(line) && !containsForbiddenStrings(line) {
			niceStrings++
		}
	}

	fmt.Println(niceStrings)
}

func containsThreeVowels(line string) bool {

	vowels := 0
	for _, vowel := range []string{"a", "e", "i", "o", "u"} {
		vowels += strings.Count(line, vowel)
	}

	return vowels >= 3
}

func containsLetterPair(line string) bool {

	for i := 0; i < len(line)-1; i++ {
		if line[i] == line[i+1] {
			return true
		}
	}
	return false
}

func containsForbiddenStrings(line string) bool {

	for _, forbiddenString := range []string{"ab", "cd", "pq", "xy"} {
		if strings.Contains(line, forbiddenString) {
			return true
		}
	}
	return false
}
