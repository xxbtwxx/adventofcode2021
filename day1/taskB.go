package day1

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Solution for day 1 problem B",
		Run: func(cmd *cobra.Command, args []string) {
			taskB()
		},
	})
}

const sumSize = 3

func taskB() {
	depthValues := depthReader()

	elements := make([]int, 0, sumSize+1)
	increments := 0
	for depthValue := range depthValues {
		elements = append(elements, depthValue)
		if len(elements) < sumSize+1 {
			continue
		}

		if elements[0] < elements[len(elements)-1] {
			increments++
		}

		elements = elements[1:]
	}

	fmt.Println(increments)
}
