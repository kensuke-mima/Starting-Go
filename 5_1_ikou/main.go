package main

import (
	"fmt"
	"math"
)

/* int型をとりint型を返すコールバック型 */
type Callback func(i int) int //func(i int) int　にCallbackというエイリアスを定義

func Sum(ints []int, callback Callback) int {
	var sum int
	for _, i := range ints {
		sum += i
	}
	fmt.Println("sum:", sum)
	//ここまでで1+2+・・・+5で、15
	return callback(sum) //callbackにsumを入れる。callbackは2倍する関数なので
}

type Point2 struct{ X, Y int }

/* *Point型のメソッドRender */
//メソッド定義では、関数とは異なりfuncとメソッド名の間に「レシーバー（Receiver）」の型とその変数名が必要になります。この場合は、*Point型の変数pがレシーバーを表しています。
func (p *Point2) Render() {
	fmt.Printf("<%d,%d>\n", p.X, p.Y)
}

/* 2点間の距離を求めるメソッド */
//Point2型のDistanceメソッドを呼ぶとき、引数にPoint2型が必要
func (p *Point2) Distance(dp *Point2) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y
	return math.Sqrt(float64(x*x + y*y))
}

func main() {
	//
	fmt.Println("\n■■■ポインタ■■■")
	//省略

	//
	fmt.Println("\n■■■構造体■■■")
	fmt.Println("\n■■■エイリアス■■■")
	/* typeによるさまざまなエイリアス */
	type (
		IntPair     [2]int
		Strings     []string
		AreaMap     map[string][2]float64
		IntsChannel chan []int
	)
	pair := IntPair{1, 2}
	strs := Strings{"Apple", "Banana", "Cherry"}
	amap := AreaMap{"Tokyo": {35.689488, 139.691706}}
	ich := make(IntsChannel)
	fmt.Println(pair, strs, amap, ich)

	//
	fmt.Println("\n■■■/* int型をとりint型を返すコールバック型 */■■■")
	n := Sum(
		[]int{1, 2, 3, 4, 5},
		func(i int) int {
			return i * 2
		},
	)
	//15 * 2　で30になる。
	fmt.Println(n) /* n == 30 */

	//
	fmt.Println("\n■■■構造体の定義■■■")
	type Point struct {
		X int
		Y int
	}
	var pt Point
	fmt.Println(pt.X, pt.Y)
	pt.X = 10
	pt.Y = 8
	fmt.Println(pt.X, pt.Y)

	//
	fmt.Println("\n■■■フィールド定義の詳細■■■")
	type T struct {
		int
		float64
		string
	}
	t := T{1, 3.14, "文字列"}
	fmt.Println(t.int, t.float64, t.string)

	//
	fmt.Println("\n■■■メソッド■■■")
	// type Point2 struct{ X, Y int }
	// /* *Point2型のメソッドRender */
	// func (p *Point2) Render() {
	//    fmt.Printf("<%d,%d>\n", p.X, p.Y)
	// }
	p := &Point2{X: 5, Y: 12}
	/* メソッドRenderの呼び出し */
	p.Render() //<5,12>

	//
	fmt.Println("\n■■■メソッド2■■■")
	/* 2点間の距離を求めるメソッド */
	//Point2型のDistanceメソッドを呼ぶとき、引数にPoint2型が必要
	// func (p *Point2) Distance(dp *Point2) float64 {
	// 	x, y := p.X-dp.X, p.Y-dp.Y
	// 	return math.Sqrt(float64(x*x + y*y))
	// }
	p2 := &Point2{X: 0, Y: 0}
	p2A := p2.Distance(&Point2{X: 1, Y: 1}) // == 1.4142135623730951
	fmt.Println(p2A)

	//
	fmt.Println("\n■■■型のコンストラクタ■■■")
	//NewUser パッケージ外公開
	//newuser パッケージ内

	//
	fmt.Println("\n■■■レシーバーとポインタ型■■■")

	//
	fmt.Println("\n■■■フィールととメソッドの可視性■■■")
	//公開　　大文字
	//非公開　小文字

	//
	fmt.Println("\n■■■スライスと構造体■■■")

	//
	fmt.Println("\n■■■マップと構造体■■■")

}
