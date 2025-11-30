package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
	/* ここで, math.Sqrtは引数が 0 未満の場合 NaN を返すため虚数を返してほしい場合は上記のような string で表記するのが早い
	ただ, あくまで sqrt の結果を string で出力しているだけなので, この値を用いて計算をすると行ったことができないことに注意
	*/
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
