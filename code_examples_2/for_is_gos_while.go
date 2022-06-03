package main

import "fmt"

func main() {
	sum := 1
	for sum < 4 {
		sum += sum
	}
	fmt.Println(sum)
}
/*
 2 = 1+1
 4 = 2+2
 6 = 3+3

 */
