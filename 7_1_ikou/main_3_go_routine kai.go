package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("1st Goroutine")
			time.Sleep(time.Millisecond * 100)
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("2nd Goroutine")
			time.Sleep(time.Millisecond * 100)
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			fmt.Println("3rd Goroutine")
			time.Sleep(time.Millisecond * 100)
		}
	}()
	for {
	}
}
