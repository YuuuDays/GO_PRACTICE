package main

import "testing"

/*
テストを書く際のルール
xxx_test.goのようなファイル名
関数名は必ず"Test"から始まる
テスト関数の引数は "t *testing.T"
import "testing"の記載が必要
*/

func TestHello(t *testing.T) {
	assertMes := func(t testing.TB, got, want string) {
		t.Helper() //ヘルパー関数としてマークするメソッド
		if got != want {
			t.Errorf("got: (%q) want: (%q)", got, want) //[%q]は文字列を""で囲み、安全にエスケープする
		}
	}

	//t.Runはどのテストコードを実行しているか表示してくれる
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertMes(t, got, want)

		// got := Hello("Chris")
		// want := "Hello, Chris"

		// if got != want {
		// 	t.Errorf("got %q want %q", got, want)
		// }
	})
	t.Run("これはどうなの", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertMes(t, got, want)

	})
}
