package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe   bool   = false
	MaxInt uint64 = 1<<64 - 1
	// 1<<64 は 2^64 を表している (<< は bit シフト演算) つまり, uint64 が扱える上限値の2^64 - 1 が MaxIntに入っている
	z complex128 = cmplx.Sqrt(-5 + 12i)
	// 前述の通り, cmplx内の関数といったようなパッケージから関数を引用するときには頭文字が大文字になることに注意!!
)

func main() {
	fmt.Printf("Type: %T Value:  %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
	//ここで, %T と大文字にしないと bool 値 (=全くの別物) が出力されてしまうので注意!! おそらく fmt パッケージから型が何かを引っ張ってきているため
}
