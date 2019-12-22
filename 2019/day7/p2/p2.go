package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/fahlmant/adventofcode/2019/day7/localintcode"
)

func main() {

	var instructionStrings []string
	max := 0

	file, err := os.Open("../input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		instructionStrings = strings.Split(line, ",")
	}

	instructions := make([]int, len(instructionStrings))
	ins1 := make([]int, len(instructionStrings))
	ins2 := make([]int, len(instructionStrings))
	ins3 := make([]int, len(instructionStrings))
	ins4 := make([]int, len(instructionStrings))
	ins5 := make([]int, len(instructionStrings))

	for i, v := range instructionStrings {
		instructions[i], _ = strconv.Atoi(v)
	}

	copy(ins1, instructions)
	copy(ins2, instructions)
	copy(ins3, instructions)
	copy(ins4, instructions)
	copy(ins5, instructions)
	//possiblePhaseSettings := buildCombinationsList(5, 10)
	possiblePhaseSettings := getPossibleCombos(5)

	for _, setting := range possiblePhaseSettings {

		fmt.Printf("Running phase setting: %+v\n", setting)
		//Run first iteration with phase settings
		

		computerA := localintcode.Computer{PC: 0, Offset: 0, Input: []int{setting[0], 0}, Output: []int{}, Instructions: ins1}
		computerA.RunProgram()

		computerB := localintcode.Computer{PC: 0, Offset: 0, Input: []int{setting[1]}, Output: []int{}, Instructions: ins2}	
		computerB.RunProgram()

		computerC := localintcode.Computer{PC: 0, Offset: 0, Input: []int{setting[2]}, Output: []int{}, Instructions: ins3}
		computerC.RunProgram()

		computerD := localintcode.Computer{PC: 0, Offset: 0, Input: []int{setting[3]}, Output: []int{}, Instructions: ins4}
		computerD.RunProgram()

		computerE := localintcode.Computer{PC: 0, Offset: 0, Input: []int{setting[4]}, Output: []int{}, Instructions: ins5}
		computerE.RunProgram()

		for computerE.Instructions[computerE.PC] != 99 {
			computerA.Input = computerE.Output
			computerA.RunProgram()

			computerB.Input = computerA.Output
			computerB.RunProgram()

			computerC.Input = computerB.Output
			computerC.RunProgram()

			computerD.Input = computerC.Output
			computerD.RunProgram()

			computerE.Input = computerD.Output
			computerE.RunProgram()

		}

		if computerE.Output[0] > max {
			max = computerE.Output[0]
		}
		fmt.Println(computerE.Output)
	}

	fmt.Println(max)

}

func buildCombinationsList(low, high int) [][]int {

	rand.Seed(time.Now().Unix())
	var combinationList [][]int

	x := 0
	for len(combinationList) < (factorial(high) - factorial(low)) {
		fmt.Printf("iteration: %d with %d of %d\n", x, len(combinationList), (factorial(high) - factorial(low)))
		var combination string
		for len(combination) < 5 {

			if len(combination) < 1 {
				combination = strconv.Itoa(rand.Intn(high-low) + low)
			}

			nextNum := rand.Intn(high-low) + low
			if !strings.Contains(combination, strconv.Itoa(nextNum)) {
				combination = strings.Join([]string{combination, strconv.Itoa(nextNum)}, "")
			}
		}

		for _, item := range combinationList {
			if equal(item, stringToIntArray(combination)) {

			}
		}
		if !sliceInSlice(combinationList, stringToIntArray(combination)) {
			combinationList = append(combinationList, stringToIntArray(combination))
		}
		x++
	}

	return combinationList
}

func stringToIntArray(value string) []int {

	var final []int
	for _, v := range value {
		num, _ := strconv.Atoi(string(v))
		final = append(final, num)
	}
	return final
}

func sliceInSlice(slice [][]int, value []int) bool {
	for _, v := range slice {
		if equal(value, v) {
			return true
		}
	}
	return false
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func factorial(n int) (result int) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

func getPossibleCombos(offset int) [][]int {
	combos := [][]int{}
	for a := 0 + offset; a < 5+offset; a++ {
		for b := 0 + offset; b < 5+offset; b++ {
			for c := 0 + offset; c < 5+offset; c++ {
				for d := 0 + offset; d < 5+offset; d++ {
					for e := 0 + offset; e < 5+offset; e++ {
						tester := map[int]bool{}
						tester[a] = true
						tester[b] = true
						tester[c] = true
						tester[d] = true
						tester[e] = true
						if len(tester) == 5 {
							combos = append(combos, []int{a, b, c, d, e})
						}
					}
				}
			}
		}
	}
	return combos
}
