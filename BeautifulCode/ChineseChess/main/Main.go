package main

import (
	"Algorithms/BeautifulCode/chinesechesse/chinesechess"
	"flag"
	"fmt"
)

var rowN = flag.Int("row", 3, "the number of rows")
var columnN = flag.Int("column", 3, "the number of columns")

func main() {
	flag.Parse()
	positions := chinesechess.GetValidPosition(*rowN, *columnN)
	for index, aPositions := range positions[0] {
		fmt.Println(aPositions, positions[1][index])
	}
}
