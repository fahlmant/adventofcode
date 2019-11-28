package main

import (
	"bufio"
	"fmt"
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
		total += calculateNeededPaper(v)
	}

	fmt.Println(total)
}

func calculateNeededPaper(line string) int {

	res := strings.Split(line, "x")
	length, _ := strconv.Atoi(res[0])
	width, _ := strconv.Atoi(res[1])
	height, _ := strconv.Atoi(res[2])

	side1 := length * width
	side2 := width * height
	side3 := height * length

	extra := findMin([]int{side1, side2, side3})

	return (2 * side1) + (2 * side2) + (2 * side3) + extra
}

func findMin(values []int) int {

	min := values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}

	return min
}
