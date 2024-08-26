package main

import (
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

func BenchmarkRepeat(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}
	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}
