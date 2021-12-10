package day04

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	Cmd.AddCommand(&cobra.Command{
		Use:   "b",
		Short: "Solution for day 4 problem B",
		Run: func(cmd *cobra.Command, args []string) {
			taskB()
		},
	})
}

func taskB() {
	draw, cards := inputReader()

	bingoValue := 0
	for num := range draw {
		for index, card := range cards {
			card.mark(num)
			if card.isBingo {
				bingoValue = card.undrawnValue() * num
				delete(cards, index)
			}
		}
	}

	fmt.Println(bingoValue)
}
