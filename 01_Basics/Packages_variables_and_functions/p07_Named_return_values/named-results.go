package main

import "fmt"

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// int 型の数字をx : y = 4 : 5 となるように分けて return している
	return
	/* return に何も書かずに戻すことを naked return という この場合, 関数の定義通り x, y の値を自動で返す
	可読性が悪くなるので短い関数以外では明示するのが良い
	明示する場合,
	return x, y
	となる このとき, 関数を定義する際に型は明記しているので表記はしない (C++ とおなじような形になる)
	*/
}

func main() {
	fmt.Println(split(17))
}
