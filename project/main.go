package main

import (
	"fmt"
	"project/calc"
)

var a, b float64
var c string

func main() {
	a = readFloat()
	c = readRune()
	b = readFloat()

	calculator := calc.NewCalc()
	result := calculator.Calculate(a, b, c)

	fmt.Println(result)
}

func readFloat() float64 {
	var number float64
	fmt.Print("Введите число: ")
	_, err := fmt.Scanln(&number)
	if err != nil {
		fmt.Println(fmt.Sprintf("Не получилось прочитать число: %f", err))

	}
	return number
}

func readRune() string {
	var symbol string
	fmt.Println("Введите знак выражения: ")
	_, err := fmt.Scanln(&symbol)
	if err != nil {
		fmt.Println(fmt.Sprintf("Не получилось прочитать знак выражения: %q", err))
	}
	return symbol
}
