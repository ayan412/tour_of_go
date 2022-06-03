package main

import "fmt"

func main() {
	sum, limit := 1, 100
	for sum < limit {
		fmt.Printf("sum before iteration: %d\nsum < limit %t\n", sum, sum < limit)
		sum += sum
		fmt.Printf("sum after iteration: %d\nsum < limit: %t\n\n", sum, sum < limit)
	}
}
