package main

import (
	"fmt"
	"math"
)

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
	/*ここの return は if のスコープ外なので v を使えないことに注意
	もし v を else 処理の際に使いたい場合には中に入れる必要がある
	*/
}

func main() {
	fmt.Println(
		pow(3, 2, 10),
		pow(3, 3, 10),
		/*　ここで, fmt.Println は内部の pow の処理がすべて終わってから実行されるため
		fmt.Println(9, fmt.printf("%g >= %g\n", 27, 10), 10)
		となり, また内部の処理が優先されるため
		27 >= 10
		9 10
		といった出力がされる
		*/
	)
}
