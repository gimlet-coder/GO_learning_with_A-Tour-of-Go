package main

import (
	"fmt"
	"strings"
)

func main() {
	board := [][]string{
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
		[]string{"-", "-", "-"},
	}
	/* 初期状態の枠には - をプロットしている 初期化子があるため, []string の型指定は省略可能 (VSCodeで波線が出ているところ)
	つまり,
	hoge = [][]string{
		{ , , }
		{ , , }
		{ , , }
	} みたいにかける
	*/

	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][1] = "X"
	// マルバツゲームみたいなことをやってるけど何故か X が先行
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
		// 行ごとに貼り付ける
	}
}
