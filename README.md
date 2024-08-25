# GO_PRACTICE
勉強メモ

# テストコードの書き方
テストを書く際のルール
xxx_test.goのようなファイル名
関数名は必ず"Test"から始まる
テスト関数の引数は "t *testing.T"
import "testing"の記載が必要

# ベンチマークTest出力結果見方
[内容]
goos: windows
goarch: amd64
pkg: github.com/YuuHikida/GO_PRACTICE/Iteration
cpu: Intel(R) Core(TM) i5-14400F
=== RUN   BenchmarkRepeat
BenchmarkRepeat
BenchmarkRepeat-16
15996181                73.69 ns/op           16 B/op          4 allocs/op
PASS
ok      github.com/YuuHikida/GO_PRACTICE/Iteration      2.424s

[意味]
BenchmarkRepeat   : ベンチマークされた関数の名前
BenchmarkRepeat-16: このベンチマークが16スレッド（または16の並行性）で実行されたことを示す

---
15996181: このベンチマーク関数がテスト期間中に実行された回数です。
73.69 ns/op: 操作ごとに平均73.69ナノ秒がかかりました。
16 B/op: 操作ごとに平均16バイトのメモリが割り当てられました。
4 allocs/op: 操作ごとに平均4回のメモリ割り当て（アロケーション）が発生しました。
(ちなみに上記の回数は1回呼び出すあたりにかかったナノ秒&とメモリ割り当てになる)

---
PASS: ベンチマークテストが正常にパスしたことを示します。
ok: ベンチマークテストが成功したことを示します。
github.com/YuuHikida/GO_PRACTICE/Iteration: テストされたパッケージの場所。
2.424s: テスト全体が完了するのにかかった総時間です。
