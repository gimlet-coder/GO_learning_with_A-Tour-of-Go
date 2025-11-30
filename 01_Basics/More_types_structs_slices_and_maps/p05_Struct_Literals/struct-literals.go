package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	// 下のように X: 1, Y: 2 としても良いし, 省略もできる
	v2 = Vertex{X: 1}
	/* X に 1 が入るが Y が省略されている場合, 暗黙的に Y : 0 が設定される
	ただし, 上のように X: を省略してしまうと X にかかっているのか Y にかかっているのかわからない上に引数の数も一致しないのでエラーとなる
	*/
	v3 = Vertex{}
	// この場合, 各成分 X, Y は 0 となる
	p = &Vertex{1, 2}
	// 構造体のポインタも渡すことができる なお, 前セクションの通り(*p).X としなくても p.Xで参照することができる
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
