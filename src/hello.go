package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

// func Greet(writer io.Writer, name string) {
// 	fmt.Fprintf(writer, "Hello, %s", name)
// }

// func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
// 	Greet(w, "world")
// }

const finalWord = "Go!"
const countdownStart = 3

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
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)

}

func main() {
	sleeper := &SpySleeper{}
	Countdown(os.Stdout, sleeper)
}
