package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	black       = 0
	white       = 1
	transparent = 2
)

func main() {

	layerStrings := make([][]string, 1)
	image := make([][]int, 6)

	for i := 0; i < 6; i++ {
		image[i] = make([]int, 25)
		for j := 0; j < 25; j++ {
			image[i][j] = 2
		}
	}

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
			layerStrings = append(layerStrings, splitSubN(layer, 25))
		}
	}

	for _, layer := range layerStrings {
		image = overlayImage(image, layer)
	}

	for _, item := range image {
		fmt.Println(item)
	}
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

func overlayImage(image [][]int, layer []string) [][]int {

	i := 0
	for _, item := range layer {
		for j := 0; j < 25; j++ {
			num, _ := strconv.Atoi(string(item[j]))
			if image[i][j] == 2 {
				image[i][j] = num
			}
		}
		i++
	}

	return image
}
