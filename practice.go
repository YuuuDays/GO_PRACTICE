package main

import (
	"fmt"
	"reflect"
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

func main() {
	//slice_confirmation()
	type_confirmation()

}
