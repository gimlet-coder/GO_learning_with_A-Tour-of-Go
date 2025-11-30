package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9 // 本来は (*p).X と表記するべきだが, Go では省略できる (すごい)
	fmt.Println(v)
}
