package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	//rand.Intn は疑似乱数を返す おそらく(10) で 0~10 と値を制限している (1からかわからなかったが、複数回実行して0が出ることを確認済)
}
