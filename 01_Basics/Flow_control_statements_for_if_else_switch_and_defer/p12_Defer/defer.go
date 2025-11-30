package main

import "fmt"

func main() {
	defer fmt.Println("World")
	/*defer で渡した引数そのものはすぐに処理されるが, 実行されるのは最後なので部分的に計算前の値を保持したいときやあらかじめ最後に行いたい処理が決まっているときなどに有効
	また, defer を複数行書いた場合は, 上に書いたものが最後に実行されることに注意
	*/
	fmt.Println("hello")
}
