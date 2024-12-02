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

	var reports [][]int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineSplit := strings.Split(line, " ")

		var report []int
		for i := range lineSplit {
			num, err := strconv.Atoi(lineSplit[i])
			if err != nil {
				panic(err)
			}
			report = append(report, num)
		}

		reports = append(reports, report)
	}

	for _, r := range reports {

		// Assume the report is true unless we find a reason for it to be false
		safe := true
		increasing := true
		for j := range r {
			if !safe {
				break
			}
			// If we're on the last number, we're done
			if j == len(r)-1 {
				break
			}

			result := r[j] - r[j+1]
			// If we're on the first number, we need to check if it's increasing or decreasing
			if j == 0 {
				if result > 0 {
					increasing = false
				}
			}

			// Needs to increase/decrase by at most 3 and at least 1
			if absInt(result) > 3 || absInt(result) < 1 {
				safe = false
				break
			}

			if r[j] > r[j+1] && increasing {
				safe = false
				break
			}
			if r[j] < r[j+1] && !increasing {
				safe = false
				break
			}
		}

		if safe {
			total += 1
		}
	}

	fmt.Println(total)
}

func absInt(x int) int {

	if x < 0 {
		return -x
	}
	return x
}
