package calc

import "fmt"

type calculator struct{}

func NewCalc() calculator {
	return calculator{}
}

//(calc *calculator)

func (calc *calculator) Calculate(a, b float64, c string) float64 {
	var total float64
	switch c {
	case "+":
		total = add(a, b)
	case "-":
		total = sub(a, b)
	case "*":
		total = multi(a, b)
	case "/":
		total = div(a, b)
	default:
		fmt.Println("Ошибка!")
	}
	return total
}

func add(a, b float64) float64 {
	return a + b
}

func sub(a, b float64) float64 {
	return a - b
}

func multi(a, b float64) float64 {
	return a * b
}

func div(a, b float64) float64 {
	if b == 0 {
		panic("На ноль делить нельзя!")
	}
	return a / b
}
