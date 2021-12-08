package day8

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Solution for day 8 problem A",
		Run: func(cmd *cobra.Command, args []string) {
			taskA()
		},
	})
}

func taskA() {
	simpleDigits := 0
	for input := range inputParser() {
		for _, digit := range input.digits {
			switch len(digit) {
			case 2, 3, 4, 7:
				simpleDigits++
			}
		}

	}

	fmt.Println(simpleDigits)

}
