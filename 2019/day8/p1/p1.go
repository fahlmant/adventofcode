package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {

	layerStrings := make([]string, 1)
	minZeros := 0
	var index int

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		layers := splitSubN(line, 25*6)
		for _, layer := range layers {
			layerStrings = append(layerStrings, layer)
		}
	}

	for i, layer := range layerStrings {
		numZeros := countValueInLayer(layer, 0)
		if minZeros == 0 || numZeros < minZeros {
			minZeros = numZeros
			index = i
		}
	}

	fmt.Println(countValueInLayer(layerStrings[index], 1) * countValueInLayer(layerStrings[index], 2))

}

func splitSubN(s string, n int) []string {
	sub := ""
	subs := []string{}

	runes := bytes.Runes([]byte(s))
	l := len(runes)
	for i, r := range runes {
		sub = sub + string(r)
		if (i+1)%n == 0 {
			subs = append(subs, sub)
			sub = ""
		} else if (i + 1) == l {
			subs = append(subs, sub)
		}
	}

	return subs
}

func countValueInLayer(layer string, value int) int {

	count := 0

	for i := 0; i < len(layer); i++ {
		currentNum, _ := strconv.Atoi(string(layer[i]))
		if currentNum == value {
			count++
		}
	}

	return count
}
