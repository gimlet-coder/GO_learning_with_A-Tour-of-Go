package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2] // 0, 1番目の要素が入る
	b := names[1:3] // 1, 2番目の要素が入る
	fmt.Println(a, b)

	b[0] = "XXX"
	/* ここで, b[0] に当たる names[1] = "Paul" が "XXX" に置き換わる
	つまり, names の参照をしているだけなので参照元にも影響が生じる
	対策するには値として初期化子を入れる必要がある
	b[0] の値を共有している a[1] にも同様に影響が生じる
	*/
	fmt.Println(a, b)
	fmt.Println(names)
}
