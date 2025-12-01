package main

import "fmt"

func main() {
	a := make([]int, 5)
	/* make でスライスを作成できる int を float64にすればリテラルの型も変更できる 初期化子がないことから自動的に 0 が入る
	make([]型, 長さ, 容量) と書くが, 容量は省略することもできる その場合 長さ = 容量 となる
	*/
	printSlice("a", a)

	b := make([]int, 0, 5) // 長さが 0 の容量が 5 のスライスなので, 空のスライスで 5 だけ長さが伸ばせるスライスができる
	printSlice("b", b)

	c := b[:2] // b の容量は 5 なので, 長さを 2 にしても問題ない 初期化子がないことから各々のリテラルは 0 となる
	printSlice("c", c)

	d := c[2:5] // 長さは 5 - 2 = 3 となり, 開始ポインタが 2 番目 (0-index) となる 容量も開始ポインタを変更したことから 5 - 2 = 3 となる
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len = %d cap = %d %v\n", s, len(x), cap(x), x)
}
