package pra_interface

import "fmt"

// fildの変数の先頭を大文字にすることで外部からアクセス可能
type Blog struct {
	Title   string
	Content string
}

func (b Blog) GetFullArticle() string {
	return b.Title + "\n" + "----------" + "\n" + b.Content
}

type Blog2 struct {
	Title     string
	Paragraph []string
}

func (b Blog2) GetFullArticle() string {
	article := b.Title + "\n" + "------------" + "\n"

	for _, paragraph := range b.Paragraph {
		article += paragraph + "\n\n" // 段落の間は2重に改行する
	}

	return article
}

//同じですやん...
// func DisplayBlog(b Blog) {
// 	fmt.Println(b.GetFullArticle())
// }

// func DisplayBlog2(b Blog2) {
// 	fmt.Println(b.GetFullArticle())
// }

//というわけでinterfaceを使うべ
// 説明:GetFullArticle()という、引数を取らずstring型を出力するメソッドがある型は、何でも入っていいよ
type BlogInterface interface {
	GetFullArticle() string
}

//2つ上のprint関数を簡潔にすべ
func DisplayBlog(b BlogInterface) {
	fmt.Println(b.GetFullArticle())
}

/*注意点★
以下はエラー。インターフェースはインターフェース型の関数しか保証せず、フィールド変数は使えない。
インターフェースには違うフィールド変数が入る可能性があるため=今回だとBlogとBlog2の構造体のフィールドが違うから
*/
func DameDisplayBlog(b BlogInterface) {
	//   if b.title == "" { // b.titleはアウト
	//     return
	//   }

	fmt.Println(b.GetFullArticle())
}
