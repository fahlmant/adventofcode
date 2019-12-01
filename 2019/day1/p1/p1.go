package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {

	var total int

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		mass, _ := strconv.ParseFloat(line, 64)
		total += int(math.Floor(mass/3) - 2)
	}

	fmt.Println(total)
}
