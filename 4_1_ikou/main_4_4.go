package main

import (
	"fmt"
)

func main() {
	fmt.Println("\n■■■マップ■■■")
	m := make(map[int]string)
	m[1] = "US"
	m[81] = "Japan"
	fmt.Println(m)

	m2 := make(map[string]string)
	m2["Yamada"] = "Taro"
	m2["Sato"] = "Hanako"
	fmt.Println(m2)

	fmt.Println("\n■■■マップの中に、スライスを作成■■■")
	m3 := map[int][]int{
		1: []int{1},
		2: []int{1, 2},
		3: []int{1, 2, 3},
	}
	fmt.Println(m3)

	fmt.Println("\n■■■マップの中に、マップを作成■■■")
	m4 := map[int]map[float64]string{
		1: {3.14: "円周率"},
		2: {5.55: "テスト"},
	}
	fmt.Println(m4)

	fmt.Println("\n■■■要素の参参照■■")
	fmt.Println(m3[1])
	fmt.Println(m3[2])
	s, ok := m3[1]     //okには、要素の存在有無でtrue,falseが入る
	fmt.Println(s, ok) //[1] true
	s, ok = m3[4]
	fmt.Println(s, ok) //[] false

}
