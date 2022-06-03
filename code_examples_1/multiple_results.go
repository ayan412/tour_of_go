package main

import "fmt"

func swap(x, y, z string) (string, string, string) { /* здесь 3 string задает количестов переменных*/
	return z, y,x
}

func main() {
	a, b, c := swap("hello", "world", "my")
	fmt.Println(a, b, c)
}
// меняет местами введенные значения
