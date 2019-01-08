package FlapJack

import "fmt"

func Sort(data []int) {
	for i := len(data) - 1; i >= 0; i-- {
		var maxIndex int
		sorted := true
		for j := 0; j <= i; j++ {
			if data[j] > data[maxIndex] {
				maxIndex = j
			} else if data[j] < data[maxIndex] {
				sorted = false
			}
		}
		if sorted {
			return
		}
		if maxIndex == i {
			continue
		}
		swap(data, maxIndex)
		swap(data, i)
	}
}

func swap(data []int, maxIndex int) {
	fmt.Printf("需翻转前%d张饼\n", maxIndex+1)
	for i := 0; i < (maxIndex+1)/2; i++ {
		data[i], data[maxIndex-i] = data[maxIndex-i], data[i]
	}
	fmt.Println(data)
}
