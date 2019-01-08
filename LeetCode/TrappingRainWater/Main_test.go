package TrappingRainWater

import (
	"fmt"
	"testing"
)

func TestTrap(t *testing.T) {
	fmt.Println(trap([]int{4, 3, 3, 9, 3, 0, 9, 2, 8, 3}))
}

func BenchmarkTrap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		trap([]int{5, 2, 1, 2, 1, 5})
	}
}
