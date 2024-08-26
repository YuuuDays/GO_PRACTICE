package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	// numbers := [...]int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	want := 15

	if got != want {
		t.Errorf("got :(%q), want:(%q)", got, want)
	}

	t.Run("aaaa", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got:(%q),want:(%q)", got, want)
		}
	})
}

func TestSumAll(t *testing.T) {

	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	// スライスの内容は比較演算で比較できない
	// if got != want {
	//     t.Errorf("got %v want %v", got, want)
	// }

	// reflect.DeepEqualは同じかどうか判定してくれる
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got:(%v),want:(%v)", got, want)
	}
}

func BenchmarkRepeat(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func BenchmarkRepeat_B(b *testing.B) {
	numbers := SumAll([]int{1, 2}, []int{0, 9})
	for i := 0; i < b.N; i++ {
		SumAll(numbers)
	}
}
