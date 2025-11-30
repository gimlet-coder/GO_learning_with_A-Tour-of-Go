package main

import "fmt"

var c, python, java bool

/*最後に変数の型を明示しているので, C++ 流に表記するならば
bool c, python, java
となる var は変数を定義するときに明示するもの
*/

func main() {
	var i int
	fmt.Println(i, c, python, java)
	// int の初期値は 0 となっており, bool の初期値は false となっていることが実行結果からわかる
}
