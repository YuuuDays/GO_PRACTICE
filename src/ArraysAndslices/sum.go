package arrays

import (
	"fmt"
	"reflect"
)

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

/*
go可変長引数関数の作り方
func 関数名(引数名 ...型)
*/
func SumAll(args ...[]int) []int {
	fmt.Println("型:", reflect.TypeOf(args)) // ->[][]int ...スライスの中のスライス

	lengthOfNumbers := len(args)
	fmt.Println(lengthOfNumbers)

	sums := make([]int, lengthOfNumbers)

	for i, numbers := range args {
		sums[i] = Sum(numbers)
	}

	return sums
}
