package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	/*default:
		// windows, plan9
		fmt.Printf("%s.\n", os)
	*/
	}
}
/*
Go взял откуда-то системные параметры и вставил и вывод
A switch statement is a shorter way to write a sequence of if - else statements.
It runs the first case whose value is equal to the condition expression.
 */