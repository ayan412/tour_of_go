package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x<0 {
		return sqrt(-x) + "i" // задание в случае отриц-го х; комплексное число
	}
	return fmt.Sprint(math.Sqrt(x)) // задание вывода х в остальных случаях
}

func main() {
	fmt.Println(sqrt(2), sqrt(4)) // сам вывод
}
