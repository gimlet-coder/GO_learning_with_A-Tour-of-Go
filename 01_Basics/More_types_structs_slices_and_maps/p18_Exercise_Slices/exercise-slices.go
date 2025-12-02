package main

import "golang.org/x/tour/pic"

// @breaf 長さ dy の slice に各要素が 8bit の符号なし整数で長さ dx のスライスを割り当てたものを返す
func Pic(dx, dy int) [][]uint8 {
	s := make([][]uint8, dy) // まず全体の箱を作る このとき, dy 行あるようにする
	for i := range s {
		s[i] = make([]uint8, dx) // s のそれぞれの行に 長さ dx のスライスを当てる
	}
	for i := range s {
		for j := range s[i] {
			s[i][j] = uint8(i ^ ^j) // i ^ ^j は i と (j のビット反転) の排他的論理和となる
		}
	}
	return s
}

func main() {
	pic.Show(Pic)
}
