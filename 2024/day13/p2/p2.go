package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Button struct {
	plusX int
	plusY int
}

type Prize struct {
	x int
	y int
}

type Machine struct {
	buttonA Button
	buttonB Button
	prize   Prize
}

func main() {

	total := 0

	file, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(file), "\n\n")

	var machines []Machine
	for _, i := range input {
		m := buildMachine(i)
		machines = append(machines, m)
	}

	for _, m := range machines {
		total += calculateMinimumTokens(m)
	}

	fmt.Println(total)
}

func calculateMinimumTokens(m Machine) int {

	// a1 = buttonA.x a2 = buttonA.y b1 = buttonB.x b2 = buttonB.y
	// a1 * timesAPressed + b1 * timesBPressed = x
	// a2 * timesAPressed + b2 * timesBpressed = y

	// Eliminate timesBPressed by multiplying top by b2 and bottom by b1
	// b2*a1*timesAPressed + b2*b1*timesBPressed = x*b2
	// b1*a2*timesApressed + b1*b2*timesBPressed = y*b1

	// x*b2 - y*b1 = b2*a1*timesAPressed - b1*a2*timesAPressed
	// x*b2 - y*b1 = (b2*a1 - b1*a2) timesAPressed
	//Solve for timesAPressed
	// timesAPressed = (x*b2 - y*b1) / (b2*a1 - b1*a2)

	// The reverse can be done to solve for timesBPressed
	// timesBPressed = (x*a2 - y*a1) / (b1*a2 - b2*a1)

	timesAPressed := ((m.prize.x * m.buttonB.plusY) - (m.prize.y * m.buttonB.plusX)) / ((m.buttonB.plusY * m.buttonA.plusX) - (m.buttonB.plusX * m.buttonA.plusY))
	timesBPressed := ((m.prize.x * m.buttonA.plusY) - (m.prize.y * m.buttonA.plusX)) / ((m.buttonB.plusX * m.buttonA.plusY) - (m.buttonB.plusY * m.buttonA.plusX))

	// If the number of times each button is pressed adds up to the prize from the solution above, we have a valid solution
	if (m.buttonA.plusX*timesAPressed+m.buttonB.plusX*timesBPressed == m.prize.x) && (m.buttonA.plusY*timesAPressed+m.buttonB.plusY*timesBPressed == m.prize.y) {
		// It takes 3 tokesn to press A
		return timesBPressed + (3 * timesAPressed)
	}
	return 0
}

// Ewww
// I probably don't need all this, and it can probably be done in a much cleaner way
func buildMachine(input string) Machine {

	// Set up split and regexes
	machineSplit := strings.Split(input, "\n")
	regexButtonX := regexp.MustCompile(`X\+(.*?),`)
	regexButtonY := regexp.MustCompile(`Y\+(.*)`)

	// A button
	matchesAX := regexButtonX.FindStringSubmatch(machineSplit[0])
	matchesAY := regexButtonY.FindStringSubmatch(machineSplit[0])

	buttonAX, err := strconv.Atoi(matchesAX[1])
	if err != nil {
		panic(err)
	}
	buttonAY, err := strconv.Atoi(matchesAY[1])
	if err != nil {
		panic(err)
	}

	buttonA := Button{plusX: buttonAX, plusY: buttonAY}

	// B button
	matchesBX := regexButtonX.FindStringSubmatch(machineSplit[1])
	matchesBY := regexButtonY.FindStringSubmatch(machineSplit[1])

	buttonBX, err := strconv.Atoi(matchesBX[1])
	if err != nil {
		panic(err)
	}
	buttonBY, err := strconv.Atoi(matchesBY[1])
	if err != nil {
		panic(err)
	}

	buttonB := Button{plusX: buttonBX, plusY: buttonBY}

	// Prize
	regexPrizeX := regexp.MustCompile(`X\=(.*?),`)
	regexPrizeY := regexp.MustCompile(`Y\=(.*)`)
	matchesPX := regexPrizeX.FindStringSubmatch(machineSplit[2])
	matchesPY := regexPrizeY.FindStringSubmatch(machineSplit[2])
	prizeX, err := strconv.Atoi(matchesPX[1])
	if err != nil {
		panic(err)
	}
	prizeY, err := strconv.Atoi(matchesPY[1])
	if err != nil {
		panic(err)
	}
	prize := Prize{x: prizeX + 10000000000000, y: prizeY + 10000000000000}

	m := Machine{buttonA: buttonA, buttonB: buttonB, prize: prize}

	return m
}
