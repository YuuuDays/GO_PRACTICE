package main

import "fmt"

// ↓これは構造体を使った変数の宣言,インスタンス

// var coordinate struct {
// 	lat float64
// 	lng float64
// }

//構造体自体を定義してるらしい
// 型名は大文字で始めるとよい、大文字で始めるとエクスポートされる
type coordinate struct {
	lat float64
	lng float64
}

func main() {
	// coordinate.lat = 3.141491919
	// coordinate.lng = 6.6464646464
	// ↓これが宣言としては一般的
	var s coordinate
	s.lat = 3.14
	s.lng = 6.666
	fmt.Println(s.lat, s.lng)
	fmt.Println(s)
}
