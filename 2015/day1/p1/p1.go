package main

import (
	"bufio"
	"fmt"

	"log"
	"os"
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
		lineByte := []byte(line)
		for _, b := range lineByte {
			if string(b) == "(" {
				total++
			} else if string(b) == ")" {
				total--
			}
		}
	}

	fmt.Println(total)
}
