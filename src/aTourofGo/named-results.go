package main

import (
	"fmt"
	"net/http"
	"time"
)

/*2000年後の僕へ
go run xxx.go
上記だと単体のgoファイルしかビルドされないぞ
go run .
とか
go run *.go
でまとめてコンパイルしよう
*/
// naked return とかいうものですぅ
func CallTwoArgs(sum int) (x, y int) {
	x = sum * 2
	y = sum / 2
	return
}

var (
	ToBe bool = false
	moji rune = 'い'
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context() // リクエストの生存監視

		select {
		case <-time.After(3 * time.Second):
			fmt.Println(w, "処理完了")
		case <-ctx.Done(): //クライアントとが途中離脱するとここにはいる
			fmt.Println("中断理由:", ctx.Err())
			return
		}
	})
	fmt.Println("server start:8080")
	http.ListenAndServe(":8080", nil)
}
