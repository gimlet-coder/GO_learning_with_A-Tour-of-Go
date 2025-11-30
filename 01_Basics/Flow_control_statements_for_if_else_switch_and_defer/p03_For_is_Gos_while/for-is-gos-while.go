package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		/* 必ずしも初期化と後処理するステートメントは書く必要はないため, while 文のような振る舞いをさせることも可能
		ちなみに Go には while 文が存在しないため, 多言語の while と同じような処理がしたい場合はこの通りに書く必要がある
		*/
		sum += sum
	}
	fmt.Println(sum)
}
