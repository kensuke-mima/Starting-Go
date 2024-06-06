package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//
	fmt.Println("\n■■■log.Fatal■■■")
	_, err := os.Open("foo")
	if err != nil {
		fmt.Println("■■■err出力■■■")
		log.Fatal(err)
		fmt.Println("■■■err出力ここまで■■■")

	}

}
