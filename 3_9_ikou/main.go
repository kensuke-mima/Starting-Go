package main

import "fmt"

func main() {

	//3.9 interface{} と nil
	var x interface{}
	fmt.Printf("%#v\n", x) //<nil>
	x = 1
	fmt.Printf("%#v\n", x) //1
	x = '山'
	fmt.Printf("%#v\n", x) //23665

	//3.11 関数
	fmt.Println(plus(1, 2))
	fmt.Println(div(5, 2)) //5÷2 2余り1

	//戻り値の破棄
	q, _ := div(5, 2) //2
	println(q)
	_, r := div(5, 2) //1
	println(r)
	//_, _ := div(5, 2) //コンパイルエラー

	//無名関数
	f := func(x, y int) int { return x + y }
	f(2, 3) //5
	fmt.Printf("\n%T\n", f)
	fmt.Printf("%v\n", f(2, 3))
	fmt.Printf("%#v\n", f)
	fmt.Printf("%#v\n", func(x, y int) int { return x + y }(2, 3))
	fmt.Printf("%T\n\n", func(x, y int) int { return x + y })

	//名前付き関数と無名関数
	var a = plusAlias(10, 5)
	fmt.Printf("plusAlias %v\n\n", a)

	//3.11 関数を返す関数
	f2 := returnFunc()
	f2() //⇒これがプリントする関数I'm a function
	//直接呼出し
	returnFunc()()

	//3.11 関数を引数にとる関数
	callFunction(func() {
		fmt.Println("I'm a function HIKISU") //このprint文の関数を引数として渡している
	})

}

//3.11 関数
func plus(x, y int) int {
	return x + y
}

//3.11 名前付き関数と無名関数
var plusAlias = plus

func div(a, b int) (int, int) {
	q := a / b
	r := a % b
	return q, r
}

//3.11 関数を返す関数
func returnFunc() func() {
	return func() {
		fmt.Println("I'm a function")
	}
}

//3.11 関数を引数にとる関数
func callFunction(f func()) {
	f() //引数でもらった func f　を実行している。
}
