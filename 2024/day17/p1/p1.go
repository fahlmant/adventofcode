package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ThreeBitComputer struct {
	a, b, c int
	ip      int
	prog    []int
}

func main() {

	file, err := os.ReadFile("../input")
	if err != nil {
		panic(err)
	}

	input := strings.Split(string(file), "\n\n")

	registersInput := strings.Split(input[0], "\n")
	registerA, err := strconv.Atoi(strings.Split(registersInput[0], ": ")[1])
	if err != nil {
		panic(err)
	}
	registerB, err := strconv.Atoi(strings.Split(registersInput[1], ": ")[1])
	if err != nil {
		panic(err)
	}
	registerC, err := strconv.Atoi(strings.Split(registersInput[2], ": ")[1])
	if err != nil {
		panic(err)
	}

	var program []int
	programString := strings.Split(input[1], ": ")
	for _, v := range strings.Split(programString[1], ",") {
		vNum, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		program = append(program, vNum)
	}

	computer := ThreeBitComputer{a: registerA, b: registerB, c: registerC, ip: 0, prog: program}
	computer.Run()
}

func (c *ThreeBitComputer) Run() {

	getComboOperand := func(op int) int {

		switch op {
		case 4:
			return c.a
		case 5:
			return c.b
		case 6:
			return c.c
		}
		return 0
	}
	for c.ip < len(c.prog)-1 {
		switch c.prog[c.ip] {
		case 0:
			//fmt.Println("adv")
			denom := c.prog[c.ip+1]
			if c.prog[c.ip+1] >= 4 {
				denom = getComboOperand(c.prog[c.ip+1])
			}
			c.a >>= denom
			c.ip += 2
		case 1:
			//fmt.Println("bxl")
			c.b ^= c.prog[c.ip+1]
			c.ip += 2
		case 2:
			//fmt.Println("bst")
			num := c.prog[c.ip+1]
			if c.prog[c.ip+1] >= 4 {
				num = getComboOperand(c.prog[c.ip+1])
			}
			c.b = num % 8
			c.ip += 2
		case 3:
			//fmt.Println("jnz")
			if c.a == 0 {
				c.ip += 2
			} else {
				c.ip = c.prog[c.ip+1]
			}
		case 4:
			//fmt.Println("bxc")
			c.b ^= c.c
			c.ip += 2
		case 5:
			//fmt.Println("out")
			num := c.prog[c.ip+1]
			if c.prog[c.ip+1] >= 4 {
				num = getComboOperand(c.prog[c.ip+1])
			}
			num = num % 8
			fmt.Printf("%d,", num)
			c.ip += 2
		case 6:
			//fmt.Println("bdv")
			denom := c.prog[c.ip+1]
			if c.prog[c.ip+1] >= 4 {
				denom = getComboOperand(c.prog[c.ip+1])
			}
			c.b = c.a >> denom
			c.ip += 2
		case 7:
			//fmt.Println("cdv")
			denom := c.prog[c.ip+1]
			if c.prog[c.ip+1] >= 4 {
				denom = getComboOperand(c.prog[c.ip+1])
			}
			c.c = c.a >> denom
			c.ip += 2
		default:
			fmt.Println("Unrecognized instruction")
		}
	}
	fmt.Printf("\n")
}
