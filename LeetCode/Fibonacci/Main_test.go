package Fibonacci

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	fmt.Println(fibonacci(100))
}

func TestFibonacciDP(t *testing.T) {
	fmt.Println(fibonacciDP(100))
}
