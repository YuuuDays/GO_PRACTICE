# GO_PRACTICE
勉強メモ

# テストコードの書き方
テストを書く際のルール
xxx_test.goのようなファイル名
関数名は必ず"Test"から始まる
テスト関数の引数は "t *testing.T"
import "testing"の記載が必要

# ベンチマークTest出力結果見方
[コマンド]
go test -bench="ベンチマーク関数の名前"
例:
go test -bench=BenchmarkRepeat_B 
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
(補足1:ちなみに上記の回数は1回呼び出すあたりにかかったナノ秒&とメモリ割り当てになる)
(補足2:1ナノ秒は1秒の10億分の1)
---
PASS: ベンチマークテストが正常にパスしたことを示します。
ok: ベンチマークテストが成功したことを示します。
github.com/YuuHikida/GO_PRACTICE/Iteration: テストされたパッケージの場所。
2.424s: テスト全体が完了するのにかかった総時間です。

# 配列関連
・配列(Array)は固定されたサイズを持つ、故に
    [5]int ,[4]intは異なる型になる
・スライス(Slice)スライスはサイズ可変で配列より柔軟

・make関数
    基本形
    make(T, size , capacity)
    T       ... 作成するデータ構造の型(スライス、マップ、チャンネル)
    size    ... スライスやチャネルの初期サイズ
    capacity... スライスやチャネルの容量（スライスには省略可)

例：
slice := make([]int, 5)          // 長さ5のスライスを作成、全要素はゼロ値で初期化される
sliceWithCap := make([]int, 5, 10) // 長さ5、容量10のスライスを作成

長さを調べるときは"len()",容量を調べるときは"cap()",(sliceの場合)配列への追加は"append"

"長さ"はスライスやmapに実際に格納されているデータ数、"容量"はその配列の最大数

append関数は、第一引数にslice配列を、第二引数以降には要素を入れる事で
    第一引数の配列の後ろへ追加する

    numbers := []int{1, 2, 3}
    numbers = append(numbers, 4, 5, 6)

    fmt.Println(numbers) // 出力: [1 2 3 4 5 6]

    moreNumbers := []int{7, 8, 9}
    numbers = append(numbers, moreNumbers...)

    fmt.Println(numbers) // 出力: [1 2 3 4 5 6 7 8 9]

[スライスのスライス操作]
    例: tail := numbers[1:]
	[1:]はスライスのINDEX番号'1'から最後までを切り出す
	つまり[]int{1,2,3,4}の時、tailには{2,3,4}が入る

