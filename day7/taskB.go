package day7

import (
	"fmt"
	"math"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Solution for day 7 problem B",
		Run: func(cmd *cobra.Command, args []string) {
			taskB()
		},
	})
}

func taskB() {
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
			posDelta := i - pos
			if posDelta < 0 {
				posDelta *= -1
			}
			fuelCostPerCrab := posDelta * (posDelta + 1) / 2

			currPosFuel += fuelCostPerCrab * count
		}

		if currPosFuel <= minFuel {
			minFuel = currPosFuel
		}
	}

	fmt.Println(minFuel)
}
