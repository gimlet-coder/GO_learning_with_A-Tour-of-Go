package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to the playground! ")

	fmt.Println("The time is", time.Now())
	// どうして Now で Web上だと 2009年の11月30日固定なの? Apple の9:41 みたいに特別な意味がありそう
}
