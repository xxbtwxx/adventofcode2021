package day11

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Solution for day 11 problem B",
		Run: func(cmd *cobra.Command, args []string) {
			taskB()
		},
	})
}

func taskB() {
	energyLevels := energyLevelsReader()

	for iteration := 1; ; iteration++ {
		flashed := map[int]map[int]struct{}{}
		iterationFlashes := 0
		for i := 0; i < len(energyLevels); i++ {
			for j := 0; j < len(energyLevels[i]); j++ {
				iterationFlashes += computeFlashes(i, j, flashed, &energyLevels)
			}
		}

		if iterationFlashes == len(energyLevels)*len(energyLevels[0]) {
			fmt.Println(iteration)
			return
		}
	}

}
