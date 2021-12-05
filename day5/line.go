package day5

type line struct {
	start  point
	end    point
	points []point
}

func (l *line) computePoints(includeDiagonals bool) {
	if l.isHorizontal() {
		start, end := 0, 0
		if l.start.y < l.end.y {
			start = l.start.y
			end = l.end.y
		} else {
			start = l.end.y
			end = l.start.y
		}

		for i := start; i <= end; i++ {
			l.points = append(l.points, point{x: l.start.x, y: i})
		}
	} else if l.isVertical() {
		start, end := 0, 0
		if l.start.x < l.end.x {
			start = l.start.x
			end = l.end.x
		} else {
			start = l.end.x
			end = l.start.x
		}

		for i := start; i <= end; i++ {
			l.points = append(l.points, point{x: i, y: l.start.y})
		}
	} else if includeDiagonals {
		pointsCount := l.start.x - l.end.x
		if pointsCount < 0 {
			pointsCount *= -1
		}

		for i := 0; i <= pointsCount; i++ {
			pointX, pointY := l.start.x, l.start.y
			if l.start.x > l.end.x {
				pointX -= i
			} else {
				pointX += i
			}

			if l.start.y > l.end.y {
				pointY -= i
			} else {
				pointY += i
			}

			l.points = append(l.points, point{x: pointX, y: pointY})
		}
	}
}

func (l line) isHorizontal() bool {
	return l.start.x == l.end.x
}

func (l line) isVertical() bool {
	return l.start.y == l.end.y
}

type point struct {
	x int
	y int
}
