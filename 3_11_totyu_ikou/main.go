package main

import "fmt"

//3.11 クロージャ関数
func lator() func(string) string {
	//一つ前に与えられた文字列を保存するための変数(store)
	var store string
	//引数に文字列(next)を取り、文字列sを返す関数、を返す
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

func integers() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	f := lator() //この関数の変数化により、store変数がローカル変数からクロージャに属する変数となる

	fmt.Println(f("Golang"))   //⇒"""
	fmt.Println(f("is"))       //"Golang"
	fmt.Println(f("awesome!")) //"is"

	//なぜ、fでインスタンス化するのか
	//これだと、latorがインスタンス化されていないのでstoreに値が残っていないので、
	//何回やっても、初回の"""が返却される
	fmt.Println(lator()("test"))  //⇒"""
	fmt.Println(lator()("test2")) //⇒"""
	fmt.Println(lator()("test3")) //⇒"""

	//ジェネレータの実装
	ints := integers()
	fmt.Println(ints())      //⇒"1"
	fmt.Println(ints())      //⇒"2"
	fmt.Println(ints())      //⇒"3"
	otherInts := integers()  //別関数にすると、変数の値はもちろん異なる。
	fmt.Println(otherInts()) //⇒"1"

}
