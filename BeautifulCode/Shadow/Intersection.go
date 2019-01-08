package Shadow

import (
	"errors"
	"sort"
)

func getIntersectionNum(leftPoints, rightPoints Points) (int, error) {
	if len(leftPoints) != len(rightPoints) {
		return 0, errors.New("左右交点个数不一致")
	}
	pointsWithLineID := make(map[int]Points)
	parsePointsByLineID(pointsWithLineID, leftPoints)
	sort.Sort(rightPoints)
	ys := make([]float64, len(rightPoints))
	for index, point := range rightPoints {
		ys[index] = pointsWithLineID[point.lineID][0].Y
	}
	var num int
	for index, y := range ys {
		for i := index + 1; i < len(ys); i++ {
			if y > ys[i] {
				num++
			}
		}
	}
	return num, nil
}

func parsePointsByLineID(pointsWithLineID map[int]Points, points Points) {
	for _, point := range points {
		pointsWithLineID[point.lineID] = append(pointsWithLineID[point.lineID], point)
	}
}
