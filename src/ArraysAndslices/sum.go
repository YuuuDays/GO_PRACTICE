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
	lengthOfNumbers := len(args)            //->2
	fmt.Println(lengthOfNumbers)
	var sums []int

	for _, numbers := range args {
		sums = append(sums, Sum(numbers))
	}

	return sums
}

func SumAllTails(args ...[]int) []int {
	var sums []int
	for _, numbers := range args {
		// [1:] ←これはindex1から末尾まで
		tail := numbers[1:]
		// appendは単純に第二引数の値を追加
		sums = append(sums, Sum(tail))
	}

	return sums
}
