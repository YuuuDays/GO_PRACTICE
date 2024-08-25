package main

import "fmt"

const englishHelloPrefix = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return englishHelloPrefix + name
}

func main() {

	fmt.Println(Hello("word"))
}

/*
var mes string = "Hello"
var mes = "Hello"
mee := "Hello"
or
[var mes string
mes = "Hello"]
*/
