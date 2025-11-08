package main

import (
	"fmt"
	"io"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

/*
io.Writer は共通ルール(インターフェイス)
| 使えるもの                 | 役割           |
| --------------------- | ------------ |
| `*bytes.Buffer`       | メモリ上に文字をためる♡ |
| `os.Stdout`           | ターミナルに表示♡    |
| `http.ResponseWriter` | Webレスポンスに出力♡ |
| `strings.Builder`     | 文字結合して溜める♡   |
| ファイル (`*os.File`)     | ファイルに書く♡     |
*/
func Countdown(out io.Writer) {
	fmt.Fprint(out, 3)
}

func main() {
	// Countdown()
}
