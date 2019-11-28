package main

import (
	"bufio"
	"fmt"
	"sort"
	"strings"

	"log"
	"os"
	"strconv"
)

func main() {

	var lines [1000]string
	total := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		lines[i] = scanner.Text()
		i++
	}

	for _, v := range lines {
		total += calculateNeededRibbon(v)
	}

	fmt.Println(total)
}

func calculateNeededRibbon(line string) int {

	res := strings.Split(line, "x")
	length, _ := strconv.Atoi(res[0])
	width, _ := strconv.Atoi(res[1])
	height, _ := strconv.Atoi(res[2])

	s := []int{length, width, height}
	sort.Ints(s)
	cubicFt := s[0] * s[1] * s[2]
	around := 2*s[0] + 2*s[1]

	return cubicFt + around
}
