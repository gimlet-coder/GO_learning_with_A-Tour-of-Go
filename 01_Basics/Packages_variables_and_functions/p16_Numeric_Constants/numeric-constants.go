package main

import "fmt"

const (
	Big = 1 << 100
	/* 1 を 左に 100回 ビットシフトすることで巨大な数を作っている
	これは 1 の後ろに 0 が 100 個続いている 2進数を表しているので 2^100 のことである
	*/
	Small = Big >> 99
	/* 反対に, 100 回左にシフトした数を 99 回右にシフトすると 1 回左にシフトした数と一致するので
	1 << 1 と同じ意味になり 2^1 = 2 となる
	*/
	// Big も Small もどちらも型がない定数であることに注目 引用される際に必要に応じた型に変化することができる
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
