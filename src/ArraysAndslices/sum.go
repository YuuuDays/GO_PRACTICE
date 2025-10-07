package arrays

func Sum(arrays [5]int) int {
	sum := 0
	for _, num := range arrays {
		sum += num
	}
	return sum
}
