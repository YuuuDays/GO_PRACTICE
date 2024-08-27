package main

func Sum(numbers []int) int {
	sum := 0

	// "_"は配列の添え字を無視している
	for _, number := range numbers {
		sum += number
	}

	// 添え字を表示するver↓
	// for x, number := range numbers {
	// 	sum += number
	// 	fmt.Printf("(%d)", x)
	// }
	return sum
}

func SumAll(numbersToSum ...[]int) (sums []int) {
	countNum := len(numbersToSum)
	sumSlice := make([]int, countNum)

	for i, num := range numbersToSum {
		sumSlice[i] += Sum(num)
	}
	return sumSlice
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		// [1:]はスライスのINDEX番号'1'から最後までを切り出す
		// つまり[]int{1,2,3,4}の時、tailには{2,3,4}が入る
		//　スライスのスライス操作
		tail := numbers[1:]
		sums = append(sums, Sum(tail))
	}

	return sums
}
