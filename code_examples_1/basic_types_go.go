package main

import (
	"fmt"
	"math/cmplx"
)
var (
	ToBe   bool = false
	MaxInt uint64 = 1<<64 - 1
	z	   complex128 = cmplx.Sqrt(-5 + 12i)
)
// В var объявлены переменные. Далее объявлен их вывод: сначала Тип переменных, затем значение переменных
func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z )
}

/*
%v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
%T	a Go-syntax representation of the type of the value
%%	a literal percent sign; consumes no value
 */
