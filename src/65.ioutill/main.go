package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	//入力を読み込む
	f, _ := os.Open("foo.txt")
	bs, _ := io.ReadAll(f)
	fmt.Println(string(bs))

	//ファイルに書き込み
	if err := os.WriteFile("bar.txt", bs, 0666); err != nil {
		log.Fatalln(err)
	}
}