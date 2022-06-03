package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 { // pow = x в степени n; второй float задает тип выходной переменной т.е. v
	if v := math.Pow(x, n); v < lim { // переменная v "живёт" только в области данной функции
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 200),
		)
}