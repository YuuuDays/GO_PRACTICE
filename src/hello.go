package main

import (
	"fmt"
	"time"
)

func Hello(name string) string {
	return "Hello," + name
}

func main() {
	// fmt.Println(Hello())
	fmt.Println("time is money", time.Now())
}
