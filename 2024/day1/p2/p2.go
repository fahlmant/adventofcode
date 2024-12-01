package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	total := 0

	file, err := os.Open("../../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var leftList []int
	var rightList []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, "   ")

		// Get left number
		leftInt, err := strconv.Atoi(lineSplit[0])
		if err != nil {
			panic(err)
		}
		leftList = append(leftList, leftInt)

		// Get right number
		rightInt, err := strconv.Atoi(lineSplit[1])
		if err != nil {
			panic(err)
		}
		rightList = append(rightList, rightInt)
	}

	sort.Ints(leftList)
	sort.Ints(rightList)

	if len(leftList) != len(rightList) {
		panic("Lists are not the same length")
	}

	for _, v := range leftList {
		total += v * numOfOccurances(v, rightList)
	}

	fmt.Println(total)
}

func numOfOccurances(x int, list []int) int {

	total := 0
	for i := range list {
		if list[i] == x {
			total += 1
		}
	}
	return total
}
