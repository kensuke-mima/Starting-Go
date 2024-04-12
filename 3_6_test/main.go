package main

import "fmt"

func main() {
	var n int

	var x, y, z int
	x, y, z = 1, 2, 3
	fmt.Println(x, y, z)

	/*

		var (
			x1, y1 int
			name   string
		)
	*/

	n = 5
	fmt.Println(n)

	i := 1 //int型の変数iを定義して１を代入（暗黙でint型）
	fmt.Println(i)

	fmt.Println("func one", one())
}

func one() int {
	fmt.Println("func one nakami")
	return 1
}
