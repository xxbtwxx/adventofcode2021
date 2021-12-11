package day11

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Solution for day 11 problem A",
		Run: func(cmd *cobra.Command, args []string) {
			taskA()
		},
	})
}

func taskA() {
	energyLevels := energyLevelsReader()
	flashes := 0

	iterations := 100
	for ; iterations != 0; iterations-- {
		flashed := map[int]map[int]struct{}{}
		for i := 0; i < len(energyLevels); i++ {
			for j := 0; j < len(energyLevels[i]); j++ {
				flashes += computeFlashes(i, j, flashed, &energyLevels)
			}
		}
	}

	fmt.Println(flashes)
}

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
