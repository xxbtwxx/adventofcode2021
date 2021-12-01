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

func taskB() {
	depthValues := depthReader()

	elements := make([]int, 0, 4)
	increments := 0
	for depthValue := range depthValues {
		elements = append(elements, depthValue)
		if len(elements) < 4 {
			continue
		}

		if sum(elements[:len(elements)-1]) < sum(elements[1:]) {
			increments++
		}

		elements = elements[1:]
	}

	fmt.Println(increments)
}

func sum(elements []int) (s int) {
	for _, el := range elements {
		s += el
	}

	return
}
