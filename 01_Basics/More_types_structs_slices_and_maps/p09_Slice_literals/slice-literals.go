package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13} //スライスのリテラルは長さを省略して定義する
	// 配列は後からサイズを変えられないが, スライスならば実質的に要素を追加することで可変長にできる
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}

	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	// struct を噛ませた後に定義することもできるので複数の入れ子構造が可能
	fmt.Println(s)
}
