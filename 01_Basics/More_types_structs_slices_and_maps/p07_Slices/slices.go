package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	// 配列内の 1 ~ 3 番目の要素を切り取って扱うことができる 半開区間[low, high) であることと, 0-based index であることに注意
	fmt.Println(s)
}
