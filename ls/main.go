package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	currentDir, err := os.Getwd() // カレントディレクトリ取得
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Current Directory", currentDir)

	// ReadDir関数は対象のディレクトリを読み込み、
	// ファイルの名前、情報を取得し、
	// ファイルの名前でソートして返す関数
	files, err := os.ReadDir("../")
	if err != nil {
		log.Fatal(err)
	}

	// 取得したファイル名を表示
	// _はインデックス
	for _, file := range files {
		fmt.Println(file.Name())
	}
}
