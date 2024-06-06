package main

import (
	"fmt"
	"os"
)

// エラーハンドリング
func main() {
	fmt.Println("\n■■■deferというよりはerrの実装テスト■■■")
	fmt.Println("ファイルを開くことができれば、deferが効く。開けなければreturnされて効かない")
	file, err := os.Open("/path/to/file")
	if err != nil { //errがnilではなければ　＝　ファイルが開けなくerrだったのでerrに入っている
		fmt.Println("これは呼ばれる")
		fmt.Printf("errの内容：「%v」", err)
		// ファイルのオープンに失敗したらreturn
		return
		fmt.Println("呼ばれなければファイルオープン失敗している")
	}
	defer file.Close()

}
