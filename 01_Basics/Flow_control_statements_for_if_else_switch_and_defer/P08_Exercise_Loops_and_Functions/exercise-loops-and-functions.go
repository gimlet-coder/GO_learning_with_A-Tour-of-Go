package main

import (
	"fmt"
	"math"
)

// @breaf ニュートン法を用いて sqrt(x) の近似値を求める関数
func Sqrt(x float64) float64 {
	z := 1.0
	const times = 10 //ループの規定回数
	for count := 0; count < times; count++ {
		temp := z           //対比して後ろで比較する
		const Bound = 1e-15 // 計算を止めるための境目を入力
		z -= (z*z - x) / (2 * z)
		if math.Abs(temp-z) < Bound { // 絶対値が収まるという書き方でないと temp - z が負になった瞬間満たしてしまうので注意
			fmt.Printf("End, calc. time is %v. The current approx. of sqrt(%g) is %g \n", count, x, z)
			break
		}
	}
	return z
}

//実行結果から, ニュートン法はかなり早く収束することが確認できる

func main() {
	fmt.Println(Sqrt(2))
}
