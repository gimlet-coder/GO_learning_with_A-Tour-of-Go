package main

import "fmt"

func add(x int, y int) int { // 変数は後ろで型を定義する void関数ならば func hoge(void) と返り値を省略する main 関数参照
	// C++ における int x, y のように x, y int とまとめて型を定義することも可能
	return x + y
}

func main() { // Go における main 関数では C++ と異なり return 0 や return 1 がないから func main() int と書かない
	// 明示したい場合には os.Exit(1) などを用いる このとき, import "os" が必要
	fmt.Println(add(42, 13))
}
