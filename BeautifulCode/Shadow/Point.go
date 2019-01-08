package Shadow

func NewPoint(lineID int, x, y float64) point {
	return point{lineID, x, y}
}

type Points []point

func (p Points) Len() int {
	return len(p)
}

func (p Points) Less(i, j int) bool {
	return p[i].Y < p[j].Y
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

type point struct {
	lineID int
	X, Y   float64
}
