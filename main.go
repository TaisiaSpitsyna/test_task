package main

import (
	"fmt"
	"strconv"
	"strings"
)

func addition(a, b int) int {
	return a + b
}
func subtraction(a, b int) int {
	return a - b
}
func multiplication(a, b int) int {
	return a * b
}
func quotient(a, b int) int {
	return a / b
}

var romanSymbol = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}
var intValues = map[int]string{
	1000: "M",
	900:  "CM",
	500:  "D",
	400:  "CD",
	100:  "C",
	90:   "XC",
	50:   "L",
	40:   "XL",
	10:   "X",
	9:    "IX",
	5:    "V",
	4:    "IV",
	1:    "I",
}

func romanToInt(a string) int {
	first, ok := romanSymbol[rune(a[0])]
	if !ok {
		return -1
	}
	if len(a) == 1 {
		return first
	}
	second, ok := romanSymbol[rune(a[1])]
	if !ok {
		return -1
	}
	if second > first {
		next := romanToInt(a[2:])
		if next == -1 {
			return -1
		}
		return (second - first) + next
	}
	next := romanToInt(a[1:])
	if next == -1 {
		return -1
	}
	return first + next
}
func intToRoman(a int) string {
	var roman strings.Builder
	for value, symbol := range intValues {
		for a >= value {
			roman.WriteString(symbol)
			a -= value
		}
	}

	return roman.String()
}

func main() {
	var choice string
	var a string
	var b string
	var x int
	var y int
	var isroman bool
	var result int
	isroman = false
	fmt.Println("Введите выражение:")
	fmt.Scanln(&a, &choice, &b)
	x, err1 := strconv.Atoi(a)
	y, err2 := strconv.Atoi(b)
	if err1 != nil {
		if err2 != nil {
			x = romanToInt(a)
			y = romanToInt(b)
			isroman = true
			if y == -1 || x == -1 {
				fmt.Println("Недопустимый формат чисел")
				return
			}
		} else {
			fmt.Println("Числа разного формата")
			return
		}
	} else if err2 != nil {
		fmt.Println("Числа разного формата")
		return
	}
	if x < 1 || x > 10 || y < 1 || y > 10 {
		fmt.Println("Числа не входят в диапазон")
		return
	}
	if choice == "+" {
		result = addition(x, y)
	} else if choice == "-" {
		result = subtraction(x, y)
	} else if choice == "*" {
		result = multiplication(x, y)
	} else if choice == "/" {
		result = quotient(x, y)
	} else {
		fmt.Println("Ошибка")
		return
	}
	if isroman && result <= 0 {
		fmt.Println("Ошибка, в римской системе нет отрицательных чисел.")
		return
	}
	if isroman {
		fmt.Println(intToRoman(result))
	} else {
		fmt.Println(result)
	}

}
