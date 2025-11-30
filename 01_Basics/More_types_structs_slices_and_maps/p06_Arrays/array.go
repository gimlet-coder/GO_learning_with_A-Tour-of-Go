package main

import "fmt"

func main() {
	var a [2]string
	/* 配列を扱うとき, C++ でいうところの vector<string> a[2] は上記のように書く
	また, 配列のサイズを後から変えることはできないことに注意!!
	*/
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	// まとめて表記して初期化することもできる
	fmt.Println(primes)
}
