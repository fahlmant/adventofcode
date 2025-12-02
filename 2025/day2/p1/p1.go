package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	total := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		ranges := strings.Split(line, ",")
		for _, r := range ranges {
			split := strings.Split(r, "-")
			lowerNum, err := strconv.Atoi(split[0])
			if err != nil {
				log.Fatal(err)
			}
			upperNum, err := strconv.Atoi(split[1])
			if err != nil {
				log.Fatal(err)
			}

			for i := lowerNum; i <= upperNum; i++ {
				iStr := strconv.Itoa(i)
				if len(iStr)%2 != 0 {
					continue
				}
				if !isValid(iStr) {
					total += i
				}

			}

		}

	}

	fmt.Println(total)
}

func isValid(s string) bool {
	leftHalf := s[0 : len(s)/2]
	rightHalf := s[len(s)/2:]

	return !(leftHalf == rightHalf)
}
