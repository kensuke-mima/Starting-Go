package main

import (
	"fmt"
	"os"
)

func main() {
	//
	fmt.Println("\n■■■ファイル操作■■■")
	/* ファイル名を指定して新規作成 */
	f, _ := os.Create("foo.txt")
	/* 新規作成したファイルのステータス */
	fi, _ := f.Stat()
	fi.Name()                          // == "foo.txt"
	fi.Size()                          // == 0
	fi.IsDir()                         // == false
	f.Write([]byte("Hello, World!\n")) // ファイルに[]byte型の内容を書き込み
	f.WriteAt([]byte("Golang"), 7)     // オフセットを指定して書き込み
	f.Seek(0, os.SEEK_END)             // ファイルの末尾にオフセットを移動
	f.WriteString("Yeah!")             // 文字列をファイルに書き込み
	/* 作成されるファイルの内容 */
	// Hello, Golang
	// Yeah!

}
