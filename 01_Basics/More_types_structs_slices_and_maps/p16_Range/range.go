package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		// range でスライスは 2 つの変数を返すので, i と v を用意している (最初は index で2つ目は値 value を返すので i, v とした)
		//ループ範囲はスライスのサイズに一致する
		fmt.Printf("2 ** %d = %d,\n", i, v)
	}
}
