package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
/* func (имя_параметра тип_получателя) имя_метода (параметры) (типы_возвращаемых_результатов){
тело_метода
} */
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
// In this example, the Abs method has a receiver of type Vertex named v.
