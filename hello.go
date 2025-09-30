package main

import (
	"fmt"
	"time"
)

func Hello() string {
	return "Hello,World!"
}

func main() {
	fmt.Println(Hello())
	fmt.Println("time is money", time.Now())
}
