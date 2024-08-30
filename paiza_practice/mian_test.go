package main

import (
	"testing"
)

func TestCaluculateTotal(t *testing.T) {
	num := 10
	discount := 1735
	inputs := []string{"906", "1498", "240", "934", "872", "688", "1566", "756", "618", "314"}

	expected := 8392 // ここに期待する合計値を入力
	result := CaluculateTotal(num, discount, inputs)

	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}

}

// func BenchmarkRepeat(b *testing.B) {
// 	numbers := []int{1, 2, 3, 4, 5}
// 	for i := 0; i < b.N; i++ {
// 		Sum(numbers)
// 	}
// }
