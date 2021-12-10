package day07

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
	for pos := range positionReader() {
		crabsAtPosition[pos]++

	}

	minFuel := math.MaxInt64
	for crabPosition := range crabsAtPosition {
		currPosFuel := 0
		for pos, count := range crabsAtPosition {
			fuelCostToMove := (crabPosition - pos) * count
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
