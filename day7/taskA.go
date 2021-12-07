package day7

import (
	"fmt"
	"math"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Solution for day 7 problem A",
		Run: func(cmd *cobra.Command, args []string) {
			taskA()
		},
	})
}

func taskA() {
	crabsAtPosition := map[int]int{}
	minPos := math.MaxInt64
	maxPos := math.MinInt64

	for pos := range positionReader() {
		crabsAtPosition[pos]++
		if pos > maxPos {
			maxPos = pos
		} else if pos < minPos {
			minPos = pos
		}
	}

	minFuel := math.MaxInt64

	for i := minPos; i < maxPos; i++ {
		currPosFuel := 0
		for pos, count := range crabsAtPosition {
			fuelCostToMove := (i - pos) * count
			if fuelCostToMove < 0 {
				fuelCostToMove *= -1
			}

			currPosFuel += fuelCostToMove
		}

		if currPosFuel < minFuel {
			minFuel = currPosFuel
		}
	}

	fmt.Println(minFuel)
}
