package main

import "fmt"

func swap(x, y string) (string, string) {
	// Go において, 引数のコピーを関数内で扱う このことにより, C++のような参照渡しではないので元の値に干渉することは基本できない
	return y, x
}

func main() {
	a, b := swap("hello", "world") // 変数の初期化には := を用いる このとき, いきなり後ろに関数を呼び出してその結果を反映させることも可能
	/*分解して書くならば,
		var a string
		var b string
		a, b = swap("hello", "world")
	となる
	:= を用いることにより, 型推論で a と b は string に決定される
	*/
	fmt.Println(a, b)
}
