package main

import "fmt"

const Pi = 3.14

/* 関数の外で定数や変数を定義するとき, 様々な場所で用いたい場合は頭文字を大文字にする
また, 数学的に意味がある数や Eps = 0.01 のように価値のある定数には大文字にする慣習がある
*/

func main() {
	const World = "世界"
	// World はローカル変数なので world と表記するので十分だが, 慣習的に頭文字を大文字にするプログラマが多いためこの表記を採用した
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	/* Println では フォーマット指定子(%v など)を使えない
	Printf ならば C 言語のように fmt.Printf("Happy %v Day \n", Pi) と書くことができる
	ただし Println のように自動で改行はしてくれないため, 同様な出力がしたい場合は最後に \n を付ける必要がある
	*/
	const Truth = true
	fmt.Println("Go rules?", Truth)
}
