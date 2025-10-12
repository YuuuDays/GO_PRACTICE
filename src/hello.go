package main

import arrays "github.com/YuuHikida/GO_PRACTICE/src/ArraysAndslices"

const englishHelloPrefix = "Hello,"

func Hello(name string) string {
	if name == "" {
		name = " World"

	}
	return englishHelloPrefix + name
}

func main() {
	// fmt.Println(Hello())
	//fmt.Println("time is money", time.Now())
	arrays.SumAllTails([]int{1, 2}, []int{2, 3})
}
