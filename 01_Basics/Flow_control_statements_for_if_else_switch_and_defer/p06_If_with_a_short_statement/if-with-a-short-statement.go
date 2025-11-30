package main

import (
	"fmt"
	"math"
)

// @breaf べき乗を計算する関数 ただし, lim を超えない範囲で計算し超えてしまうようであれば lim の値を返す
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		/* if 文においても for 文と同じようにローカル変数を定義することができる
		この場合, if のスコープを抜けたときに v にアクセスできなくなる
		*/
		return v
	}
	return lim
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 10),
	)
}
