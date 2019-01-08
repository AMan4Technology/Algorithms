package Shadow

func GetShadowNum(leftPoints, rightPoints Points) (int, error) {
	intersectionNum, err := getIntersectionNum(leftPoints, rightPoints)
	if err != nil {
		return 0, err
	}
	return 1 + len(leftPoints) + intersectionNum, nil
}
