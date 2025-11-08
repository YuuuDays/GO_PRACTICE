package main

import (
	"bytes"
	"testing"
)

/*
buffer には
Countdown 関数が書き込んだ「バイト列（文字列）」がメモリ上に溜まっている
*/
func TestCountdown(t *testing.T) {
	// goの標準パッケージbytesにある「バイト配列を内蔵した可変の入れ物」
	//&bytes.Buffer{} としているのは「そのバッファのポインタ」を作ってる
	buffer := &bytes.Buffer{}

	Countdown(buffer)

	// buffer.String() を呼ぶと、その溜まったバイト列が Go の string として取り出せる
	got := buffer.String()
	want := "3"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
