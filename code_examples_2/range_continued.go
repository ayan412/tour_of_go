package main

import "fmt"

func main() {
	pow := make([]int, 10) // создание среза pow заполненного 10-ю нулями
	fmt.Println(pow) // печать среза
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i - For example, 1 << 5 is "1 times 2, 5 times" or 32.
	fmt.Println("Значение среза:", uint(i)) // печать значений среза
	}
	for _, value:= range pow {
		fmt.Printf("%d\n", value)
	}
}
