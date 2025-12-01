package main

import "fmt"

func main() {
	var s []int //これではすっからかんのスライス s となる つまり長さも容量も 0 となる
	fmt.Println(s, len(s), cap(s))
	if s == nil { // すっからかんのスライスは nil と言われている ラテン語の nihil 由来 (意味は 0, なにもない 英語で言うところのempty)
		fmt.Println("nil!")
	}
}
