package main

import "testing"

/*
テストを書く際のルール
xxx_test.goのようなはファイル名
関数名は必ず"Test"から始まる
テスト関数の引数は "t *testing.T"
import "testing"の記載が必要
*/

func TestMain(t *testing.T) {
	got := Hello("Chris")
	want := "Hello, Chris"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
