package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numberA string
	fmt.Print("Input:")
	fmt.Fscan(os.Stdin, &numberA)
	operationparts := strings.Split(numberA, " ")
	var a = strconv.Atoi(operationparts[0])
	var b = strconv.Atoi(operationparts[2])
	if err == nil {
		if operationparts[1] == "+" {
			Addition(operationparts[0], operationparts[2])
		}
		if operationparts[1] == "-" {
			Subtraction(operationparts[0], operationparts[2])
		}
		if operationparts[1] == "*" {
			Multiplication(operationparts[0], operationparts[2])
		}
		if operationparts[1] == "/" {
			Division(operationparts[0], operationparts[2])
		}

	}

	roman := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	toroman := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
	}

}

func Addition(numberA, numberB int) int {
	sum := numberA + numberB
	return sum
}

func Subtraction(numberA, numberB int) int {
	difference := numberA - numberB
	return difference
}

func Multiplication(numberA, numberB int) int {
	product := numberA * numberB
	return product
}

func Division(numberA, numberB int) int {
	quotient := numberA / numberB
	return quotient
}
