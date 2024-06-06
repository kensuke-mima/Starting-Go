package main

import (
	"fmt"
	"os"
)

func main() {
	//
	fmt.Println("\n■■■コマンドライン引数■■■")
	/* os.Argsの要素数を表示 */
	fmt.Printf("length=%d\n", len(os.Args))
	/* os.Argsの内容をすべて出力 */
	for _, v := range os.Args {
		fmt.Println(v)
	}

}
