package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		/*C++ のように for のスコープ内のみで使う変数は i:= 0 (もしくは i := int(0)) と明示する
		また, var i int と事前に定義してから for i = 0; i < 10; i++ とするのも可能だが,
		スコープ外で i を操作できてしまう点も踏まえて,
		スコープ内のみで完結するならば慣習的に i := 0 とするのが自然だろう
		*/
		sum += i
	}
	fmt.Println(sum)
}
