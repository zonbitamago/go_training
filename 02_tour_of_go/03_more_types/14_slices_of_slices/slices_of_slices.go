package main

import (
	"fmt"
	"strings"
)

func main() {
	// マルバツゲームのボードを定義
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 選択された手順
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s \n", strings.Join(board[i], " "))
	}
}
