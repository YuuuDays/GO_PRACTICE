package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func SumAll(x []int, y []int) []int {
	return []int{}
}
