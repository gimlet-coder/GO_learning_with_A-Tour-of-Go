package main

import "fmt"

func main() {
	v := 42
	// この v は後ろの初期化子から型推論しているため, 3.14 や 2 + 3i といったように変更すればそのように型が決まる
	fmt.Printf("v is of type %T\n", v)
}
