package SlidingWindowMaximum

import (
	"fmt"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	fmt.Println(maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
}

func BenchmarkMaxSlidingWindow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxSlidingWindow([]int{9, 8, 7, 6, 5, 4, 3, 2}, 3)
	}
}
