// 同じ名前空間ならば以下の記述で呼び出せる
// package mypackage

// 	func 関数名() {
// 		Myfunction()
// 	}

// packageの名前は同じ階層に存在するディレクトリのgoファイルと同じ名前にしなければならない

package main

import (
	"fmt"

	"github.com/YuuHikida/GO_PRACTICE/integers"
)

// 関数外では短縮変数宣言は使用できない
// x, y := 1, 2
var x, y int = 1, 2

func main() {
	result := integers.Add(x, y)
	fmt.Println(result)
}
