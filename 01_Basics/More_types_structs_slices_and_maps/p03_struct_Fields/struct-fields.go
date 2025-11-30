package main

import "fmt"

type Vertex struct {
	X int
	Y int // ここも, X, Y int と表すことができる
}

func main() {
	v := Vertex{1, 2}
	v.X = 4 // ドットを用いることでアクセスすることができるので, 2次配列や3次配列のような可読性の悪いものよりも優れている
	fmt.Println(v.X, v.Y)
	// 上部では X しか触っていないため, 当然 Y の値は変化していないことが実行結果からわかる
}
