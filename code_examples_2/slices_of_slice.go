package main

import (
	"fmt"
	"strings"
)

func main() {
	// задание формы для игры крестики-нолики
	board := [][]string{ // двумерный массив
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// Ходы игроков по индексам
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"
 // вывод
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}