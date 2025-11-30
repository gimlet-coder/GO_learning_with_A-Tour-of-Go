package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i // ここで, p は i のアドレスを参照しているので, 結果的に i の中身を変えることができる
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j // この時点で, p と i のつながりは完全になくなったので以降の処理で i の値に影響はない
	*p = *p / 37
	fmt.Println(j)
}
