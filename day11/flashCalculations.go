package day11

func computeFlashes(i, j int, flashMap map[int]map[int]struct{}, energyLevels *[][]int) int {
	if outOfBounds(i, j, len(*energyLevels)) {
		return 0
	}

	if flashMap[i] == nil {
		flashMap[i] = map[int]struct{}{}
	}

	if _, flashed := flashMap[i][j]; (*energyLevels)[i][j] == 0 && flashed {
		return 0
	}

	flashes := 0
	if (*energyLevels)[i][j] == 9 {
		(*energyLevels)[i][j] = 0
		flashes++
		flashMap[i][j] = struct{}{}

		flashes += computeFlashes(i+1, j, flashMap, energyLevels)
		flashes += computeFlashes(i-1, j, flashMap, energyLevels)
		flashes += computeFlashes(i, j+1, flashMap, energyLevels)
		flashes += computeFlashes(i, j-1, flashMap, energyLevels)
		flashes += computeFlashes(i+1, j+1, flashMap, energyLevels)
		flashes += computeFlashes(i+1, j-1, flashMap, energyLevels)
		flashes += computeFlashes(i-1, j+1, flashMap, energyLevels)
		flashes += computeFlashes(i-1, j-1, flashMap, energyLevels)
	} else {
		(*energyLevels)[i][j]++
	}

	return flashes
}

func outOfBounds(i, j, bound int) bool {
	if i < 0 || j < 0 || i > bound-1 || j > bound-1 {
		return true
	}

	return false
}
