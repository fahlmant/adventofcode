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
		if containsSemiRepleatedLetter(line) && containsDoublePairs(line) {
			niceStrings++
		}
	}

	fmt.Println(niceStrings)
}

func containsSemiRepleatedLetter(line string) bool {

	for i := 0; i < len(line)-2; i++ {
		if line[i] == line[i+2] {
			return true
		}
	}

	return false

}

func containsDoublePairs(line string) bool {

	for i := 0; i < len(line)-2; i++ {
		if strings.Count(line, line[i:i+2]) >= 2 {
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
