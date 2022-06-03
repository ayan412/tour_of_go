package main

import "fmt"

type Vertex struct { // определение структуры
	Lat, Long float64
}

var m map[string]Vertex //определение структуры, но почему VERTEX, а не float64 и т.п.??

func main()  {
	m = make(map[string]Vertex) // создание отображения?? Функция make представляет альтернативный вариант создания отображения
	m["Bell Labs"] = Vertex{ // заполнение отображения ключом и значением
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"]) // печать по ключу
}
