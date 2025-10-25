package main

import "fmt"

// naked return とかいうものですぅ
func CallTwoArgs(sum int) (x, y int) {
	x = sum * 2
	y = sum / 2
	return
}

func main() {
	fmt.Println(CallTwoArgs(10))
}
