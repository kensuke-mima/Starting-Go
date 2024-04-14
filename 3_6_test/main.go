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

	//3.7
	//型変換キャスト
	n1 := int(1)
	b := byte(n)
	i64 := int64(n)
	fmt.Println("KATAHENKAN", n1, b, i64)

	//rune型（シングルクォート）
	r := '松'
	fmt.Printf("%v\n", r) //=>26494

	//3.8 配列型
	a := [5]int{1, 2, 3, 4, 5}
	fmt.Printf("3.8 HAIRETSU %v", a) //=>[1,2,3,4,5]

}

func one() int {
	fmt.Println("func one nakami")
	return 1
}
