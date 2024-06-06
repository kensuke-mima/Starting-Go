package main

import (
	"fmt"
	"time"
)

func receiver(ch <-chan int) {
	for {
		i := <-ch
		fmt.Println(i)
	}
}

func receive(name string, ch <-chan int) {
	for {
		i, ok := <-ch
		if ok == false {
			//受信できなくなったら終了（chがクローズしたら終了）
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "is done")
}

func main() {
	fmt.Println("\n■■■チャネル■■■")
	//Goのチャネルは単なるキューではなく、ゴルーチン間のデータの共有のために用意されています。
	ch := make(chan int)
	go receiver(ch)

	i := 0
	for i < 10 {
		ch <- i
		i++
	}

	//close
	fmt.Println("\n■■■close■■■")
	ch1 := make(chan int)
	close(ch1)
	i1, ok := <-ch1 //closeせずにprintすると「fatal error: all goroutines are asleep - deadlock!」
	fmt.Println(i1, ok)

	//
	fmt.Println("\n■■■ゴルーチンとcloseの実装■■■")
	ch2 := make(chan int, 20)

	//3つのゴルーチンを実行
	go receive("1st goroutine", ch2)
	go receive("2nd goroutine", ch2)
	go receive("3rd goroutine", ch2)

	//0から10の数字をチャネルに送信すると、1,2,3いずれかのゴルーチンが受信して表示する
	i2 := 0
	for i2 < 10 {
		ch2 <- i2
		i2++
	}
	close(ch2)

	//ゴルーチンの完了を3秒待つ
	time.Sleep(3 * time.Second)

	//
	fmt.Println("\n■■■select■■■")
	//ch1からの受信が成功した場合の処理
	//ch2からの受信が成功した場合の処理
	//という具合にケースごとの処理を行うのがselect。case文のようなもの。

}
