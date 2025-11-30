package main

import "fmt"

var i, j int = 1, 2

//C++ 流に書くならば int i = 1, j = 2 だが, まとめて書くことができる

func main() {
	var c, python, java = true, false, "no!"
	/* 初期化する段階で, その初期化子が持つ型が自動的に変数の型となる
	つまり,
	var x, y = 1, 2.5
	println(x + y)
	とすると int と float64 の足し算となってしまい Go ではエラーが発生してしまうので注意が必要
	ちなみに, float64 は C++ における double と同じ bit数を明示する書き方をする
	つまり, C++ における float は go において float32 と表記する
	*/
	fmt.Println(i, j, c, python, java)
}
