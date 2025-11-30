package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}
	/* 表記自体は C++ と似ているものの, それぞれの case の処理において break が自動で付帯しているため 1 つの処理しか行われないことに注意
	例: switch hoge:= 10; hoge{
	case hoge > 0
		fmt.Println(hoge, " is positive number")
	case hoge * hoge >= 0
		fmt.Println(hoge, "is real number")
	}
	と表記しても最初の case が処理されて終わってしまう
	*/
}
