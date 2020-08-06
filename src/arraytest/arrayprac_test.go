package arraytest

import (
	"testing"
)

// %v 默认输出格式
func TestSum(t *testing.T) {
	numbers := [5]int{1, 2, 3, 4, 5}
	got := Sum(numbers)
	want := 20

	if want != got {
		t.Errorf("got %d want %d given, %v", got, want, numbers)
	}
}

func TestSizeOf(t *testing.T) {
	var slice = []int{}
	SizeOf(slice)
}
