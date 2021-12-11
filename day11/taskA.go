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
