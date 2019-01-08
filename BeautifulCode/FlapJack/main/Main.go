package main

import (
	"Algorithms/BeautifulCode/flapjacklcode/flapjack"
	"fmt"
)

func main() {
	data := []int{5, 4, 3, 2, 1, 1, 1, 6}
	fmt.Println(data)
	flapjack.Sort(data)
	fmt.Println(data)
}
