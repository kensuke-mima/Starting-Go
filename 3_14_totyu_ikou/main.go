package main

import (
	"fmt"
)

func testRecover(src interface{}) {
	defer func() {
		if x := recover(); x != nil {
			switch v := x.(type) {
			case int:
				fmt.Printf("panic復旧: int=%v\n", v)
			case string:
				fmt.Printf("panic復旧: string=%v\n", v)
			default:
				fmt.Printf("panic復旧: unknown\n")
			}
		}
	}() //この()は無名関数の時に付けるんだったけかな
	panic(src)
	return //return来る前にパニックしてるな
}

func main() {
	fmt.Println("\n■■■panicとrecover■■■")
	testRecover(128)
	testRecover("hogehoge")
	testRecover([...]int{1, 2, 3})
	//testRecover() //xをnilにして、if文に入らないようにしたかったが、エラーで実行できない

	fmt.Println("\n■■■go■■■")

}
