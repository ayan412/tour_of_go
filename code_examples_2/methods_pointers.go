package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) { // метод Scale
	v.X = v.X * f
	v.Y = v.Y * f

	fmt.Println("in scale:", v.Abs())
}

func main() {
	v := Vertex{3, 4} // src
	v.Scale(10)
	fmt.Println("in main: ", v.Abs()) // 10*3 + 10*4 и затем квадр.корень из этой суммы
}
