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
		for i, b := range lineByte {
			if string(b) == "(" {
				total++
			} else if string(b) == ")" {
				total--
			}
			if total == -1 {
				fmt.Println(i + 1)
				os.Exit(0)
			}
		}
	}
}
