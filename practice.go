package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"time"

	"golang.org/x/sync/errgroup"
)

// 構造体(スライスに関して)
func slice_confirmation() {
	type User struct {
		Name string
		Age  int
	}

	//ポインタ宣言
	u := []*User{
		&User{Name: "torou", Age: 10},
	}

	// ポインタぢゃない
	u_2 := []User{
		User{Name: "spygear", Age: 192},
	}

	fmt.Println("変更前(ポインタ)", u[0].Name, ":", u[0].Age) //変更前(ポインタ) torou : 10
	u[0].Name = "TANAKA"
	fmt.Println("変更後(ポインタ)", u[0].Name, ":", u[0].Age) //変更後(ポインタ) TANAKA : 10

	fmt.Println("変更前(ポインタ)", u_2[0].Name, ":", u_2[0].Age) //変更前(ポインタ) torou : 10
	u_2[0].Name = "EIDA"
	fmt.Println("変更後(ポインタ)", u_2[0].Name, ":", u_2[0].Age) //変更後(ポインタ) TANAKA : 10

}

/*-------------------------------------------------------------------------------*/
/*Go言語ではインターフェースや構造体、そしてメソッドの定義はトップレベルでしかできない*/
/*-------------------------------------------------------------------------------*/
type Storage interface {
	Save(data string)
}

// FileStorage 構造体
type FileStorage struct{}

// FileStorageのSaveメソッド（インターフェースの実装）
func (fs FileStorage) Save(data string) {
	fmt.Println("Saving to file:", data)
}

/*↑ 関数の中で関数を定義してはならないので
関数内で構造体メソッド定義を行うとエラー*/

func type_confirmation() {
	var number int = 42
	fmt.Println("number->", reflect.TypeOf(number))

	//低結合性と柔軟性
	var storage Storage = FileStorage{}
	storage.Save("OOENTKYK!")

	fmt.Println("storage->", reflect.TypeOf(storage)) // main.FileStorage

}

/*ただのhttp.ListenAndServer*/
func httpPractice_1() {
	err := http.ListenAndServe(
		":18080",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello,%s!", r.URL.Path[1:])
		}),
	)

	if err := run(context.Background()); err != nil {
		fmt.Printf("falied to terminate server :%v", err)
	}

	fmt.Println(err)
}

func run(ctx context.Context) error {
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "Hello,%s", r.URL.Path[1:])
		}),
	}
	/*
		WithContextを使うことで、コンテキストがゴルーチン間で共有され、
		ゴルーチン全体が同じキャンセルやタイムアウトの制御を受けます。
	*/
	eg, ctx := errgroup.WithContext(ctx)

	// 別ゴルーチンでHTTPサーバーを起動する
	eg.Go(func() error {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("failed to close :%+v", err)
			return err
		}
		return nil
	})

	// チャンネルからの通知を待機する
	<-ctx.Done()
	if err := s.Shutdown(context.Background()); err != nil {
		log.Printf("failed to shutdown:%+v", err)
	}
	// GOメソッドで起動した別後ルーチンの終了を待つ
	return eg.Wait()
}

func errgroup_practice() {
	g := new(errgroup.Group)
	var urls = []string{
		"http://www.golang.org/",
		"http://www.google.com/",
		"http://www.エラー出るよ.com/",
	}

	for _, url := range urls {
		url = url
		/* 通常ゴルーチンは "go 関数名()"
		しかしerrgroupは"g.Go( 関数名() )"で起動
		*/
		g.Go(func() error {
			resp, err := http.Get(url)
			if err == nil {
				fmt.Println("取得完了")
				resp.Body.Close()
			}
			return err
		})
		// g.Wait()は全てのゴルーチンが終了するまで待つ
		if err := g.Wait(); err != nil {
			fmt.Printf("err出たよ:%v", err)
			return
		}
		fmt.Println("--------------------")
	}
}

/*
	contexの使い方練習

以下解説
1.非同期ゴルーチン開始 → for文が回り続ける（default で「ゴルーチンが実行中です」を表示）。
2.メインで time.Sleep(5 * time.Second) → 5秒間ゴルーチンが動き続ける。
3.cancel() 呼び出し → ctx.Done() のチャネルが閉じる。
4.case <-ctx.Done() が実行され、ゴルーチンが終了。
*/
func contex_practice_1() {
	/* "ctx"はctx.Done() というチャネルでキャンセル通知を待つことができる */
	/* "cancel" キャンセル操作を行う関数。ctx.Done()が「完了した」という通知を送る*/
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() //記載した関数が終わり次第cancelを行ってくれる(明示的に書く必要はないが習慣として書く)

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("ゴルーチンがキャンセル")
				return
			default:
				fmt.Println("ゴルーチンが実行中です")
				time.Sleep(time.Second)
			}

		}
	}()

	fmt.Println("cancelのtypeは～～？->", reflect.TypeOf(cancel)) // context.CancelFunc

	// メインプログラムを5秒間実行し、ゴルーチンに実行時間を与える
	time.Sleep(5 * time.Second)

	// ゴルーチンをキャンセル
	cancel()

	// キャンセルメッセージを確認するためにさらに1秒待つ
	time.Sleep(2 * time.Second)

	fmt.Println(ctx)

}

// func main() {
// 	//slice_confirmation()
// 	//type_confirmation()
// 	//httpPractice_1()

// 	/* ゴルーチン(errgroupテスト)*/
// 	//errgroup_practice()

// 	contex_practice_1()

// }
