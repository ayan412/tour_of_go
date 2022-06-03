package main

import "fmt"

type Vertex struct {
	X int
	Y int
}
func main() {
	v := Vertex{1, 2}
	v.X = 5
	v.Y = 8
	fmt.Println(v.X, v.Y) // вывод без фигурных скобок
	fmt.Println(v) // вывод с фигурными скобками
}
// доступность через точку