package main

import (
	"fmt"
	"math"
)
// Суть примера в том, чтобы показать, что необходимо преобразовывать типы переменных при каждом их использовании
func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
/*
То же, но короче:
x, y := 3, 4
f := math.Sqrt(float64(x*x + y*y))
z := uint(f)
*/

