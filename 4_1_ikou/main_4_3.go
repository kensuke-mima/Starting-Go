package main

import (
	"fmt"
)

// 複数のint型引数を受け取って、その合計値を返却する関数
func sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println("\n■■■スライス■■■")
	s := make([]int, 10) //スライスを定義
	fmt.Println(s, len(s))

	var a [10]int //配列を定義
	fmt.Println(a)

	//
	fmt.Println("\n■■■cap■■■")
	s2 := make([]int, 5, 10) //要素数5, 容量10
	fmt.Println(s2, len(s2), cap(s2))

	//
	fmt.Println("\n■■■スライスを生成するリテラル■■■")
	s3 := []int{1, 2, 3, 4, 5} //要素数5, 容量5のスライスを定義
	fmt.Println(s3, len(s3), cap(s3))

	//
	fmt.Println("\n■■■簡易スライス式■■■")
	a4 := [5]int{1, 2, 3, 4, 5} //配列
	s4 := a4[0:2]
	fmt.Println(s4, len(s4), cap(s4))

	//
	fmt.Println("\n■■■append■■■")
	s5 := []int{1, 2, 3} //要素数3, 容量3のスライスを定義
	fmt.Println(s5, len(s5), cap(s5))
	s5 = append(s5, 4) //要素数４、容量６になる
	fmt.Println(s5, len(s5), cap(s5))

	//
	fmt.Println("\n■■■copy■■■")
	//省略

	//
	fmt.Println("\n■■■スライスと可変長引数■■■")
	//省略
	fmt.Printf("%v\n", sum(1, 2, 3))
	fmt.Printf("%v\n", sum())
	s6 := []int{1, 2, 3}
	fmt.Printf("%v\n", sum(s6...)) //sum(s6)とするとエラーになる。...を付加するとスライスの要素数を可変長引数とする

	//
	fmt.Println("\n■■■参照型としてのスライス■■■")
	//配列は値渡し
	//スライスは参照渡し
	var (
		a7 [3]int
		s7 []int
	)
	fmt.Printf("%v\n", a7)
	fmt.Printf("%v\n", s7)
	fmt.Println(s7 == nil) //⇒true。nilを作成したければ、空のスライスを作ればよい

}
