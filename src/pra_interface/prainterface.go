package pra_interface

// fildの変数の先頭を大文字にすることで外部からアクセス可能
type Blog struct {
	Title   string
	Content string
}

func (b Blog) GetFullArticle() string {
	return b.Title + "\n" + "----------" + "\n" + b.Content
}
