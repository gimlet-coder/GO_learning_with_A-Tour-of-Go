package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3
	c, python, java := true, false, "no!"
	/* := を使うことで var と方の明示が不要となる ただし, 関数外では使用できないので注意が必要 (あくまでその関数内のみで扱うローカル変数のみに適用するべき)
	一方で, Go に置いては可読性を上げる観点から極力簡素な表現が好まれるので, := が使用できる場合には用いることが推奨される
	*/

	fmt.Println(i, j, k, c, python, java)
}
