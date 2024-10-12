package main

import "fmt"

func main() {
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
