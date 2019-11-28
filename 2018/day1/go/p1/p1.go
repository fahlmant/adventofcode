package main

import (
	"bufio"
	"fmt"

	"log"
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
		num, _ := strconv.Atoi(line[1:])
		if string(line[0]) == "+" {
			total += num
		} else if string(line[0]) == "-" {
			total -= num
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(total)
}
