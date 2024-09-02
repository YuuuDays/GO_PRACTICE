package main

import "fmt"

// ↓これは構造体を使った変数の宣言,インスタンス
// 無名構造体
// var coordinate struct {
// 	lat float64
// 	lng float64
// }

//構造体自体を定義してるらしい
// 型名は大文字で始めるとよい、大文字で始めるとエクスポートされる
type Coordinate struct {
	lat float64
	lng float64
}

func main() {
	// coordinate.lat = 3.141491919
	// coordinate.lng = 6.6464646464
	// ↓これが宣言としては一般的
	var s Coordinate
	s.lat = 3.14
	s.lng = 6.666
	fmt.Println(s.lat, s.lng)
	fmt.Println(s)
	fmt.Println("-------------------------------")

	sibuya := Coordinate{lat: 1.111111, lng: 2.22222222}
	fmt.Println(sibuya.lat)
	fmt.Println(sibuya)

	fmt.Println("-------------------------------")

	tanaka := Coordinate{4.444444, 5.5555555}
	fmt.Println(tanaka.lat)
	fmt.Println(tanaka)
}
