package main

import (
	"fmt"

	"github.com/YuuHikida/GO_PRACTICE/src/pra_interface"
)

const englishHelloPrefix = "Hello,"

func Hello(name string) string {
	if name == "" {
		name = " World"

	}
	return englishHelloPrefix + name
}

func main() {
	blog := pra_interface.Blog{
		Title:   "taitorudesu!!",
		Content: "naiyoudessssssssssssss",
	}

	blog2 := pra_interface.Blog2{
		Title:     "タイトルです。",
		Paragraph: []string{"いち", "にい", "さん・・・？"},
	}

	fmt.Println(blog.GetFullArticle())
	fmt.Println("--------------")
	fmt.Println(blog2.GetFullArticle())
}
