package main

import (
	"fmt"
	"time"
)

const englishHelloPrefix = "Hello,"

func Hello(name string) string {
	if name == "" {
		return englishHelloPrefix + " World"

	}
	return englishHelloPrefix + name
}

func main() {
	// fmt.Println(Hello())
	fmt.Println("time is money", time.Now())
}
