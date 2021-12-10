package day05

type floorMap map[point]int

func (fm floorMap) dangerousPointsCount() int {
	points := 0

	for _, dangerousLevel := range fm {
		if dangerousLevel >= 2 {
			points++
		}
	}

	return points
}
