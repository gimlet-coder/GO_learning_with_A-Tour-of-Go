package main

import "fmt"

func main() {
	pow := make([]int, 10) // スライスを int のゼロ値で初期化して長さと容量を 10 に設定
	for i := range pow {   // index や 値の必要な部分だけを使うこともできる その際不要な部分は _ とするが, 値が不要な場合は後ろの部分なので省略可能
		pow[i] = 1 << uint(i) // 符号なし整数 i だけ 1 を左にビットシフトする つまり 2^i を計算して代入している
	}
	for _, value := range pow { // index が不要なので _ で代用して書く
		fmt.Printf("%d \n", value)
	}
}
