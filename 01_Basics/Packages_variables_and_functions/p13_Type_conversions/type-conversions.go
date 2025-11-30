package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	/* 前項目でやったように,
	x, y := 3, 4
	f := math.Sqrt(flort64(x*x + y*y))
	と表記することもできる (むしろ, Go の慣習的にこちらのほうが望ましいが, 今回は型変換がテーマなので意図的に明示している)
	当然, f は float64 と定義しているので, int で定義している x, y をそのまま使うことはできないためキャストしている
	また, math.Sqrt の引数には int を渡せないことにも注意が必要
	*/
	var z uint = uint(f)
	// f の値を z に代入したいとき, 型が異なるので castしたい型(変数) という書き方で型変換を行う
	fmt.Println(x, y, z)
	// 内容としては, ピタゴラス数を出力している
}
