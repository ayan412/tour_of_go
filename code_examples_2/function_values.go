package main

import (
	"fmt"
	"math"
)

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 { // зачем 2-й тип float64??? тип возвращаемого значения
		return math.Sqrt(x*x + y*y) // возведение в степень и квадратный корень
	}
	fmt.Println(hypot(5, 12)) //
	fmt.Println(compute(hypot)) // возведение в степень и квадратный корень
	fmt.Println(compute(math.Pow)) // возведение в степень 4

}