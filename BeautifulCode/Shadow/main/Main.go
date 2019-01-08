package main

import (
	"Algorithms/BeautifulCode/shadowfulcode/shadow"
	"fmt"
)

func main() {
	left := shadow.Points{
		shadow.NewPoint(1, 0, 1),
		shadow.NewPoint(2, 0, 2),
	}
	right := shadow.Points{
		shadow.NewPoint(2, 1, 1),
		shadow.NewPoint(1, 1, 2),
	}
	fmt.Println(shadow.GetShadowNum(left, right))
}
