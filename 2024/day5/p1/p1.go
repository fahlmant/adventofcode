package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Rule struct {
	start int
	end   int
}

func main() {

	total := 0

	file, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(file), "\n\n")

	rulesList := input[0]
	updatesList := input[1]

	rulesListSplit := strings.Split(rulesList, "\n")
	updatesListSplit := strings.Split(updatesList, "\n")

	var rules []Rule

	for i := range rulesListSplit {
		pages := strings.Split(string(rulesListSplit[i]), "|")
		start, err := strconv.Atoi(pages[0])
		if err != nil {
			panic(err)
		}
		end, err := strconv.Atoi(pages[1])
		if err != nil {
			panic(err)
		}
		rules = append(rules, Rule{start: start, end: end})
	}

	for i := range updatesListSplit {
		updateSplit := strings.Split(string(updatesListSplit[i]), ",")
		var updatesAsInt []int
		for _, value := range updateSplit {
			page, err := strconv.Atoi(value)
			if err != nil {
				panic(err)
			}
			updatesAsInt = append(updatesAsInt, page)
		}

		valid := true
		for _, rule := range rules {
			valid = validateUpdate(updatesAsInt, rule)
			if !valid {
				break
			}
		}
		if valid {
			total += updatesAsInt[len(updatesAsInt)/2]
		}
	}

	fmt.Println(total)
}

// Validate the update according to the rules
func validateUpdate(update []int, rule Rule) bool {

	if !intInSlice(update, rule.start) || !intInSlice(update, rule.end) {
		return true
	}
	firstFound := false
	for i := range update {
		if update[i] == rule.start {
			firstFound = true
		}
		if update[i] == rule.end {
			return firstFound
		}
	}

	return true
}

func intInSlice(list []int, target int) bool {
	for i := range list {
		if list[i] == target {
			return true
		}
	}

	return false
}
