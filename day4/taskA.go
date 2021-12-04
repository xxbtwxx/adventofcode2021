package day4

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "a",
		Short: "Solution for day 4 problem A",
		Run: func(cmd *cobra.Command, args []string) {
			taskA()
		},
	})
}

func taskA() {
	draw, cards := inputReader()

	for num := range draw {
		for _, card := range cards {
			card.mark(num)
			if card.isBingo {
				fmt.Println(card.undrawnValue() * num)
				return
			}
		}
	}
}
