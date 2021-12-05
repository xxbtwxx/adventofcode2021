package day5

import (
	"fmt"
)

func task(includeDiagonals bool) {
	lines := lineReader()

	floorMap := floorMap{}
	for line := range lines {
		line.computePoints(includeDiagonals)
		for _, linePoints := range line.points {
			floorMap[linePoints]++
		}
	}

	fmt.Println(floorMap.dangerousPointsCount())
}
