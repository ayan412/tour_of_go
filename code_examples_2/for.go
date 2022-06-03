package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
// присвоение sum, i значение 0; i меньше 10; i++ - прибавление с шагом 1;
//sum += " adds the value of the right operand to a variable and assigns the result to the variable."
//--->т.е. присвоение значения переменой справа (i) переменой слева (sum)
/*
The basic for loop has three components separated by semicolons:
the init statement: executed before the first iteration
the condition expression: evaluated before every iteration
the post statement: executed at the end of every iteration
 */