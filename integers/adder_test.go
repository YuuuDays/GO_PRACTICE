package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected:(%d), sum:(%d)", expected, sum)
	}
}

// サンプル関数
// Goのテストツール'ExampleXxx()'関数内の'// Output:'という特別なコメントを自動的に読み取る
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output:6
}
