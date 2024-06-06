package main

import (
	"fmt"
	"runtime"
)

func init() {
	fmt.Println("\n■■■init■■■")
	fmt.Println("init()")
}

func main() {
	fmt.Println("\n■■■go■■■")
	go fmt.Println("Yeah!") //これが追加されると、これが1になり
	//time.Sleep(1)                                            //sleepすると1になる。
	fmt.Printf("NumCPU: %d\n", runtime.NumCPU())             //
	fmt.Printf("NumGoroutine: %d\n", runtime.NumGoroutine()) //これが1から2になる
	fmt.Printf("Version: %s\n", runtime.Version())

}
