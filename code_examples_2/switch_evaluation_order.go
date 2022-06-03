package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday() // вызов пакета time, затем времени и дня недели
	switch time.Saturday { // установка на субботу
	case today + 0:
		fmt.Println("Today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In 2 days")
	case today + 4:
		fmt.Println("In 4 days")
	default:
		fmt.Println("Too far away")
	}
}

