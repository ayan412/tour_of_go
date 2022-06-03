package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX" // индекс указывает на 1-ый элемент среза т.е. на "Paul", это значение будет заменено в срезе и списке
	fmt.Println(a, b)
	fmt.Println(names)
}