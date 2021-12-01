package day1

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Solution for day 1 problem A",
		Run: func(cmd *cobra.Command, args []string) {
			taskA()
		},
	})
}

func taskA() {
	depthValues := depthReader()

	previousValue := <-depthValues
	increments := 0
	for depthValue := range depthValues {
		if depthValue > previousValue {
			increments++
		}

		previousValue = depthValue
	}

	fmt.Println(increments)
}
