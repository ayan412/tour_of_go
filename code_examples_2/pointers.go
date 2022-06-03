package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p) // вывод значения указателя.
	*p = 21
	fmt.Println(i) // вывод нового значения указателя.

	p = &j
	*p = *p / 37
	fmt.Println(j) // вывод результата деления 2701 на 37.
}