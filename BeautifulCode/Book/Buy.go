package Book

import (
	"errors"
	"math"
	"sort"
	"sync"
)

var soleWithN = map[int]float64{
	1: 0,
	2: 0.05,
	3: 0.1,
	4: 0.2,
	5: 0.25,
}

func BuyBooks(x ...int) float64 {
	return buyBooks(x)
}

func buyBooks(y []int) float64 {
	wg := sync.WaitGroup{}
	ch := make(chan float64)
	sort.Ints(y)
	for number := range soleWithN {
		copyY := make([]int, len(y))
		copy(copyY, y)
		result, err := buyBook(number, copyY)
		if err != nil {
			continue
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			ch <- result + buyBooks(copyY)
		}()
	}
	min := math.Inf(1)
	go func() {
		defer close(ch)
		wg.Wait()
	}()
	for value := range ch {
		min = math.Min(min, value)
	}
	if math.IsInf(min, 1) {
		return 0
	}
	return min
}

func buyBook(n int, y []int) (float64, error) {
	if y[len(y)-n] == 0 {
		return -1, errors.New("invalid")
	}
	for i := 0; i < n; i++ {
		y[len(y)-1-i] = y[len(y)-1-i] - 1
	}
	return float64(n*8) * (1 - soleWithN[n]), nil
}

func getMin(data []float64) float64 {
	if len(data) == 0 {
		return 0
	}
	min := math.MaxFloat64
	for _, value := range data {
		min = math.Min(min, value)
	}
	if min == math.MaxFloat64 {
		return 0
	}
	return min
}
