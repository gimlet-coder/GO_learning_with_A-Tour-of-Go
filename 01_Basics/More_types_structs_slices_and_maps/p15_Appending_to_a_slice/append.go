package main

import "fmt"

func main() {
	var s []int   // nil のスライス s を定義
	printSlice(s) // nil なので 長さも容量も 0 で中身がない配列が出力される

	s = append(s, 0)
	// スライス s に対して 0 というリテラルを追加する nil にも append はできるので一旦 nil でスライスを作成してからちびちび増やすことができる
	printSlice(s)

	s = append(s, 1)
	// 現状のスライス s は {0} だが, これにリテラル 1 を追加したので {0, 1} となる 長さも容量も必要に応じて取り直すので 2 となる
	printSlice(s)

	s = append(s, 2, 3, 4)
	// スライス s = {0, 1} に リテラル 2, 3, 4 を追加する 複数まとめて追加することもできる 同様に, 必要になった分だけ長さと容量も取り直して 2 + 3 = 5 となる
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len = %d cap = %d %v \n", len(s), cap(s), s)
}
