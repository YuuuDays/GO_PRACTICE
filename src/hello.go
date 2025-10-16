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

	fmt.Println(blog.GetFullArticle())
}
