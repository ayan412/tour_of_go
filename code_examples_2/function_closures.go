package main

import "fmt"

/* определение функции adder и 2-й??? функции func, к-рая возвращает целое число т.е. описывает ф-ю в return
Функция как результат другой функции
 */
func adder() func(int) int { // т.к. adder пуст, то в итоге ф-я adder возвр-т целое число
	sum := 0
	return func(x int) int {
		sum += x // sum = sum + x = sum; 0+0=0;0+1=1;1+2=3;3+3=6;6+4=10;10+5=15;15+6=21
		return sum
	}
}

func main() {
	//pos, neg := adder(), adder() // одновременное присвоение значений 2-м переменным
	pos := adder()
	neg := adder()
	for i := 0; i < 10; i++ { // цикл до 10 с увеличением на +1
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}