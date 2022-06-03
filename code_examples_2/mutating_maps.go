package main

import "fmt"

func main()  {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value: ", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value: ", m["Answer"])

	//delete(m, "Answer")
	//fmt.Println("The value: ", m["Answer"])

	val, ok := m["Answer"] // типа if
	fmt.Println("The value: ", val, "Present?", ok)
	/*
	Если значение по заданному ключу имеется в отображении, то переменная ok будет равна true,
	а переменная val будет хранить полученное значение.
	Если переменная ok равна false, то значения по ключу в отображении нет.
	 */



}