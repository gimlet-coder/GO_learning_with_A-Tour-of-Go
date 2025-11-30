package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	/* 条件がない switch は for{} のように true を省略して書くことができる
	if(){} else if(){}... と続けるよりも, go であれば switch{case hoge: ... case huga: ...} と書くほうが可読性が上がりとても良い
	*/
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
