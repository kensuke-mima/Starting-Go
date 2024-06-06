package main

import (
	"fmt"
)

const ONE = 1

func one() (int, int) {
	const TWO = 2
	return ONE, TWO
}

// 2個目の戻り値がエラーのになるようにしたい。。
func doSomething() (string, string) {
	return "no probrem", "error"
}

func anything(a interface{}) {
	fmt.Println(a)
}

func main() {
	x, y := one()
	fmt.Println(x, y)

	//if文
	if true {
		fmt.Println("true1")
	}
	if false {
		fmt.Println("false1")
	}

	//簡易文付き
	if a, b := 1, 2; a < b {
		fmt.Println("true2")
	}
	a, b := 3, 5
	if n := a * b; n%3 == 0 {
		fmt.Println("true3")
	}

	//errorハンドリングのやり方がわからない。
	// if _, err := doSomething(); err != nil {
	// 	//doSomethingがエラーの場合の処理
	// 	fmt.Println("error")
	// }

	//for文
	i := 0
	for {
		fmt.Println(i)
		i++
		if i == 3 {
			break
		}
	}
	j := 0
	for j < 3 {
		fmt.Println(j)
		j++
	}

	// for ［配列のインデックス］, ［配列の要素］ := range ［配列型］ {
	// 	……
	// }
	fruits := [3]string{"apple", "banana", "cherry"}
	for i, s := range fruits {
		//i : 文字列配列のインデックス
		//s : インデックスに対応した文字列の値
		fmt.Printf("fruits[%d]=%s\n", i, s)
	}

	//switch
	switch n := 2; n {
	case 1, 3, 5, 7, 9:
		fmt.Printf("%d is odd\n", n)
	case 2, 4, 6, 8, 10:
		fmt.Printf("%d is even\n", n)
	}

	// func anything(a interface{}) {
	// 	fmt.Println(a)
	// }
	fmt.Println("型アサーション")
	anything(1)
	anything(3.14)
	anything(" 日本語 ")
	anything([...]int{1, 2, 3})

	//型アサーション　変数一つ
	var x3 interface{} = 3
	i3 := x3.(int)
	//f := x3.(float64) //エラー発生する。panic: interface conversion: interface {} is int, not float64
	fmt.Println(i3)

	//型アサーション書き方　その２　変数2つ
	//2つの変数へ代入するような形式で記述します。
	//型アサーションが成功した場合は、1番目の変数にはその値が、2番目の変数にはtrueが代入されます。
	var x4 interface{} = 3.14
	i4, isInt := x4.(int)
	fmt.Println(i4, isInt)
	f, isFloat64 := x4.(float64)
	fmt.Println(f, isFloat64)

	fmt.Println("\n interface{}型と型アサーション　の使い方 x5")
	// var x5 interface{}          // x is nil
	// var x5 interface{} = 3      // x is integer
	var x5 interface{} = "test" // test
	if x5 == nil {
		fmt.Println("x is nil")
	} else if i5, isInt := x5.(int); isInt {
		fmt.Printf("x is integer : %d\n", i5)
	} else if s5, isString := x5.(string); isString {
		fmt.Println(s5)
	} else {
		fmt.Println("unsupported type")
	}

	fmt.Println("\n■■■goto■■■")
	goto L
	fmt.Println("B") //処理されない
L:
	fmt.Println("C")

	fmt.Println("\n■■■ラベル付き文■■■")
	//省略

	fmt.Println("\n■■■defer■■■")
	defer fmt.Println("defer dummy file close")

	fmt.Println("\n■■■deferというよりはerrの実装テスト■■■")
	fmt.Println("ファイルを開くことができれば、deferが効く。開けなければreturnされて効かない")
	// file, err := os.Open("/path/to/file")
	// if err != nil { //errがnilではなければ　＝　ファイルが開けなくerrだったのでerrに入っている
	// 	fmt.Println("これは呼ばれる")
	// 	fmt.Printf("errの内容：「%v」", err)
	// 	// ファイルのオープンに失敗したらreturn
	// 	return
	// 	fmt.Println("呼ばれなければファイルオープン失敗している")
	// }
	// defer file.Close()

}
